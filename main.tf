variable "region" {
  default = "us-east-1"
}

provider "aws" {
  region = "${var.region}"
}

variable "availability_zone" {
  default = "us-east-1"
}

variable "ec2_ami_id" {
  default = "ami-04902260ca3d33422" # amzn-ami-hvm-2018.03.0.20180508-x86_64-gp2
}

variable "instance_type" {
  default = "t2.micro"
}

variable "vpc_id" {
  default = ""
}

variable "subnet_id" {
  default = ""
}

variable "pem_key_name" {
  default = ""
}

variable "your_name" {
  default = ""
}

output "url" {
   value = "https://18.207.204.47"
}
