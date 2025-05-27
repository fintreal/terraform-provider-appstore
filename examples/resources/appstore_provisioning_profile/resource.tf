resource "appstore_provisioning_profile" "this" {
  name       = "Terraform Test Provisioning Profile"
  bundle_identifier_id = "..."
  type = "IOS_APP_STORE"
  certificate_ids = ["..."]
}
