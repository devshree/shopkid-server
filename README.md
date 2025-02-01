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

## Setup

1. Install dependencies:
```
bash
go mod init kids_shop
go get github.com/gorilla/mux
go get github.com/lib/pq
go get github.com/joho/godotenv
```
2. Set up PostgreSQL database:
bash
createdb kids_shop
psql kids_shop < schema.sql

3. Configure environment variables:
   - Copy `.env.example` to `.env`
   - Update the values in `.env` with your database credentials

4. Run the server:
bash
go run .

## Environment Variables

Create a `.env` file with the following variables:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=kids_shop
```

## License

MIT

