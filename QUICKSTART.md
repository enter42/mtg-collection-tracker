# Quick Start Guide

## 1. Prerequisites Check

Before starting, ensure you have:
- ✅ Go 1.21 or higher: `go version`
- ✅ Docker installed (recommended): `docker --version`
- ✅ OR MySQL 8.0+ installed: `mysql --version`

## 2. Installation (5 minutes)

### Using Docker (Easiest)

```bash
# Clone the repository
git clone https://github.com/enter42/mtg-collection-tracker.git
cd mtg-collection-tracker

# Install Go dependencies
go mod download

# Start MySQL in Docker
docker compose up -d

# Copy environment configuration
cp .env.docker .env

# Wait 10 seconds for MySQL to initialize
sleep 10

# Run the application
go run cmd/server/main.go
```

### Using Existing MySQL

```bash
# Clone the repository
git clone https://github.com/enter42/mtg-collection-tracker.git
cd mtg-collection-tracker

# Install Go dependencies
go mod download

# Create database and user
mysql -u root -p << EOF
CREATE DATABASE mtg_collection;
CREATE USER 'mtguser'@'localhost' IDENTIFIED BY 'mtgpass';
GRANT ALL PRIVILEGES ON mtg_collection.* TO 'mtguser'@'localhost';
FLUSH PRIVILEGES;
EOF

# Copy and edit environment file
cp .env.example .env
# Edit .env with your database credentials

# Run the application
go run cmd/server/main.go
```

## 3. Access the Application

Open your browser and navigate to:
```
http://localhost:8080
```

## 4. First-Time Setup

### Create Your Account
1. Click "Register here" on the login page
2. Enter a username and password
3. Confirm your password
4. Click "Register"

### Login
1. Enter your username and password
2. Click "Login"
3. You'll be redirected to your card collection (empty at first)

### Add Your First Card

1. Click the blue "Add Card" button (top right)
2. Fill in the card details:
   - **Card Name**: e.g., "Black Lotus"
   - **Card Image URL**: (optional) e.g., "https://example.com/image.jpg"
   - **Set Code**: e.g., "LEA"
   - **Collector Number**: e.g., "232"
   - **Language**: e.g., "English"
   - **Quantity**: e.g., "1"
   - **Buying Price**: e.g., "50000.00" (in THB)
   - **Bought Date**: Select from date picker
   - **Sell Date**: (optional) Leave blank if not sold
3. Click "Add Card"
4. You'll see your card in the collection list

## 5. Managing Your Collection

### Search Cards
- Use the search bar to find cards by name, set code, or collector number
- Click "Search" to filter results

### Edit a Card
- Click the yellow pencil icon (✏️) in the Actions column
- Update any fields
- Click "Update Card"

### Delete a Card
- Click the red trash icon (🗑️) in the Actions column
- Confirm the deletion
- Card will be permanently removed

### Browse Pages
- If you have more than 20 cards, use pagination at the bottom
- Click page numbers or Previous/Next buttons

## 6. Tips & Tricks

### Image URLs
- Find card images on sites like Scryfall or Gatherer
- Right-click on the image and copy the URL
- Paste into the "Card Image URL" field

### Organizing
- Use consistent naming for sets (e.g., "DMU", "BRO", "ONE")
- Add all details for better tracking
- Use the search function to find specific cards quickly

### Price Tracking
- Record prices in Thai Baht (THB)
- Update "Sell Date" when you sell a card
- Use this to track your collection's value

### Bulk Entry
- The form remembers your language preference
- Quantity defaults to 1 (adjust for playsets)
- Leave sell date blank for cards you still own

## 7. Troubleshooting

### Can't Connect to Database
**Docker users:**
```bash
docker compose ps  # Check if MySQL is running
docker compose logs  # Check for errors
docker compose restart  # Restart if needed
```

**MySQL users:**
```bash
mysql -u mtguser -pmtgpass mtg_collection  # Test connection
```

### Port Already in Use
Edit `.env` file and change `SERVER_PORT`:
```
SERVER_PORT=3000  # or any available port
```

### Forgot Password
Currently, password reset is not implemented. You'll need to:
1. Access MySQL directly
2. Delete the user: `DELETE FROM users WHERE username='yourname';`
3. Register again

### Application Won't Start
```bash
# Check Go installation
go version

# Check dependencies
go mod tidy

# Rebuild
go build -o bin/server cmd/server/main.go
./bin/server
```

## 8. Stopping the Application

### Stop the Server
Press `Ctrl+C` in the terminal where the application is running

### Stop Docker MySQL (if used)
```bash
docker compose down  # Stops and removes containers
# OR
docker compose stop  # Stops but keeps containers
```

### Keep Data Safe
The Docker setup uses a volume (`mysql_data`) to persist your data even after stopping containers. Your data will be safe!

## 9. Updating the Application

```bash
# Pull latest changes
git pull origin main

# Update dependencies
go mod tidy

# Restart the application
go run cmd/server/main.go
```

## 10. Next Steps

- Explore the FEATURES.md file for detailed feature documentation
- Check README.md for architecture and API details
- Add all your MTG cards!
- Track your collection value over time
- Use search and filter to organize your collection

## Support

For issues or questions:
- Check the README.md file
- Review FEATURES.md for detailed documentation
- Open an issue on GitHub

Happy collecting! 🎴
