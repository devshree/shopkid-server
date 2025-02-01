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

### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install postgresql postgresql-contrib golang-go

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
brew install postgresql go

# Start PostgreSQL service
brew services start postgresql
```

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
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=kids_shop
```

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

