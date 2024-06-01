resource "discord_member_roles" "jake" {
  user_id   = var.user_id
  server_id = var.server_id
  role {
    role_id = var.role_id_to_add
  }
  role {
    role_id  = var.role_id_to_always_remove
    has_role = false
  }
}
