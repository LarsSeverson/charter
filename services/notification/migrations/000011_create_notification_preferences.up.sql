-- Each user's overrides of notification_defaults. When several rows match,
-- the most specific wins: kind, then category, then everything; within those,
-- a named channel beats all channels. Mandatory kinds ignore this table.
create table notification_preferences (
    id               uuid primary key,
    user_id          uuid not null,
    -- null = not limited to one category
    category_code    text references notification_categories (code),
    -- null = not limited to one kind
    kind_code        text references notification_kinds (code),
    -- null = every channel
    channel_code     text references notification_channels (code),
    preference_level preference_level not null,
    origin           preference_source not null,
    created_at       timestamptz not null,
    updated_at       timestamptz not null,

    -- One preference per user, scope and channel. Nulls count as equal, so
    -- two "all channels" rows for the same scope are rejected.
    unique nulls not distinct (user_id, category_code, kind_code, channel_code),

    -- A preference applies to a category or a kind, never both; neither means
    -- it applies to everything.
    check (category_code is null or kind_code is null)
);

-- Example rows, from the user's settings actions:
--
--   "Turn off everything"          null      null             null   none
--   "...except Comments"           comments  null             null   all
--   Unsubscribe from reply emails  null      comment_replied  email  none

grant select, insert, update, delete
on table notification_preferences
to notification_runtime;
