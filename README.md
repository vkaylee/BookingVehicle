## DEMO
### Install mongodb via docker
```
docker run -d -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=123456 -p 27017:27017 mongo
```
### Run app
```
go get -v -d ./...
```
```
go run main.go
```
### Functions
#### Sign Up (Post method)
- End point
```
http://localhost:1234/auth/signUp
```
- Contents
```
{
	"firstName":"LEE",
	"lastName":"Tuan",
	"email":"tuanlm1989@gmail.com",
	"password":"123456"
}
```
#### Sign In (Post method)
- End point
```
http://localhost:1234/auth/signIn
```
- Contents
```
{
	"email":"tuanlm1989@gmail.com",
	"password":"123456"
}
```
