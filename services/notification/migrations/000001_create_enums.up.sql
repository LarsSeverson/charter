-- Who caused an event: a user, or the system itself.
create type actor_type as enum (
    'user',
    'system'
);

-- How much of something a user wants to receive. 'direct' means only
-- notifications where the recipient was targeted personally (see
-- notifications.direct).
create type preference_level as enum (
    'all',
    'direct',
    'none'
);

-- Where a preference came from.
create type preference_source as enum (
    'settings',
    'unsubscribe_link',
    'system'
);

-- Lifecycle of a digest: collecting, being sent, done, failed, or closed with
-- nothing to send because everything in it was already read.
create type digest_status as enum (
    'open',
    'sending',
    'sent',
    'failed',
    'empty'
);

-- Lifecycle of a single delivery on a channel.
create type delivery_status as enum (
    'pending',
    'digested',
    'sent',
    'failed',
    'suppressed'
);

-- What decided whether a delivery is sent: the kind being mandatory, the
-- kind's default, a user preference, the channel being turned off
-- platform-wide, or the user having read it in the app before it was sent.
create type delivery_decision as enum (
    'mandatory',
    'default',
    'preference',
    'channel_disabled',
    'already_read'
);
