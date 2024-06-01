package discord

import (
	"context"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/go-playground/colors.v1"
)

func dataSourceDiscordColor() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscordColorRead,
		Description: "A simple helper to get the integer representation of a hex or rgb color",
		Schema: map[string]*schema.Schema{
			"hex": {
				ExactlyOneOf: []string{"hex", "rgb"},
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The hex color code. One of these must be present",
			},
			"rgb": {
				ExactlyOneOf: []string{"hex", "rgb"},
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The RGB color (format: `rgb(R, G, B)`). One of these must be present",
			},
			"dec": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The integer representation of the passed color",
			},
		},
	}
}

func ConvertToInt(hex string) (int64, error) {
	hex = strings.Replace(hex, "0x", "", 1)
	hex = strings.Replace(hex, "0X", "", 1)
	hex = strings.Replace(hex, "#", "", 1)

	return strconv.ParseInt(hex, 16, 64)
}

func dataSourceDiscordColorRead(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	var hex string
	if v, ok := d.GetOk("hex"); ok {
		if clr, err := colors.ParseHEX(v.(string)); err != nil {
			return diag.Errorf("Failed to parse hex %s: %s", v.(string), err.Error())
		} else {
			hex = clr.String()
		}
	}
	if v, ok := d.GetOk("rgb"); ok {
		if clr, err := colors.ParseRGB(v.(string)); err != nil {
			return diag.Errorf("Failed to parse rgb %s: %s", v.(string), err.Error())
		} else {
			hex = clr.ToHEX().String()
		}
	}

	if intColor, err := ConvertToInt(hex); err != nil {
		return diag.Errorf("Failed to parse hex %s: %s", hex, err.Error())
	} else {
		d.SetId(strconv.Itoa(int(intColor)))
		d.Set("dec", int(intColor))

		return diags
	}
}
