data "discord_server" "discord_api" {
  server_id = "81384788765712384"
}

output "discord_api_region" {
  value = data.discord_server.discord_api.region
}
