terraform {
  required_providers {
    zitadel = {
      source  = "zitadel/zitadel"
      version = "1.0.0-alpha.14"
    }
  }
}

provider "zitadel" {
  domain           = "localhost"
  insecure         = "true"
  port             = "8080"
  project          = "170832731415117995"
  jwt_profile_file = "local-token"
}
