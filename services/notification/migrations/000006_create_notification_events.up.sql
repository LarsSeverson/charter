-- Things that happened in other services that may notify users.
create table notification_events (
    id              uuid primary key,
    source_service  text not null,            -- producing service
    source_event_id text not null,            -- producer's ID, for idempotency
    kind_code       text not null references notification_kinds (code),
    actor_type      actor_type not null,
    actor_id        uuid,                     -- NULL when actor_type = 'system'
    -- what the event is about in the producing service, e.g. a comment ID;
    -- with source_service, finds every notification about one thing
    subject_id      text not null,
    -- ids and hints code needs to render and route, e.g. {"file_id": "f9"};
    -- never the source of truth
    metadata        jsonb not null,
    occurred_at     timestamptz not null,
    received_at     timestamptz not null,

    -- A producer's event is recorded once, even if it is delivered twice.
    unique (source_service, source_event_id),

    -- ensure system does not have invalid actor_id
    check ((actor_type = 'system') = (actor_id is null))
);

grant select, insert, delete
on table notification_events
to notification_runtime;
