module gitlab.id.vin/vincart/golib-sample-internal

go 1.14

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.3.3
	github.com/swaggo/swag v1.7.4
	gitlab.id.vin/vincart/golib v0.9.4
	gitlab.id.vin/vincart/golib-data v0.7.0
	gitlab.id.vin/vincart/golib-gin v0.5.0
	gitlab.id.vin/vincart/golib-message-bus v0.1.0
	gitlab.id.vin/vincart/golib-migrate v0.0.1
	gitlab.id.vin/vincart/golib-sample-adapter v0.0.0-00010101000000-000000000000
	gitlab.id.vin/vincart/golib-sample-core v0.0.0-00010101000000-000000000000
	gitlab.id.vin/vincart/golib-security v0.8.4
	gitlab.id.vin/vincart/golib-test v0.2.2
	go.uber.org/fx v1.13.1
	gorm.io/gorm v1.22.3
)

replace (
	gitlab.id.vin/vincart/golib-sample-adapter => ./../adapter
	gitlab.id.vin/vincart/golib-sample-core => ./../core
)
