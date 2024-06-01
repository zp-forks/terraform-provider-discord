resource "discord_news_channel" "general" {
  name      = "general"
  server_id = var.server_id
  position  = 0
}
