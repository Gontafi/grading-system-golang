<h1>Project Setup Guide</h1>

1. Clone the repository:
```
git clone https://github.com/Gontafi/grading-system-golang/
cd grading-system-golang
```
2. Build and run the application using Docker:
```
docker-compose up --build
```
3. Do migrations:
```
migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:6543/postgres?sslmode=disable" -verbose up
```
4. Simple post method to auth(sign up, sign in):
```
curl -X POST -H "Content-Type: application/json" -d '{
    "username": "example_username",
    "password": "example_password",
    "role_id": 1,  
    "name": "John",
    "surname": "Doe"
}' http://localhost:8080/auth/sign-up/

curl -X POST -H "Content-Type: application/json" -d '{
    "username": "example_username",
    "password": "example_password"
}' http://localhost:8080/auth/sign-in/
```
Postman collections will be tomorrow...
