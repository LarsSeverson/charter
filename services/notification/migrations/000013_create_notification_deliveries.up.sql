-- Each attempt to send a notification out on a channel, and its result.
create table notification_deliveries (
    id                  uuid primary key,
    notification_id     uuid not null references notifications (id),
    channel_code        text not null references notification_channels (code),
    -- push or chat destination; null for email and for suppressed deliveries
    destination_id      uuid references notification_destinations (id),
    digest_id           uuid references notification_digests (id),
    status              delivery_status not null,
    -- 'mandatory', 'default', 'preference', 'channel_disabled',
    -- 'no_destination'
    decided_by          text not null,
    -- the rule that decided, if any
    preference_id       uuid references notification_preferences (id),
    attempts            int not null,
    next_attempt_at     timestamptz,
    last_error          text,
    provider_message_id text,
    sent_at             timestamptz,
    created_at          timestamptz not null,
    updated_at          timestamptz not null,

    -- One delivery per notification, channel and destination. Nulls count as
    -- equal, so an email (no destination) is delivered only once.
    unique nulls not distinct (notification_id, channel_code, destination_id)
);

grant select, insert, update, delete
on table notification_deliveries
to notification_runtime;
