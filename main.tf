variable "org_network" {default="V2Net2"}
variable "edge_gw" {default="EG2"}
provider "pyvcloudvcd" {
  user                 = "user1"
  password             = "Admin!23"
  org                  = "O1"
  ip                   = "10.112.83.27"
  vdc                  = "V03"
  max_retry_timeout      = "30"
  allow_unverified_ssl = "true"
}



resource "pyvcloudvcd_server" "my-server" {
username = "user1"
password = "Admin!23"
org	= "O1"
}


