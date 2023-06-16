package mailchimp

import (
	"context"

	"github.com/hanzoai/gochimp3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableMailchimpAuthorizedApp(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mailchimp_authorized_app",
		Description: "Get a list of an account's registered, connected applications.",
		List: &plugin.ListConfig{
			Hydrate: listAuthorizedApps,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getAuthorizedApp,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID for the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the application.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "A short description of the application.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "users",
				Description: "An array of usernames for users who have linked the app.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listAuthorizedApps(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_authorized_app.listAuthorizedApps", "client_error", err)
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

	for {
		apps, err := client.GetAuthorizedApps(&params)
		if err != nil {
			logger.Error("mailchimp_authorized_app.listAuthorizedApps", "query_error", err)
			return nil, err
		}

		for _, list := range apps.Apps {
			d.StreamListItem(ctx, list)
		}

		if len(apps.Apps) == 0 || len(apps.Apps) < int(maxLimit) {
			break
		}
		if apps.TotalItems > 0 {
			params.Offset += int(maxLimit)
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAuthorizedApp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

	// List id should not be empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_authorized_app.getAuthorizedApp", "client_error", err)
		return nil, err
	}

	params := gochimp3.BasicQueryParams{}
	if d.EqualsQuals["status"] != nil && d.EqualsQualString("status") != "" {
		params.Status = d.EqualsQualString("status")
	}

	list, err := client.GetAuthroizedApp(id, &params)
	if err != nil {
		logger.Error("mailchimp_authorized_app.getAuthorizedApp", "query_error", err)
		return nil, err
	}

	return list, nil
}
