create table channels (
  code text primary key,
  status text not null,

  create_time timestamptz not null default now(),
  update_time timestamptz not null default now(),

  constraint channels_code_check check (
    char_length(code) between 1 and 64
    and code ~ '^[a-z][a-z0-9_]*$'
  ),

  constraint channels_status_check (
    statis in ('active', 'disabled')
  )
);

insert into channels (code, status)
values
  ('email', 'active'),
  ('web_push', 'active');

grant select, insert, update, delete
on table channels
to notification_runtime;