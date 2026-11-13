package main

import (
	"github.com/turbot/steampipe-plugin-mailchimp/mailchimp"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: mailchimp.Plugin})
}
