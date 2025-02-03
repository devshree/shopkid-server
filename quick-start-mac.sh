#!/bin/bash

# Exit on any error
set -e

# Check for parameters
CLEAN_DB=""
WITH_SAMPLE_DATA=""
for arg in "$@"
do
    if [ "$arg" == "cleanAllDB" ]; then
        CLEAN_DB=cleanAllDB
    elif [ "$arg" == "withSampleData" ]; then
        WITH_SAMPLE_DATA=true
    fi
done

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

# Make the setup-db.sh script executable
chmod +x setup-db.sh

# Run database setup
./setup-db.sh $CLEAN_DB

# Insert sample data if requested
if [ "$WITH_SAMPLE_DATA" = true ]; then
    echo "📝 Inserting sample data..."
    PGPASSWORD=kidshop psql -U kidshop -d kids_shop -f mockData.sql
    if [ $? -eq 0 ]; then
        echo "✅ Sample data inserted successfully"
    else
        echo "❌ Failed to insert sample data"
        exit 1
    fi
fi

# Install Go dependencies
echo "📦 Installing Go dependencies..."
# Check if go.mod exists
if [ -f "go.mod" ]; then
    echo "✅ Go module already initialized"
    go get -u ./...  # Update existing dependencies
else
    go mod init kids_shop
    go get github.com/gorilla/mux
    go get github.com/lib/pq
    go get github.com/joho/godotenv
    go get github.com/gorilla/handlers
fi

# Create .env file
echo "⚙️ Creating .env file..."
# Check if .env exists
if [ -f ".env" ]; then
    echo "🔄 Updating .env file..."
    # Backup existing .env
    cp .env .env.backup
fi
 
# Source the environment variables
set -a # automatically export all variables
source .env
set +a

# Verify database connection
echo "🔍 Verifying database connection..."
if PGUSER=kidshop PGDATABASE=kids_shop psql -c "SELECT 1" > /dev/null 2>&1; then
    echo "✅ Database connection successful"
else
    echo "❌ Database connection failed"
    echo "Error: Unable to connect to database 'kids_shop' as user 'kidshop'"
    echo "Try running with cleanAllDB: ./quick-start-mac.sh cleanAllDB"
    exit 1
fi

echo "✨ Setup complete!"
echo "🚀 To start the server, run: go run ."
echo "🌍 Server will be available at http://localhost:8080" 

# Install required Go tools
echo "Installing Go tools..."

go install golang.org/x/tools/gopls@latest
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Download dependencies
echo "Downloading dependencies..."
go mod tidy

# Start the server
echo "🔄 Starting server..."
go run .

echo "✅ Server started successfully"


