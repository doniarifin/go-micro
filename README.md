# go-micro

## Requirements

- Go 1.23+
- MySQL
- Caddy
- RabbitMQ

## Development Requirements

- [golang-migrate](https://github.com/golang-migrate/migrate) (for database migrations)

## Quick Start

### Clone this repo

```
git clone https://github.com/doniarifin/go-micro.git
cd go-micro
```

### Copy and rename config.yaml

Setup your DB in config.yaml

```
cp ./data/config.yaml.copy ./data/config.yml
```

### Install depedencies

```
go mod tidy
```

### Golang migrate

you can manually run command for golang migration. you can see the [docs here](https://github.com/golang-migrate/migrate).

```
migrate -database ${DB_URL} -path db/migrations up
```

#### example:

```
migrate -database 'mysql://user:pass@tcp(localhost:3306)/golang' -path db/migrations up
```

### Run Caddy

Ensure you have installed Caddy

go to Terminal tab > Run Task > then select `Run Caddy`

or open terminal run this command

```
caddy run --config=data/Caddyfile
```

this caddy running as reserve proxy so your app will run in port that you set in `data/Caddyfile`

### Start service

If you use vscode just press `F5`

or go to Debug Tab > then select Run All Services

or run Makefile command

```
make all
```
