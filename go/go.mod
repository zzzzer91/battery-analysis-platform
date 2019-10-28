module battery-analysis-platform

go 1.13

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/gin-contrib/sessions v0.0.0-20190814041826-3f766bdeb0d1
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/gocelery/gocelery v0.0.0-20191008025616-26dc580ba02e
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gorilla/websocket v1.4.0
	github.com/jinzhu/gorm v1.9.10
	github.com/kr/pretty v0.1.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shirou/gopsutil v2.18.12+incompatible
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.1.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/satori/go.uuid v1.2.0 => github.com/gofrs/uuid v3.2.0+incompatible
