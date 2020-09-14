# 5pmGebetServer

## startup
Set the following environment variables:
- PORT: port, the server is run on
- MASTERKEY: api key for accessing the admin interface at ´/backend.html?key=MASTERKEY´, the MASTERKEY currntly must be an integer

after frist start of the server the `pray.db` file is generated. this is a sqlite3 file.
the provided master key will be inserted into the database.
this can be used to insert a Type-0-Key (Master Key) into the server

## reset
to reset the database and server just delete the `pray.db` and redo the startup procedure and insert new master key


## API's
* [/api/](doc/api_.md)
* [/api/key](doc/api_key.md)
* [/api/view](doc/api_view.md)

## Deployment
### using docker
**BUILD**
```shell script
docker build -t 5pm 
```

**RUN**
```shell script
docker run -d --name 5pm --env PORT=8080 --env MASTERKEY=12345 -p 8080:8080 5pm
```
