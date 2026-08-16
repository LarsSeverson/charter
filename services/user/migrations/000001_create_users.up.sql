create table users (
    id uuid primary key,
    status text not null check (
        status in ('pending', 'active', 'suspended', 'closed')
    ),
    create_time timestamptz not null default now(),
    update_time timestamptz not null default now(),
    delete_time timestamptz default null,
    expire_time timestamptz default null,
    purge_time timestamptz default null
);

grant select, insert, update, delete
on table users
to user_runtime;
