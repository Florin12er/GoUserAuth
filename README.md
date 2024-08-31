# GoUserAuth

GoUserAuth is a user authentication API built with Go, leveraging Gin framework, PostgreSQL database, and JWT for secure authentication.

## Features

- User registration
- User login
- JWT-based authentication
- Secure password hashing
- PostgreSQL database integration

## Technologies Used

- Go
- Gin Web Framework
- PostgreSQL
- JWT (JSON Web Tokens)
- bcrypt for password hashing

## Prerequisites

- Go (version 1.x or later)
- PostgreSQL
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Florin12er/GoUserAuth.git
```

2. Navigate to the project directory:
```bash
cd GoUserAuth
```

3. Install dependencies:
```
go mod tidy
```

4. Set up your PostgreSQL database and update the connection string in the code.

5. Run the application:

go run main.go
text

## API Endpoints

- `POST /api/register`: Register a new user
- `POST /api/login`: Login and receive JWT token
- `GET /api/home`: Access protected route (requires JWT)

## Usage

1. Register a new user:

POST /api/register
{
"name": "John Doe",
"email": "john@example.com",
"password": "securepassword"
}


2. Login to receive JWT token:

POST /api/login
{
"email": "john@example.com",
"password": "securepassword"
}
text

3. Use the received token in the Authorization header for protected routes:

GET /api/home
Authorization: Bearer <your_jwt_token>


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).

This README provides a comprehensive overview of your project, including its features, technologies used, setup instructions, and basic usage guide. You can further customize it by adding more specific details about your implementation or any unique features of your authentication system.
