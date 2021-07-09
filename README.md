# Discord Terraform Provider

This is a fork of [aequasi/terraform-provider-discord](https://github.com/aequasi/terraform-provider-discord). We ran into some problems with this provider and decided to fix them with this opinionated version.

https://registry.terraform.io/providers/Chaotic-Logic/discord/latest

## Building the provider
### Development
```sh
go mod vendor
make
```

### Release
```
go mod vendor
export GPG_FINGERPRINT="D081560F57E59EDA7CB369BE2FFBD6BE37B85C17"
goreleaser release --skip-publish
```

## Resources

* discord_category_channel
* discord_channel_permission
* discord_invite
* discord_member_roles
* discord_message
* discord_role
* discord_role_everyone
* discord_server
* discord_text_channel
* discord_voice_channel

## Data

* discord_color
* discord_local_image
* discord_permission

## Todo

#### Data Sources

* discord_category_channel
* discord_text_channel
* discord_voice_channel
