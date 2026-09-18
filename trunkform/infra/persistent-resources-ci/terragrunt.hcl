include "root" {
  path = find_in_parent_folders("remote_state.hcl")
}

terraform {
  source = "../modules//persistent-resources"
}

inputs = {
  environment = "persistent-resources-ci"
}
