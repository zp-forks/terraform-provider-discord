package discord

import (
	"github.com/andersfylling/disgord"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
)

func resourceDiscordInvite() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInviteCreate,
		ReadContext:   resourceInviteRead,
		DeleteContext: resourceInviteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"channel_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"max_age": {
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Default:  86400,
			},
			"max_uses": {
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"temporary": {
				Type:     schema.TypeBool,
				ForceNew: true,
				Optional: true,
			},
			"unique": {
				Type:     schema.TypeBool,
				ForceNew: true,
				Optional: true,
			},
			"code": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceInviteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Client

	channelId := getId(d.Get("channel_id").(string))

	if invite, err := client.Channel(channelId).CreateInvite(&disgord.CreateInvite{
		MaxAge:    d.Get("max_age").(int),
		MaxUses:   d.Get("max_uses").(int),
		Temporary: d.Get("temporary").(bool),
		Unique:    d.Get("unique").(bool),
	}); err != nil {
		return diag.Errorf("Failed to create a invite: %s", err.Error())
	} else {
		d.SetId(invite.Code)
		d.Set("code", invite.Code)

		return diags
	}
}

func resourceInviteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Client

	if invite, err := client.Invite(d.Id()).Get(false); err != nil {
		d.SetId("")
	} else {
		d.Set("code", invite.Code)
	}

	return diags
}

func resourceInviteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Client

	if _, err := client.Invite(d.Id()).Delete(); err != nil {
		return diag.FromErr(err)
	} else {
		return diags
	}
}
