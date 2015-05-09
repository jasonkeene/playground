module "consul" {
    source = "github.com/hashicorp/consul/terraform/aws"

    key_name = "${var.key_name}"
    key_path = "${var.key_path}"
    region = "${var.region}"
    servers = "3"
}
