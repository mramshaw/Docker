package main

import (
    "fmt"
    "log"
    "net/http"

    "gopkg.in/redis.v3"
)

var client *redis.Client // Singleton

func init() {

    client = redis.NewClient(&redis.Options{
        Addr: "rdb:6379",
        Password: "", // no password set
        DB: 0,        // use default DB
    })
}

func main() {

    log.Println("redis web server is now serving requests...")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        res, err := client.Ping().Result()
        if err != nil {
            panic(err)
        }
        fmt.Fprintf(w, "Hello, response to ping is %q\n", res)
    })

    http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

	err := client.Set("key1", "value1", 0).Err() // no expiration
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "key1 = %q\n", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Fprintf(w, "key2 does not exist\n")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Fprintf(w, "key2 = %q\n", val2)
	}
        // Output: key1 value1
        // key2 does not exist
    })

    log.Fatal(http.ListenAndServe(":5000", nil))
}
