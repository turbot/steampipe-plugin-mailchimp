package mailchimp

import (
	"context"

	"github.com/hanzoai/gochimp3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableMailchimpAutomationQueue(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mailchimp_automation_queue",
		Description: "Get information about a classic automation email queue.",
		List: &plugin.ListConfig{
			ParentHydrate: listAutomations,
			Hydrate:       listAutomationQueues,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "email_id",
					Require: plugin.Required,
				},
				{
					Name:    "workflow_id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "workflow_id", "email_id"}),
			Hydrate:    getAutomationQueue,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the automation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "email_address",
				Description: "The list member's email address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email_id",
				Description: "A string that uniquely identifies an email in an Automation workflow.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("EmailID"),
			},
			{
				Name:        "list_id",
				Description: "A string that uniquely identifies a list.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ListID"),
			},
			{
				Name:        "next_send",
				Description: "The date and time of the next send for the workflow email in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "workflow_id",
				Description: "A string that uniquely identifies an automation workflow.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listAutomationQueues(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	workflowId := h.ParentItem.(*gochimp3.Automation).ID
	var workflowEmailId string

	if d.EqualsQuals["workflow_id"] != nil && d.EqualsQualString("workflow_id") != workflowId {
		return nil, nil
	}
	if d.EqualsQualString("email_id") != "" {
		workflowEmailId = d.EqualsQualString("email_id")
	} else {
		return nil, nil
	}

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_automation_queue.listAutomationQueues", "client_error", err)
		return nil, err
	}

	automationQueues, err := client.GetAutomationQueues(workflowId, workflowEmailId)
	if err != nil {
		logger.Error("mailchimp_automation_queue.listAutomationQueues", "query_error", err)
		return nil, err
	}

	for _, automationQueue := range automationQueues.Queues {
		d.StreamListItem(ctx, &automationQueue)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAutomationQueue(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")
	workflowId := d.EqualsQualString("workflow_id")
	workflowEmailId := d.EqualsQualString("email_id")

	// Email id and workflow id should not be empty
	if id == "" || workflowId == "" || workflowEmailId == "" {
		return nil, nil
	}

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_automation_queue.getAutomationQueue", "client_error", err)
		return nil, err
	}

	automationQueue, err := client.GetAutomationQueue(workflowId, workflowEmailId, id)
	if err != nil {
		logger.Error("mailchimp_automation_queue.getAutomationQueue", "query_error", err)
		return nil, err
	}

	return automationQueue, nil
}
