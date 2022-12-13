terraform {
  backend "s3" {
    bucket = "ai-events-tfstate-integration"
    key    = "ai-events-streams/rhosak/terraform.tfstate"
  }
  required_providers {
    rhoas = {
      source  = "registry.terraform.io/redhat-developer/rhoas"
      version = "0.3"
    }
  }
}
