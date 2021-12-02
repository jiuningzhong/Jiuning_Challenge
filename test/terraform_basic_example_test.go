package test

import (
        "testing"

        "github.com/gruntwork-io/terratest/modules/terraform"
        "github.com/stretchr/testify/assert"
        "crypto/tls"
        "fmt"
        http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
)

func TestTerraformHelloWorldExample(t *testing.T) {
        t.Parallel()

        expectedUrl := "https://18.207.204.47"

        terraformOptions := &terraform.Options{
                // Set the path to the Terraform code that will be tested.
                TerraformDir: "../",
                VarFiles:     []string{"test/http_test_input.tfvars"},
        }

        // Clean up resources with "terraform destroy" at the end of the test.
        defer terraform.Destroy(t, terraformOptions)

        // Run "terraform init" and "terraform apply". Fail the test if there are any errors.
        terraform.InitAndApply(t, terraformOptions)


        // Run `terraform output` to get the values of output variables and check they have the expected values.
        url := terraform.Output(t, terraformOptions, "url")
        assert.Equal(t, expectedUrl, url)

        // Perform an HTTP request on the resource and ensure we get a 200.
        tlsConfig := tls.Config{}
        statusCode, body := http_helper.HttpGet(t, fmt.Sprintf("%s", url), &tlsConfig)

        assert.Equal(t, 200, statusCode)
        assert.NotNil(t, body)
}
