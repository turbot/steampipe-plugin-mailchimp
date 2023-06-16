package mailchimp

import (
	"context"

	"github.com/hanzoai/gochimp3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableMailchimpCampaign(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mailchimp_campaign",
		Description: "Mailchimp Campaign.",
		List: &plugin.ListConfig{
			Hydrate:    listCampaigns,
			KeyColumns: plugin.OptionalColumns([]string{"create_time", "send_time", "status", "type"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "id",
					Require: "required",
				},
				{
					Name:    "status",
					Require: "optional",
				},
			},
			Hydrate: getCampaign,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "A string that uniquely identifies this campaign.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "archive_url",
				Description: "The link to the campaign's archive version in ISO 8601 format.",
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
				Description: "The total number of emails sent for this campaign.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "long_archive_url",
				Description: "The original link to the campaign's archive version.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "needs_block_refresh",
				Description: "Determines if the campaign needs its blocks refreshed by opening the web-based campaign editor. Deprecated and will always return false.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "send_time",
				Description: "The date and time a campaign was sent.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "status",
				Description: "The current status of the campaign.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "Type of the campaign.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "web_id",
				Description: "The ID used in the Mailchimp web application.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("WebID"),
			},

			// JSON fields

			{
				Name:        "campaign_content",
				Description: "The HTML, plain-text, and template content for your Mailchimp campaigns.",
				Hydrate:     getCampaignContent,
				Transform:   transform.FromValue(),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "delivery_status",
				Description: "Updates on campaigns in the process of sending.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "recipients",
				Description: "List settings for the campaign.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "report_summary",
				Description: "For sent campaigns, a summary of opens, clicks, and e-commerce data.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "settings",
				Description: "Settings for the campaign including the subject line, from name, reply-to address, and more.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tracking",
				Description: "The tracking options for a campaign.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listCampaigns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_campaign.listCampaigns", "client_error", err)
		return nil, err
	}

	params := gochimp3.CampaignQueryParams{}
	if d.EqualsQuals["status"] != nil && d.EqualsQualString("status") != "" {
		params.Status = d.EqualsQualString("status")
	}
	if d.EqualsQuals["type"] != nil && d.EqualsQualString("type") != "" {
		params.Type = d.EqualsQualString("type")
	}

	campaigns, err := client.GetCampaigns(&params)
	if err != nil {
		logger.Error("mailchimp_campaign.listCampaigns", "query_error", err)
		return nil, err
	}

	for _, campaign := range campaigns.Campaigns {
		d.StreamListItem(ctx, &campaign)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getCampaign(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

	// Campaign id should not be empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_campaign.getCampaign", "client_error", err)
		return nil, err
	}

	params := gochimp3.BasicQueryParams{}
	if d.EqualsQuals["status"] != nil && d.EqualsQualString("status") != "" {
		params.Status = d.EqualsQualString("status")
	}

	campaign, err := client.GetCampaign(id, &params)
	if err != nil {
		logger.Error("mailchimp_campaign.getCampaign", "query_error", err)
		return nil, err
	}

	return campaign, nil
}

func getCampaignContent(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := h.Item.(*gochimp3.CampaignResponse).ID

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_campaign.getCampaignContent", "client_error", err)
		return nil, err
	}

	params := gochimp3.BasicQueryParams{}
	campaignContent, err := client.GetCampaignContent(id, &params)
	if err != nil {
		logger.Error("mailchimp_campaign.getCampaignContent", "query_error", err)
		return nil, err
	}

	return campaignContent, nil
}
