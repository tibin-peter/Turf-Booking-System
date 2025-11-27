CREATE TABLE IF NOT EXISTS booking(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    turf_id INT NOT NULL REFERENCES turs(id)  ON DELETE CASCADE,

    date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,

    amount INTEGER NOT NULL,

    payment_method VARCHAR(20) DEFAULT 'cash',
    payment_status VARCHAR(20) DEFAULT  'success',

    status VARCHAR(20) DEFAULT 'booked',

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);