# library-api
basic authentication system in Go<br>

# Setup
Install dependencies:<br>
1. [Go](https://go.dev/doc/install)
2. [MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)

Then setup MySQL and start the server using `sudo systemctl start mysql`<br>

Clone the repo:
```bash
git clone https://github.com/Dami-18/library-api.git
cd library-api
```

Copy .env template:
```bash
cp .env.template .env
```
Add your mysql root password in the generated .env file

Setting up database:
```bash
chmod +x init_db.sh
bash init_db.sh
```

Installing go dependencies:
```go
go mod tidy
```

Build and run executable:
```go
go build -o main
./main
```