resource "discord_role" "moderator" {
  server_id   = var.server_id
  name        = "Moderator"
  permissions = data.discord_permission.moderator.allow_bits
  color       = data.discord_color.blue.dec
  hoist       = true
  mentionable = true
  position    = 5
}
