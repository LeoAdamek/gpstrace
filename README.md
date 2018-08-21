GPSTrace
========

A (very) simple app for devices to post their position to and users to view a history of that location on a map.


Running
-------

### Configuration (Toml)

````toml
[postgres]
url = "postgresql://user:password@host/db?sslmode=disable"

[http]
listen = ":8080"
````

### Run

    ./gpstrace
    

Adding Data
-----------

Note: All headers are optional, the server will only accept `application/x-www-form-urlencoded` data.

```http request
POST localhost:8080/data
Content-Type: application/x-www-form-urlencoded

ent=Some Device Name&lat=N 1 2 3.45&lng=W 1 2 3.45
```


Reading Data
------------

```http request
GET localhost:8080/data
```

