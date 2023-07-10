package mailchimp

import (
	"context"
	"errors"
	"os"

	"github.com/hanzoai/gochimp3"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connectMailchimp(ctx context.Context, d *plugin.QueryData) (*gochimp3.API, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "mailchimp"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*gochimp3.API), nil
	}

	// Default to using env vars (#2)
	apiKey := os.Getenv("MAILCHIMP_API_KEY")

	// But prefer the config (#1)
	mailchimpConfig := GetConfig(d.Connection)
	if mailchimpConfig.APIKey != nil {
		apiKey = *mailchimpConfig.APIKey
	}

	if apiKey == "" {
		// Credentials not set
		return nil, errors.New("api_key must be configured")
	}

	client := gochimp3.New(apiKey)
	if client != nil {
		d.ConnectionManager.Cache.Set(cacheKey, client)
	}

	return client, nil
}
