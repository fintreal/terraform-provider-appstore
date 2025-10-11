terraform {
  required_providers {
    appstore = {
      source  = "fintreal/appstore"
    }
  }
}

provider "appstore" {
  appstore_key           = "..."
  appstore_key_id        = "..."
  appstore_key_issuer_id = "..."
}
