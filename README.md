Add a .env file cmd/main path
structure of env file:
DB_USER=mysql_username
DB_PASSWORD=mysql_password
DB_NAME=sql_database_name

Steps to run:
Open terminal
go to project folder path where all the files are present. It should be like .\mysql-db
run following commands
go get .
go mod tidy
cd cmd/main
go run .


