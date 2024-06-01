data "discord_local_image" "logo" {
  file = "logo.png"
}

resource "discord_server" "server" {
  // ...
  icon_data_uri = data.discord_local_image.logo.data_uri
}
