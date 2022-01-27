package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/lucky3028/discord-terraform/discord"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: discord.Provider})
}
