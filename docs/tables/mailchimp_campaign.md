---
title: "Steampipe Table: mailchimp_campaign - Query Mailchimp Campaigns using SQL"
description: "Allows users to query Mailchimp Campaigns, specifically the details of individual campaigns, providing insights into campaign performance and engagement metrics."
---

# Table: mailchimp_campaign - Query Mailchimp Campaigns using SQL

Mailchimp campaigns are a key resource within the Mailchimp platform, designed to create, send, and track email marketing campaigns. These campaigns can be categorized into various types such as regular, automated, RSS, A/B testing, and more. Each campaign carries crucial information about the content, audience, schedule, performance, and other related details.

## Table Usage Guide

The `mailchimp_campaign` table provides insights into Mailchimp campaigns within the Mailchimp email marketing platform. As a marketing analyst or data scientist, explore campaign-specific details through this table, including content, audience, schedule, and performance metrics. Utilize it to uncover information about campaigns, such as those with high engagement rates, the performance of various campaign types, and the verification of campaign schedules.

## Examples

### Basic info
Explore which marketing campaigns have been most effective by analyzing their status, type, and the number of emails sent. This can help you optimize your future campaigns by understanding what has worked well in the past.

```sql+postgres
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

```sql+sqlite
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

### Get the list settings for each campaign
Explore the configuration of each marketing campaign to understand the targeted recipient group, including the specific list and segment used, as well as the total number of recipients. This can help optimize your marketing efforts by identifying the most effective recipient groups.

```sql+postgres
select
  id,
  title,
  recipients ->> 'list_id' as recipients_list_id,
  recipients ->> 'list_name' as recipients_list_name,
  recipients ->> 'segment_text' as recipients_segment_text,
  recipients ->> 'recipient_count' as recipients_recipient_count
from
  mailchimp_campaign;
```

```sql+sqlite
select
  id,
  title,
  json_extract(recipients, '$.list_id') as recipients_list_id,
  json_extract(recipients, '$.list_name') as recipients_list_name,
  json_extract(recipients, '$.segment_text') as recipients_segment_text,
  json_extract(recipients, '$.recipient_count') as recipients_recipient_count
from
  mailchimp_campaign;
```

### Get the settings for each campaign
Explore the various settings for each marketing campaign to gain insights into their configuration and assess the elements within each one, such as authentication, automatic features, and social media integration. This can help in understanding the campaign structure and optimizing future campaigns.

```sql+postgres
select
  id,
  title,
  settings ->> 'authenticate' as authenticate,
  settings ->> 'auto_footer' as auto_footer,
  settings ->> 'auto_tweet' as auto_tweet,
  settings ->> 'drag_and_drop' as drag_and_drop,
  settings ->> 'fb_comments' as fb_comments,
  settings ->> 'folder_id' as folder_id,
  settings ->> 'from_name' as from_name,
  settings ->> 'inline_css' as inline_css,
  settings ->> 'preview_text' as preview_text,
  settings ->> 'reply_to' as reply_to,
  settings ->> 'subject_line' as subject_line,
  settings ->> 'template_id' as template_id,
  settings ->> 'timewarp' as timewarp,
  settings ->> 'to_name' as to_name,
  settings ->> 'use_conversation' as use_conversation
from
  mailchimp_campaign;
```

```sql+sqlite
select
  id,
  title,
  json_extract(settings, '$.authenticate') as authenticate,
  json_extract(settings, '$.auto_footer') as auto_footer,
  json_extract(settings, '$.auto_tweet') as auto_tweet,
  json_extract(settings, '$.drag_and_drop') as drag_and_drop,
  json_extract(settings, '$.fb_comments') as fb_comments,
  json_extract(settings, '$.folder_id') as folder_id,
  json_extract(settings, '$.from_name') as from_name,
  json_extract(settings, '$.inline_css') as inline_css,
  json_extract(settings, '$.preview_text') as preview_text,
  json_extract(settings, '$.reply_to') as reply_to,
  json_extract(settings, '$.subject_line') as subject_line,
  json_extract(settings, '$.template_id') as template_id,
  json_extract(settings, '$.timewarp') as timewarp,
  json_extract(settings, '$.to_name') as to_name,
  json_extract(settings, '$.use_conversation') as use_conversation
from
  mailchimp_campaign;
```

