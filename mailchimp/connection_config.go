package mailchimp

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type mailchimpConfig struct {
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &mailchimpConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) mailchimpConfig {
	if connection == nil || connection.Config == nil {
		return mailchimpConfig{}
	}
	config, _ := connection.Config.(mailchimpConfig)
	return config
}
