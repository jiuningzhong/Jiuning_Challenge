package test

import (
	"os"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestWebServer(t *testing.T) {
	vpcID := os.Getenv("VPC_ID")
	subnetID := os.Getenv("SUBNET_ID")
	pemKeyName := os.Getenv("PEM_KEY_NAME")
	yourName := os.Getenv("YOUR_NAME")

	terraformOptions := &terraform.Options{
		// The path to where your Terraform code is located
		TerraformDir: "./tf",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"vpc_id":       vpcID,
			"subnet_id":    subnetID,
			"pem_key_name": pemKeyName,
			"your_name":    yourName,
		},
	}
	// At the end of the test, run `terraform destroy`
	defer terraform.Destroy(t, terraformOptions)
	// Run `terraform init` and `terraform apply`
	terraform.InitAndApply(t, terraformOptions)
	// Run `terraform output` to get the value of an output variable
	url := terraform.Output(t, terraformOptions, "url")
	// Verify that we get back a 200 OK with the expected text. It
	// takes ~1 min for the Instance to boot, so retry a few times.
	status := 200
	text := "Hello, World"
	retries := 3
	sleep := 3 * time.Second
	http_helper.HttpGetWithRetry(t, url, status, text, retries, sleep)
}
