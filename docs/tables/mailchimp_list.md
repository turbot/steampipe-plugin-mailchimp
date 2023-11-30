---
title: "Steampipe Table: mailchimp_list - Query Mailchimp Lists using SQL"
description: "Allows users to query Mailchimp Lists, providing insights into subscriber lists and their associated details."
---

# Table: mailchimp_list - Query Mailchimp Lists using SQL

Mailchimp Lists are a fundamental component of the Mailchimp platform. They represent a collection of contacts and subscribers that users can target and manage for marketing campaigns. Lists contain detailed information about subscribers, including their email addresses, subscription status, and other associated data.

## Table Usage Guide

The `mailchimp_list` table provides insights into Lists within Mailchimp. As a marketing professional or business owner, explore list-specific details through this table, including subscriber count, campaign statistics, and associated metadata. Utilize it to uncover information about your Lists, such as those with high unsubscribe rates, the growth of your lists over time, and the performance of your marketing campaigns.

## Examples

### Basic info
Gain insights into your Mailchimp lists by identifying their creation dates and visibility settings. This can help you understand the evolution of your email marketing efforts and assess the accessibility of your lists.

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
Explore the default settings of each marketing campaign to understand the sender's email and name, the subject line, and the language used. This could be beneficial for assessing consistency in branding or identifying areas for personalization.

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
Explore which audience members are associated with specific companies and locations. This is useful for tailoring marketing campaigns or communications to specific geographical areas or business sectors.

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
Explore the performance of each audience segment by evaluating statistics such as total contacts, unsubscribe rate, and campaign engagement. This information can be used to understand audience behavior and optimize your marketing strategies.

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