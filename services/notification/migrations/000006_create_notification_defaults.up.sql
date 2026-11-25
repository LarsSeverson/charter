-- How each kind is sent on each channel when a user has no preference.
create table notification_defaults (
    kind_code        text not null references notification_kinds (code),
    channel_code     text not null references notification_channels (code),
    preference_level preference_level not null,
    digest_window    interval,                   -- NULL = send immediately
    primary key (kind_code, channel_code)
);

-- noqa: disable=LT01,LT05
insert into notification_defaults (kind_code, channel_code, preference_level, digest_window)
values
    ('comment_created',         'email', 'all',  interval '15 minutes'),
    ('comment_created',         'push',  'none', null),
    ('comment_created',         'chat',  'none', null),
    ('comment_replied',         'email', 'all',  interval '15 minutes'),
    ('comment_replied',         'push',  'all',  null),
    ('comment_replied',         'chat',  'none', null),
    ('user_mentioned',          'email', 'all',  null),
    ('user_mentioned',          'push',  'all',  null),
    ('user_mentioned',          'chat',  'none', null),
    ('file_invited',            'email', 'all',  null),
    ('file_invited',            'push',  'all',  null),
    ('file_invited',            'chat',  'none', null),
    ('access_requested',        'email', 'all',  null),
    ('access_requested',        'push',  'all',  null),
    ('access_requested',        'chat',  'none', null),
    ('access_request_resolved', 'email', 'all',  null),
    ('access_request_resolved', 'push',  'all',  null),
    ('access_request_resolved', 'chat',  'none', null);
-- noqa: enable=LT01,LT05

grant select
on table notification_defaults
to notification_runtime;
