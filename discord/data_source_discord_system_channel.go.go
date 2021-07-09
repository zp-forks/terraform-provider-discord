package discord

import (
    "context"
    "github.com/andersfylling/disgord"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscordSystemChannel() *schema.Resource {
    return &schema.Resource{
        ReadContext: dataSourceDiscordSystemChannelRead,
        Schema: map[string]*schema.Schema{
            "server_id": {
                Type:         schema.TypeString,
                Required:     true,
            },
            "system_channel_id": {
                Type:     schema.TypeInt,
                Computed: true,
            },
        },
    }
}

func dataSourceDiscordSystemChannelRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    var err error
    var server *disgord.Guild
    client := m.(*Context).Client

    if v, ok := d.GetOk("server_id"); ok {
        server, err = client.GetGuild(ctx, getId(v.(string)))
        if err != nil {
            return diag.Errorf("Failed to fetch server %s: %s", v.(string), err.Error())
        }
    }

    d.Set("system_channel_id", server.SystemChannelID)

    return diags
}
