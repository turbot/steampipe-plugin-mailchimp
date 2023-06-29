# Table: mailchimp_root

The root resource links to all other resources available in the API. Calling the root directory also returns details about the Mailchimp user account.

## Examples

### Basic info

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