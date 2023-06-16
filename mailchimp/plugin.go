package mailchimp

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-mailchimp"

// Plugin creates this (mailchimp) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultGetConfig: &plugin.GetConfig{},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"mailchimp_authorized_app": tableMailchimpAuthorizedApp(ctx),
			"mailchimp_campaign":       tableMailchimpCampaign(ctx),
			"mailchimp_list":           tableMailchimpList(ctx),
		},
	}

	return p
}
