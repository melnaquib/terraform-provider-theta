terraform {
  required_providers {
    mock = {
      source = "integralist/mock"
    }
  }
}

provider "mock" {
  foo = "example_value"
  #
  # if 'foo' wasn't set here by us, then the value would default to the value 
  # assigned to the environment variable 'MOCK_FOO' or the default value of nil
  # if the environment variable wasn't set.
}
 
resource "mock_theta_edge_function" "terraform_lambda_func20" {
filename                       = "thetafn/thetafn.js"
function_name                  = "thetafn33"
handler                        = "main"
runtime                        = "denotheta.2"
}
