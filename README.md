# 5pmGebetServer

## startup
Set the following environment variables:
- PORT: port, the server is run on
- MASTERKEY: api key for accessing the admin interface at ´/backend.html?key=MASTERKEY´, the MASTERKEY currntly must be an integer

after frist start of the server the `pray.db` file is generated. this is a sqlite3 file.
the provided master key will be inserted into the database.
this can be used to insert a Type-0-Key (Master Key) into the server

## reset
to reset the database and server just delete the ```pray.db``` and redo the startup procedure and insert new master key


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

### using docker-compose
docker-compose can manage docker deployments and allows the use of a configuration file instead of command line parameters.

first, build the docker image as described above.
then edit the ```docker-compose.yml``` and change the image tag to the new one.

afterwards, run
```shell script
MASTERKEY=12345 docker-compose up -d
```

this will create the docker process and monitor it.
if it stops, it gets restarted.
it also mounts ```/var/5pm``` from the host to ```/data``` inside the container, to ensure the database survices a restart/update.

to update the application, build the new image using a different tag, change the image tag in the ```docker-compose.yml``` and then run
```shell script
docker-compose up -d
```
again. this will stop the old container and start a new one with the updated image.
