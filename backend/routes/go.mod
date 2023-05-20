module routes

go 1.20

replace models => ../models

replace configuration => ../configuration

replace crud => ../crud

require (
	configuration v0.0.0-00010101000000-000000000000
	crud v0.0.0-00010101000000-000000000000
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	gorm.io/gorm v1.25.1
	models v0.0.0-00010101000000-000000000000
)

require (
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/driver/mysql v1.5.0 // indirect
)
