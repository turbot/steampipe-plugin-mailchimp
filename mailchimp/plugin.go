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
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "account_id",
				Hydrate: getAccountId,
			},
		},
		TableMap: map[string]*plugin.Table{
			"mailchimp_authorized_app":   tableMailchimpAuthorizedApp(ctx),
			"mailchimp_automation_email": tableMailchimpAutomationEmail(ctx),
			"mailchimp_automation_queue": tableMailchimpAutomationQueue(ctx),
			"mailchimp_automation":       tableMailchimpAutomation(ctx),
			"mailchimp_batch_operation":  tableMailchimpBatchOperation(ctx),
			"mailchimp_campaign_folder":  tableMailchimpCampaignFolder(ctx),
			"mailchimp_campaign":         tableMailchimpCampaign(ctx),
			"mailchimp_list":             tableMailchimpList(ctx),
			"mailchimp_root":             tableMailchimpRoot(ctx),
			"mailchimp_store":            tableMailchimpStore(ctx),
			"mailchimp_template_folder":  tableMailchimpTemplateFolder(ctx),
			"mailchimp_template":         tableMailchimpTemplate(ctx),
		},
	}

	return p
}
