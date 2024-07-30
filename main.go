package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/melnaquib/terraform-provider-theta/theta"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: theta.Provider,
	})
}
