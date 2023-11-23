Biblioteca para migrations:

sudo curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | sudo apt-key add -
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
apt-get update
apt-get install -y migrate

criar migrations:
migrate create -ext=sql -dir=sql/migrations -seq init

executar migrations:
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

rollback migrations:
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

Uma biblioteca muito utilizada para utilização de SQL puro é o SQLX, mas não é ele que vamos mexer.

vamos usar SQLC:
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest