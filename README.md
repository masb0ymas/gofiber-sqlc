## Fiber x Sqlc

### Migrations
you can install golang-migrate is here:
[Golang Migrate Cli](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

Generate migrations with command:
```sh
migrate create -ext sql -dir ./src/database/migrations -seq create_initial_table
```

Running migrations with command:
```sh
make db-migrate
```