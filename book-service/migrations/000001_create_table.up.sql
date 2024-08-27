DO $$ 
BEGIN 
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'writing_type_enum') THEN
        CREATE TYPE writing_type_enum AS ENUM(
            'latin', 
            'cyrillic'
        );
    END IF;   
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
        CREATE TYPE role_enum AS ENUM(
            'admin', 
            'user',
            'provider'
        );
    END IF;    
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'working_styles_enum') THEN
        CREATE TYPE working_styles_enum AS ENUM(
            'remote', 
            'offline',
            'hybrid'
        );
    END IF;   
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'working_types_enum') THEN
        CREATE TYPE working_types_enum AS ENUM(
            'full_time', 
            'part_time'
        );
    END IF;  
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vacancy_status') THEN
        CREATE TYPE vacancy_status AS ENUM(
            'active', 
            'non-active'
        );
    END IF;         
END $$;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    is_email_verified BOOLEAN DEFAULT false,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(13),
    date_of_birth DATE,
    role role_enum NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- Trigger function to update the updated_at column
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger for the users table
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Create the tokens table
CREATE TABLE IF NOT EXISTS tokens (
    token_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    access_token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at VARCHAR(255),
    expires_at VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS authors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    biography TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS publishers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone_number VARCHAR(15),
    address VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS translators (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS categories (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS languages (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS books (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    provider_id UUID REFERENCES users(id) NOT NULL,
    publisher_id UUID REFERENCES publishers(id) NOT NULL,
    category_id UUID REFERENCES categories(id) NOT NULL,
    translator_id UUID REFERENCES translators(id) NOT NULL,
    author_id UUID REFERENCES authors(id) NOT NULL,
    language_id UUID REFERENCES languages(id) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    published_year DATE NOT NULL,
    total_pages INTEGER NOT NULL CHECK (total_pages > 0),
    price DECIMAL(10, 2) NOT NULL CHECK (price >= 0),
    stock INTEGER NOT NULL CHECK (stock >= 0),
    image_url VARCHAR(255),
    writing_type writing_type_enum NOT NULL,
    view_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS vacancies (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    publisher_id UUID REFERENCES publishers(id) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status vacancy_status NOT NULL DEFAULT 'active',
    salary_from NUMERIC CHECK (salary_from >= 0),
    salary_to NUMERIC CHECK (salary_to >= 0),
    working_styles working_styles_enum,
    working_types working_types_enum,
    phone_number VARCHAR(13) NOT NULL,
    location VARCHAR(255),
    view_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    CONSTRAINT check_salary_range CHECK (salary_to >= salary_from)
);

ALTER TABLE users
ADD CONSTRAINT users_unique_tg UNIQUE (email, phone_number, deleted_at);

ALTER TABLE publishers
ADD CONSTRAINT publishers_unique_tg UNIQUE (email, phone_number, deleted_at);