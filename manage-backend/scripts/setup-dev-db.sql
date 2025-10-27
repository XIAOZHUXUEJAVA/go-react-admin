-- Development database setup script
-- Run this script to create the development database

-- Create development database
CREATE DATABASE go_manage_starter;

-- Create test database (if not exists)

-- Grant permissions (adjust username as needed)
GRANT ALL PRIVILEGES ON DATABASE go_manage_starter TO xiaozhu;

-- You can run this script with:
-- psql -U postgres -h localhost -f scripts/setup-dev-db.sql