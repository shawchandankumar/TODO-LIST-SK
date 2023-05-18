module routes

go 1.20

replace models => ../models

require models v0.0.0-00010101000000-000000000000

require github.com/go-sql-driver/mysql v1.7.1 // indirect
