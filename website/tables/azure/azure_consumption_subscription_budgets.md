# Table: azure_consumption_subscription_budgets

This table shows data for Azure Consumption Subscription Budgets.

https://learn.microsoft.com/en-us/rest/api/consumption/budgets/list?tabs=HTTP#budget

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|e_tag|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|