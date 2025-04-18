## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#25](https://github.com/turbot/steampipe-plugin-mailchimp/pull/25))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#25](https://github.com/turbot/steampipe-plugin-mailchimp/pull/25))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#23](https://github.com/turbot/steampipe-plugin-mailchimp/pull/23))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#23](https://github.com/turbot/steampipe-plugin-mailchimp/pull/23))

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/install/steampipe.sh), as a [Postgres FDW](https://steampipe.io/install/postgres.sh), as a [SQLite extension](https://steampipe.io/install/sqlite.sh) and as a standalone [exporter](https://steampipe.io/install/export.sh).
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension.
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-mailchimp/blob/main/docs/LICENSE).

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server enacapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#18](https://github.com/turbot/steampipe-plugin-mailchimp/pull/18))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#10](https://github.com/turbot/steampipe-plugin-mailchimp/pull/10))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#8](https://github.com/turbot/steampipe-plugin-mailchimp/pull/8))
- Recompiled plugin with Go version `1.21`. ([#8](https://github.com/turbot/steampipe-plugin-mailchimp/pull/8))

## v0.0.3 [2023-07-10]

_Bug fixes_

- Fixed the plugin's config argument to use `api_key` instead of `mailchimp_api_key` to align with the API documentation. ([#4](https://github.com/turbot/steampipe-plugin-mailchimp/pull/4))

## v0.0.2 [2023-06-29]

_Bug fixes_

- Fixed the incorrect table document references in the plugin to correctly render the example queries on the hub. ([#2](https://github.com/turbot/steampipe-plugin-mailchimp/pull/2))

## v0.0.1 [2023-06-29]

_What's new?_

- New tables added
  - [mailchimp_authorized_app](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_authorized_app)
  - [mailchimp_automation_email](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_automation_email)
  - [mailchimp_automation_queue](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_automation_queue)
  - [mailchimp_automation](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_automation)
  - [mailchimp_batch_operation](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_batch_operation)
  - [mailchimp_campaign_folder](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_campaign_folder)
  - [mailchimp_campaign](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_campaign)
  - [mailchimp_list](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_list)
  - [mailchimp_root](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_root)
  - [mailchimp_store](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_store)
  - [mailchimp_template_folder](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_template_folder)
  - [mailchimp_template](https://hub.steampipe.io/plugins/turbot/mailchimp/tables/mailchimp_template)
