#!/bin/sh

# create song
echo "creating song"
curl -X POST -d '{"album": "one x", "author": "Three Days Grace", "genre": "alternative rock", "title": "Animal I Have Become"}' localhost:8080/song

# get song by id
echo "get song 5"
curl localhost:8080/song/5

# get all songs
echo "get all songs"
curl localhost:8080/song

# delete song
echo "**delete song 2"
curl -X DELETE localhost:8080/song/2

# update song
echo "**update song 10"
curl -X PUT -d '{"album": "One X", "author": "Three Days Grace", "genre": "Alternative Rock", "title": "Animal I Have Become"}' localhost:8080/song/10
