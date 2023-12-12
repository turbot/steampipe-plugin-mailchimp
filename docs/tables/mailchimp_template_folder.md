---
title: "Steampipe Table: mailchimp_template_folder - Query Mailchimp Template Folders using SQL"
description: "Allows users to query Mailchimp Template Folders, providing insights into the template folders used for organizing campaign templates."
---

# Table: mailchimp_template_folder - Query Mailchimp Template Folders using SQL

Mailchimp Template Folders are a feature within Mailchimp that allows users to organize and manage their campaign templates. These folders help in maintaining a structured and easily navigable template library. They are particularly beneficial for users with a large number of templates, enabling them to categorize templates based on various criteria such as campaign type, design, content, and more.

## Table Usage Guide

The `mailchimp_template_folder` table provides insights into the template folders in Mailchimp. If you are a marketing professional or a campaign manager, this table is extremely beneficial for you as it allows you to explore details about your template folders, including their names, creation dates, and associated templates. Utilize this table to manage and organize your campaign templates more effectively and efficiently.

## Examples

### Basic info
Explore which Mailchimp template folders have the highest count to optimize your email marketing efforts. This can help in identifying popular templates and strategizing your marketing campaigns accordingly.

```sql+postgres
select
  id,
  name,
  count
from
  mailchimp_template_folder;
```

```sql+sqlite
select
  id,
  name,
  count
from
  mailchimp_template_folder;
```

### List templates in each folder
Explore which templates are associated with each folder in your Mailchimp account. This is useful for organizing and managing your email marketing campaigns.

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