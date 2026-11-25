-- Who caused an event: a user, or the system itself.
create type actor_type as enum (
    'user',
    'system'
);

-- How much of something a user wants to receive. 'direct' means only
-- notifications whose reason is direct (see notification_reasons.is_direct).
create type preference_level as enum (
    'all',
    'direct',
    'none'
);

-- What a preference applies to: everything, a category, or a single kind.
create type preference_scope as enum (
    'global',
    'category',
    'kind'
);

-- Where a preference came from.
create type preference_source as enum (
    'settings',
    'unsubscribe_link',
    'system'
);

-- Lifecycle of a digest: collecting, being sent, done, or failed.
create type digest_status as enum (
    'open',
    'sending',
    'sent',
    'failed'
);

-- Lifecycle of a single delivery on a channel.
create type delivery_status as enum (
    'pending',
    'digested',
    'sent',
    'failed',
    'suppressed'
);
