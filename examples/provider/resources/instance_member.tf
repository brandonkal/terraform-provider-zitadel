resource "zitadel_instance_member" "instance_member" {
  user_id = zitadel_human_user.human_user.id
  roles   = ["IAM_OWNER"]
}