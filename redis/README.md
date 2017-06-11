#Redis

    In console A)

        $ sudo docker run --name redis --rm mramshaw4docs/redis
        
            OR (to run in background)
            
        $ sudo docker run --name redis -d --rm mramshaw4docs/redis

    In console B)

        $ sudo docker run --link redis:rdb --name rediswebserver --rm mramshaw4docs/rediswebserver
        2017/06/06 03:44:36 redis web server is now serving requests...

            OR (to run in background)
        
        $ sudo docker run --link redis:rdb --name rediswebserver -d --rm mramshaw4docs/rediswebserver

    In console C)

        $ sudo docker inspect rediswebserver | grep -i ipaddress

        [Note the IPAddress for use below.]
        
    Test redis 'ping" with:
    
        http://{IPAddress}:5000/
        
        Result: Hello, response to ping is "PONG"

    Test redis with:
    
        http://{IPAddress}:5000/test
        
        Result: key1 = "value1"
                key2 does not exist

    Can also test with redis client:

        $ sudo docker inspect redis | grep -i ipaddress
        
        [Note the IPAddress for use below.]
    
        $ redis-cli -h {IPAddress}
        172.17.0.2:6379> get key1
        "value1"
        172.17.0.2:6379> get key2
        (nil)
        172.17.0.2:6379> ttl key1
        (integer) -1
        172.17.0.2:6379> info
        # Server
        redis_version:3.2.9
        < ... >
        >exit
        $

    Final cleanup:
    
        $ sudo docker stop redis rediswebserver
        redis
        rediswebserver
        $
        
