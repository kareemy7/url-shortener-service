# url-shortener-service
 A URL shortener service built with Go, Gin, JWT for authentication, and GORM with MySQL for database management. This application supports user registration, login, URL shortening, and redirection.

## Features

- **Shorten a URL**: Convert long URLs into short, easily shareable links.
- **Retrieve Original URL**: Use short codes to retrieve the original URL.
- **User Authentication**: Secure access with email and password, using JWT for authentication.
- **Data Storage**: Store URL mappings and user data in a MySQL database.

## Getting Started

### Prerequisites

- Go (1.18+)
- MySQL
- Git

### Installation

1. **Clone the Repository**

 ```bash
   git clone https://github.com/kareemy7/url-shortener-service.git
   cd url-shortener-service
 ```

2. **Setup Environment**

   Create a `.env` file in the root directory and add the following variables:

   ```plaintext
   DSN=your-mysql-dsn
   SECRET_KEY=your-secret-key
   PORT=your-port
   ```

   Replace `your-mysql-dsn` with your MySQL connection string, `your-secret-key` with a strong secret key for JWT, and `your-port` with the port number you want the server to listen on.

3. **Install Dependencies**

   ```bash
   go mod tidy
   ```

4. **Run the Application**

   ```bash
   go run main.go
   ```

   The server will start and listen on the specified port.

### Usage

- **Register a User**
  
  `POST /register`
  
  Request body:
  ```json
  {
    "Email": "user@example.com",
    "Password": "password123"
  }
  ```

- **Login**

  `POST /login`
  
  Request body:
  ```json
  {
    "Email": "user@example.com",
    "Password": "password123"
  }
  ```

  Response will include a JWT token.

- **Shorten a URL**

  `POST /shorten` (Requires authentication)
  
  Request body:
  ```json
  {
    "OriginalURL": "https://example.com"
  }
  ```

- **Retrieve Original URL**

  `GET /:shortCode`
  
  Example: `GET /abc123` will redirect to the original URL.

- **Validate Authentication**

  `GET /validate` (Requires authentication)
  
  Returns user information if the token is valid.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

If you want to contribute to this project, please fork the repository and submit a pull request. Ensure that your changes are well-documented and tested.

## Acknowledgements

- Go
- Gin
- Gorm
- MySQL
- JWT
    

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

- **Author:** Karim Elbassiouny
- **Email:** karimelbassiouny@gmail.com
- **GitHub:** [kareemy7](https://github.com/kareemy7)