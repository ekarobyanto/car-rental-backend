-- Delete seed data in reverse order (using UUID values)
DELETE FROM rental_transactions WHERE renter_id IN (
    SELECT id FROM renters WHERE id_card_number LIKE '3201%'
);

DELETE FROM renters WHERE id_card_number LIKE '3201%';

DELETE FROM cars WHERE license_plate LIKE 'B-%';

DELETE FROM users WHERE email IN ('admin@carental.com', 'staff@carental.com');
