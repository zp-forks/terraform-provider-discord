data "discord_permission" "member" {
  view_channel     = "allow"
  send_messages    = "allow"
  use_vad          = "deny"
  priority_speaker = "deny"
}
data "discord_permission" "moderator" {
  allow_extends    = data.discord_permission.member.allow_bits
  deny_extends     = data.discord_permission.member.deny_bits
  kick_members     = "allow"
  ban_members      = "allow"
  manage_nicknames = "allow"
  view_audit_log   = "allow"
  priority_speaker = "allow"
}
resource "discord_role" "member" {
  // ...
  permissions = data.discord_permission.member.allow_bits
}
resource "discord_role" "moderator" {
  // ...
  permissions = data.discord_permission.moderator.allow_bits
}
resource "discord_channel_permission" "general_mod" {
  type         = "role"
  overwrite_id = discord_role.moderator.id
  allow        = data.discord_permission.moderator.allow_bits
  deny         = data.discord_permission.moderator.deny_bits
}
