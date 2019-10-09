# redcoins

## Description

API to register users and their bitcoins transactions (buy/sell)

## Technologies

* Go Lang
* MySQL
* Docker

## How to deploy

### 1. Get project

Clone or download the project

`git clone https://github.com/giovanni-rc/redcoins.git`

### 2. Creat .ENV

Need to create a .env file on root folder with the following format:

```
DB_NAME=redcoinsdb
DB_USER=redcoins
DB_PASS=redcoins
DB_HOST=db
DB_PORT=3306
DB_ROOT=p@ssw0rd!
PORT=8080
TOKEN_PASSWORD=JwtredcoinsTokenEncrypt
TOKEN_COINMARKET=5sdfsdcc2-fdf14-Ygjnvdf-80ae-02rj5651231260c
```

Where:

Variable | Value
---------|---------
DB_NAME  | Database name to be created
DB_USER  | Database user to be created
DB_PASS  | Database password for user <DB_USER>
DB_HOST  | Name of database host (use 'db' case docker-compose be used)
DB_PORT  | Database port
DB_ROOT  | Password for root user on database
PORT     | Application port (use '8080' case Dockerfile be used)
TOKEN    | Token for JWT use
TOKEN_COINMARKET | Token created on CoinMarket site to use API

### 3. Build Containers

Execute the following command to build app and database containers

`docker-compose up --build`

## How to use

All requests need to add on Header the parameter Auth, the value will be the token provide for the application when create a user or execute a login

### 1. Create User

Endpoint: `http://<SERVER_IP>:8080/redcoins/api/user/new`

Method: POST

Resquest Example:
```
{
	"email" : "joao@email.com",
	"password" : "joao123",
	"name" : "Joao Silva",
	"birthday" : "1995-09-04"
}
```

### 2. Login

Endpoint: `http://<SERVER_IP>:8080/redcoins/api/login`

Method: POST

Resquest Example:
```
{
	"email" : "joao@email.com",
	"password" : "joao123"
}
```

### 3. Reset Password

Endpoint: `http://<SERVER_IP>:8080/redcoins/api/user/reset_password`

Method: POST

Resquest Example:
```
{
	"email" : "joao@email.com",
	"password" : "joao12345"
}
```

### 4. Create Operation

Endpoint: `http://<SERVER_IP>:8080/redcoins/api/operation/new`

Method: POST

Resquest Example:

```
{
	"qty" : 2,
	"date" : "2018-09-23T17:49:31Z",
	"user_id" : 1,
	"type" : 1
}
```

On type use 0 for buy and 1 for sell

### 5. Operations by Date

Endpoint: `http://<SERVER_IP>:8080/redcoins/api/operation/get_by_date/<date>`

Method: GET

Resquest Example:

```
http://<SERVER_IP>:8080/redcoins/api/operation/get_by_date/2018-09-23
```

### 6. Operations by User


Endpoint: `http://<SERVER_IP>:8080/redcoins/api/operation/get_by_user/<user_id>`

Method: GET

Resquest Example:

```
http://<SERVER_IP>:8080/redcoins/api/operation/get_by_user/2

```

