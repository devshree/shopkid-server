#!/bin/bash
set -e
# Check for cleanAllDB parameter to clean all databases
CLEAN_DB=false
for arg in "$@"
do
    if [ "$arg" == "cleanAllDB" ]; then
        CLEAN_DB=true
    fi
done

# Drop database if cleanAllDB is true
if [ "$CLEAN_DB" = true ]; then
    echo "🗑️  Cleaning database..."
    
    # Terminate all connections to the database
    echo "Terminating active connections..."
    psql postgres -c "
        SELECT pg_terminate_backend(pg_stat_activity.pid)
        FROM pg_stat_activity
        WHERE pg_stat_activity.datname IN ('kids_shop', 'kidshop')
        AND pid <> pg_backend_pid();" || true
    
    # Wait a moment for connections to close
    sleep 2

    # Drop databases first
    dropdb --if-exists kids_shop
    dropdb --if-exists kidshop

    # Then drop the user
    psql postgres -c "DROP OWNED BY kidshop;" || true
    dropuser --if-exists kidshop
    
    echo "✅ Database cleaned"
fi

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
if [ "$CLEAN_DB" = false ] ; then
    echo "✅ Database tables already exist"
else
    echo "Importing schema as kidshop user..."
    PGUSER=kidshop PGDATABASE=kids_shop psql < schema.sql
    echo "✅ Schema imported successfully"
fi

echo "🎉 Database setup complete!"