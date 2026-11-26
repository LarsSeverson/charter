-- Notifications batched for one destination and sent as one message.
create table notification_digests (
    id                  uuid primary key,
    -- where the digest is sent; the user and channel come from here
    destination_id      uuid not null references notification_destinations (id),
    status              digest_status not null,
    -- set by the delivery that opens the digest; later ones join without
    -- moving it
    closes_at           timestamptz not null,
    sent_at             timestamptz,
    provider_message_id text,
    created_at          timestamptz not null
);

-- At most one open digest per destination, so everything batched for an
-- address in the same window goes out as one message.
create unique index one_open_digest
on notification_digests (destination_id)
where status = 'open';

grant select, insert, update, delete
on table notification_digests
to notification_runtime;
