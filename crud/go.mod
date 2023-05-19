module crud

go 1.20

replace configuration => ../configuration

replace models => ../models

require (
	configuration v0.0.0-00010101000000-000000000000
	models v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/driver/mysql v1.5.0 // indirect
	gorm.io/gorm v1.25.1 // indirect
)
