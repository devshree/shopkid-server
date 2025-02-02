#!/bin/bash

# Create database user and database
echo "🗄️ Setting up database..."

# Ensure we're using postgres database for admin commands
export PGDATABASE=postgres

# Switch to postgres user to create new user and database
echo "Creating database user 'kidshop'..."
createuser -s kidshop || echo "✅ User already exists"

echo "Creating database 'kids_shop'..."
psql -U kidshop postgres -c 'CREATE DATABASE kids_shop;' || echo "✅ Database already exists"

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