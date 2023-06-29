# Table: mailchimp_list

Mailchimp list is a database containing information about your contacts to whom you send emails in bulk using Mailchimp.

## Examples

### Basic info

```sql
select
  id,
  name,
  date_created,
  visibility
from
  mailchimp_list;
```

### Get the campaign defaults of each audience

```sql
select
  id,
  campaign_defaults ->> 'from_email' as from_email,
  campaign_defaults ->> 'from_name' as from_name,
  campaign_defaults ->> 'subject' as subject,
  campaign_defaults ->> 'language' as language
from
  mailchimp_list;
```

### Get the contact information of each audience

```sql
select
  id,
  contact ->> 'company' as company,
  contact ->> 'address1' as address1,
  contact ->> 'address2' as address2,
  contact ->> 'city' as city,
  contact ->> 'state' as state,
  contact ->> 'zip' as zip,
  contact ->> 'country' as country,
  contact ->> 'phone' as phone
from
  mailchimp_list;
```

### Get the stats of each audience

```sql
select
  id,
  stats ->> 'member_count' as member_count,
  stats ->> 'total_contacts' as total_contacts,
  stats ->> 'unsubscribe_count' as unsubscribe_count,
  stats ->> 'cleaned_count' as cleaned_count,
  stats ->> 'member_count_since_send' as member_count_since_send,
  stats ->> 'unsubscribe_count_since_send' as unsubscribe_count_since_send,
  stats ->> 'cleaned_count_since_send' as cleaned_count_since_send,
  stats ->> 'campaign_count' as campaign_count,
  stats ->> 'campaign_last_sent' as campaign_last_sent,
  stats ->> 'merge_field_count' as merge_field_count,
  stats ->> 'avg_sub_rate' as avg_subscribe_rate,
  stats ->> 'avg_unsub_rate' as avg_unsubscribe_rate,
  stats ->> 'target_sub_rate' as target_subscribe_rate,
  stats ->> 'open_rate' as open_rate,
  stats ->> 'click_rate' as click_rate,
  stats ->> 'last_sub_date' as last_subscribe_date,
  stats ->> 'last_unsub_date' as last_unsubscribe_date
from
  mailchimp_list;
```