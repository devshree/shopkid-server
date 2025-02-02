#!/bin/bash

# Exit on any error
set -e

# Check for cleanAllDB parameter
CLEAN_DB=false
for arg in "$@"
do
    if [ "$arg" == "cleanAllDB" ]; then
        CLEAN_DB=true
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

# Drop database if cleanAllDB is true
if [ "$CLEAN_DB" = true ]; then
    echo "🗑️  Cleaning database..."
    dropdb --if-exists kids_shop
    dropuser --if-exists kidshop
    echo "✅ Database cleaned"
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

# Ensure we're using postgres database for admin commands
export PGDATABASE=postgres

# Switch to postgres user to create new user and database
echo "Creating database user 'kidshop'..."
createuser -s kidshop || echo "✅ User already exists"

echo "Creating database 'kids_shop'..."
createdb kids_shop || echo "✅ Database already exists"

# Grant privileges
echo "Granting privileges..."
psql postgres -c "ALTER DATABASE kids_shop OWNER TO kidshop"
psql kids_shop -c "ALTER SCHEMA public OWNER TO kidshop"

# Import schema
echo "📝 Importing database schema..."
# Check if tables exist and cleanAllDB is not true
if [ "$CLEAN_DB" = false ] && PGDATABASE=kids_shop psql -c "\dt" | grep -q 'products\|cart_items'; then
    echo "✅ Database tables already exist"
else
    echo "Importing schema as kidshop user..."
    PGUSER=kidshop PGDATABASE=kids_shop psql < schema.sql
    echo "✅ Schema imported successfully"
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

# Always create/update .env to ensure correct values
cat > .env << EOL
DB_HOST=localhost
DB_PORT=5432
DB_USER=kidshop
DB_PASSWORD=
DB_NAME=kids_shop
EOL
echo "✅ .env file updated"

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

# Start the server
echo "🔄 Starting server..."
go run .

echo "✅ Server started successfully"


