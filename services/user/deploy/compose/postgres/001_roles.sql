\set ON_ERROR_STOP on

set password_encryption = 'scram-sha-256';

create role user_runtime with
nologin
nosuperuser
nocreatedb
nocreaterole
noreplication
nobypassrls;

create role user_app with
login
password 'user_app_development'
nosuperuser
nocreatedb
nocreaterole
noreplication
nobypassrls
in role user_runtime;

create role user_migrator with
login
password 'user_migrator_development'
nosuperuser
nocreatedb
nocreaterole
noreplication
nobypassrls;

revoke all privileges on database users from public;

grant connect on database users to user_runtime;
grant connect on database users to user_migrator;

revoke all privileges on schema public from public;

grant usage, create on schema public to user_migrator;
grant usage on schema public to user_runtime;

alter role user_app set search_path = public;
alter role user_migrator set search_path = public;

alter default privileges
for role user_migrator
in schema public
revoke execute on functions from public;

comment on role user_runtime is
'Non-login role containing user-service runtime database permissions';

comment on role user_app is
'Development login used by the running user service';

comment on role user_migrator is
'Development login used exclusively for user-service migrations';
