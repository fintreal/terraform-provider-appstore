terraform {
  required_providers {
    eas = {
      source  = "fintreal/appstore"
    }
  }
}

provider "eas" {
  appstore_key           = "APPSTORE_KEY"
  appstore_key_id        = "APPSTORE_KEY_ID"
  appstore_key_issuer_id = "APPSTORE_KEY_ISSUER_ID"
}
