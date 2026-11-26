-- How each kind is sent on each channel when a user has no preference.
create table notification_defaults (
    kind_code        text not null references notification_kinds (code),
    channel_code     text not null references notification_channels (code),
    preference_level preference_level not null,
    digest_window    interval,                   -- NULL = send immediately
    primary key (kind_code, channel_code)
);

-- Example rows. Producers register their own; nothing is seeded here.
-- Replies are batched by email but pushed at once; mentions go out at once
-- everywhere except chat; password changes are mandatory, so these rows pick
-- the channels and user preferences are ignored.
--
--   kind_code         channel_code  preference_level  digest_window
--   ----------------  ------------  ----------------  -------------
--   comment_replied   email         all               15 minutes
--   comment_replied   push          all               null
--   comment_replied   chat          none              null
--   user_mentioned    email         all               null
--   user_mentioned    push          all               null
--   user_mentioned    chat          none              null
--   password_changed  email         all               null
--   password_changed  push          none              null
--   password_changed  chat          none              null

grant select
on table notification_defaults
to notification_runtime;
