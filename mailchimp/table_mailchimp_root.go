package mailchimp

import (
	"context"

	"github.com/hanzoai/gochimp3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableMailchimpRoot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mailchimp_root",
		Description: "The root directory returns details about the Mailchimp user account.",
		List: &plugin.ListConfig{
			Hydrate: listRoots,
		},
		Columns: []*plugin.Column{
			{
				Name:        "account_id",
				Description: "The Mailchimp account ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AccountID"),
			},
			{
				Name:        "account_name",
				Description: "The name of the account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "The account email address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_login",
				Description: "The date and time of the last login for this account in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "pro_enabled",
				Description: "Whether the account includes Mailchimp Pro.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role",
				Description: "The user role for the account.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "total_subscribers",
				Description: "The total number of subscribers across all lists in the account.",
				Type:        proto.ColumnType_INT,
			},

			// JSON fields
			{
				Name:        "contact",
				Description: "Information about the account contact.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "industy_stats",
				Description: "The average campaign statistics for all campaigns in the account's specified industry.",
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

func listRoots(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_root.listRoots", "client_error", err)
		return nil, err
	}

	for {
		root, err := client.GetRoot(&gochimp3.BasicQueryParams{})
		if err != nil {
			logger.Error("mailchimp_root.listRoots", "query_error", err)
			return nil, err
		}
		d.StreamListItem(ctx, root)

	}
}
