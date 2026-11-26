-- One notification per user per event. This is the user's inbox.
create table notifications (
    id           uuid primary key,
    event_id     uuid not null references notification_events (id),
    recipient_id uuid not null,
    -- why the producer chose this recipient, e.g. 'mentioned', 'file_owner';
    -- opaque to this service, only shown to the user
    reason       text not null,
    -- true when the recipient was targeted personally; used by preference
    -- 'direct'
    direct       boolean not null,
    read_at      timestamptz,
    archived_at  timestamptz,
    created_at   timestamptz not null,

    -- An event lands in a user's inbox at most once.
    unique (event_id, recipient_id)
);

grant select, insert, update, delete
on table notifications
to notification_runtime;
