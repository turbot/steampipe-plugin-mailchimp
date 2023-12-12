---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/mailchimp.svg"
brand_color: "#FFE01B"
display_name: "Mailchimp"
short_name: "mailchimp"
description: "Steampipe plugin to query audiences, automation workflows, campaigns, and more from Mailchimp."
og_description: "Query Mailchimp with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/mailchimp-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Mailchimp + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[Mailchimp](https://mailchimp.com) is a marketing automation and email marketing platform.

List details of your Mailchimp campaign:

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

### Credentials

| Item        | Description                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Mailchimp requires an [API key](https://mailchimp.com/developer/marketing/guides/quick-start/#generate-your-api-key/) for all requests.                                                               |
| Permissions | API keys have the same permissions as the user who creates them, and if the user permissions change, the API key permissions also change.                                                             |
| Radius      | Each connection represents a single Mailchimp Installation.                                                                                                                                           |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/mailchimp.spc`)<br />2. Credentials specified in environment variables, e.g., `MAILCHIMP_API_KEY`.                     |

### Configuration

Installing the latest mailchimp plugin will create a config file (`~/.steampipe/config/mailchimp.spc`) with a single connection named `mailchimp`:

```hcl
connection "mailchimp" {
  plugin = "mailchimp"

  # Mailchimp API key for requests. Required.
  # Generate your API Key as per: https://mailchimp.com/developer/marketing/guides/quick-start/#generate-your-api-key/
  # This can also be set via the `MAILCHIMP_API_KEY` environment variable.
  # api_key = "08355689e3e6f9fd0f5630362b16b1b5-us21"
}
```

Alternatively, you can also use the standard Mailchimp environment variables to obtain credentials **only if other argument (`api_key`) is not specified** in the connection:

```sh
export MAILCHIMP_API_KEY=q8355689e3e6f9fd0f5630362b16b1b5-us21
```
