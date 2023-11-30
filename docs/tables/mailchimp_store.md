---
title: "Steampipe Table: mailchimp_store - Query Mailchimp Stores using SQL"
description: "Allows users to query Mailchimp Stores, providing detailed information about each store including the store name, currency code, domain, primary locale, and more."
---

# Table: mailchimp_store - Query Mailchimp Stores using SQL

Mailchimp Stores is a resource within the Mailchimp service that allows users to connect their online store to Mailchimp. Once connected, users can create targeted campaigns, automate helpful product follow-ups, and send back-in-stock messaging. It provides a centralized way to manage and monitor the performance of your online store.

## Table Usage Guide

The `mailchimp_store` table provides insights into the stores connected within Mailchimp. As an e-commerce manager or digital marketer, explore store-specific details through this table, including the store name, currency code, domain, and primary locale. Utilize it to manage and monitor the performance of your online store, enabling targeted campaigns and effective product follow-ups.

## Examples

### Basic info
Explore the fundamental information related to your online store, such as its creation date, domain, platform, and primary locale. This can help you gain insights into your store's operational details and understand its configuration better.

```sql
select
  id,
  name,
  created_at,
  currency_code,
  domain,
  money_format,
  platform,
  primary_locale
from
  mailchimp_store;
```

### Get contact info of each store
Explore which stores you have contact information for, allowing you to reach out for business communications or updates. This can be particularly beneficial when managing customer relationships or conducting marketing outreach.

```sql
select
  id,
  name,
  email_address,
  phone,
  address,
  timezone
from
  mailchimp_store;
```

### Get details of the audience associated with each store
Explore the relationship between different stores and their associated audiences. This can help in understanding the reach of each store and strategizing marketing efforts accordingly.

```sql
select
  s.id as store_id,
  s.name as store_name,
  l.id as list_id,
  l.name as list_name,
  l.date_created as list_date_created,
  l.visibility as list_visibility
from
  mailchimp_store s,
  mailchimp_list l
where
  s.list_id = l.id;
```

### List stores created in the last 30 days
Discover the segments that have recently added stores in the past 30 days. This is particularly useful for tracking growth and understanding recent market activity.

```sql
select
  id,
  name,
  created_at,
  currency_code,
  domain,
  money_format,
  platform,
  primary_locale
from
  mailchimp_store
where
  created_at >= now() - interval '30' day;
```

### List stores which haven't been updated in the last 10 days
Explore which stores have not been updated recently to identify potential areas for review or maintenance. This is useful for ensuring your store information is current and accurate, which can enhance customer experience and business operations.

```sql
select
  id,
  name,
  created_at,
  currency_code,
  domain,
  money_format,
  platform,
  primary_locale
from
  mailchimp_store
where
  updated_at <= now() - interval '10' day;
```