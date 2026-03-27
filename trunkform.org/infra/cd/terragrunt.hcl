include "root" {
  path = find_in_parent_folders("remote_state.hcl")
}

terraform {
  source = "../modules//registered_domain"
}

inputs = {
  domain_name = "trunkform.org"
  environment = "cd"
}
