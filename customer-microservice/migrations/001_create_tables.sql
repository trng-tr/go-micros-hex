-- Active: 1768501849859@@127.0.0.1@5433@goapp1db
    CREATE TABLE IF NOT EXISTS addresses (
        id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        street_number VARCHAR(10) NOT NULL,
        street_name VARCHAR(100) NOT NULL,
        zip_code VARCHAR(10) NOT NULL,
        city VARCHAR(100) NOT NULL,
        region VARCHAR(100) NOT NULL,
        country VARCHAR(100) NOT NULL,
        complement VARCHAR(100) DEFAULT NULL
    )

    CREATE TABLE IF NOT EXISTS customers (
        id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        firstname VARCHAR(100) NOT NULL,
        lastname VARCHAR(100) NOT NULL,
        genda VARCHAR(1) NOT NULL,
        email VARCHAR(100) NOT NULL UNIQUE,
        phone_number VARCHAR(20) NOT NULL UNIQUE,
        status VARCHAR(9) NOT NULL,
        address_id BIGINT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP DEFAULT NULL,
        CONSTRAINT fk_customers FOREIGN KEY (address_id) REFERENCES addresses(id) ON DELETE SET NULL,
        CONSTRAINT check_customer_genda CHECK(genda IN('F','M')),
        CONSTRAINT check_customer_status CHECK(status IN('ACTIVE','SUSPENDED','DELETED'))
    )