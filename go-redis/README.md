# go-redis

Working with Redis Databases in Go


```shell
$ redis-cli
127.0.0.1:6379> ping
PONG
127.0.0.1:6379> keys *
(empty array)
127.0.0.1:6379> set name lesenelir
OK
127.0.0.1:6379> get name
"lesenelir"
127.0.0.1:6379> hset id name lesenelir
(integer) 1
127.0.0.1:6379> hget id name
"lesenelir"
127.0.0.1:6379> hkeys id
1) "name"
127.0.0.1:6379> set ai-space:user:1 "{id: 1, name: lesenelir}"
OK
127.0.0.1:6379> keys *
1) "ai-space:user:1"
2) "name"
127.0.0.1:6379> hset ai-space:user:2 2 "{id: 2, name: lesenelir}"
(integer) 1
127.0.0.1:6379> keys *
1) "ai-space:user:1"
2) "name"
3) "ai-space:user:2"
127.0.0.1:6379> hset ai-space:user:3 id 3 name lesenelir age 18
(integer) 3
127.0.0.1:6379> lpush users 1 2 3
(integer) 3
127.0.0.1:6379> rpush users 4 5 6
(integer) 6
127.0.0.1:6379> lpop users
"3"
127.0.0.1:6379> rpop users
"6"
127.0.0.1:6379> lrange users 1 3
1) "1"
2) "4"
3) "5"
127.0.0.1:6379> sadd s1 a b c d
(integer) 4
127.0.0.1:6379> smembers s1
1) "a"
2) "b"
3) "c"
4) "d"
127.0.0.1:6379> zadd sorted-set-key1 50 user1 70 user2 30 user3
(integer) 3
127.0.0.1:6379> zrank sorted-set-key1 user1
(integer) 1
127.0.0.1:6379> zrange sorted-set-key1 0 -1
1) "user3"
2) "user1"
3) "user2"
127.0.0.1:6379> zrevrange sorted-set-key1 0 -1
1) "user2"
2) "user1"
3) "user3"
127.0.0.1:6379> zscore sorted-set-key1 user1
"50"
127.0.0.1:6379> exit
```
