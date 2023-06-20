# Table: mailchimp_template_folder

A template is an HTML file used to create the layout and basic design for a campaign. Templates can be organized using folders.

## Examples

### Basic info

```sql
select
  id,
  name,
  count
from
  mailchimp_template_folder;
```

### List templates in each folder

```sql
select
  f.id as folder_id,
  f.name as folder_name,
  t.id as template_id,
  t.title as template_title
from
  mailchimp_template t
  left join mailchimp_template_folder f
    on t.folder_id = f.id;
```