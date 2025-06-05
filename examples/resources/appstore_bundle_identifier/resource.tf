resource "appstore_bundle_identifier" "this" {
  name       = "Terraform Test Bundle Identifier"
  identifier = "com.terraform.test"
  capabilities = ["APPLE_ID_AUTH"]
}
