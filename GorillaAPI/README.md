
---

# Movie API

A simple RESTful API for managing a collection of movies. This API allows users to create, read, update, and delete movie records.

## Technologies Used

- Go (Golang)
- Gorilla Mux (for routing)
- JSON (for data interchange)

## Endpoints

### Get All Movies

- **URL**: `/movies`
- **Method**: `GET`
- **Response**: A JSON array of all movies.

### Get a Movie

- **URL**: `/movies/{id}`
- **Method**: `GET`
- **URL Params**: 
  - `id` - The ID of the movie you want to retrieve.
- **Response**: A JSON object representing the movie, or an error if not found.

### Create a Movie

- **URL**: `/movies`
- **Method**: `POST`
- **Request Body**: A JSON object representing the new movie.
  ```json
  {
    "title": "Movie Title",
    "isbn": "123456",
    "director": {
      "firstname": "First",
      "lastname": "Last"
    }
  }
  ```
- **Response**: The created movie object in JSON format.

### Update a Movie

- **URL**: `/movies/{id}`
- **Method**: `PUT`
- **URL Params**:
  - `id` - The ID of the movie to update.
- **Request Body**: A JSON object with updated movie data.
- **Response**: The updated movie object in JSON format.

### Delete a Movie

- **URL**: `/movies/{id}`
- **Method**: `DELETE`
- **URL Params**:
  - `id` - The ID of the movie to delete.
- **Response**: A JSON array of the remaining movies after deletion.

## Getting Started

1. **Clone the repository**:
   ```bash
   git clone <Portfolio>
   ```

2. **Navigate to the project directory**:
   ```bash
   cd <Porfolio/GorillaAPI>
   ```

3. **Install dependencies**:
   ```bash
   go get
   ```

4. **Run the application**:
   ```bash
   go run main.go
   ```

5. **Access the API**: Open Postman or your preferred API client and make requests to `http://localhost:8080/movies`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Gorilla Mux](https://github.com/gorilla/mux) for routing support in Go.

---
