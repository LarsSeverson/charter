-- Groups of notification kinds that users manage together.
create table notification_categories (
    code         text primary key,
    display_name text not null,
    description  text,
    created_at   timestamptz not null
);

-- noqa: disable=LT01,LT05
insert into notification_categories (code, display_name, description, created_at)
values
    ('comments', 'Comments', 'New comments, replies, and mentions',         now()),
    ('invites',  'Invites',  'Invitations to files, projects, and teams',   now()),
    ('requests', 'Requests', 'Access, upgrade, plugin, and merge requests', now()),
    ('account',  'Account',  'Security, billing, and account changes',      now());
-- noqa: enable=LT01,LT05

grant select
on table notification_categories
to notification_runtime;
