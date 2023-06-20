# Table: mailchimp_campaign

Any distributed content, that's created and measured in Mailchimp, including regular emails, automations, landing pages, and ads.

## Examples

### Basic info

```sql
select
  id,
  title,
  content_type,
  create_time,
  emails_sent,
  send_time,
  status,
  type,
from
  mailchimp_campaign;
```

### Get the list settings for each campaign

```sql
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

### Get the settings for each campaign

```sql
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

### Get the tracking options for each campaign

```sql
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

### Get the report summary for each campaign

```sql
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

### List campaigns having delivery status enabled

```sql
select
  id,
  content_type,
  create_time,
  emails_sent,
  status,
  type,
from
  mailchimp_campaign
where
  delivery_status_enabled;
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