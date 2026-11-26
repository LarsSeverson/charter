-- Groups of notification kinds that users manage together.
create table notification_categories (
    code         text primary key,
    display_name text not null,
    description  text,
    created_at   timestamptz not null
);

-- Example rows. Producers register their own; nothing is seeded here.
--
--   code       display_name  description
--   ---------  ------------  ------------------------------------------
--   comments   Comments      New comments, replies, and mentions
--   requests   Requests      Access requests and their outcomes
--   account    Account       Security, billing, and account changes

grant select
on table notification_categories
to notification_runtime;
