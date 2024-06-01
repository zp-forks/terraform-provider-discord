resource "discord_role_everyone" "everyone" {
  server_id   = var.server_id
  permissions = data.discord_permission.everyone.allow_bits
}
