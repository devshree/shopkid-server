# Kids Shop API

A RESTful API for a children's clothing and toy shop built with Go and PostgreSQL.

## Features

- Product management (list, get, create)
- Shopping cart functionality (view, add, remove items)
- PostgreSQL database integration
- CORS support for frontend integration

## API Endpoints

Detailed API documentation is available in OpenAPI/Swagger format. You can view it by:

1. Copy the contents of `swagger.yaml` to [Swagger Editor](https://editor.swagger.io/)
2. Or use a local Swagger UI server:
```bash
docker run -p 8081:8080 -e SWAGGER_JSON=/swagger.yaml -v $(pwd):/swagger swaggerapi/swagger-ui
```
Then visit http://localhost:8081

### Available Endpoints
- `GET /api/products` - List all products
- `GET /api/products/{id}` - Get a specific product
- `POST /api/products` - Create a new product
- `GET /api/cart` - View cart contents
- `POST /api/cart/add` - Add item to cart
- `DELETE /api/cart/remove/{id}` - Remove item from cart

## Prerequisites

### PostgreSQL Setup

The application uses PostgreSQL with the following default configuration:
- Database name: `kids_shop`
- Database user: `kidshop`
- No password (for local development)

### Install Go using GVM
```bash
# Install GVM prerequisites
sudo apt-get install curl git mercurial make binutils bison gcc build-essential

# Install GVM
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

# Source GVM
source ~/.gvm/scripts/gvm

# Install Go 1.4 (needed to build newer versions)
gvm install go1.4 -B
gvm use go1.4

# Install latest stable version
gvm install go1.21.6
gvm use go1.21.6 --default

# Verify installation
go version
```

Add Go to your PATH by adding these lines to your ~/.zshrc or ~/.bashrc:
```bash
export GOROOT=$HOME/.gvm/gos/go1.21.6
export GOPATH=$HOME/go
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

Then source your shell configuration:
```bash
source ~/.zshrc  # or source ~/.bashrc
```

### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install postgresql postgresql-contrib

# Start PostgreSQL service
sudo service postgresql start

# Enable PostgreSQL to start on boot
sudo systemctl enable postgresql
```

### macOS
1. Install Homebrew first:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

2. After installing Homebrew, you might need to add it to your PATH. Run these commands:
```bash
echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zshrc
source ~/.zshrc
```

3. Then install PostgreSQL and Go:
```bash
brew install postgresql@14

# Start PostgreSQL service
brew services start postgresql@14

# If PostgreSQL is already running, restart it
brew services restart postgresql@14

# Check PostgreSQL service status
brew services list | grep postgresql
```

Note: The `@14` specifies PostgreSQL version 14. Adjust this number if you installed a different version.

### Windows
1. Download and install Go from [golang.org](https://golang.org/dl/)
2. Download and install PostgreSQL from [postgresql.org](https://www.postgresql.org/download/windows/)
3. PostgreSQL service should start automatically. If not:
   - Open Services (Win + R, type 'services.msc')
   - Find 'PostgreSQL' service
   - Right-click and select 'Start'
   - Right-click again, select 'Properties', and set 'Startup type' to 'Automatic'

## Setup

1. Install dependencies:
```bash
go mod init kids_shop
go get github.com/gorilla/mux
go get github.com/lib/pq
go get github.com/joho/godotenv
```

2. Set up PostgreSQL database:
```bash
createdb kids_shop
psql kids_shop < schema.sql
```

3. Configure environment variables:
   - Copy `.env.example` to `.env`
   - Update the values in `.env` with your database credentials

4. Run the server:
```bash
go run .
```

## Environment Variables

Create a `.env` file with the following variables:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=kidshop
DB_PASSWORD=
DB_NAME=kids_shop

# Logging Configuration
ENABLE_REQUEST_LOGGING=ON

# CORS Configuration
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173,http://yourdomain.com
CORS_ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
CORS_ALLOWED_HEADERS=Content-Type,Authorization
CORS_ALLOW_CREDENTIALS=true
CORS_MAX_AGE=86400
```

### PostgreSQL Setup

The quick start script will handle the database setup automatically, but if you need to do it manually:

```bash
# Check PostgreSQL installation
which psql
psql --version

# Create the database user
createuser -s kidshop

# Create the database
createdb kids_shop

# Set ownership
psql postgres -c "ALTER DATABASE kids_shop OWNER TO kidshop"
psql kids_shop -c "ALTER SCHEMA public OWNER TO kidshop"

# Import the schema
PGUSER=kidshop PGDATABASE=kids_shop psql < schema.sql

# Verify database creation
psql -l | grep kids_shop
```

Note: For local development, the application is configured to use PostgreSQL without a password.

## Quick Start Script for macOS

To quickly set up the project on macOS, run the following script:

```bash
# For normal setup
./quick-start-mac.sh

# To clean and recreate database
./quick-start-mac.sh cleanAllDB
```

The script will:
- Install Homebrew if not present
- Install and start PostgreSQL
- Set up the database (kids_shop) and user (kidshop)
- Import the schema
- Initialize Go module and install dependencies
- Create the .env file with default database credentials
- Start the server automatically

### Script Options

- `cleanAllDB`: Drops and recreates the database and user
  - Removes existing database and user
  - Creates fresh database with schema
  - Useful for resetting to a clean state

### Troubleshooting

If you encounter database connection issues:
1. Check if PostgreSQL is running: `brew services list | grep postgresql`
2. Try cleaning and recreating the database: `./quick-start-mac.sh cleanAllDB`
3. Verify database exists: `psql -l | grep kids_shop`
4. Check user permissions: `psql -U kidshop -d kids_shop -c "\du"`

The server will start automatically after setup. Access the API at http://localhost:8080

