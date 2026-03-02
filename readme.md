# Go API Tutorial

A simple Go-based API service demonstrating middleware, dependency injection with interfaces, and mock database implementation.

## Features
- **Routing**: Uses `go-chi` for lightweight routing.
- **Logging**: Uses `logrus` for structured logging.
- **Middleware**: Custom authorization middleware.
- **Database**: Mock database implementation with interface-based design.

## Getting Started

### Prerequisites
- Go 1.26 or higher

### Installation
1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the API
```bash
go run cmd/api/main.go
```
The server will start on `localhost:8000`.

## API Usage

### Get Coin Balance
Returns the coin balance for a specific user.

- **URL**: `/account/coins`
- **Method**: `GET`
- **Query Params**:
  - `username` (string): The username to query.
- **Headers**:
  - `Authorization` (string): Required. Valid tokens are listed below.

#### Test Credentials
| Username | AuthToken |
| :--- | :--- |
| `andy` | `123ABC` |
| `bob` | `456DEF` |
| `marie` | `789GHI` |

#### Example Request
```bash
curl -H "Authorization: 123ABC" "http://localhost:8000/account/coins?username=andy"
```

## Credits
Tutorial by [Alex Mux](https://www.youtube.com/watch?v=8uiZC0l4Ajw)