terraform {
  required_providers {
    appstore = {
      source  = "fintreal/appstore"
      version = "~> 1.4"
    }
  }
}

provider "appstore" { }

resource "appstore_bundle_identifier" "this" {
  name       = "Terraform Test Bundle Identifier"
  identifier = "com.fintreal.terraform.test"
  platform   = "UNIVERSAL"
  capabilities = ["APPLE_ID_AUTH"]
}
