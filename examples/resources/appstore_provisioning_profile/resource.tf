terraform {
  required_providers {
    appstore = {
      source  = "fintreal/appstore"
      version = "~> 1.4"
    }
  }
}

provider "appstore" { }

resource "appstore_provisioning_profile" "this" {
  name       = "Test Terraform Provisioning Profile"
  bundle_identifier_id = "S65568LTPR"
  type = "IOS_APP_STORE"
  certificate_ids = ["JVLG7LVRRL"]
}
