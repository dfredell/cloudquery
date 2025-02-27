# Table: stripe_treasury_credit_reversals

This table shows data for Stripe Treasury Credit Reversals.

https://stripe.com/docs/api/treasury/credit_reversals

The primary key for this table is **id**.

## Relations

This table depends on [stripe_treasury_financial_accounts](stripe_treasury_financial_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|amount|`int64`|
|created|`timestamp[us, tz=UTC]`|
|currency|`utf8`|
|financial_account|`utf8`|
|hosted_regulatory_receipt_url|`utf8`|
|livemode|`bool`|
|metadata|`json`|
|network|`utf8`|
|object|`utf8`|
|received_credit|`utf8`|
|status|`utf8`|
|status_transitions|`json`|
|transaction|`json`|