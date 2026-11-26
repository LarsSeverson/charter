-- Each send of a notification to one destination, and its result. Retries
-- update the same row. A user with no destination on a channel gets no row
-- for that channel.
create table notification_deliveries (
    id                  uuid primary key,
    notification_id     uuid not null references notifications (id),
    -- where it goes; the channel comes from here
    destination_id      uuid not null references notification_destinations (id),
    -- the digest that sends it; once digested, the digest owns the outcome
    digest_id           uuid references notification_digests (id),
    status              delivery_status not null,
    decided_by          delivery_decision not null,
    -- the preference that decided, if any; cleared if the preference is
    -- later deleted
    preference_id       uuid references notification_preferences (id)
    on delete set null,
    attempts            int not null,
    next_attempt_at     timestamptz,
    last_error          text,
    provider_message_id text,
    sent_at             timestamptz,
    created_at          timestamptz not null,
    updated_at          timestamptz not null,

    -- One delivery per notification and destination.
    unique (notification_id, destination_id),

    -- A digested delivery belongs to a digest, and only digested ones do. It
    -- stays 'digested' after the digest is sent; the digest holds the result.
    check ((status = 'digested') = (digest_id is not null)),

    -- A sent delivery records when it was sent, and only sent ones do.
    check ((status = 'sent') = (sent_at is not null))
);

grant select, insert, update, delete
on table notification_deliveries
to notification_runtime;
