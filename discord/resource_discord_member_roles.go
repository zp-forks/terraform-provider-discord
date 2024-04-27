package discord

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
)

type RoleSchema struct {
	RoleId  string `json:"role_id"`
	HasRole bool   `json:"has_role"`
}

func convertToRoleSchema(v interface{}) (*RoleSchema, error) {
	var roleSchema *RoleSchema
	j, _ := json.MarshalIndent(v, "", "    ")
	err := json.Unmarshal(j, &roleSchema)

	return roleSchema, err
}

func resourceDiscordMemberRoles() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMemberRolesCreate,
		ReadContext:   resourceMemberRolesRead,
		UpdateContext: resourceMemberRolesUpdate,
		DeleteContext: resourceMemberRolesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"server_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"role": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"has_role": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
			},
		},
	}
}

func resourceMemberRolesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*Context).Session

	serverId := d.Get("server_id").(string)
	userId := d.Get("user_id").(string)

	if _, err := client.GuildMember(serverId, userId, discordgo.WithContext(ctx)); err != nil {
		return diag.Errorf("Could not get member %s in %s: %s", userId, serverId, err.Error())
	}

	d.SetId(generateTwoPartId(serverId, userId))

	diags = append(diags, resourceMemberRolesRead(ctx, d, m)...)
	diags = append(diags, resourceMemberRolesUpdate(ctx, d, m)...)

	return diags
}

func resourceMemberRolesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Session

	// parse server ID and userID out of the ID:
	var serverId, userId string
	sId, uId, err := parseTwoIds(d.Id())
	if err != nil {
		log.Default().Printf("Unable to parse IDs out of the resource ID. Falling back on legacy config behavior.")
		serverId = d.Get("server_id").(string)
		userId = d.Get("user_id").(string)
	} else {
		serverId = sId
		userId = uId
	}

	member, err := client.GuildMember(serverId, userId, discordgo.WithContext(ctx))
	if err != nil {
		return diag.Errorf("Could not get member %s in %s: %s", userId, serverId, err.Error())
	}

	items := d.Get("role").(*schema.Set).List()
	roles := make([]*RoleSchema, 0, len(items))

	for _, r := range items {
		v, _ := convertToRoleSchema(r)
		if hasRole(member, v.RoleId) {
			roles = append(roles, &RoleSchema{RoleId: v.RoleId, HasRole: true})
		} else {
			roles = append(roles, &RoleSchema{RoleId: v.RoleId, HasRole: false})
		}
	}
	d.Set("role", roles)

	return diags
}

func resourceMemberRolesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Session

	serverId := d.Get("server_id").(string)
	userId := d.Get("user_id").(string)

	member, err := client.GuildMember(serverId, userId, discordgo.WithContext(ctx))
	if err != nil {
		return diag.Errorf("Could not get member %s in %s: %s", userId, serverId, err.Error())
	}

	oldRole, newRole := d.GetChange("role")
	oldItems := oldRole.(*schema.Set).List()
	items := newRole.(*schema.Set).List()

	roles := member.Roles

	for _, r := range items {
		v, _ := convertToRoleSchema(r)
		memberHasRole := hasRole(member, v.RoleId)
		// If it's supposed to have the role, and it does, continue
		if memberHasRole && v.HasRole {
			continue
		}
		// If it's supposed to have the role, and it doesn't, add it
		if v.HasRole && !memberHasRole {
			roles = append(roles, v.RoleId)
		}
		// If it's not supposed to have the role, and it does, remove it
		if !v.HasRole && memberHasRole {
			roles = removeRoleById(roles, v.RoleId)
		}
	}

	// If the change removed the role, and the user has it, remove it
	for _, r := range oldItems {
		v, _ := convertToRoleSchema(r)
		if wasRemoved(items, v) && v.HasRole {
			roles = removeRoleById(roles, v.RoleId)
		}
	}

	if _, err := client.GuildMemberEdit(serverId, userId, &discordgo.GuildMemberParams{
		Roles: &roles,
	}, discordgo.WithContext(ctx)); err != nil {
		return diag.Errorf("Failed to edit member %s: %s", userId, err.Error())
	}

	return diags
}

func wasRemoved(items []interface{}, v *RoleSchema) bool {
	for _, i := range items {
		item, _ := convertToRoleSchema(i)
		if item.RoleId == v.RoleId {
			return false
		}
	}

	return true
}

func resourceMemberRolesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Session
	serverId := d.Get("server_id").(string)
	userId := d.Get("user_id").(string)

	member, err := client.GuildMember(serverId, userId, discordgo.WithContext(ctx))
	if err != nil {
		return diag.Errorf("Could not get member %s in %s: %s", userId, serverId, err.Error())
	}

	items := d.Get("role").(*schema.Set).List()
	roles := member.Roles

	for _, r := range items {
		v, _ := convertToRoleSchema(r)
		hasRole := hasRole(member, v.RoleId)
		// if it's supposed to have the role, and it does, remove it
		if hasRole && v.HasRole {
			roles = removeRoleById(roles, v.RoleId)
		}
	}

	if _, err := client.GuildMemberEdit(serverId, userId, &discordgo.GuildMemberParams{
		Roles: &roles,
	}, discordgo.WithContext(ctx)); err != nil {
		return diag.Errorf("Failed to delete member roles %s: %s", userId, err.Error())
	}

	return diags
}
