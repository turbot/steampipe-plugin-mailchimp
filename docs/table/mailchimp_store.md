# Table: mailchimp_store

 Mailchimp store helps you sell products directly from your website. It has the tools you need to bring your business online and start making sales.

## Examples

### Basic info

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