module github.com/cgalimberti/2099-CleanArch/ordersystem

go 1.18

require (
	github.com/99designs/gqlgen v0.14.0
	github.com/cgalimberti/2099-CleanArch/configs v0.0.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/streadway/amqp v0.0.0-20210922163529-1b1c8c8c5c5e
	google.golang.org/grpc v1.39.0
)

replace github.com/cgalimberti/2099-CleanArch/configs => ../configs
