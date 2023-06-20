package mailchimp

import (
	"context"

	"github.com/hanzoai/gochimp3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableMailchimpBatchOperation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mailchimp_batch_operation",
		Description: "Get a summary of batch requests that have been made.",
		List: &plugin.ListConfig{
			Hydrate: listBatchOperations,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getBatchOperation,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the batch request.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "completed_at",
				Description: "The date and time when all operations in the batch request completed in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "errored_operations",
				Description: "The number of completed operations that returned an error.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "finished_operations",
				Description: "The number of completed operations. This includes operations that returned an error.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "response_body_url",
				Description: "The URL of the gzipped archive of the results of all the operations.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the batch call.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "submitted_at",
				Description: "The date and time when the server received the batch request in ISO 8601 format.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "total_operations",
				Description: "The total number of operations to complete as part of this batch request. For GET requests requiring pagination, each page counts as a separate operation.",
				Type:        proto.ColumnType_INT,
			},
		},
	}
}

//// LIST FUNCTION

func listBatchOperations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_batch_operation.listBatchOperations", "client_error", err)
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

	last := 0

	for {
		batchOperations, err := client.GetBatchOperations(&params)
		if err != nil {
			logger.Error("mailchimp_batch_operation.listBatchOperations", "query_error", err)
			return nil, err
		}

		for _, batchOperation := range batchOperations.BatchOperations {
			d.StreamListItem(ctx, &batchOperation)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		last = params.Offset + len(batchOperations.BatchOperations)
		if last >= batchOperations.TotalItems {
			return nil, nil
		} else {
			params.Offset = last
		}
	}
}

//// HYDRATE FUNCTIONS

func getBatchOperation(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

	// List id should not be empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectMailchimp(ctx, d)
	if err != nil {
		logger.Error("mailchimp_batch_operation.getBatchOperation", "client_error", err)
		return nil, err
	}

	params := gochimp3.BasicQueryParams{}

	batchOperation, err := client.GetBatchOperation(id, &params)
	if err != nil {
		logger.Error("mailchimp_batch_operation.getBatchOperation", "query_error", err)
		return nil, err
	}

	return batchOperation, nil
}
