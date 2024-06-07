resource "discord_category_channel" "chatting" {
  name      = "Chatting"
  server_id = var.server_id
  position  = 0
}
