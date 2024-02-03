# packform-backend
### These application consist of 1 :
   - api
   - cli

### How to run api application:
   - Local Development
       - install golang with version `1.21.1`
       - copy `.env.sample` to `.env`, and fill with your own setting
       - set `DB_HOST=localhost`, to read db host from your local db
       - set `APP_PLATFORM=local`, to include gorm auto migration
       - set `POPULATE_DATA_FROM=cli` (default configuration), used to populate data via cli you can read cli instruction, if you set `POPULATE_DATA_FROM=api` it means directly insert data when api server start which can make error every time you rebuild golang, so set `POPULATE_DATA_FROM=cli` after you run `POPULATE_DATA_FROM=api`, it will skip insert data every server run. populate can only run one once whether api or cli.
       - run `go mod tidy` for install all related library
       - run `go run cmd/app/main.go` or `go build -o bin/app cmd/app/main.go`
       - if `go build` then you should run `./bin/app` for running api
       - set url variable in postman `http://localhost:8080`
       - import postman collection `PackformServer.postman_environment.json` to your own postman

   - Docker Compose
       - set `DB_HOST=postgres`, to read db host from docker service
       - set `APP_PLATFORM=docker`, to skip gorm auto migration, because using `init.sql` to create table
       - set `POPULATE_DATA_FROM=cli` (default configuration), used to populate data via cli you can read cli instruction, if you set `POPULATE_DATA_FROM=api` it means directly insert data when api server start which can make error every time you rebuild golang, so set `POPULATE_DATA_FROM=cli` after you run `POPULATE_DATA_FROM=api`, it will skip insert data every server run. populate can only run one once whether api or cli.
       - run `docker compose up`, then waiting until finish
       - set url variable in postman `http://localhost:8080` or `http://0.0.0.0:8080`
       - import postman collection `PackformServer.postman_environment.json` to your own postman

### How to run cli application:
   - Local Development
       - install golang with version `1.21.1`
       - copy `.env.sample` to `.env`, and fill with your own setting.
       - run `go mod tidy` for install all related library
       - set `DB_HOST=localhost`, to read db host from your local db
       - set `POPULATE_DATA_FROM=cli` (default configuration), used to populate data via cli you can read cli instruction, if you set `POPULATE_DATA_FROM=api` it means directly insert data when api server start which can make error every time you rebuild golang, so set `POPULATE_DATA_FROM=cli` after you run `POPULATE_DATA_FROM=api`, it will skip insert data every server run. populate can only run one once whether api or cli.
       - run `go run cmd/app/main.go` or `go build -o bin/app cmd/app/main.go`
       - run `./bin/app import-csv -d [destination_table] -f [csv_files]` for populate csv file to postgres, please run it sequentially command below:
           - will be inserted to tbl_companies
           ```sh 
           ./bin/app import-csv -d companies -f files/csv/Test\ task\ -\ Postgres\ -\ customer_companies.csv
           ``` 
           - will be inserted to tbl_customers
           ```sh 
           ./bin/app import-csv -d customers -f files/csv/Test\ task\ -\ Postgres\ -\ customers.csv
           ``` 
           - will be inserted to tbl_orders
           ```sh 
           ./bin/app import-csv -d orders -f files/csv/Test\ task\ -\ Postgres\ -\ orders.csv
           ```
           - will be inserted to tbl_order_items
           ```sh
           ./bin/app import-csv -d order_items -f files/csv/Test\ task\ -\ Postgres\ -\ order_items.csv
           ```
           - will be inserted to tbl_order_item_deliveries
           ```sh
           ./bin/app import-csv -d order_item_deliveries -f files/csv/Test\ task\ -\ Postgres\ -\ deliveries.csv
           ```
       - then login to your postgres db, you will be see 5 tables under db name from .env
       - note:
         - set your postgres timezone to `UTC`, you can run `show timezone;` for seeing current timezone, then run `set timezone="UTC";` to set to UTC timezone
         - data will be inserted not in order because import data using go channel and go routine
   
   - Docker Compose
       - set `DB_HOST=postgres`, to read db host from docker service
       - set `APP_PLATFORM=docker`, to skip gorm auto migration, because using `init.sql` to create table
       - set `POPULATE_DATA_FROM=cli` (default configuration), used to populate data via cli you can read cli instruction, if you set `POPULATE_DATA_FROM=api` it means directly insert data when api server start which can make error every time you rebuild golang, so set `POPULATE_DATA_FROM=cli` after you run `POPULATE_DATA_FROM=api`, it will skip insert data every server run. populate can only run one once whether api or cli.
       - run command `docker compose exec -it app bash`, then run the command below in order
           - will be inserted to tbl_companies
           ```sh 
           app import-csv -d companies -f files/csv/Test\ task\ -\ Postgres\ -\ customer_companies.csv
           ``` 
           - will be inserted to tbl_customers
           ```sh 
           app import-csv -d customers -f files/csv/Test\ task\ -\ Postgres\ -\ customers.csv
           ``` 
           - will be inserted to tbl_orders
           ```sh 
           app import-csv -d orders -f files/csv/Test\ task\ -\ Postgres\ -\ orders.csv
           ```
           - will be inserted to tbl_order_items
           ```sh
           app import-csv -d order_items -f files/csv/Test\ task\ -\ Postgres\ -\ order_items.csv
           ```
           - will be inserted to tbl_order_item_deliveries
           ```sh
           app import-csv -d order_item_deliveries -f files/csv/Test\ task\ -\ Postgres\ -\ deliveries.csv
           ```
       