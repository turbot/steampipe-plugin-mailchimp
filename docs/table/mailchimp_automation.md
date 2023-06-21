# Table: mailchimp_automation

Mailchimp's classic automations feature lets you build a series of emails that send to subscribers when triggered by a specific date, activity, or event.

## Examples

### Basic info

```sql
select
  id,
  title,
  create_time,
  start_time,
  emails_sent,
  status
from
  mailchimp_automation;
```

### Get recipient settings of an automation

```sql
select
  id,
  recipients ->> 'list_id' as recipient_list_id,
  recipients -> 'segment_options' ->> 'conditions' as recipient_segment_conditions,
  recipients -> 'segment_options' ->> 'match' as recipient_segment_match,
  recipients -> 'segment_options' ->> 'saved_segment_id' as saved_segment_id
from
  mailchimp_automation;
```

### Get settings of an automation

```sql
select
  id,
  settings ->> 'authenticate' as authenticate,
  settings ->> 'auto_footer' as auto_footer,
  settings ->> 'from_name' as from_name,
  settings ->> 'inline_css' as inline_css,
  settings ->> 'reply_to' as reply_to,
  settings ->> 'title' as title,
  settings ->> 'to_name' as to_name,
  settings ->> 'use_conversation' as use_conversation
from
  mailchimp_automation;
```

### Get tracking options of an automation

```sql
select
  id,
  tracking -> 'capsule' ->>'notes' as tracking_capsule_notes,
  tracking ->> 'clicktale' as tracking_clicktale,
  tracking ->> 'ecomm360' as tracking_ecomm360,
  tracking ->> 'goal_tracking' as tracking_goal_tracking,
  tracking ->> 'google_analytics' as tracking_google_analytics,
  tracking -> 'highrise' ->> 'campaign' as tracking_highrise_campaign,
  tracking -> 'highrise' ->> 'notes' as tracking_highrise_notes,
  tracking ->> 'html_clicks' as tracking_html_clicks,
  tracking ->> 'opens' as tracking_opens,
  tracking -> 'salesforce' ->> 'campaign' as tracking_salesforce_campaign,
  tracking -> 'salesforce' ->> 'notes' as tracking_salesforce_notes,
  tracking ->> 'text_clicks' as tracking_text_clicks
from
  mailchimp_automation;
```

### Get trigger settings of an automation

```sql
select
  id,
  title,
  trigger_settings -> 'runtime' ->> 'days' as trigger_runtime_days,
  trigger_settings -> 'runtime' ->> 'hours' as trigger_runtime_hours,
  trigger_settings ->> 'workflow_emails_count' as trigger_workflow_emails_count,
  trigger_settings ->> 'workflow_title' as trigger_workflow_title,
  trigger_settings ->> 'workflow_type' as trigger_workflow_type
from
  mailchimp_automation;
```