---
title: "Steampipe Table: mailchimp_automation_email - Query Mailchimp Automation Emails using SQL"
description: "Allows users to query Automation Emails in Mailchimp, specifically the email details in an automation workflow, providing insights into email marketing campaigns and automation workflows."
---

# Table: mailchimp_automation_email - Query Mailchimp Automation Emails using SQL

Mailchimp Automation is a feature within Mailchimp that allows marketers to create targeted emails that send when triggered by a specific date, event, or contact's activity. It provides a way to build and manage automated marketing campaigns, including the ability to send a series of emails to contacts who meet certain conditions. Mailchimp Automation helps users stay informed about the performance of their marketing campaigns and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `mailchimp_automation_email` table provides insights into automation emails within Mailchimp. As a marketing professional, explore email-specific details through this table, including the status, send time, and associated metadata. Utilize it to uncover information about emails, such as those with high engagement rates, the timing of emails, and the verification of send conditions.

## Examples

### Basic info
Explore the timing, status, and associated workflow of your automated emails within Mailchimp. This allows you to gain insights into your email marketing strategy's effectiveness and make necessary adjustments.

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
Explore the delay settings of an automated email to understand how long it waits before taking action. This can be useful in optimizing the timing of your email campaigns.

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
Explore the specific settings of automated emails to understand the recipient details and segmentation conditions. This can be beneficial in tailoring your marketing strategies by analyzing the recipient's segmentation conditions and preferences.

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
Analyze the settings of an automated email to understand its various features such as authentication, auto posting on Facebook, auto footer, drag and drop functionality, Facebook comments, sender's name, inline CSS, reply-to address, subject line, and title.

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