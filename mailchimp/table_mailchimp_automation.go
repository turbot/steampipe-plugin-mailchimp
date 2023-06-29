package mailchimp

import (
	"context"

	"github.com/hanzoai/gochimp3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableMailchimpAutomation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mailchimp_automation",
		Description: "Get a summary of an account's classic automations.",
		List: &plugin.ListConfig{
			Hydrate:    listAutomations,
			KeyColumns: plugin.OptionalColumns([]string{"status"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getAutomation,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the automation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "create_time",
				Description: "The date and time the automation was created in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "emails_sent",
				Description: "The total number of emails sent for the automation.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "start_time",
				Description: "The date and time the automation was started in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "status",
				Description: "The current status of the automation.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON Columns
			{
				Name:        "recipients",
				Description: "List settings for the automation.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "removed_subscribers",
				Description: "A list of subscribers removed from the automation workflow.",
				Hydrate:     getAutomationRemovedSubscribers,
				Transform:   transform.FromValue(),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "report_summary",
				Description: "A summary of opens and clicks for sent campaigns.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "settings",
				Description: "The settings for the automation workflow.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tracking",
				Description: "The tracking options for the automation.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "trigger_settings",
				Description: "Available triggers for Automation workflows.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the automation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Settings.Title"),
			},
		},
	}
}

//// LIST FUNCTION

func listAutomations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_automation.listAutomations", "connection_error", err)
		return nil, err
	}

	params := gochimp3.BasicQueryParams{}
	if d.EqualsQuals["status"] != nil {
		params.Status = d.EqualsQualString("status")
	}

	automations, err := client.GetAutomations(&params)
	if err != nil {
		logger.Error("mailchimp_automation.listAutomations", "api_error", err)
		return nil, err
	}

	for _, automation := range automations.Automations {
		d.StreamListItem(ctx, &automation)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAutomation(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

	// List id should not be empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_automation.getAutomation", "connection_error", err)
		return nil, err
	}

	automation, err := client.GetAutomation(id)
	if err != nil {
		logger.Error("mailchimp_automation.getAutomation", "api_error", err)
		return nil, err
	}

	return automation, nil
}

func getAutomationRemovedSubscribers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := h.Item.(*gochimp3.Automation).ID

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_automation.getAutomationRemovedSubscribers", "connection_error", err)
		return nil, err
	}

	automationEmails, err := client.GetAutomationRemovedSubscribers(id)
	if err != nil {
		logger.Error("mailchimp_automation.getAutomationRemovedSubscribers", "api_error", err)
		return nil, err
	}

	return automationEmails.Subscribers, nil
}
