data "zitadel_org_idp_saml" "default" {
  org_id = data.zitadel_org.default.id
  id     = "123456789012345678"
}
