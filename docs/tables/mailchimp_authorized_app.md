---
title: "Steampipe Table: mailchimp_authorized_app - Query Mailchimp Authorized Apps using SQL"
description: "Allows users to query Mailchimp Authorized Apps, providing details about the apps which have been authorized to access a Mailchimp account."
---

# Table: mailchimp_authorized_app - Query Mailchimp Authorized Apps using SQL

Mailchimp Authorized Apps are applications that have been granted permission to access a Mailchimp account. This access is granted by the account owner, and can be revoked at any time. Authorized Apps can perform various actions on the account, limited by the permissions set by the account owner.

## Table Usage Guide

The `mailchimp_authorized_app` table provides insights into the apps authorized to access a Mailchimp account. As an account owner or administrator, you can use this table to manage and monitor app access to your Mailchimp account. It can be used to review app details, including permissions, access levels, and associated metadata.

## Examples

### Basic info
Explore which authorized applications are connected to your Mailchimp account, gaining insights into their names and descriptions for better management and oversight.

```sql
select
  id,
  name,
  description
from
  mailchimp_authorized_app;
```

### List users who have linked the app
Explore which users have connected their accounts with the application, useful for understanding user engagement and app utilization.

```sql
select
  id,
  name,
  description,
  u as user
from
  mailchimp_authorized_app,
  jsonb_array_elements_text(users) u;
```