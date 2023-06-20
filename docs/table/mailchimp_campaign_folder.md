# Table: mailchimp_campaign_folder

Campaign folders are used to organize your regular emails, ads, automations and other projects in Mailchimp.

## Examples

### Basic info

```sql
select
  id,
  name,
  count
from
  mailchimp_campaign_folder;
```

### List campaigns in each folder

```sql
select
  f.id as folder_id,
  f.name as folder_name,
  c.id as campaign_id,
  c.title as campaign_title
from
  mailchimp_campaign c
  left join mailchimp_campaign_folder f
    on c.settings ->> 'folder_id' = f.id;
```