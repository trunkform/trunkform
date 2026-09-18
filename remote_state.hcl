locals {
  remote_url   = trimspace(run_cmd("--terragrunt-quiet", "git", "config", "--get", "remote.origin.url"))
  organization = regex(".+[:/]+([^/]+)/[^/]+(?:\\.git)?$", local.remote_url)[0]
  repository   = replace(basename(get_repo_root()), ".nosync", "")
  relative     = get_path_from_repo_root()
  account      = can(regex("[/ -]cd", local.relative)) ? "245760921574" : can(regex("[/ -]ci", local.relative)) ? "810322435582" : null # trunkform-cd : trunkform-ci : null
  region       = can(regex("[/ -]cd", local.relative)) ? "us-east-2" : can(regex("[/ -]ci", local.relative)) ? "us-east-2" : null       # use2         : use2         : null

  # The mcp/ directory was renamed to trunkform/ (see trunkform/CHANGELOG.md). Existing S3 state
  # was written under the "mcp/" key prefix; map it back so we don't have to migrate state.
  relative_parts = split("/", local.relative)
  state_relative  = local.relative_parts[0] == "trunkform" ? join("/", concat(["mcp"], slice(local.relative_parts, 1, length(local.relative_parts)))) : local.relative

  key = "state/${local.organization}/${local.repository}/${local.state_relative}/terraform.tfstate"
}

remote_state {
  backend = "s3"
  generate = {
    path      = "backend.tf"
    if_exists = "overwrite_terragrunt"
  }
  config = {
    bucket  = "trunkform-state-${local.account}"
    key     = local.key
    region  = local.region
    encrypt = true
  }
}
