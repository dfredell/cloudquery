# Table: azure_consumption_billing_account_legacy_usage_details

This table shows data for Azure Consumption Billing Account Legacy Usage Details.

https://learn.microsoft.com/en-us/rest/api/consumption/usage-details/list?tabs=HTTP#legacyusagedetail

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|kind|`utf8`|
|properties|`json`|
|etag|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|tags|`json`|
|type|`utf8`|