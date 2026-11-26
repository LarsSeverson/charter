-- Every type of notification the platform can send.
create table notification_kinds (
    -- 'comment_replied', 'access_requested'
    code          text primary key,
    category_code text not null references notification_categories (code),
    -- label shown to users in notification settings
    display_name  text not null,
    -- sent on every channel its defaults allow, ignoring user preferences
    mandatory     boolean not null,
    description   text,
    created_at    timestamptz not null
);

-- Example rows. Producers register their own; nothing is seeded here.
--
--   code              category_code  display_name      mandatory
--   ----------------  -------------  ----------------  ---------
--   comment_replied   comments       Comment replies   false
--   user_mentioned    comments       Mentions          false
--   access_requested  requests       Access requests   false
--   password_changed  account        Password changes  true

grant select
on table notification_kinds
to notification_runtime;
