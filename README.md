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
- MySQL 5.7 or higher (or Docker for easy setup)

### Installation

#### Option 1: Using Docker (Recommended)

1. Clone the repository:
```bash
git clone https://github.com/enter42/mtg-collection-tracker.git
cd mtg-collection-tracker
```

2. Install dependencies:
```bash
go mod download
```

3. Start MySQL using Docker Compose:
```bash
docker compose up -d
```

4. Copy the Docker environment file:
```bash
cp .env.docker .env
```

5. Wait for MySQL to be ready (about 10-15 seconds), then run the application:
```bash
go run cmd/server/main.go
```

#### Option 2: Using Existing MySQL Installation

1. Clone the repository:
```bash
git clone https://github.com/enter42/mtg-collection-tracker.git
cd mtg-collection-tracker
```

2. Install dependencies:
```bash
go mod download
```

3. Create MySQL database and user:
```bash
mysql -u root -p
```

```sql
CREATE DATABASE mtg_collection;
CREATE USER 'mtguser'@'localhost' IDENTIFIED BY 'mtgpass';
GRANT ALL PRIVILEGES ON mtg_collection.* TO 'mtguser'@'localhost';
FLUSH PRIVILEGES;
EXIT;
```

4. Copy the example environment file and configure it:
```bash
cp .env.example .env
```

5. Edit `.env` file with your database credentials:
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=mtguser
DB_PASSWORD=mtgpass
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

## Quick Start with Docker

For the fastest setup, run the provided setup script:

```bash
./setup.sh
```

This will:
- Start MySQL in a Docker container
- Configure the environment
- Provide instructions to run the application

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

## Database Schema

The application automatically creates the following tables:

### Users Table
- `id` - Primary key
- `username` - Unique username
- `password` - Bcrypt hashed password
- `created_at`, `updated_at`, `deleted_at` - Timestamps

### Cards Table
- `id` - Primary key
- `user_id` - Foreign key to users table
- `card_name` - Name of the card
- `card_image_url` - URL to card image
- `set_code` - MTG set code
- `collector_number` - Collector number
- `language` - Card language
- `quantity` - Number of copies
- `buying_price` - Purchase price in THB
- `bought_date` - Purchase date
- `sell_date` - Sale date (if sold)
- `created_at`, `updated_at`, `deleted_at` - Timestamps

## API Routes

### Public Routes
- `GET /` - Redirect to login
- `GET /login` - Login page
- `POST /login` - Login submission
- `GET /register` - Registration page
- `POST /register` - Registration submission

### Protected Routes (Requires Authentication)
- `GET /logout` - Logout
- `GET /cards` - List all cards with pagination and search
- `GET /cards/add` - Add card form
- `POST /cards/add` - Create new card
- `GET /cards/edit/:id` - Edit card form
- `POST /cards/edit/:id` - Update card
- `POST /cards/delete/:id` - Delete card

## Development

### Code Structure

The application follows Clean Architecture principles:

- **Domain Layer** (`internal/domain/`): Contains business entities and repository interfaces
- **Use Case Layer** (`internal/usecase/`): Contains business logic
- **Infrastructure Layer** (`internal/infrastructure/`): Contains database implementations
- **Handler Layer** (`internal/handler/`): Contains HTTP handlers and middleware
- **Templates** (`templates/`): Contains HTML templates

### Running Tests

```bash
make test
```

### Building

```bash
make build
./bin/server
```

## Troubleshooting

### Database Connection Issues

If you get database connection errors:
1. Ensure MySQL is running: `sudo service mysql status` or `docker compose ps`
2. Verify credentials in `.env` file
3. Check if database exists: `mysql -u mtguser -p -e "SHOW DATABASES;"`

### Port Already in Use

If port 8080 is already in use, change `SERVER_PORT` in `.env` file to a different port.

## License

MIT