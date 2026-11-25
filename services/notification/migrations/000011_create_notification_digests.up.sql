-- Notifications grouped together to be sent as one message.
create table notification_digests (
    id                  uuid primary key,
    user_id             uuid not null,
    channel_code        text not null references notification_channels (code),
    -- what is grouped, e.g. 'comments:file:…'
    digest_key          text not null,
    status              digest_status not null,
    closes_at           timestamptz not null,
    sent_at             timestamptz,
    provider_message_id text,
    created_at          timestamptz not null
);

-- At most one open digest per user, channel and key, so notifications
-- arriving at the same time join one batch instead of opening two.
create unique index one_open_digest
on notification_digests (user_id, channel_code, digest_key)
where status = 'open';

grant select, insert, update, delete
on table notification_digests
to notification_runtime;
