resource "discord_channel_permission" "chatting" {
  channel_id   = var.channel_id
  type         = "role"
  overwrite_id = var.role_id
  allow        = data.discord_permission.chatting.allow_bits
}