### Get the tracking options for each campaign
This query is useful for understanding the various tracking options available for each of your campaigns. It provides insights into the different analytics tools being used, which can help optimize your marketing strategies.

```sql+postgres
select
  id,
  tracking ->> 'clicktale' as clicktale,
  tracking ->> 'ecomm360' as ecomm360,
  tracking ->> 'goal_tracking' as goal_tracking,
  tracking ->> 'google_analytics' as google_analytics,
  tracking ->> 'html_clicks' as html_clicks,
  tracking ->> 'opens' as opens,
  tracking ->> 'text_clicks' as text_clicks
from
  mailchimp_campaign;
```

```sql+sqlite
select
  id,
  json_extract(tracking, '$.clicktale') as clicktale,
  json_extract(tracking, '$.ecomm360') as ecomm360,
  json_extract(tracking, '$.goal_tracking') as goal_tracking,
  json_extract(tracking, '$.google_analytics') as google_analytics,
  json_extract(tracking, '$.html_clicks') as html_clicks,
  json_extract(tracking, '$.opens') as opens,
  json_extract(tracking, '$.text_clicks') as text_clicks
from
  mailchimp_campaign;
```

### Get the report summary for each campaign
Gain insights into the performance of each marketing campaign by examining key metrics such as open rates, click rates, and ecommerce revenue. This allows you to evaluate the effectiveness of each campaign and make data-driven decisions for future marketing strategies.

```sql+postgres
select
  id,
  report_summary ->> 'opens' as opens,
  report_summary ->> 'unique_opens' as unique_opens,
  report_summary ->> 'open_rate' as open_rate,
  report_summary ->> 'clicks' as clicks,
  report_summary ->> 'subscriber_clicks' as subscriber_clicks,
  report_summary ->> 'click_rate' as click_rate,
  report_summary -> 'ecommerce' ->> 'total_orders' as ecommerce_total_orders,
  report_summary -> 'ecommerce' ->> 'total_spent' as ecommerce_total_spent,
  report_summary -> 'ecommerce' ->> 'total_revenue' as ecommerce_total_revenue
from
  mailchimp_campaign;
```

```sql+sqlite
select
  id,
  json_extract(report_summary, '$.opens') as opens,
  json_extract(report_summary, '$.unique_opens') as unique_opens,
  json_extract(report_summary, '$.open_rate') as open_rate,
  json_extract(report_summary, '$.clicks') as clicks,
  json_extract(report_summary, '$.subscriber_clicks') as subscriber_clicks,
  json_extract(report_summary, '$.click_rate') as click_rate,
  json_extract(report_summary, '$.ecommerce.total_orders') as ecommerce_total_orders,
  json_extract(report_summary, '$.ecommerce.total_spent') as ecommerce_total_spent,
  json_extract(report_summary, '$.ecommerce.total_revenue') as ecommerce_total_revenue
from
  mailchimp_campaign;
```

### List campaigns having delivery status enabled
Explore which campaigns are currently active by examining their delivery status. This can be beneficial for assessing ongoing marketing strategies and ensuring all campaigns are functioning as expected.

```sql+postgres
select
  id,
  content_type,
  create_time,
  emails_sent,
  status,
  type
from
  mailchimp_campaign
where
  delivery_status_enabled;
```

```sql+sqlite
select
  id,
  content_type,
  create_time,
  emails_sent,
  status,
  type
from
  mailchimp_campaign
where
  delivery_status_enabled = 1;
```

### List campaigns in each folder
Explore which marketing campaigns are organized under each folder, allowing you to understand the structure of your marketing efforts and identify potential areas for reorganization.

```sql+postgres
select
  f.id as folder_id,
  f.name as folder_name,
  c.id as campaign_id,
  c.title as campaign_title
from
  mailchimp_campaign c
  left join mailchimp_campaign_folder f on c.settings ->> 'folder_id' = f.id;
```

```sql+sqlite
select
  f.id as folder_id,
  f.name as folder_name,
  c.id as campaign_id,
  c.title as campaign_title
from
  mailchimp_campaign c
  left join mailchimp_campaign_folder f on json_extract(c.settings, '$.folder_id') = f.id;
```