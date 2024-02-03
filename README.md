# packform-backend
These application consist of 1 :
   - api
   - cli

How to run api application:
   - Local Development
       - install golang with version `1.21.1`
       - copy `.env.sample` to `.env`, and fill with your own setting
       - run `go mod tidy` for install all related library
       - run `go run cmd/app/main.go` or `go build -o bin/app cmd/app/main.go` for building binary, by default will run on port `8080`
       - if you run `go build` then you should run `./bin/app`
       - import postman collection `PackformServer.postman_environment.json` to your own postman
   - Docker Compose
       - change `DB_HOST` value to `postgres` which taken from docker compose service
       - change `APP_PLATFORM` value to `docker`
       - run `docker compose up`, then waiting until finish
       - first run cli application for populate data
       - import postman collection `PackformServer.postman_environment.json` to your own postman

How to run cli application:
   - Local Development
       - install golang with version `1.21.1`
       - copy `.env.sample` to `.env`, and fill with your own setting.
       - run ```sh go mod tidy``` for install all related library
       - run ```sh go run cmd/app/main.go``` or ```sh go build -o bin/app cmd/app/main.go``` for building binary, by default will run on port `8080`
       - run ```sh ./bin/app import-csv -d [destination_table] -f [csv_files]``` for populate csv file to postgres, please run it sequentially command below:
           - ```sh ./bin/app import-csv -d companies -f files/csv/Test\ task\ -\ Postgres\ -\ customer_companies.csv``` -> will be inserted to `tbl_companies`
           - ```sh ./bin/app import-csv -d customers -f files/csv/Test\ task\ -\ Postgres\ -\ customers.csv``` -> will be inserted to `tbl_customers`
           - ```sh ./bin/app import-csv -d orders -f files/csv/Test\ task\ -\ Postgres\ -\ orders.csv``` -> will be inserted to `tbl_orders`
           - ```sh ./bin/app import-csv -d order_items -f files/csv/Test\ task\ -\ Postgres\ -\ order_items.csv``` -> will be inserted to `tbl_order_items`
           - ```sh ./bin/app import-csv -d order_item_deliveries -f files/csv/Test\ task\ -\ Postgres\ -\ deliveries.csv``` -> will be inserted to `tbl_order_item_deliveries`
       - then login to your postgres db, you will be see 5 tables under db name from .env
       - note:
         - set your postgres timezone to `UTC`, you can run `show timezone;` for seeing current timezone, then run `set timezone="UTC";` to set to UTC timezone
         - data will be inserted not in order because import data using go channel and go routine
   
   - Docker Compose
       - change `APP_PLATFORM` value to `docker`
       - run command ```sh docker compose exec -it app bash```, then run the command below in order
             
           - ```sh app import-csv -d companies -f files/csv/Test\ task\ -\ Postgres\ -\ customer_companies.csv``` -> will be insert to `tbl_companies`
           - ```sh app import-csv -d customers -f files/csv/Test\ task\ -\ Postgres\ -\ customers.csv``` -> will be insert to `tbl_customers`
           - ```sh app import-csv -d orders -f files/csv/Test\ task\ -\ Postgres\ -\ orders.csv``` -> will be insert to `tbl_orders`
           - ```sh app import-csv -d order_items -f files/csv/Test\ task\ -\ Postgres\ -\ order_items.csv``` -> will be insert to `tbl_order_items`
           - ```sh app import-csv -d order_item_deliveries -f files/csv/Test\ task\ -\ Postgres\ -\ deliveries.csv``` -> will be insert to `tbl_order_item_deliveries`
       