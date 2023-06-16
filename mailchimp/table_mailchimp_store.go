package mailchimp

import (
	"context"

	"github.com/hanzoai/gochimp3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableMailchimpStore(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mailchimp_store",
		Description: "List of E-commerce Stores.",
		List: &plugin.ListConfig{
			Hydrate: listStores,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getStore,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the store.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The date and time the store was created in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "currency_code",
				Description: "The three-letter ISO 4217 code for the currency that the store accepts.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "domain",
				Description: "The store domain. The store domain must be unique within a user account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email_address",
				Description: "The email address for the store.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "list_id",
				Description: "The unique identifier for the list that's associated with the store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ListID"),
			},
			{
				Name:        "money_format",
				Description: "The currency format for the store.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "phone",
				Description: "The store phone number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "platform",
				Description: "The e-commerce platform of the store.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "primary_locale",
				Description: "The primary locale for the store.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "timezone",
				Description: "The timezone for the store.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "updated_at",
				Description: "The date and time the store was last updated in ISO 8601 format.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "address",
				Description: "The store address.",
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

func listStores(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_store.listStores", "client_error", err)
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

	params := gochimp3.ExtendedQueryParams{
		Count:  int(maxLimit),
		Offset: 0,
	}

	last := 0

	for {
		stores, err := client.GetStores(&params)
		if err != nil {
			logger.Error("mailchimp_store.listStores", "query_error", err)
			return nil, err
		}

		for _, store := range stores.Stores {
			d.StreamListItem(ctx, &store)
		}

		last = params.Offset + len(stores.Stores)
		if last >= stores.TotalItems {
			return nil, nil
		} else {
			params.Offset = last
		}
	}
}

//// HYDRATE FUNCTIONS

func getStore(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

	// List id should not be empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_store.getStore", "client_error", err)
		return nil, err
	}

	params := gochimp3.BasicQueryParams{}

	store, err := client.GetStore(id, &params)
	if err != nil {
		logger.Error("mailchimp_store.getStore", "query_error", err)
		return nil, err
	}

	return store, nil
}
