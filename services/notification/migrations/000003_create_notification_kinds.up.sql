-- Every type of notification the platform can send.
create table notification_kinds (
    -- 'comment_replied', 'access_requested'
    code          text primary key,
    category_code text not null references notification_categories (code),
    subject_type  text not null,               -- 'comment', 'access_request'
    actionable    boolean not null,
    mandatory     boolean not null,   -- ignores preferences
    description   text,
    created_at    timestamptz not null
);

-- noqa: disable=LT01,LT05
insert into notification_kinds (code, category_code, subject_type, actionable, mandatory, description, created_at)
values
    ('comment_created',         'comments', 'comment',        false, false, null, now()),
    ('comment_replied',         'comments', 'comment',        false, false, null, now()),
    ('user_mentioned',          'comments', 'comment',        false, false, null, now()),
    ('file_invited',            'invites',  'invite',         true,  false, null, now()),
    ('access_requested',        'requests', 'access_request', true,  false, null, now()),
    ('access_request_resolved', 'requests', 'access_request', false, false, null, now());
-- noqa: enable=LT01,LT05

grant select
on table notification_kinds
to notification_runtime;
