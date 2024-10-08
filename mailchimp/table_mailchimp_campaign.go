package mailchimp

import (
	"context"
	"time"

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
			Hydrate: listCampaigns,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:       "create_time",
					Operators:  []string{">", ">=", "<", "<=", "="},
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:       "send_time",
					Operators:  []string{">", ">=", "<", "<=", "="},
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:    "status",
					Require: plugin.Optional,
				},
				{
					Name:    "type",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "id",
					Require: plugin.Required,
				},
				{
					Name:    "status",
					Require: plugin.Optional,
				},
			},
			Hydrate: getCampaign,
		},
		Columns: commonColumns([]*plugin.Column{
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
				Name:        "delivery_status_enabled",
				Description: "Updates on campaigns in the process of sending.",
				Transform:   transform.FromField("DeliveryStatus.Enabled"),
				Type:        proto.ColumnType_BOOL,
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

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the campaign.",
				Transform:   transform.FromField("Settings.Title"),
				Type:        proto.ColumnType_STRING,
			},
		}),
	}
}

//// LIST FUNCTION

func listCampaigns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_campaign.listCampaigns", "connection_error", err)
		return nil, err
	}

	// Limiting the results
	maxLimit := int32(1000)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	params := gochimp3.CampaignQueryParams{
		ExtendedQueryParams: gochimp3.ExtendedQueryParams{
			Count:  int(maxLimit),
			Offset: 0,
		},
	}
	if d.EqualsQuals["status"] != nil {
		params.Status = d.EqualsQualString("status")
	}
	if d.EqualsQuals["type"] != nil {
		params.Type = d.EqualsQualString("type")
	}
	if d.Quals["create_time"] != nil {
		for _, q := range d.Quals["create_time"].Quals {
			timestamp := q.Value.GetTimestampValue().AsTime().Format(time.RFC3339)
			timestampAdd := q.Value.GetTimestampValue().AsTime().Add(time.Second).Format(time.RFC3339)
			switch q.Operator {
			case ">=", ">":
				params.SinceCreateTime = timestamp
			case "<":
				params.BeforeCreateTime = timestamp
			case "<=":
				params.BeforeCreateTime = timestampAdd
			case "=":
				params.SinceCreateTime = timestamp
				params.BeforeCreateTime = timestampAdd
			}
		}
	}
	if d.Quals["send_time"] != nil {
		for _, q := range d.Quals["send_time"].Quals {
			timestamp := q.Value.GetTimestampValue().AsTime().Format(time.RFC3339)
			timestampAdd := q.Value.GetTimestampValue().AsTime().Add(time.Second).Format(time.RFC3339)
			switch q.Operator {
			case ">=", ">":
				params.SinceSendTime = timestamp
			case "<":
				params.BeforeSendTime = timestamp
			case "<=":
				params.BeforeSendTime = timestampAdd
			case "=":
				params.SinceSendTime = timestamp
				params.BeforeSendTime = timestampAdd
			}
		}
	}

	last := 0

	for {
		campaigns, err := client.GetCampaigns(&params)
		if err != nil {
			logger.Error("mailchimp_campaign.listCampaigns", "api_error", err)
			return nil, err
		}

		for _, campaign := range campaigns.Campaigns {
			d.StreamListItem(ctx, &campaign)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		last = params.Offset + len(campaigns.Campaigns)
		if last >= campaigns.TotalItems {
			return nil, nil
		} else {
			params.Offset = last
		}
	}
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
		logger.Error("mailchimp_campaign.getCampaign", "connection_error", err)
		return nil, err
	}

	params := gochimp3.BasicQueryParams{}
	if d.EqualsQuals["status"] != nil && d.EqualsQualString("status") != "" {
		params.Status = d.EqualsQualString("status")
	}

	campaign, err := client.GetCampaign(id, &params)
	if err != nil {
		logger.Error("mailchimp_campaign.getCampaign", "api_error", err)
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
		logger.Error("mailchimp_campaign.getCampaignContent", "connection_error", err)
		return nil, err
	}

	params := gochimp3.BasicQueryParams{}
	campaignContent, err := client.GetCampaignContent(id, &params)
	if err != nil {
		logger.Error("mailchimp_campaign.getCampaignContent", "api_error", err)
		return nil, err
	}

	return campaignContent, nil
}
