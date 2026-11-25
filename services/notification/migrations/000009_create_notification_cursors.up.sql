-- When each user last opened their inbox. Anything newer is unseen.
create table notification_cursors (
    user_id      uuid primary key,
    last_seen_at timestamptz not null
);

grant select, insert, update, delete
on table notification_cursors
to notification_runtime;
