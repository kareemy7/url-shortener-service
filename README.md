# url-shortener-service
 A URL shortener service built with Go, Gin, JWT for authentication, and GORM with MySQL for database management. This application supports user registration, login, URL shortening, and redirection.

## Features

- **User Registration**: Allows users to create an account with email and password.
- **User Login**: Provides a JWT token upon successful login.
- **URL Shortening**: Shortens URLs and associates them with user accounts.
- **URL Retrieval**: Redirects users to the original URL based on the short code.
- **JWT Authentication**: Secure endpoints with JSON Web Tokens (JWT).

## Technologies Used

- **Go**: Programming language used for backend development.
- **Gin**: Web framework for building the API.
- **GORM**: ORM library for interacting with the MySQL database.
- **JWT**: For user authentication and session management.
- **MySQL**: Database for storing user and URL data.

## Installation

### Prerequisites

- Go (1.16 or later)
- MySQL
- Git

### Setup

1.  **Clone the repository:**
    
    ```bash
    `git clone https://github.com/your-username/url-shortener-service.git`
    ```
2.  **Navigate to the project directory:**
    
    ```bash
    `cd url-shortener-service`
    ```
3.  **Install dependencies:**
    
  ```  bash
    `go mod tidy`
```
4.  **Configure the database:**
    
    - Ensure you have MySQL installed and running.
    - Update the `dsn` in `main.go` with your MySQL database credentials.
5.  **Run the application:**
    
    bash
    `go run main.go`
    
    The server will start on `http://localhost:8080`.
    

## API Endpoints

- **POST /register**: Register a new user.
    
    - Request Body: `{ "email": "user@example.com", "password": "password" }`
    - Response: `{ "message": "Registration successful" }`
- **POST /login**: Log in and get a JWT token.
    
    - Request Body: `{ "email": "user@example.com", "password": "password" }`
    - Response: `{ "token": "your_jwt_token" }`
- **POST /shorten**: Shorten a URL (requires JWT token in Authorization header).
    
    - Request Body: `{ "original_url": "http://example.com" }`
    - Response: `{ "short_url": "http://localhost:8080/short_code" }`
- **GET /**
    
    : Redirect to the original URL.
    
    - Example: `http://localhost:8080/abc123`

## Contributing

1.  **Fork the repository**
    
2.  **Create a new branch:**
    
    ```bash
    `git checkout -b feature-branch`
    ```
3.  **Make your changes**
    
4.  **Commit your changes:**
    
 ```   bash
    `git commit -m "Add new feature"`
 ```   
5.  **Push to the branch:**
    
  ```  bash
    `git push origin feature-branch`
 ```
6.  **Create a pull request**
    

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

- **Author:** Karim Elbassiouny
- **Email:** karimelbassiouny@gmail.com
- **GitHub:** [kareemy7](https://github.com/kareemy7)