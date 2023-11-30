---
title: "Steampipe Table: mailchimp_root - Query MailChimp Account using SQL"
description: "Allows users to query MailChimp Accounts, specifically the account details about the authenticated MailChimp user."
---

# Table: mailchimp_root - Query MailChimp Account using SQL

MailChimp is a marketing automation platform and an email marketing service. It provides a platform for users to create, send, and manage email newsletters. The platform also offers various integrations and features to help manage and talk to their audience and customers.

## Table Usage Guide

The `mailchimp_root` table provides insights into the authenticated MailChimp user's account details. As a marketing professional or business owner, explore account-specific details through this table, including contact details, industry type, and plan details. Utilize it to uncover information about your account, such as the number of audiences, the number of campaigns, and the date of the last login.

## Examples

### Basic info
Explore the general information of your Mailchimp account, such as the account name, email, last login date, pro status, role, and total subscribers. This can be useful for understanding the account's status and user engagement level.

```sql
select
  account_id,
  account_name,
  email,
  last_login,
  pro_enabled,
  role,
  total_subscribers
from
  mailchimp_root;
```

### Get contact details of the account
This query allows you to gain insights into the contact details associated with each account. This can be beneficial for tasks such as account management, marketing outreach, or customer service inquiries.

```sql
select
  account_id,
  account_name,
  contact ->> 'addr1' as address1,
  contact ->> 'addr2' as address2,
  contact ->> 'city' as city,
  contact ->> 'company' as company,
  contact ->> 'country' as country,
  contact ->> 'state' as state,
  contact ->> 'zip' as zip
from
  mailchimp_root;
```

### Get the industry's average campaign statistics of the account
Explore the average campaign statistics within a specific industry to understand performance trends such as open rate, bounce rate, and click rate. This can help in benchmarking your account's performance against industry averages.

```sql
select
  account_id,
  account_name,
  industry_stats ->> 'open_rate' as open_rate,
  industry_stats ->> 'bounce_rate' as bounce_rate,
  industry_stats ->> 'click_rate' as click_rate
from
  mailchimp_root;
```

### Get the details of the users who havent't logged in in the last 30 days
Discover users who have been inactive for the past month. This query is useful for identifying potential user churn or inactive accounts that may require re-engagement efforts.

```sql
select
  account_id,
  account_name,
  email,
  last_login,
  role,
  total_subscribers
from
  mailchimp_root
where
  last_login <= now() - interval '30' day;
```

### Get the details of the users who use Mailchimp Pro version
Discover the segments of users who have opted for the Pro version of Mailchimp. This is beneficial for understanding user preferences and tailoring marketing strategies accordingly.

```sql
select
  account_id,
  account_name,
  email,
  last_login,
  role,
  total_subscribers
from
  mailchimp_root
where
  pro_enabled;
```