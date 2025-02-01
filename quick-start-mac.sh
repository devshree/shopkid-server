#!/bin/bash

# Exit on any error
set -e

echo "🚀 Starting Kids Shop API setup..."

# Check if Homebrew is installed
if ! command -v brew &> /dev/null; then
    echo "📦 Installing Homebrew..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zshrc
    source ~/.zshrc
else
    echo "✅ Homebrew already installed"
fi

# Install PostgreSQL
echo "📦 Installing PostgreSQL..."
brew install postgresql@14

# Start PostgreSQL service
echo "🔄 Starting PostgreSQL service..."
brew services restart postgresql@14

# Wait for PostgreSQL to start
sleep 3

# Create database user and database
echo "🗄️ Setting up database..."
createuser -s $(whoami) || echo "✅ User already exists"
createdb kids_shop || echo "✅ Database already exists"

# Import schema
echo "📝 Importing database schema..."
psql kids_shop < schema.sql

# Install Go dependencies
echo "📦 Installing Go dependencies..."
go mod init kids_shop
go get github.com/gorilla/mux
go get github.com/lib/pq
go get github.com/joho/godotenv

# Create .env file
echo "⚙️ Creating .env file..."
cat > .env << EOL
DB_HOST=localhost
DB_PORT=5432
DB_USER=$(whoami)
DB_PASSWORD=
DB_NAME=kids_shop
EOL

echo "✨ Setup complete! Run 'go run .' to start the server" 