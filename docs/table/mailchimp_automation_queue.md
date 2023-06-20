# Table: mailchimp_automation_queue

Mailchimp's classic automations feature lets you build a series of emails that send to subscribers when triggered by a specific date, activity, or event. Automation queues are the list member queues for classic automation emails.

The `mailchimp_automation_queue` table can be used to query information about any queue, and **you must specify the automation email id** in the where or join clause using the `email_id` column.

## Examples

### Basic info

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