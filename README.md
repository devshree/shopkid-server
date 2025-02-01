# Kids Shop API

A RESTful API for a children's clothing and toy shop built with Go and PostgreSQL.

## Features

- Product management (list, get, create)
- Shopping cart functionality (view, add, remove items)
- PostgreSQL database integration

## API Endpoints

- `GET /api/products` - List all products
- `GET /api/products/{id}` - Get a specific product
- `POST /api/products` - Create a new product
- `GET /api/cart` - View cart contents
- `POST /api/cart/add` - Add item to cart
- `DELETE /api/cart/remove/{id}` - Remove item from cart

## Prerequisites

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
DB_USER=your_system_username    # On macOS, this is your system username
DB_PASSWORD=                    # Leave empty for local development
DB_NAME=kids_shop
```

### PostgreSQL Setup

After installing PostgreSQL, you need to create a database user. On macOS:

```bash
# Check PostgreSQL installation
which psql
psql --version

# Create a PostgreSQL user with your system username
createuser -s $(whoami)

# Create the database
createdb kids_shop

# Import the schema
psql kids_shop < schema.sql

# Verify database creation
psql -l | grep kids_shop
```

Note: For local development on macOS, you typically don't need a password for PostgreSQL when using your system username.

The project structure should look like this:

kids_shop/
├── .env
├── .env.example
├── .gitignore
├── README.md
├── main.go
├── db.go
├── handlers.go
├── models.go
├── schema.sql
└── go.mod


## Quick Start Script for macOS

To quickly set up the project on macOS, run the following script:

```bash
./quick-start-mac.sh
```

The script will:
- Install Homebrew if not present
- Install and start PostgreSQL
- Set up the database and user
- Import the schema
- Initialize Go module and install dependencies
- Create the .env file with your system username

After running the script, you can start the server with `go run .`

