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

func tableMailchimpList(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mailchimp_list",
		Description: "Get information about all lists in the account.",
		List: &plugin.ListConfig{
			Hydrate: listLists,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:       "date_created",
					Operators:  []string{">", ">=", "<", "<=", "="},
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getList,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "A string that uniquely identifies this list.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the list.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "beamer_address",
				Description: "The list's email beamer address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "date_created",
				Description: "The date and time that this list was created in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "email_type_option",
				Description: "Whether the list supports multiple formats for emails.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "list_rating",
				Description: "An auto-generated activity score for the list (0-5).",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "notify_on_subscribe",
				Description: "The email address to send subscribe notifications to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "notify_on_unsubscribe",
				Description: "The email address to send unsubscribe notifications to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "permission_reminder",
				Description: "The permission reminder for the list.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "subscribe_url_long",
				Description: "The full version of this list's subscribe form (host will vary).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SubscribeURLLong"),
			},
			{
				Name:        "subscribe_url_short",
				Description: "The full version of this list's subscribe form (host will vary).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SubscribeURLShort"),
			},
			{
				Name:        "use_archive_bar",
				Description: "Whether campaigns for this list use the Archive Bar in archives by default.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "visibility",
				Description: "Legacy - visibility settings are no longer used Possible values: 'pub' or 'prv'.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields

			{
				Name:        "campaign_defaults",
				Description: "Default values for campaigns created for this list.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "contact",
				Description: "Contact information displayed in campaign footers to comply with international spam laws.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "modules",
				Description: "Any list-specific modules installed for this list.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "stats",
				Description: "Stats for the list. Many of these are cached for at least five minutes.",
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

func listLists(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_list.listLists", "client_error", err)
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

	params := gochimp3.ListQueryParams{
		ExtendedQueryParams: gochimp3.ExtendedQueryParams{
			Count:  int(maxLimit),
			Offset: 0,
		},
	}
	if d.Quals["date_created"] != nil {
		for _, q := range d.Quals["date_created"].Quals {
			timestamp := q.Value.GetTimestampValue().AsTime().Format(time.DateTime)
			timestampAdd := q.Value.GetTimestampValue().AsTime().Add(time.Second).Format(time.DateTime)
			switch q.Operator {
			case ">=", ">":
				params.SinceDateCreated = timestamp
			case "<":
				params.BeforeDateCreated = timestamp
			case "<=":
				params.SinceDateCreated = timestampAdd
			case "=":
				params.SinceDateCreated = timestamp
				params.BeforeDateCreated = timestampAdd
			}
		}
	}

	last := 0

	for {
		lists, err := client.GetLists(&params)
		if err != nil {
			logger.Error("mailchimp_list.listLists", "query_error", err)
			return nil, err
		}

		for _, list := range lists.Lists {
			d.StreamListItem(ctx, list)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		last = params.Offset + len(lists.Lists)
		if last >= lists.TotalItems {
			return nil, nil
		} else {
			params.Offset = last
		}
	}
}

//// HYDRATE FUNCTIONS

func getList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

	// List id should not be empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_list.getList", "client_error", err)
		return nil, err
	}

	params := gochimp3.BasicQueryParams{}

	list, err := client.GetList(id, &params)
	if err != nil {
		logger.Error("mailchimp_list.getList", "query_error", err)
		return nil, err
	}

	return list, nil
}
