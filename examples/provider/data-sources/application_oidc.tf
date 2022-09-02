data zitadel_application_oidc oidc_application {
  depends_on = [data.zitadel_org.org, data.zitadel_project.project]

  org_id     = data.zitadel_org.org.id
  project_id = data.zitadel_project.project.id
  app_id         = "177073626925760515"
}

output oidc_application {
  value = data.zitadel_application_oidc.oidc_application
}