data "discord_member" "jake" {
  server_id = "81384788765712384"
  user_id   = "103559217914318848"
}

output "jakes_username_and_discrim" {
  value = "${data.discord_member.jake.username}#${data.discord_member.jake.discriminator}"
}
