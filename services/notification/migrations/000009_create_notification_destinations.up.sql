-- Where a user can be reached on each channel: email addresses, push devices
-- and chat accounts. Email addresses are synced from the user service.
create table notification_destinations (
    id           uuid primary key,
    user_id      uuid not null,
    channel_code text not null references notification_channels (code),
    -- the system that delivers: 'ses', 'apns', 'fcm', 'slack'
    provider     text not null, -- noqa: RF04
    -- what the provider needs to reach this destination, unique within the
    -- provider: an email address, a device token, or 'T123/U456' for a Slack
    -- user in a workspace
    address      text not null,
    -- hints for building the message, e.g. {"platform": "ios"}
    metadata     jsonb not null,
    -- no longer used: the address bounced, the token died, or it moved to
    -- another user
    revoked_at   timestamptz,
    -- when the client last registered this destination; apps re-register on
    -- every launch, so an old value means the destination is probably dead
    last_seen_at timestamptz not null,
    created_at   timestamptz not null
);

-- An active address belongs to one user. When a token shows up on another
-- account, the old row is revoked and a new one is created, so past
-- deliveries keep pointing at the right user.
create unique index one_active_destination
on notification_destinations (provider, address)
where revoked_at is null;

grant select, insert, update, delete
on table notification_destinations
to notification_runtime;
