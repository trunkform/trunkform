include "root" {
  path = find_in_parent_folders("remote_state.hcl")
}

terraform {
  source = "../../modules//www"
}

dependency "registered_domain" {
  config_path = "../"
}

inputs = {
  account     = "245760921574"
  domain_name = "trunkform.ai"
  environment = "cd"
  region      = "us-east-2"
  zone_id     = dependency.registered_domain.outputs.zone_id
}
