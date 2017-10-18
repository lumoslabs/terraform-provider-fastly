package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/lumoslabs/terraform-provider-fastly/fastly"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fastly.Provider})
}
