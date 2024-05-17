package mailchimp

import (
	"context"

	"github.com/hanzoai/gochimp3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableMailchimpAutomationEmail(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mailchimp_automation_email",
		Description: "Get a summary of the emails in a classic automation workflow.",
		List: &plugin.ListConfig{
			ParentHydrate: listAutomations,
			Hydrate:       listAutomationEmails,
			KeyColumns:    plugin.OptionalColumns([]string{"workflow_id"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "workflow_id"}),
			Hydrate:    getAutomationEmail,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the automation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "archive_url",
				Description: "The link to the campaign's archive version.",
				Transform:   transform.FromField("ArchiveURL"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "content_type",
				Description: "How the campaign's content is put together.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_time",
				Description: "The date and time the campaign was created in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "emails_sent",
				Description: "The total number of emails sent for the campaign.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "position",
				Description: "The position of an automation email in a workflow.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "send_time",
				Description: "The date and time a campaign was sent in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "start_time",
				Description: "The date and time the campaign was started in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "status",
				Description: "The current status of the automation.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "workflow_id",
				Description: "A string that uniquely identifies an automation workflow.",
				Transform:   transform.FromField("WorkflowID"),
				Type:        proto.ColumnType_STRING,
			},

			// JSON Columns
			{
				Name:        "delay",
				Description: "The delay settings for the automation email.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "recipients",
				Description: "List settings for the campaign.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "report_summary",
				Description: "A summary of opens and clicks for sent campaigns.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "settings",
				Description: "Settings for the campaign including the email subject, from name, and from email address.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "social_card",
				Description: "The preview for the campaign, rendered by social networks like Facebook and Twitter.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tracking",
				Description: "The tracking options for the automation.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "trigger_settings",
				Description: "Available triggers for AutomationEmail workflows.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the automation email.",
				Transform:   transform.FromField("Settings.Title"),
				Type:        proto.ColumnType_STRING,
			},
		}),
	}
}

//// LIST FUNCTION

func listAutomationEmails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := h.Item.(*gochimp3.Automation).ID

	if d.EqualsQuals["workflow_id"] != nil && d.EqualsQualString("workflow_id") != id {
		return nil, nil
	}

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_automation_email.listAutomationEmails", "connection_error", err)
		return nil, err
	}

	automationEmails, err := client.GetAutomationEmails(id)
	if err != nil {
		logger.Error("mailchimp_automation_email.listAutomationEmails", "api_error", err)
		return nil, err
	}

	for _, automationEmail := range automationEmails.Emails {
		d.StreamListItem(ctx, &automationEmail)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAutomationEmail(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")
	workflowId := d.EqualsQualString("workflow_id")

	// Email id and workflow id should not be empty
	if id == "" || workflowId == "" {
		return nil, nil
	}

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_automation_email.getAutomationEmail", "connection_error", err)
		return nil, err
	}

	automationEmail, err := client.GetAutomationEmail(workflowId, id)
	if err != nil {
		logger.Error("mailchimp_automation_email.getAutomationEmail", "api_error", err)
		return nil, err
	}

	return automationEmail, nil
}
