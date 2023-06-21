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
  id,
  title,
  content_type,
  create_time,
  emails_sent,
  send_time,
  status,
  type
from
  mailchimp_campaign;
```

```
+------------+------------------------------------+--------------+---------------------------+-------------+-----------+--------+------------------+
| id         | title                              | content_type | create_time               | emails_sent | send_time | status | type             |
+------------+------------------------------------+--------------+---------------------------+-------------+-----------+--------+------------------+
| f739729f66 | We're here to help you get started | template     | 2023-06-16T17:51:52+05:30 | <null>      | <null>    | save   | automation-email |
+------------+------------------------------------+--------------+---------------------------+-------------+-----------+--------+------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/mailchimp/tables)**

## Quick start

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

  # Mailchimp API key for requests. Required.
  # Generate your API Key as per: https://mailchimp.com/developer/marketing/guides/quick-start/#generate-your-api-key/
  # This can also be set via the `MAILCHIMP_API_KEY` environment variable.
  # mailchimp_api_key = "08355689e3e6f9fd0f5630362b16b1b5-us21"
}
```

Alternatively, you can also use the standard Mailchimp environment variables to obtain credentials **only if other arguments (`mailchimp_api_key`) is not specified** in the connection:

```sh
export MAILCHIMP_API_KEY=q8355689e3e6f9fd0f5630362b16b1b5-us21
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-mailchimp
- Community: [Slack Channel](https://steampipe.io/community/join)
