variable "region" {
  default = "ap-northeast-1"
}

provider "aws" {
  region = "${var.region}"
}

variable "availability_zone" {
  default = "ap-northeast-1c"
}

variable "ec2_ami_id" {
  default = "ami-92df37ed" # amzn-ami-hvm-2018.03.0.20180508-x86_64-gp2
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

resource "aws_security_group" "web_server" {
  name   = "${var.your_name}"
  vpc_id = "${var.vpc_id}"

  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"

    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 8080
    to_port   = 8080
    protocol  = "tcp"

    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 443
    to_port   = 443
    protocol  = "tcp"

    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.your_name}"
  }
}

resource "aws_instance" "web_server" {
  ami                         = "${var.ec2_ami_id}"
  availability_zone           = "${var.availability_zone}"
  ebs_optimized               = false
  instance_type               = "${var.instance_type}"
  monitoring                  = false
  key_name                    = "${var.pem_key_name}"
  subnet_id                   = "${var.subnet_id}"
  vpc_security_group_ids      = ["${aws_security_group.web_server.id}"]
  associate_public_ip_address = true
  source_dest_check           = true

  # Run a "Hello, World" web server on port 8080
  user_data = <<-EOF
              #!/bin/bash
              echo "Hello, World" > index.html
              python -m SimpleHTTPServer 8080 &
              EOF

  tags {
    Name = "${var.your_name}"
  }

  provisioner "remote-exec" {
    inline = [
      "while [ ! -f /var/lib/cloud/instance/boot-finished ]; do echo -e 'Waiting for cloud-init...'; sleep 1; done",
    ]

    connection {
      type        = "ssh"
      user        = "ec2-user"
      private_key = "${file("~/.ssh/${var.pem_key_name}.pem")}"
    }
  }

  provisioner "local-exec" {
    command = <<EOT
    echo 'Instance is running!'
EOT
  }
}

output "ssh" {
  value = "ssh -i ~/.ssh/${var.pem_key_name}.pem ec2-user@${aws_instance.web_server.public_ip}"
}

output "url" {
  value = "http://${aws_instance.web_server.public_ip}:8080"
}
