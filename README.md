# RESTAURANT API
Restaurant API created with GO

## Pre-requisites

1. Understanding GO Fundamentals <br>
[Tutorials Playlist](https://youtube.com/playlist?list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw)

2. Learn Gorilla/Mux and GORM <br>
[Gorilla/Mux Documentation](https://pkg.go.dev/github.com/gorilla/mux) <br>
[GORM Documentation](https://gorm.io/docs/)

3. Familiarize with MySQL <br>
[Documentation](https://dev.mysql.com/doc/)

4. Understanding REST API Concepts <br>
[Explanation](https://medium.com/jagoanhosting/perbedaan-antara-api-rest-api-dan-restful-api-6a66d655a6c2) <br>
[RESTful APIs in 100 Seconds](https://youtu.be/-MTSQjw5DrM)

5. Familiarize with Postman Documentation <br>
[Postman Documentation Tutorial](https://www.softwaretestinghelp.com/postman-api-documentation/)

## Setup

### 1. Set the `.env` file with credentials needed.
```sh
# Port of the Server
PORT=<PORT>

# Database Information
DB_URL=<DB_URL>

# JWT Information
JWT_KEY=<JWT_KEY>
```

### 2. Create MySQL Database
#### Local
```sh
# Login to MySQL User
# You may use root as the user
mysql -u <MySQL User> -p

# Enter password on given prompt
Password: <MySQL Password>

# Create Database
CREATE DATABSE <MySQL Database Name>; 
```

#### Cloud
[Create Free MySQL Database Online](https://youtu.be/TMGHOW8Hzvw)

### 3. Run Modules Installation
```sh
# Install all modules used in the project
go get ./...
```

### 4. Run API
#### Run (without build)
```sh
go run .\main.go
```

#### Build (convert to .exe format)
```sh
go build .\main.go

.\main.exe
```

## Contributor

- Vincent Pradipta (xhfmvls)