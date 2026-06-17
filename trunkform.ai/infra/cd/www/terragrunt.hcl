include "root" {
  path = find_in_parent_folders("remote_state.hcl")
}

terraform {
  source = "../../modules//www"
}

dependency "zone" {
  config_path = "../zone"
}

inputs = {
  account     = "245760921574"
  domain_name = "trunkform.ai"
  environment = "cd"
  region      = "us-east-2"
  zone_id     = dependency.zone.outputs.zone_id
}
