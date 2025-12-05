CREATE TABLE IF NOT EXISTS rental_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    renter_id UUID NOT NULL,
    car_id UUID NOT NULL,
    rental_start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    rental_end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    total_rental_cost DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'booked' CHECK (status IN ('booked', 'in-progress', 'completed', 'cancelled')),
    car_condition_on_return TEXT,
    penalty_fee DECIMAL(10, 2) DEFAULT 0,
    final_total_payment DECIMAL(10, 2),
    actual_return_date TIMESTAMP WITH TIME ZONE,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_renter FOREIGN KEY (renter_id) REFERENCES renters(id) ON DELETE RESTRICT,
    CONSTRAINT fk_car FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE RESTRICT
);

CREATE INDEX idx_rental_transactions_renter_id ON rental_transactions(renter_id);
CREATE INDEX idx_rental_transactions_car_id ON rental_transactions(car_id);
CREATE INDEX idx_rental_transactions_status ON rental_transactions(status);
CREATE INDEX idx_rental_transactions_deleted_at ON rental_transactions(deleted_at);
CREATE INDEX idx_rental_transactions_dates ON rental_transactions(rental_start_date, rental_end_date);
