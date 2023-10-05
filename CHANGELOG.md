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
