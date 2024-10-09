Make sure, that you create .env file with your params:
```
LOG_LEVEL=
HTTP_HOST=
HTTP_PORT=

PG_DB_SERVER=
PG_DB_PORT=
PG_DB_NAME=
PG_DB_USER=
PG_DB_PASS=
PG_DB_SSLMODE=
```

Swagger generation
```
swagger generate server -f swagger/swagger.yaml -t internal/api --exclude-main
```

Running PostgreSQL database
```
sudo docker run --name=bomond-tennis -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d postgres
```

Running SQL migrations
```
goose "host=localhost port=5436 dbname=bomond user=postgres password=qwerty sslmode=disable" up   
```