package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/polds/imgbase64"
	"golang.org/x/net/context"
)

func resourceDiscordWebhook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWebhookCreate,
		ReadContext:   resourceWebhookRead,
		UpdateContext: resourceWebhookUpdate,
		DeleteContext: resourceWebhookDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"channel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"avatar_url": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"avatar_data_uri"},
			},
			"avatar_data_uri": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"avatar_url"},
			},
			"avatar_hash": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"token": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"url": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"slack_url": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"github_url": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceWebhookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Session

	channelId := d.Get("channel_id").(string)

	avatar := ""
	if v, ok := d.GetOk("avatar_url"); ok {
		avatar = imgbase64.FromRemote(v.(string))
	}
	if v, ok := d.GetOk("avatar_data_uri"); ok {
		avatar = v.(string)
	}

	if webhook, err := client.WebhookCreate(channelId, d.Get("name").(string), avatar, discordgo.WithContext(ctx)); err != nil {
		return diag.Errorf("Failed to create webhook: %s", err.Error())
	} else {
		url := "https://discord.com/api/webhooks/" + webhook.ID + "/" + webhook.Token

		d.SetId(webhook.ID)
		d.Set("avatar_hash", webhook.Avatar)
		d.Set("token", webhook.Token)
		d.Set("url", url)
		d.Set("slack_url", url+"/slack")
		d.Set("github_url", url+"/github")

		return diags
	}
}

func resourceWebhookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Session

	if webhook, err := client.Webhook(d.Id(), discordgo.WithContext(ctx)); err != nil {
		d.SetId("")
	} else {
		url := "https://discord.com/api/webhooks/" + webhook.ID + "/" + webhook.Token

		d.Set("channel_id", webhook.ChannelID)
		d.Set("name", webhook.Name)
		d.Set("avatar_hash", webhook.Avatar)
		d.Set("token", webhook.Token)
		d.Set("url", url)
		d.Set("slack_url", url+"/slack")
		d.Set("github_url", url+"/github")
	}

	return diags
}

func resourceWebhookUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Session

	channelId := d.Get("channel_id").(string)
	name := d.Get("name").(string)

	avatar := ""
	if v, ok := d.GetOk("avatar_url"); ok {
		avatar = imgbase64.FromRemote(v.(string))
	}
	if v, ok := d.GetOk("avatar_data_uri"); ok {
		avatar = v.(string)
	}

	if webhook, err := client.WebhookEdit(d.Id(), name, avatar, channelId, discordgo.WithContext(ctx)); err != nil {
		return diag.Errorf("Failed to update webhook %s: %s", d.Id(), err.Error())
	} else {
		url := "https://discord.com/api/webhooks/" + webhook.ID + "/" + webhook.Token
		d.Set("channel_id", webhook.ChannelID)
		d.Set("name", webhook.Name)
		d.Set("avatar_hash", webhook.Avatar)
		d.Set("token", webhook.Token)
		d.Set("url", url)
		d.Set("slack_url", url+"/slack")
		d.Set("github_url", url+"/github")
	}

	return diags
}

func resourceWebhookDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Session

	if err := client.WebhookDelete(d.Id(), discordgo.WithContext(ctx)); err != nil {
		return diag.FromErr(err)
	} else {
		return diags
	}
}
