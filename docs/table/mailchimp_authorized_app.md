# Table: mailchimp_authorized_app

An authorized application is an app that has permission to write or view data in your Mailchimp account.

## Examples

### Basic info

```sql
select
  id,
  name,
  description
from
  mailchimp_authorized_app;
```

### List users who have linked the app

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