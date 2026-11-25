-- Why a user received a notification, and whether that counts as direct.
create table notification_reasons (
    code       text primary key,
    -- true when the user was targeted personally; used by preference 'direct'
    is_direct  boolean not null,
    created_at timestamptz not null
);

-- noqa: disable=LT01
insert into notification_reasons (code, is_direct, created_at)
values
    ('mentioned',   true,  now()),
    ('participant', true,  now()),
    ('invitee',     true,  now()),
    ('approver',    true,  now()),
    ('requester',   true,  now()),
    ('file_owner',  false, now()),
    ('follower',    false, now());
-- noqa: enable=LT01

grant select
on table notification_reasons
to notification_runtime;
