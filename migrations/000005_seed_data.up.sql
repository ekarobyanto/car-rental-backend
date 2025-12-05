-- Seed admin user
INSERT INTO users (id, email, name, password, created_at, updated_at) VALUES
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'admin@carental.com', 'Admin User', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRps.9cGLcZEiGDMVr5yUP1KUOYTa', NOW(), NOW()),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'staff@carental.com', 'Staff User', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRps.9cGLcZEiGDMVr5yUP1KUOYTa', NOW(), NOW());
-- Password for both: password123

-- Seed cars
INSERT INTO cars (id, name, brand, production_year, passenger_capacity, transmission_type, license_plate, rental_price_per_day, status, photo_url, created_at, updated_at) VALUES
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a01', 'Avanza 1.3 G', 'Toyota', 2023, 7, 'manual', 'B-1234-ABC', 250000.00, 'available', 'https://example.com/cars/avanza.jpg', NOW(), NOW()),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a02', 'Xenia 1.3 R', 'Daihatsu', 2023, 7, 'manual', 'B-5678-DEF', 240000.00, 'available', 'https://example.com/cars/xenia.jpg', NOW(), NOW()),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a03', 'Innova Reborn 2.4 G', 'Toyota', 2022, 7, 'automatic', 'B-9012-GHI', 450000.00, 'available', 'https://example.com/cars/innova.jpg', NOW(), NOW()),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a04', 'CR-V 1.5 Turbo', 'Honda', 2023, 7, 'automatic', 'B-3456-JKL', 650000.00, 'available', 'https://example.com/cars/crv.jpg', NOW(), NOW()),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a05', 'Pajero Sport Dakar', 'Mitsubishi', 2022, 7, 'automatic', 'B-7890-MNO', 750000.00, 'available', 'https://example.com/cars/pajero.jpg', NOW(), NOW()),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a06', 'Fortuner 2.4 VRZ', 'Toyota', 2023, 7, 'automatic', 'B-2345-PQR', 800000.00, 'rented', 'https://example.com/cars/fortuner.jpg', NOW(), NOW()),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a07', 'Brio RS CVT', 'Honda', 2023, 5, 'automatic', 'B-6789-STU', 200000.00, 'available', 'https://example.com/cars/brio.jpg', NOW(), NOW()),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a08', 'Agya 1.2 G', 'Toyota', 2022, 5, 'manual', 'B-0123-VWX', 180000.00, 'available', 'https://example.com/cars/agya.jpg', NOW(), NOW()),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a09', 'Jazz RS CVT', 'Honda', 2023, 5, 'automatic', 'B-4567-YZA', 350000.00, 'available', 'https://example.com/cars/jazz.jpg', NOW(), NOW()),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a10', 'Ertiga Sport', 'Suzuki', 2023, 7, 'automatic', 'B-8901-BCD', 300000.00, 'maintenance', 'https://example.com/cars/ertiga.jpg', NOW(), NOW());

-- Seed renters
INSERT INTO renters (id, name, id_card_number, phone_number, address, driving_license_number, created_at, updated_at) VALUES
('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a01', 'John Doe', '3201012345678901', '+6281234567890', 'Jl. Merdeka No. 123, Jakarta Pusat', 'SIM-1234567890', NOW(), NOW()),
('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a02', 'Jane Smith', '3201019876543210', '+6281234567891', 'Jl. Sudirman No. 456, Jakarta Selatan', 'SIM-0987654321', NOW(), NOW()),
('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a03', 'Robert Johnson', '3201015555666677', '+6281234567892', 'Jl. Thamrin No. 789, Jakarta Pusat', 'SIM-5555666677', NOW(), NOW()),
('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a04', 'Emily Davis', '3201011111222233', '+6281234567893', 'Jl. Gatot Subroto No. 321, Jakarta Selatan', 'SIM-1111222233', NOW(), NOW()),
('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a05', 'Michael Wilson', '3201014444555566', '+6281234567894', 'Jl. Kuningan No. 654, Jakarta Selatan', 'SIM-4444555566', NOW(), NOW());

-- Seed rental transactions
INSERT INTO rental_transactions (id, renter_id, car_id, rental_start_date, rental_end_date, total_rental_cost, status, created_at, updated_at) VALUES
('d0eebc99-9c0b-4ef8-bb6d-6bb9bd380a01', 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a01', 'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a06', NOW() - INTERVAL '2 days', NOW() + INTERVAL '3 days', 4000000.00, 'in-progress', NOW(), NOW()),
('d0eebc99-9c0b-4ef8-bb6d-6bb9bd380a02', 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a02', 'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a01', NOW() + INTERVAL '5 days', NOW() + INTERVAL '8 days', 750000.00, 'booked', NOW(), NOW());

-- Sample completed transaction
INSERT INTO rental_transactions (
    id, renter_id, car_id, rental_start_date, rental_end_date, 
    total_rental_cost, status, car_condition_on_return, 
    penalty_fee, final_total_payment, actual_return_date,
    created_at, updated_at
) VALUES
('d0eebc99-9c0b-4ef8-bb6d-6bb9bd380a03', 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a03', 'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a03', NOW() - INTERVAL '10 days', NOW() - INTERVAL '7 days', 1350000.00, 
'completed', 'Good condition, no damage', 0.00, 1350000.00, NOW() - INTERVAL '7 days', NOW(), NOW());
