-- Push devices and chat accounts registered for a user. Email addresses come
-- from the user service.
create table notification_destinations (
    id           uuid primary key,
    user_id      uuid not null,
    channel_code text not null references notification_channels (code),
    -- push token or chat user ID
    address      text not null,
    -- platform, chat workspace, ...
    metadata     jsonb not null,
    created_at   timestamptz not null,

    -- An address belongs to one user. A token seen on a new account moves
    -- there instead of being duplicated.
    unique (channel_code, address)
);

grant select, insert, update, delete
on table notification_destinations
to notification_runtime;

-- TODO for Lars: Think more about 'address'
