# Table: mailchimp_batch_operation

Mailchimp's batch operations are used to complete multiple operations with a single call.

## Examples

### Basic info

```sql
select
  id,
  submitted_at,
  completed_at,
  status,
  total_operations
from
  mailchimp_batch_operation;
```

### Get status of each batch operation

```sql
select
  id,
  status
  errored_operations,
  finished_operations,
  total_operations
from
  mailchimp_batch_operation;
```

### Get total time taken by each batch operation

```sql
select
  id,
  (extract(epoch from (completed_at - submitted_at))) || ' seconds' as time_taken,
  total_operations,
  errored_operations,
  finished_operations
from
  mailchimp_batch_operation;
```