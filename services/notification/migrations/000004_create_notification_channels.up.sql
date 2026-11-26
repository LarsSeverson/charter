-- External channels a notification can be sent out through.
create table notification_channels (
    code         text primary key,
    display_name text not null,
    -- turns the channel off platform-wide
    enabled      boolean not null,
    created_at   timestamptz not null
);

-- noqa: disable=LT01
insert into notification_channels (code, display_name, enabled, created_at)
values
    ('email', 'Email', true, now()),
    ('push',  'Push',  true, now()),
    ('chat',  'Chat',  true, now());
-- noqa: enable=LT01

grant select
on table notification_channels
to notification_runtime;

grant update (enabled)
on table notification_channels
to notification_runtime;
