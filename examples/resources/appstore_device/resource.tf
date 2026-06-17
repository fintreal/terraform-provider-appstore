# Register a device. Apple cannot delete devices, so `terraform destroy`
# (or removing this from config) disables the device instead of deleting it.
resource "appstore_device" "this" {
  name     = "Terraform Test Device"
  udid     = "00008120-001A742226E3601E"
  platform = "IOS"
}
