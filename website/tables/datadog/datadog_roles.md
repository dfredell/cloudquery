# Table: datadog_roles

This table shows data for Datadog Roles.

The composite primary key for this table is (**account_name**, **id**).

## Relations

The following tables depend on datadog_roles:
  - [datadog_role_permissions](datadog_role_permissions)
  - [datadog_role_users](datadog_role_users)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|attributes|`json`|
|id (PK)|`utf8`|
|relationships|`json`|
|type|`utf8`|
|additional_properties|`json`|