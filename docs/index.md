---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/mailchimp.svg"
brand_color: "#000000"
display_name: "mailchimp"
short_name: "mailchimp"
description: "Steampipe plugin to query VPN networks, devices and more from mailchimp."
og_description: "Query Mailchimp with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/mailchimp-social-graphic.png"
---

# Mailchimp + Steampipe

[Mailchimp](https://mailchimp.com) is a zero config VPN which installs on any device in minutes and manages firewall rules.

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
  plugin       = "mailchimp"

  # Required: Set your API Key and Tailnet name
  # Generate your API Key as per: https://mailchimp.com/kb/1101/api/
  api_key = "abcde-krSvfN1CNTRL-M67st8X5o1234567"
  tailnet_name = "example.com"
}
```

- `api_key` - API Key of the Tailscale account.
- `tailnet_name` - Name of your Tailnet.

Environment variables are also available as an alternate configuration method:
* `TAILSCALE_API_KEY`
* `TAILSCALE_TAILNET_NAME`

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-mailchimp
- Community: [Slack Channel](https://steampipe.io/community/join)
