# Go Auth Starter
Halo ini adalah starter kit Auhtentication sederharana menggunakan bahasa pemograman Golang. Basic dalam starter ini meliputi : 
- Register 
- Login
- Profile.

### Cara Install
- Clone Repository ini dengan cara
```
git clone https://github.com/andes2912/Go-Auth-Starter
```
- Masuk ke directory
- Lalu jalankan ``` go run main.go ```

### Documentation API
- Register http://127.0.0.1:3000/register
```
{
    "nama": "Andri Desmana",
    "phone": "12345678",
    "email": "andri@mail.com",
    "password": "testing"
}
```

- Login http://127.0.0.1:3000/login
```
{
    "email": "andri@mail.com",
    "password": "testing"
}
```
- Profile http://127.0.0.1:3000/user

### Package yang digunakan
- [GoDotEnv](https://github.com/joho/godotenv)
``` 
go get github.com/joho/godotenv
```

- [jwt-go](https://github.com/golang-jwt/jwt)
```
go get -u github.com/golang-jwt/jwt/v5
```

- [Gin Web Framework](https://gin-gonic.com/)
```
go get -u github.com/gin-gonic/gin
```

- [GORM](https://gorm.io/)
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mqsql
```
- [Go Cryptography](https://pkg.go.dev/golang.org/x/crypto#section-readme)
```
go get -u golang.org/x/crypto/bcrypt
```

- [Daemon](https://github.com/githubnemo/CompileDaemon)
```
go get github.com/githubnemo/CompileDaemon
```
