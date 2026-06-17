include "root" {
  path = find_in_parent_folders("remote_state.hcl")
}

terraform {
  source = "../modules//registered_domain"
}

inputs = {
  domain_name = "trunkform.io"
  environment = "cd"
  name_servers = [
    "ns-1157.awsdns-16.org",
    "ns-1631.awsdns-11.co.uk",
    "ns-41.awsdns-05.com",
    "ns-988.awsdns-59.net",
  ]
}
