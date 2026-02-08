include "root" {
  path = find_in_parent_folders("remote_state.hcl")
}

terraform {
  source = "../modules//mcp"
}

inputs = {
  environment = "ci"
}
