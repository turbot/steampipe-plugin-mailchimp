![image](https://hub.steampipe.io/images/plugins/turbot/mailchimp-social-graphic.png)

# Mailchimp Plugin for Steampipe

Use SQL to query audiences, automation workflows, campaigns, and more from Mailchimp

- **[Get started →](https://hub.steampipe.io/plugins/turbot/mailchimp)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/mailchimp/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-mailchimp/issues)

## Quick start

### Install

Download and install the latest Mailchimp plugin:

```bash
steampipe plugin install mailchimp
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/mailchimp#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/mailchimp#configuration).

Configure your account details in `~/.steampipe/config/mailchimp.spc`:

```hcl
connection "mailchimp" {
  plugin = "mailchimp"

  # Authentication information
  api_key = "08355689e3e6f9fd0f5630362b16b1b5-us21"
}
```

Or through environment variables:

```sh
export MAILCHIMP_API_KEY=08355689e3e6f9fd0f5630362b16b1b5-us21
```

Run steampipe:

```shell
steampipe query
```

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

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-mailchimp.git
cd steampipe-plugin-mailchimp
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/mailchimp.spc
```

Try it!

```
steampipe query
> .inspect mailchimp
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-mailchimp/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-mailchimp/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Mailchimp Plugin](https://github.com/turbot/steampipe-plugin-mailchimp/labels/help%20wanted)
