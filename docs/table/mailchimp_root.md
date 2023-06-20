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