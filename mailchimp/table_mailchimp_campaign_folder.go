package mailchimp

import (
	"context"

	"github.com/hanzoai/gochimp3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableMailchimpCampaignFolder(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mailchimp_campaign_folder",
		Description: "Get a list of an account's campaign folders.",
		List: &plugin.ListConfig{
			Hydrate: listCampaignFolders,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "A string that uniquely identifies this campaign folder.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the folder.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "count",
				Description: "The number of campaigns in the folder.",
				Type:        proto.ColumnType_INT,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listCampaignFolders(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_campaign_folder.listCampaignFolders", "connection_error", err)
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

	params := gochimp3.CampaignFolderQueryParams{
		ExtendedQueryParams: gochimp3.ExtendedQueryParams{
			Count:  int(maxLimit),
			Offset: 0,
		},
	}

	last := 0

	for {
		folders, err := client.GetCampaignFolders(&params)
		if err != nil {
			logger.Error("mailchimp_campaign_folder.listCampaignFolders", "api_error", err)
			return nil, err
		}

		for _, campaign := range folders.Folders {
			d.StreamListItem(ctx, campaign)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		last = params.Offset + len(folders.Folders)
		if last >= folders.TotalItems {
			return nil, nil
		} else {
			params.Offset = last
		}
	}
}
