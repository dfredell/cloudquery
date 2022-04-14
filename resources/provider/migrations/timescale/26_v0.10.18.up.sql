-- Autogenerated by migration tool on 2022-04-13 11:51:48

-- Resource: autoscaling.scheduled_actions
CREATE TABLE IF NOT EXISTS "aws_autoscaling_scheduled_actions" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"auto_scaling_group_name" text,
	"desired_capacity" integer,
	"end_time" timestamp without time zone,
	"max_size" integer,
	"min_size" integer,
	"recurrence" text,
	"arn" text,
	"name" text,
	"start_time" timestamp without time zone,
	"time" timestamp without time zone,
	"time_zone" text,
	CONSTRAINT aws_autoscaling_scheduled_actions_pk PRIMARY KEY(cq_fetch_date,arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_autoscaling_scheduled_actions');
