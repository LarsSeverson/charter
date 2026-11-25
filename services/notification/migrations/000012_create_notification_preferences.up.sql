-- Each user's choices about how they want to receive notifications.
create table notification_preferences (
    id               uuid primary key,
    user_id          uuid not null,
    scope_type       preference_scope not null,
    category_code    text references notification_categories (code),
    kind_code        text references notification_kinds (code),
    -- null = all channels
    channel_code     text references notification_channels (code),
    preference_level preference_level not null,
    origin           preference_source not null,
    created_at       timestamptz not null,
    updated_at       timestamptz not null,

    -- One preference per user, scope and channel. Nulls count as equal, so
    -- two "all channels" rows for the same scope are rejected.
    unique nulls not distinct (
        user_id, scope_type, category_code, kind_code, channel_code
    ),

    -- Each scope fills in exactly its own column: global uses neither,
    -- category uses only category_code, kind uses only kind_code.
    check (
        (scope_type = 'global' and category_code is null and kind_code is null)
        or (
            scope_type = 'category'
            and category_code is not null
            and kind_code is null
        )
        or (
            scope_type = 'kind'
            and kind_code is not null
            and category_code is null
        )
    )
);

grant select, insert, update, delete
on table notification_preferences
to notification_runtime;
