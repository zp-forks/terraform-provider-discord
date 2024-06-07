data "discord_local_image" "avatar" {
  file = "avatar.png"
}

resource "discord_webhook" "webhook" {
  channel_id      = var.channel_id
  name            = "Bot"
  avatar_data_uri = data.discord_local_image.avatar.data_uri
}

output "webhook-url" {
  value = nonsensitive(discord_webhook.webhook.url)
}
