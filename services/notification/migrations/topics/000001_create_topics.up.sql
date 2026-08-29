create table topics (
  code text primary key,
  status text not null,

  create_time timestamptz not null default now(),
  update_time timestamptz not null default now(),

  constraint topics_code_check check (
    char_length(code) between 1 and 64
    and code ~ '^[a-z][a-z0-9_]*$'
  ),

  constraint topics_status_check (
    statis in ('active', 'disabled')
  )
);

grant select, insert, update, delete
on table topics
to notification_runtime;