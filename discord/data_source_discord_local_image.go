package discord

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/polds/imgbase64"
)

func dataSourceDiscordLocalImage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscordLocalImageRead,
		Description: "A simple helper to get data uri of a local image",
		Schema: map[string]*schema.Schema{
			"file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The path to the file to process",
			},
			"data_uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The data uri of the `file`",
			},
		},
	}
}

func dataSourceDiscordLocalImageRead(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	file := d.Get("file").(string)

	if img, err := imgbase64.FromLocal(file); err != nil {
		return diag.Errorf("Failed to process %s: %s", file, err.Error())
	} else {
		d.Set("data_uri", img)
		d.SetId(strconv.Itoa(Hashcode(d.Get("data_uri").(string))))

		return diags
	}
}
