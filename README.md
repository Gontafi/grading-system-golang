<h2>Project Idea</h2>

The project involves creating a rating service for a school where students receive daily grades for homework and attendance in lessons. The service utilizes Go 1.21, Fiber, PGX, Golang JWT v5, Viper, Cobra, Redis, and runs in containers.

Key features include:
1. JWT-based authorization with roles (administrator, teacher, student).
2. Tracking student homework and attendance, with a maximum score of 6 (5 for homework, 1 for attendance) per lesson.
3. Periodic rating calculations for top-performing students, with the ability to filter ratings by lesson and time frame (week, month, year).
4. Utilizes Redis for caching and PostgreSQL for data storage.

<h2>Project Setup Guide</h2>

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
