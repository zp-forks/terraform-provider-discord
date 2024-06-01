data "discord_role" "mods_id" {
  server_id = "81384788765712384"
  role_id   = "175643578071121920"
}
data "discord_role" "mods_name" {
  server_id = "81384788765712384"
  name      = "Mods"
}

output "mods_color" {
  value = data.discord_role.mods_id.color
}
