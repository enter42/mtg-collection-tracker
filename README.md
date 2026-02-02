# MTG Collection Tracker

A web application for managing your Magic: The Gathering card collection.

## Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin (HTTP web framework)
- **Architecture**: Clean Architecture
- **Frontend**: Server-side rendered HTML using Go templates
- **UI Framework**: Bootstrap 5 (CDN)
- **Database**: MySQL
- **ORM**: GORM

## Features

### 1. User Authentication
- User registration
- Login/Logout functionality
- Session-based authentication

### 2. Card Collection Management
- Add cards with the following information:
  - Card name
  - Card image URL
  - Set code
  - Collector number
  - Language
  - Quantity
  - Buying price (THB)
  - Bought date
  - Sell date
- Edit existing cards
- Delete cards

### 3. Card Collection List
- View all cards in your collection
- Pagination (20 cards per page)
- Search/filter by card name, set code, or collector number
- Display card images
- Sort by creation date (newest first)

## Setup Instructions

### Prerequisites
- Go 1.21 or higher
- MySQL 5.7 or higher

### Installation

1. Clone the repository:
```bash
git clone https://github.com/enter42/mtg-collection-tracker.git
cd mtg-collection-tracker
```

2. Install dependencies:
```bash
go mod download
```

3. Create MySQL database:
```bash
mysql -u root -p -e "CREATE DATABASE mtg_collection;"
```

4. Copy the example environment file and configure it:
```bash
cp .env.example .env
```

5. Edit `.env` file with your database credentials:
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=mtg_collection
SERVER_PORT=8080
SESSION_SECRET=your-secret-key-change-this
```

### Running the Application

1. Using Go directly:
```bash
go run cmd/server/main.go
```

2. Using Makefile:
```bash
make run
```

3. Build and run binary:
```bash
make build
./bin/server
```

The application will be available at `http://localhost:8080`

## Usage

1. **Register**: Navigate to `http://localhost:8080/register` and create a new account
2. **Login**: Use your credentials to log in
3. **Add Cards**: Click "Add Card" button to add cards to your collection
4. **View Collection**: Browse your cards with pagination and search
5. **Edit/Delete**: Manage your cards using the action buttons

## Project Structure

```
mtg-collection-tracker/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── domain/
│   │   ├── entity/           # Domain entities (User, Card)
│   │   └── repository/       # Repository interfaces
│   ├── infrastructure/
│   │   ├── database/         # Database connection
│   │   └── repository/       # Repository implementations
│   ├── usecase/              # Business logic
│   └── handler/              # HTTP handlers and middleware
├── templates/
│   ├── layouts/              # Layout templates
│   └── pages/                # Page templates
├── .env.example              # Environment variables template
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## License

MIT