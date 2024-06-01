resource "discord_message" "hello_world" {
  channel_id = var.channel_id
  content    = "hello world"
}
