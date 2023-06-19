---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/mailchimp.svg"
brand_color: "#000000"
display_name: "mailchimp"
short_name: "mailchimp"
description: "Steampipe plugin to query audiences, automation workflows, campaigns, and more from Mailchimp."
og_description: "Query Mailchimp with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/mailchimp-social-graphic.png"
---

# Mailchimp + Steampipe

[Mailchimp](https://mailchimp.com) is a marketing automation and email marketing platform.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List devices which block incoming connections in your Mailchimp tailnet:

```sql
select
  name,
  device.user,
  created,
  tailnet_name
from
  mailchimp_device as device
where
  blocks_incoming_connections;
```

```
+------------------------------------+-----------+---------------------------+--------------+
| name                               | user      | created                   | tailnet_name |
+------------------------------------+-----------+---------------------------+--------------+
| francis-macbook-pro.turbot.com     | francis   | 2022-09-19T10:28:55+08:00 | testdo.com   |
| oneplus-nord2-5g.testdo.com        | keyma     | 2022-09-19T16:58:56+08:00 | testdo.com   |
| test-macbook-pro.testdo.com        | test      | 2022-09-19T10:27:55+08:00 | testdo.com   |
| ip-172-32-10-22.testdo.com         | steampipe | 2022-09-20T12:50:55+08:00 | testdo.com   |
+------------------------------------+-----------+---------------------------+--------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/mailchimp/tables)**

## Get started

### Install

Download and install the latest Mailchimp plugin:

```bash
steampipe plugin install mailchimp
```

### Configuration

Installing the latest mailchimp plugin will create a config file (`~/.steampipe/config/mailchimp.spc`) with a single connection named `mailchimp`:

```hcl
connection "mailchimp" {
  plugin = "mailchimp"

  # Required: Set your Mailchimp API Key
  # Generate your API Key as per: https://mailchimp.com/developer/marketing/guides/quick-start/#generate-your-api-key/
  # This can also be set via the `MAILCHIMP_API_KEY` environment variable
  mailchimp_api_key = "08355689e3e6f9fd0f5630362b16b1b5-us21"
}
```

- `mailchimp_api_key` - API Key of the Mailchimp account.

Environment variables are also available as an alternate configuration method:
* `MAILCHIMP_API_KEY`

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-mailchimp
- Community: [Slack Channel](https://steampipe.io/community/join)
