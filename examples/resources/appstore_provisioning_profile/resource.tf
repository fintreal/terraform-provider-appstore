resource "appstore_provisioning_profile" "this" {
  name       = "Test Terraform Provisioning Profile"
  bundle_identifier_id = "..."
  type = "IOS_APP_STORE"
  certificate_ids = ["..."]
}
