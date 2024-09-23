Swagger generation
```
swagger generate server -f swagger/swagger.yaml -t internal/api --exclude-main
```

Running database
```
sudo docker run --name=bomond-tennis -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d postgres
```

Running migrations
```
goose "host=localhost port=5436 dbname=bomond user=postgres password=qwerty sslmode=disable" up   
```