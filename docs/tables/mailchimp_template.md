---
title: "Steampipe Table: mailchimp_template - Query Mailchimp Templates using SQL"
description: "Allows users to query Mailchimp Templates, providing insights into the details of templates like their design, settings, and associated metadata."
---

# Table: mailchimp_template - Query Mailchimp Templates using SQL

Mailchimp Templates are a resource within the Mailchimp platform that allows users to design and store email layouts for repeated use. They provide a centralized way to manage and customize the appearance of your emails, including their content, style, and layout. Mailchimp Templates help you maintain consistency in your email campaigns and save time by reusing designs.

## Table Usage Guide

The `mailchimp_template` table provides insights into the templates within Mailchimp. As a Digital Marketing Specialist or Email Campaign Manager, explore template-specific details through this table, including design elements, settings, and associated metadata. Utilize it to manage and optimize your email campaigns, ensuring consistency and efficiency in your email marketing efforts.

## Examples

### Basic info
Explore which Mailchimp templates are active, who created them, and when they were created to better understand your email marketing efforts and strategies. This can help in assessing the effectiveness of different templates and their usage over time.

```sql+postgres
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

```sql+sqlite
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
Explore active templates to understand their origin and category. This can help in assessing which templates are currently in use and who created them, providing insights for potential updates or re-usage.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  category,
  created_by,
  date_created
from
  mailchimp_template
where
  active = 1;
```

### List templates using drag and drop editor
Explore which email templates are created using the drag and drop editor in Mailchimp. This can help in identifying the templates that were quickly designed without the need for coding knowledge.

```sql+postgres
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

```sql+sqlite
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
  drag_and_drop = 1;
```

### List templates of type user
Explore which user-created templates are active within your Mailchimp account, providing a quick way to assess your personalized content and its creation details. This can be particularly useful for auditing your marketing efforts or identifying templates for future campaigns.

```sql+postgres
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

```sql+sqlite
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
Explore which templates are organized in each folder in your Mailchimp account. This is useful for understanding your email marketing structure and identifying potential areas for reorganization or consolidation.

```sql+postgres
select
  f.id as folder_id,
  f.name as folder_name,
  t.id as template_id,
  t.title as template_title
from
  mailchimp_template t
  left join mailchimp_template_folder f on t.folder_id = f.id;
```

```sql+sqlite
select
  f.id as folder_id,
  f.name as folder_name,
  t.id as template_id,
  t.title as template_title
from
  mailchimp_template t
  left join mailchimp_template_folder f on t.folder_id = f.id;
```

### Get template count by category
Determine the areas in which different categories of Mailchimp templates are being used, allowing you to understand their popularity and usage frequency.

```sql+postgres
select
  category,
  count(*) as count
from
  mailchimp_template
group by
  category;
```

```sql+sqlite
select
  category,
  count(*) as count
from
  mailchimp_template
group by
  category;
```