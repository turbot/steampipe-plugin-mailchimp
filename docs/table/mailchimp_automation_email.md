# Table: mailchimp_automation_email

Mailchimp's classic automations feature lets you build a series of emails that send to subscribers when triggered by a specific date, activity, or event. Automation emails are individual emails in a classic automation workflow.

## Examples

### Basic info

```sql
select
  id,
  content_type,
  create_time,
  emails_sent,
  position,
  send_time,
  start_time,
  status,
  workflow_id
from
  mailchimp_automation_email;
```

### Get delay settings of an automation email

```sql
select
  id,
  delay ->> 'action' as delay_action,
  delay ->> 'amount' as delay_amount,
  delay ->> 'direction' as delay_direction,
  delay ->> 'type' as delay_type
from
  mailchimp_automation_email;
```

### Get recipient settings of an automation email

```sql
select
  id,
  recipients ->> 'list_id' as recipient_list_id,
  recipients -> 'segment_options' ->> 'conditions' as recipient_segment_conditions,
  recipients -> 'segment_options' ->> 'match' as recipient_segment_match,
  recipients -> 'segment_options' ->> 'saved_segment_id' as saved_segment_id
from
  mailchimp_automation_email;
```

### Get settings of an automation email

```sql
select
  id,
  settings ->> 'authenticate' as authenticate,
  settings ->> 'auto_fb_post' as auto_fb_post,
  settings ->> 'auto_footer' as auto_footer,
  settings ->> 'auto_tweet' as auto_footer,
  settings ->> 'drag_and_drop' as drag_and_drop,
  settings ->> 'fb_comments' as fb_comments,
  settings ->> 'from_name' as from_name,
  settings ->> 'inline_css' as inline_css,
  settings ->> 'reply_to' as reply_to,
  settings ->> 'subject_line' as subject_line,
  settings ->> 'title' as title
from
  mailchimp_automation_email;
```
