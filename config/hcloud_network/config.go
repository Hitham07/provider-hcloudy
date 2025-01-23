package hcloud_network

import "github.com/crossplane/upjet/pkg/config"

// Configure sets up resource configurations for multiple hcloud resources.
func Configure(p *config.Provider) {
	// Configure the hcloud_network resource
	p.AddResourceConfigurator("hcloud_network", func(r *config.Resource) {
		r.ShortGroup = "hcloud"
	})

}
