resource "discord_text_channel" "general" {
  name      = "general"
  server_id = var.server_id
  position  = 0
}
