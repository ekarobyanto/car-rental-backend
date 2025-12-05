CREATE TABLE IF NOT EXISTS cars (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    production_year INT NOT NULL,
    passenger_capacity INT NOT NULL,
    transmission_type VARCHAR(20) NOT NULL CHECK (transmission_type IN ('manual', 'automatic')),
    license_plate VARCHAR(50) UNIQUE NOT NULL,
    rental_price_per_day DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'available' CHECK (status IN ('available', 'rented', 'maintenance')),
    photo_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_cars_status ON cars(status);
CREATE INDEX idx_cars_deleted_at ON cars(deleted_at);
CREATE INDEX idx_cars_license_plate ON cars(license_plate);
