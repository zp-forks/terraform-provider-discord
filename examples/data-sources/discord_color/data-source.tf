data "discord_color" "blue" {
  hex = "#4287f5"
}

data "discord_color" "green" {
  rgb = "rgb(46, 204, 113)"
}

resource "discord_role" "blue" {
  // ...
  color = data.discord_color.blue.dec
}
resource "discord_role" "green" {
  // ...
  color = data.discord_color.green.dec
}
