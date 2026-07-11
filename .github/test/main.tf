terraform {
  required_providers {
    appstore = {
      source  = "fintreal/appstore"
    }
  }
}

provider "appstore" { }

resource "appstore_bundle_identifier" "this" {
  name       = "Terraform Test Bundle Identifier"
  identifier = "com.terraform.test.workflow"
  capabilities = ["APPLE_ID_AUTH"]
}

resource "appstore_provisioning_profile" "this" {
  name       = "Terraform Test Provisioning Profile"
  bundle_identifier_id = appstore_bundle_identifier.this.id
  type = "IOS_APP_STORE"
  certificate_ids = ["64A2UF5HD4"]
}

data "appstore_certificate" "this" {
  serial_number = "6AB7A99542D0D58517914E6EA8119579"
}