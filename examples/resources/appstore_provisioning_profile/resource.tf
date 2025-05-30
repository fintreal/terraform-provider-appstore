resource "appstore_bundle_identifier" "this" {
  name       = "Terraform Test Bundle Identifier"
  identifier = "com.terraform.test"
  platform   = "UNIVERSAL"
  capabilities = ["APPLE_ID_AUTH"]
}

resource "appstore_provisioning_profile" "this" {
  name       = "Terraform Test Provisioning Profile"
  bundle_identifier_id = appstore_bundle_identifier.this.id
  type = "IOS_APP_STORE"
  certificate_ids = ["..."]
}
