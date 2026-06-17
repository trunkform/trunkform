include "root" {
  path = find_in_parent_folders("remote_state.hcl")
}

terraform {
  source = "../../modules//zone"
}

inputs = {
  domain_name = "trunkform.ai"
  environment = "cd"
}
