resource "discord_voice_channel" "general" {
  name      = "General"
  server_id = var.server_id
  position  = 0
}
