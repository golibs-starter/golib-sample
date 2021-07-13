module gitlab.id.vin/vincart/golib-sample

go 1.14

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	gitlab.id.vin/vincart/golib v0.0.0-00010101000000-000000000000
	gitlab.id.vin/vincart/golib-gin v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace gitlab.id.vin/vincart/golib => ../golib

replace gitlab.id.vin/vincart/golib-gin => ../golib-gin
