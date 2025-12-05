CREATE TABLE IF NOT EXISTS renters (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    id_card_number VARCHAR(50) UNIQUE NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    address TEXT NOT NULL,
    driving_license_number VARCHAR(50) NOT NULL,
    id_card_photo_url TEXT,
    driving_license_photo_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_renters_id_card_number ON renters(id_card_number);
CREATE INDEX idx_renters_deleted_at ON renters(deleted_at);
