#!/bin/bash

echo "Sending some data ..."

echo '127.0.0.1 - hostname [1371731251] "GET /1234/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731252] "GET /1234/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731253] "GET /1234/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731254] "GET /1234/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731255] "GET /1234/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731261] "GET /1234/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731262] "GET /1234/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731263] "GET /1234/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731264] "GET /1234/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731265] "GET /1234/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731266] "GET /1234/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731261] "GET /1235/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731262] "GET /1235/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731263] "GET /1235/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731264] "GET /1235/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731265] "GET /1235/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
echo '127.0.0.1 - hostname [1371731266] "GET /1235/foo/bar HTTP/1.1" 200 2326' | nc -4 -u -w 1 -c localhost 4263
