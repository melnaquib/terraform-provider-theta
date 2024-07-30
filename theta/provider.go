package theta

import (
	// Documentation:
	// https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema
	//
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource:
// A 'thing' you create, and then manage (update/delete) via terraform.
//
// Data Source:
// Data you can get access to and reference within your resources.

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"foo": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("THETA_FOO", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			// Naming format...
			//
			// Map key: <provider>_<thing>
			// File:    resource_<provider>_<thing>.go
			//
			// NOTE:
			// The map key is what's documented as the 'thing' a consumer of this
			// provider would add to their terraform HCL file.
			// e.g. resource "mock_example" "my_own_name_for_this" {...}
			//
			// "theta_edge_function": resourceExample(),
			"theta_edge_function": resourceEdgeFunction(),
			"theta_subnet": resourceSubnet(),
			"theta_evm_contract": resourceEvmContract(),
			"theta_edge_vod": resourceEdgeVod(),
			"theta_edge_store": resourceEdgeStore(),
			"theta_edge_model": resourceEdgeModel(),
		},
		// DataSource is a subset of Resource.
		DataSourcesMap: map[string]*schema.Resource{
			// Naming format...
			//
			// Map key: <provider>_<thing>
			// File:    data_source_<provider>_<thing>.go
			//
			// NOTE:
			// The map key is what's documented as the 'thing' a consumer of this
			// provider would add to their terraform HCL file.
			// e.g. data_source "mock_example" "my_own_name_for_this" {...}
			//
			// In the provided Go code snippet, the `dataSource` function is used to
			// define a data source for the Terraform provider. Data sources in Terraform
			// allow you to fetch and reference external data within your Terraform
			// configurations.
			// "theta_edge_function": dataSourceExample(),
			"theta_edge_function": dataSourceEdgeFunction(),
			"theta_subnet": dataSourceSubnet(),
			"theta_evm_contract": dataSourceEvmContract(),
			"theta_edge_vod": dataSourceEdgeVod(),
			"theta_edge_store": dataSourceEdgeStore(),
			"theta_edge_model": dataSourceEdgeModel(),
		},

		// To configure the provider (i.e. create an API client)
		// then pass ConfigureFunc. The any value returned by this function
		// is stored and passed into the subsequent resources as the meta
		// parameter (this includes Data Sources as they are subsets of Resources).
		//
		// Documentation:
		// https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema#ConfigureFunc
		// https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema#ConfigureContextFunc
	}
}
