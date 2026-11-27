-- How each kind is sent on each channel when a user has no preference.
create table notification_defaults (
    kind_code        text not null references notification_kinds (code),
    channel_code     text not null references notification_channels (code),
    preference_level preference_level not null,
    -- how long to collect these into one digest; null = not batched
    digest_window    interval,
    -- how long to wait before sending, so it is skipped if the user reads it
    -- in the app first; null = send right away
    unread_delay     interval,
    primary key (kind_code, channel_code),

    -- A row waits either as a single delivery or in a digest, never both. A
    -- digest already checks what was read when it closes.
    check (unread_delay is null or digest_window is null)
);

-- Example rows. Producers register their own; nothing is seeded here.
-- Mentions buzz after a minute and email after ten if still unread. Replies
-- buzz after a minute and are batched by email. Password changes are
-- mandatory and email at once, even if already read.
--
--   kind_code         channel  level  digest_window  unread_delay
--   ----------------  -------  -----  -------------  ------------
--   user_mentioned    push     all    null           1 minute
--   user_mentioned    email    all    null           10 minutes
--   user_mentioned    chat     none   null           null
--   comment_replied   push     all    null           1 minute
--   comment_replied   email    all    15 minutes     null
--   comment_replied   chat     none   null           null
--   password_changed  email    all    null           null
--   password_changed  push     none   null           null
--   password_changed  chat     none   null           null

grant select
on table notification_defaults
to notification_runtime;
