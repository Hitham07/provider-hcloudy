package hcloud_network_subnet

import "github.com/crossplane/upjet/pkg/config"

// Configure sets up resource configurations for multiple hcloud resources.
func Configure(p *config.Provider) {

	// Configure the hcloud_network_subnet resource
	p.AddResourceConfigurator("hcloud_network_subnet", func(r *config.Resource) {
		r.ShortGroup = "hcloud"
		r.References["network_id"] = config.Reference{
			TerraformName:     "hcloud_network",
		}

	})

}
