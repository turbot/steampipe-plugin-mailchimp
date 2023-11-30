---
title: "Steampipe Table: mailchimp_automation_queue - Query Mailchimp Automation Queues using SQL"
description: "Allows users to query Mailchimp Automation Queues, specifically the emails that are waiting to be sent or are in the process of being sent as part of an automation."
---

# Table: mailchimp_automation_queue - Query Mailchimp Automation Queues using SQL

Mailchimp Automation Queues are a feature of Mailchimp, a marketing automation platform and an email marketing service. These queues contain emails that are waiting to be sent or are in the process of being sent as part of an automation. Automations are a way to send targeted emails that are triggered by a specific date, event, or contact's activity.

## Table Usage Guide

The `mailchimp_automation_queue` table provides insights into Mailchimp's automation queues. As a marketing professional or business owner, you can explore details about the queued emails in your Mailchimp automations, including their status, the time they are scheduled to send, and the email addresses they are being sent to. Utilize this table to monitor your email marketing campaigns, ensure your automations are working as expected, and identify any potential issues.

## Examples

### Basic info
Analyze the settings to understand the details of queued emails in a specific Mailchimp automation. This can be used to pinpoint the specific locations where a particular email is scheduled to be sent, providing insights into your email marketing strategy.

```sql
select
  id,
  email_id,
  email_address,
  list_id,
  next_send,
  workflow_id
from
  mailchimp_automation_queue
where
  email_id = '123abc';
```

### Check if an email is automated to be sent in the next 3 days
Gauge whether an automated email is scheduled to be dispatched within the next three days. This can be beneficial for managing communications and ensuring timely delivery of important messages.

```sql
select
  id,
  email_id,
  email_address,
  list_id,
  next_send,
  workflow_id
from
  mailchimp_automation_queue
where
  email_id = '123abc'
  and next_send <= now() - interval '3' day;
```