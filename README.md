# packform-backend
These application consist of 2 :
   - cmd/api/main.go -> it is used for api application
   - cmd/cli/main.go -> it is used for cli application

How to run api application:
   - Local Development
       - install golang with version `1.21.1`
       - copy `.env.sample` to `.env`, and fill with your own setting.
       - run `go mod tidy` for install all related library
       - run `go run cmd/api/main.go` or `go build -o bin/api cmd/api/main.go` for building binary, by default will run on port `8080`
       - if you run `go build` then you should run `./bin/api`
       - import postman collection `PackformServer.postman_environment.json` to your own postman
   - Docker Compose

How to run cli application:
   - Local Development
       - install golang with version `1.21.1`
       - copy `.env.sample` to `.env`, and fill with your own setting.
       - run `go mod tidy` for install all related library
       - run `go run cmd/cli/main.go` or `go build -o bin/cli cmd/cli/main.go` for building binary, by default will run on port `8080`
       - run `./bin/cli import -d [destination_table] -f [csv_files]` for populate csv file to postgres, please run it sequentially command below:
           - `./bin/cli import -d companies -f files/csv/Test\ task\ -\ Postgres\ -\ customer_companies.csv` -> will be insert to `tbl_companies`
           - `./bin/cli import -d customers -f files/csv/Test\ task\ -\ Postgres\ -\ customers.csv` -> will be insert to `tbl_customers`
           - `./bin/cli import -d orders -f files/csv/Test\ task\ -\ Postgres\ -\ orders.csv` -> will be insert to `tbl_orders`
           - `./bin/cli import -d order_items -f files/csv/Test\ task\ -\ Postgres\ -\ order_items.csv` -> will be insert to `tbl_order_items`
           - `./bin/cli import -d order_item_deliveries -f files/csv/Test\ task\ -\ Postgres\ -\ deliveries.csv` -> will be insert to `tbl_order_item_deliveries`
       - then login to your postgres db, you will be see 5 tables under db name from .env
       - note: please set your postgres timezone to `UTC`, you can run `show timezone;` first for see current timezone, then run `set timezone="UTC";` to set to UTC timezone