#!/bin/bash

echo "Starting MTG Collection Tracker Setup..."

# Check if Docker is available
if command -v docker &> /dev/null && command -v docker-compose &> /dev/null; then
    echo "Docker found. Starting MySQL container..."
    docker-compose up -d
    
    echo "Waiting for MySQL to be ready..."
    sleep 10
    
    echo "Copying Docker environment file..."
    cp .env.docker .env
    
    echo "Setup complete! MySQL is running in Docker."
else
    echo "Docker not found. Please set up MySQL manually."
    echo "1. Install MySQL 8.0+"
    echo "2. Create database: CREATE DATABASE mtg_collection;"
    echo "3. Create user: CREATE USER 'mtguser'@'localhost' IDENTIFIED BY 'mtgpass';"
    echo "4. Grant privileges: GRANT ALL PRIVILEGES ON mtg_collection.* TO 'mtguser'@'localhost';"
    echo "5. Copy .env.example to .env and update with your credentials"
fi

echo ""
echo "To run the application:"
echo "  go run cmd/server/main.go"
echo ""
echo "Or using Make:"
echo "  make run"
echo ""
echo "Access the application at: http://localhost:8080"
