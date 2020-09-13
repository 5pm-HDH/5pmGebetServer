# 5pmGebetServer

## startup
after frist start of the server the `pray.db` file is generated. this is a sqlite3 file.

## insert master key
to get started you have to insert a master key into the database. there a 2 options, using a gui like SQLiteBrowser <https://sqlitebrowser.org/> or via sqlite command line.
```sql
INSERT INTO authorization (auth_key,type) values("[super_secret_key]", 0)
```
this can be used to insert a Type-0-Key (Master Key) into the server

## reset
to reset the database and server just delete the `pray.db` and redo the startup stepps and insert new master key


## API's
* [/api/](doc/api_.md)
* [/api/key](doc/api_key.md)
* [/api/beamer](doc/api_beamer.md)

## Deployment
### using docker
**BUILD**
```shell script
docker build --tag 5pm --build-arg MASTER_KEY='<KEY>' --build-arg PORT=8080 .
```

**RUN**
```shell script
docker run --publish 8080:8080 --detach --name 5pm 5pm
```
