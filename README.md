# URL Shortener 📎

## Project Purpose 🎯

The URL Shortener project aims to provide a simple service that converts long URLs into shorter, more manageable links. It is designed to help users share links easily and track click statistics efficiently.

## Technologies Used 🛠️

- **Go**: The core programming language used for building the application.
- **sqlc**: SQL compiler that generates type-safe Go code from SQL queries.
- **ozzo-validation**: A Go validation library used for data validation.
- **Google UUID**: For generating unique identifiers.

## How to Use 📚

1. **Clone the Repository**:  
   ```bash
   git clone <repository-url>
   cd urlshortener
   ```

2. **Configuration**:  
   Edit the `config.json` file to setup your database configurations and application settings.

3. **Install Dependencies**:  
   Make sure you have Go installed, then run:
   ```bash
   go mod tidy
   ```

4. **Run the Application**:  
   ```bash
   go run main.go
   ```

## Routes and Handlers 🌐

- **POST /shorten**: Create a new shortened URL.
- **GET /:url_id**: Redirect to the original URL using its shortened version.
- **GET /urls**: Retrieve all shortened URLs.
- **DELETE /:url_id**: Delete a shortened URL by its ID.
- **GET /info/:url_id**: Get information about a specific shortened URL.

## Project Structure 🗂️

- **`/internal/db`**: Contains database query logic and models.
- **`/internal/api/handlers`**: Defines HTTP handlers for the URL shortener endpoints.
- **`/internal/api/utils`**: Utility functions for handling requests and responses.
- **`/internal/config`**: Configuration loading and logger setup.
- **`/internal/app`**: Application setup and initialization.
- **`/pkg/models`**: Common data models used across the application.
- **`main.go`**: The entry point of the application.

Feel free to contribute and make this project even better!