#!/bin/sh

docker exec -it music-postgres bash -c 'psql -h localhost -U postgres music'
