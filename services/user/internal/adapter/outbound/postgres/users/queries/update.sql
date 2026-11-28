update users
set
    status = $2,
    update_time = now()
where
    id = $1;
