# Table: mailchimp_template

A template is an HTML file used to create the layout and basic design for a campaign.

## Examples

### Basic info

```sql
select
  id,
  name,
  active,
  category,
  created_by,
  date_created
from
  mailchimp_template;
```

### List active templates

```sql
select
  id,
  name,
  category,
  created_by,
  date_created
from
  mailchimp_template
where
  active;
```

### List templates using drag and drop editor

```sql
select
  id,
  name,
  active,
  category,
  created_by,
  date_created
from
  mailchimp_template
where
  drag_and_drop;
```

### List templates of type user

```sql
select
  id,
  name,
  active,
  category,
  created_by,
  date_created
from
  mailchimp_template
where
  type = 'user';
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

### Get template count by category

```sql
select
  category,
  count(*) as count
from
  mailchimp_template
group by
  category;
```
