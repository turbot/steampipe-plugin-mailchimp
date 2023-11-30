---
title: "Steampipe Table: mailchimp_campaign_folder - Query Mailchimp Campaign Folders using SQL"
description: "Allows users to query Mailchimp Campaign Folders, specifically the details of each folder including id, name, and count, providing insights into campaign organization and structure."
---

# Table: mailchimp_campaign_folder - Query Mailchimp Campaign Folders using SQL

Mailchimp Campaign Folders are a feature within Mailchimp that allows users to organize their email marketing campaigns. These folders provide a way to categorize and manage campaigns for better accessibility and efficiency. Using these folders, users can group related campaigns together, making it easier to find and manage specific campaigns.

## Table Usage Guide

The `mailchimp_campaign_folder` table provides insights into Campaign Folders within Mailchimp. As a Marketing or Sales professional, explore folder-specific details through this table, including id, name, and count. Utilize it to uncover information about your campaign organization, such as how campaigns are grouped, the number of campaigns in each folder, and the overall structure of your campaigns.

## Examples

### Basic info
Discover the segments that have been organized in your Mailchimp campaigns. This query is beneficial to understand how many campaigns are housed within each segment, allowing for efficient campaign management and strategic planning.

```sql
select
  id,
  name,
  count
from
  mailchimp_campaign_folder;
```

### List campaigns in each folder
Explore which marketing campaigns are associated with each folder in your Mailchimp account. This can help you organize and manage your campaigns more effectively.

```sql
select
  f.id as folder_id,
  f.name as folder_name,
  c.id as campaign_id,
  c.title as campaign_title
from
  mailchimp_campaign c
  left join mailchimp_campaign_folder f on c.settings ->> 'folder_id' = f.id;
```