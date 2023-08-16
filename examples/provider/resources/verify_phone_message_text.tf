resource "zitadel_verify_phone_message_text" "verify_phone_en" {
  org_id   = zitadel_org.org.id
  language = "en"

  title       = "title example"
  pre_header  = "pre_header example"
  subject     = "subject example"
  greeting    = "greeting example"
  text        = "text example"
  button_text = "button_text example"
  footer_text = "footer_text example"
}