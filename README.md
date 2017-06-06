# Docker

Some experiments with Docker

Redis:

    In console A)

        $ sudo docker run --name redis mramshaw4docs/redis
        
            OR (to run in background)
            
        $ sudo docker run --name redis -d mramshaw4docs/redis

    In console B)

        $ sudo docker run --link redis:rdb --name rediswebserver mramshaw4docs/rediswebserver
        2017/06/06 03:44:36 redis web server is now serving requests...

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

    Can also test with:
    
        $ redis-cli - h {IPAddress}
