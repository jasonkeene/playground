provider "aws" {
    access_key = "${var.access_key}"
    secret_key = "${var.secret_key}"
    region = "${var.region}"
}

resource "aws_instance" "example" {
    ami = "${lookup(var.amis, var.region)}"
    instance_type = "t1.micro"
    security_groups = [
        "allow_all"
    ]
    depends_on = "aws_security_group.allow_all"

    provisioner "local-exec" {
        command = "echo ${aws_instance.example.public_ip} > file.txt"
    }
}

resource "aws_eip" "ip" {
    instance = "${aws_instance.example.id}"
}

resource "aws_security_group" "allow_all" {
  name = "allow_all"
  description = "Allow all inbound traffic"

  ingress {
      from_port = 0
      to_port = 65535
      protocol = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
  }
}
