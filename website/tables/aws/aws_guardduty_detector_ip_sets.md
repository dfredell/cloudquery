# Table: aws_guardduty_detector_ip_sets

This table shows data for Amazon GuardDuty Detector IP Sets.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetFilter.html

The composite primary key for this table is (**detector_arn**, **name**).

## Relations

This table depends on [aws_guardduty_detectors](aws_guardduty_detectors).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|detector_arn (PK)|`utf8`|
|format|`utf8`|
|location|`utf8`|
|name (PK)|`utf8`|
|status|`utf8`|
|tags|`json`|