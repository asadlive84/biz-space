-- init.sql
CREATE DATABASE auth_svc;
CREATE DATABASE shop_svc;

-- Connect to each database and install the required extension
\connect auth_svc;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
\connect shop_svc;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

