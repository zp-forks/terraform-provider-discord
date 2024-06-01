resource "discord_invite" "chatting" {
  channel_id = var.channel_id
  max_age    = 0
}
