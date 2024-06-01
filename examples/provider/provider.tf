provider "discord" {
  token = var.discord_token
}

data "discord_local_image" "logo" {
  file = "logo.png"
}

resource "discord_server" "my_server" {
  name                          = "My Awesome Server"
  region                        = "us-west"
  default_message_notifications = 0
  icon_data_uri                 = data.discord_local_image.logo.data_uri
}
