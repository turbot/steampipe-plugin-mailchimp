![image](https://hub.steampipe.io/images/plugins/turbot/mailchimp-social-graphic.png)

# Mailchimp Plugin for Steampipe

Use SQL to query audiences, automation workflows, campaigns, and more from Mailchimp

- **[Get started →](https://hub.steampipe.io/plugins/turbot/mailchimp)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/mailchimp/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-mailchimp/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install mailchimp
```

Run a query:

```sql
select
  name,
  dnssec ->> 'status',
  settings ->> 'tls_1_3'
from
  mailchimp_zone
```

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

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-mailchimp/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Mailchimp Plugin](https://github.com/turbot/steampipe-plugin-mailchimp/labels/help%20wanted)
