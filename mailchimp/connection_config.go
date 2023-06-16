package mailchimp

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type mailchimpConfig struct {
	MailchimpAPIKey *string `cty:"mailchimp_api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"mailchimp_api_key": {
		Type: schema.TypeString,
	},
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
