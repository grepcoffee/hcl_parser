module "tfplan-functions" {
  source = "../common-functions/tfplan-functions/tfplan-functions.sentinel"
}

module "tfstate-functions" {
  source = "../common-functions/tfstate-functions/tfstate-functions.sentinel"
}

module "tfconfig-functions" {
  source = "../common-functions/tfconfig-functions/tfconfig-functions.sentinel"
}

module "aws-functions" {
  source = "./aws-functions/aws-functions.sentinel"
}

policy "enforce-mandatory-tags" {
  source = "./enforce-mandatory-tags.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "protect-against-rds-instance-deletion" {
  source = "./protect-against-rds-instance-deletion.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "require-dns-support-for-vpcs" {
  source = "./require-dns-support-for-vpcs.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "require-most-recent-AMI-version" {
  source = "./require-most-recent-AMI-version.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "require-private-acl-and-kms-for-s3-buckets" {
  source = "./require-private-acl-and-kms-for-s3-buckets.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "require-vpc-and-kms-for-lambda-functions" {
  source = "./require-vpc-and-kms-for-lambda-functions.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-ami-owners" {
  source = "./restrict-ami-owners.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-assumed-roles-by-workspace" {
  source = "./restrict-assumed-roles-by-workspace.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-assumed-roles" {
  source = "./restrict-assumed-roles.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-availability-zones" {
  source = "./restrict-availability-zones.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-current-ec2-instance-type" {
  source = "./restrict-current-ec2-instance-type.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-db-instance-engines" {
  source = "./restrict-db-instance-engines.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-ec2-instance-type" {
  source = "./restrict-ec2-instance-type.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-egress-sg-rule-cidr-blocks" {
  source = "./restrict-egress-sg-rule-cidr-blocks.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-eks-node-group-size" {
  source = "./restrict-eks-node-group-size.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-iam-policy-actions" {
  source = "./restrict-iam-policy-actions.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-ingress-sg-rule-cidr-blocks" {
  source = "./restrict-ingress-sg-rule-cidr-blocks.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-ingress-sg-rule-rdp" {
  source = "./restrict-ingress-sg-rule-rdp.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-ingress-sg-rule-ssh" {
  source = "./restrict-ingress-sg-rule-ssh.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-launch-configuration-instance-type" {
  source = "./restrict-launch-configuration-instance-type.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-s3-bucket-policies" {
  source = "./restrict-s3-bucket-policies.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-sagemaker-notebooks" {
  source = "./restrict-sagemaker-notebooks.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "restrict-subnet-of-ec2-instances" {
  source = "./restrict-subnet-of-ec2-instances.sentinel"
  enforcement_level = "hard-mandatory"
}

policy "validate-providers-from-desired-regions" {
  source = "./validate-providers-from-desired-regions.sentinel"
  enforcement_level = "hard-mandatory"
}
