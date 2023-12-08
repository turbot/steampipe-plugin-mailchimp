---
title: "Steampipe Table: mailchimp_automation - Query Mailchimp Automations using SQL"
description: "Allows users to query Automations in Mailchimp, specifically the details of each automation, providing insights into email marketing campaigns and their performance."
---

# Table: mailchimp_automation - Query Mailchimp Automations using SQL

Mailchimp Automations are a feature within the Mailchimp email marketing service that allows users to create and manage automated email campaigns. These campaigns can be set up to trigger based on specific conditions, such as user behavior or predefined schedules. Automations are a powerful tool for enhancing customer engagement and optimizing marketing efforts.

## Table Usage Guide

The `mailchimp_automation` table provides insights into Automations within Mailchimp. As a digital marketer, explore automation-specific details through this table, including campaign settings, trigger events, and associated metadata. Utilize it to uncover information about your email marketing campaigns, such as their status, the number of emails sent, and the performance of each campaign.

## Examples

### Basic info
Explore the details of your email marketing campaigns, such as creation and start times, emails sent, and status, to gain insights into their performance and effectiveness. This can help you strategize and optimize future campaigns.

```sql+postgres
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

```sql+sqlite
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
Analyze the settings of an automation to understand the recipient configurations, such as the segment conditions and match options. This can be particularly useful for assessing the efficiency of your automated marketing campaigns.

```sql+postgres
select
  id,
  recipients ->> 'list_id' as recipient_list_id,
  recipients -> 'segment_options' ->> 'conditions' as recipient_segment_conditions,
  recipients -> 'segment_options' ->> 'match' as recipient_segment_match,
  recipients -> 'segment_options' ->> 'saved_segment_id' as saved_segment_id
from
  mailchimp_automation;
```

```sql+sqlite
select
  id,
  json_extract(recipients, '$.list_id') as recipient_list_id,
  json_extract(recipients, '$.segment_options.conditions') as recipient_segment_conditions,
  json_extract(recipients, '$.segment_options.match') as recipient_segment_match,
  json_extract(recipients, '$.segment_options.saved_segment_id') as saved_segment_id
from
  mailchimp_automation;
```

### Get settings of an automation
Analyze the settings to understand the configuration of an automated system. This can be useful in understanding how the system is set up, including aspects such as authentication, auto footer, reply settings, and more.

```sql+postgres
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

```sql+sqlite
select
  id,
  json_extract(settings, '$.authenticate') as authenticate,
  json_extract(settings, '$.auto_footer') as auto_footer,
  json_extract(settings, '$.from_name') as from_name,
  json_extract(settings, '$.inline_css') as inline_css,
  json_extract(settings, '$.reply_to') as reply_to,
  json_extract(settings, '$.title') as title,
  json_extract(settings, '$.to_name') as to_name,
  json_extract(settings, '$.use_conversation') as use_conversation
from
  mailchimp_automation;
```

### Get tracking options of an automation
Determine the tracking options set for an automated email marketing campaign, including various analytics and goal tracking features. This can be useful in understanding how your campaign is performing and which tracking tools are being utilized.

```sql+postgres
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

```sql+sqlite
select
  id,
  json_extract(tracking, '$.capsule.notes') as tracking_capsule_notes,
  json_extract(tracking, '$.clicktale') as tracking_clicktale,
  json_extract(tracking, '$.ecomm360') as tracking_ecomm360,
  json_extract(tracking, '$.goal_tracking') as tracking_goal_tracking,
  json_extract(tracking, '$.google_analytics') as tracking_google_analytics,
  json_extract(tracking, '$.highrise.campaign') as tracking_highrise_campaign,
  json_extract(tracking, '$.highrise.notes') as tracking_highrise_notes,
  json_extract(tracking, '$.html_clicks') as tracking_html_clicks,
  json_extract(tracking, '$.opens') as tracking_opens,
  json_extract(tracking, '$.salesforce.campaign') as tracking_salesforce_campaign,
  json_extract(tracking, '$.salesforce.notes') as tracking_salesforce_notes,
  json_extract(tracking, '$.text_clicks') as tracking_text_clicks
from
  mailchimp_automation;
```

### Get trigger settings of an automation
Explore the frequency and timing of automated tasks in a mailing system, helping you understand when and how often certain automations occur, such as sending workflow emails. This insight can assist in optimizing your communication strategy.

```sql+postgres
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

```sql+sqlite
select
  id,
  title,
  json_extract(trigger_settings, '$.runtime.days') as trigger_runtime_days,
  json_extract(trigger_settings, '$.runtime.hours') as trigger_runtime_hours,
  json_extract(trigger_settings, '$.workflow_emails_count') as trigger_workflow_emails_count,
  json_extract(trigger_settings, '$.workflow_title') as trigger_workflow_title,
  json_extract(trigger_settings, '$.workflow_type') as trigger_workflow_type
from
  mailchimp_automation;
```