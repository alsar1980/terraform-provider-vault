package main

import (
	"github.com/alsar1980/terraform-provider-vault/vault"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vault.Provider})
}
