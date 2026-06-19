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
  account       = "810322435582"
  domain_name   = "trunkform.io"
  environment   = "ci"
  force_destroy = true
  region        = "us-east-2"
  zone_id       = dependency.zone.outputs.zone_id
}
