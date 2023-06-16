package mailchimp

import (
	"context"
	"strconv"
	"time"

	"github.com/hanzoai/gochimp3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableMailchimpTemplate(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mailchimp_template",
		Description: "Get a list of an account's registered, connected applications.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "created_by",
					Require: plugin.Optional,
				},
				{
					Name:       "date_created",
					Operators:  []string{">", ">=", "<", "<=", "="},
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
				{
					Name:    "folder_id",
					Require: plugin.Optional,
				},
				{
					Name:    "type",
					Require: plugin.Optional,
				},
			},
			Hydrate: listTemplates,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getTemplate,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID for the application.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the application.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "active",
				Description: "Returns whether the template is still active.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "category",
				Description: "If available, the category the template is listed in.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_by",
				Description: "The login name for template's creator.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "date_created",
				Description: "The date and time the template was created in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "drag_and_drop",
				Description: "Whether the template uses the drag and drop editor.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "folder_id",
				Description: "The id of the folder the template is currently in.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "responsive",
				Description: "Whether the template contains media queries to make it responsive.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "share_url",
				Description: "The URL used for template sharing.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "thumbnail",
				Description: "If available, the URL for a thumbnail of the template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of template.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "template_default_content",
				Description: "The type of template.",
				Hydrate:     getTemplateDefaultContent,
				Transform:   transform.FromValue(),
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listTemplates(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_template.listTemplates", "client_error", err)
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

	params := gochimp3.TemplateQueryParams{
		ExtendedQueryParams: gochimp3.ExtendedQueryParams{
			Count:  int(maxLimit),
			Offset: 0,
		},
	}
	if d.EqualsQuals["created_by"] != nil {
		params.CreatedBy = d.EqualsQualString("created_by")
	}
	if d.EqualsQuals["folder_id"] != nil {
		params.SinceCreatedAt = d.EqualsQualString("folder_id")
	}
	if d.EqualsQuals["type"] != nil {
		params.Type = d.EqualsQualString("type")
	}
	if d.Quals["date_created"] != nil {
		for _, q := range d.Quals["date_created"].Quals {
			timestamp := q.Value.GetTimestampValue().AsTime().Format(time.DateTime)
			timestampAdd := q.Value.GetTimestampValue().AsTime().Add(time.Second).Format(time.DateTime)
			switch q.Operator {
			case ">=", ">":
				params.SinceCreatedAt = timestamp
			case "<":
				params.BeforeCreatedAt = timestamp
			case "<=":
				params.BeforeCreatedAt = timestampAdd
			case "=":
				params.SinceCreatedAt = timestamp
				params.BeforeCreatedAt = timestampAdd
			}
		}
	}

	last := 0

	for {
		templates, err := client.GetTemplates(&params)
		if err != nil {
			logger.Error("mailchimp_template.listTemplates", "query_error", err)
			return nil, err
		}

		for _, template := range templates.Templates {
			d.StreamListItem(ctx, &template)
		}

		last = params.Offset + len(templates.Templates)
		if last >= templates.TotalItems {
			return nil, nil
		} else {
			params.Offset = last
		}
	}
}

//// HYDRATE FUNCTIONS

func getTemplate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

	// List id should not be empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_template.getTemplate", "client_error", err)
		return nil, err
	}

	params := gochimp3.BasicQueryParams{}

	template, err := client.GetTemplate(id, &params)
	if err != nil {
		logger.Error("mailchimp_template.getTemplate", "query_error", err)
		return nil, err
	}

	return template, nil
}

func getTemplateDefaultContent(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := h.Item.(*gochimp3.TemplateResponse).ID

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_campaign.getTemplateDefaultContent", "client_error", err)
		return nil, err
	}

	params := gochimp3.BasicQueryParams{}
	templateContent, err := client.GetTemplateDefaultContent(strconv.Itoa(int(id)), &params)
	if err != nil {
		logger.Error("mailchimp_campaign.getTemplateDefaultContent", "query_error", err)
		return nil, err
	}

	return templateContent, nil
}
