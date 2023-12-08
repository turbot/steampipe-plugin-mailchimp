---
title: "Steampipe Table: mailchimp_batch_operation - Query Mailchimp Batch Operations using SQL"
description: "Allows users to query Mailchimp Batch Operations, specifically the status, total operations, and response body, providing insights into operation results and potential issues."
---

# Table: mailchimp_batch_operation - Query Mailchimp Batch Operations using SQL

Mailchimp Batch Operations is a feature within Mailchimp that allows you to execute multiple operations in a single API request. It enables you to perform bulk actions such as creating, updating, or deleting a large number of resources at once, saving time and reducing the number of API calls needed. Mailchimp Batch Operations simplifies the management of large datasets and enhances the efficiency of your campaigns.

## Table Usage Guide

The `mailchimp_batch_operation` table provides insights into batch operations within Mailchimp. As a data analyst or marketing professional, explore operation-specific details through this table, including status, completed operations, and error details. Utilize it to monitor the progress of your batch operations, identify failed operations, and analyze the response body for troubleshooting and optimization.

## Examples

### Basic info
Discover the segments that have completed their operations in your Mailchimp batch process. This query can help you assess the performance and status of your operations, providing essential insights for potential troubleshooting and optimization.

```sql+postgres
select
  id,
  submitted_at,
  completed_at,
  status,
  total_operations
from
  mailchimp_batch_operation;
```

```sql+sqlite
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
Determine the status of each batch operation in your Mailchimp account. This helps in understanding the progress and possible issues with your ongoing operations, aiding in efficient troubleshooting and management.

```sql+postgres
select
  id,
  status
  errored_operations,
  finished_operations,
  total_operations
from
  mailchimp_batch_operation;
```

```sql+sqlite
select
  id,
  status,
  errored_operations,
  finished_operations,
  total_operations
from
  mailchimp_batch_operation;
```

### Get total time taken by each batch operation
Explore the efficiency of your batch operations by determining the total time each one takes. This can help identify bottlenecks and optimize your processes.

```sql+postgres
select
  id,
  (extract(epoch from (completed_at - submitted_at))) || ' seconds' as time_taken,
  total_operations,
  errored_operations,
  finished_operations
from
  mailchimp_batch_operation;
```

```sql+sqlite
select
  id,
  (strftime('%s', completed_at) - strftime('%s', submitted_at)) || ' seconds' as time_taken,
  total_operations,
  errored_operations,
  finished_operations
from
  mailchimp_batch_operation;
```

### List batch operations completed in the last 10 days
Explore which batch operations have been completed in the past 10 days. This is useful to quickly assess recent activity and status, helping you to keep track of your operations.

```sql+postgres
select
  id,
  submitted_at,
  completed_at,
  status,
  total_operations
from
  mailchimp_batch_operation
where
  completed_at >= now() - interval '10' day;
```

```sql+sqlite
select
  id,
  submitted_at,
  completed_at,
  status,
  total_operations
from
  mailchimp_batch_operation
where
  completed_at >= datetime('now', '-10 day');
```