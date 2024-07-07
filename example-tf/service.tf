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
 
resource "mock_fleek_edge_function" "terraform_lambda_func20" {
filename                       = "fleekfn/fleekfn.js"
function_name                  = "fleekfn33"
handler                        = "main"
runtime                        = "denofleek.2"
}
