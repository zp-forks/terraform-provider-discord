# Managed Discord Server Resource

A resource to create a server

## Example Usage

```hcl-terraform
resource discord_managed_server my_server {
    server_id = "my-server-id"
}
```

## Argument Reference

* `server_id` (Required) The ID of the server to manage
* `name` (Optional) Name of the server
* `region` (Optional) Region of the server
* `verification_level` (Optional) Verification Level of the server
* `explicit_content_filter` (Optional) Explicit Content Filter level
* `default_message_notifications` (Optional) Default Message Notification settings (0 = all messages, 1 = mentions)
* `afk_channel_id` (Optional) Channel ID for moving AFK users to
* `af_timeout` (Optional)  many seconds before moving an AFK user
* `icon_url` (Optional) Remote URL for setting the icon of the server
* `icon_data_uri` (Optional) Data URI of an image to set the icon
* `splash_url` (Optional) Remote URL for setting the splash of the server
* `splash_data_uri` (Optional) Data URI of an image to set the splash
* `owner_id` (Optional) Owner ID of the server (Setting this will transfer ownership)
* `system_channel_id` (Optional) Channel ID for system messages

## Attribute Reference

* `icon_hash` Hash of the icon
* `splash_hash` Hash of the splash