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
        
        // Verify that we get back a 200 OK with the expected text. It
        // takes ~1 min for the Instance to boot, so retry a few times.
        status := 200
        text := "Hello World!"
        retries := 3
        sleep := 3 * time.Second
        http_helper.HttpGetWithRetry(t, url, &tlsConfig, status, text, retries, sleep)
}
