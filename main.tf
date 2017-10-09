variable "org_network" {default="V2Net2"}
variable "edge_gw" {default="EG2"}
provider "pyvcloudvcd" {
  user                 = "user1"
  password             = "Admin!23"
  org                  = "O2"
  url                  = "https://10.112.83.27:443/api"
  vdc                  = "V03"
  max_retry_timeout      = "30"
  allow_unverified_ssl = "true"
}



resource "pyvcloudvcd_server" "my-server" {
username = "user1"
password = "Admin!23"
org	= "O1"
}


