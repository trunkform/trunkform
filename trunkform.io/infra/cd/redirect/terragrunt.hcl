include "root" {
  path = find_in_parent_folders("remote_state.hcl")
}

terraform {
  source = "../../modules//redirect"
}

dependency "registered_domain" {
  config_path = "../"
}

dependency "www" {
  config_path = "../../../../trunkform.ai/infra/cd/www"
}

inputs = {
  domain_name            = "trunkform.io"
  zone_id                = dependency.registered_domain.outputs.zone_id
  cloudfront_domain_name = dependency.www.outputs.distribution_domain_name
  cloudfront_zone_id     = dependency.www.outputs.distribution_zone_id
}
