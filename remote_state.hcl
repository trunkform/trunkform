locals {
  remote_url   = trimspace(run_cmd("--terragrunt-quiet", "git", "config", "--get", "remote.origin.url"))
  organization = regex(".+[:/]+([^/]+)/[^/]+(?:\\.git)?$", local.remote_url)[0]
  repository   = replace(basename(get_repo_root()), ".nosync", "")
  relative     = get_path_from_repo_root()
  account      = can(regex("[/ -]cd", local.relative)) ? "245760921574" : can(regex("[/ -]ci", local.relative)) ? "810322435582" : null # trunkform-cd : trunkform-ci : null
  region       = can(regex("[/ -]cd", local.relative)) ? "us-east-2" : can(regex("[/ -]ci", local.relative)) ? "us-east-2" : null       # use2         : use2         : null
  key          = "state/${local.organization}/${local.repository}/${local.relative}/terraform.tfstate"
}

remote_state {
  backend = "s3"
  generate = {
    path      = "backend.tf"
    if_exists = "overwrite_terragrunt"
  }
  config = {
    bucket  = "terraform-state-${local.account}"
    key     = local.key
    region  = local.region
    encrypt = true
  }
}
