-- One notification per user per event. This is the user's inbox.
create table notifications (
    id           uuid primary key,
    event_id     uuid not null references notification_events (id),
    recipient_id uuid not null,
    reason_code  text not null references notification_reasons (code),
    read_at      timestamptz,
    archived_at  timestamptz,
    created_at   timestamptz not null,

    -- An event lands in a user's inbox at most once.
    unique (event_id, recipient_id)
);

grant select, insert, update, delete
on table notifications
to notification_runtime;
