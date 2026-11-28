select
    id,
    status
from users
where
    id = $1;
