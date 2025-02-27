# Table: oracle_blockstorage_boot_volumes

This table shows data for Oracle Block Storage Boot Volumes.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|availability_domain|`utf8`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|size_in_m_bs|`int64`|
|time_created|`timestamp[us, tz=UTC]`|
|defined_tags|`json`|
|system_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|image_id|`utf8`|
|is_hydrated|`bool`|
|vpus_per_gb|`int64`|
|size_in_g_bs|`int64`|
|volume_group_id|`utf8`|
|kms_key_id|`utf8`|
|is_auto_tune_enabled|`bool`|
|auto_tuned_vpus_per_gb|`int64`|
|boot_volume_replicas|`json`|
|autotune_policies|`json`|