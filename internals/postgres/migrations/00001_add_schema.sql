-- +goose Up

-- USERS TABLE
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- BOOKINGS TABLE
CREATE TABLE bookings (
    id TEXT PRIMARY KEY,
    seat_id TEXT NOT NULL,
    movie_id TEXT NOT NULL,
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    user_id TEXT NOT NULL,
    payment_id TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_bookings_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- PAYMENTS TABLE
CREATE TABLE payments (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    booking_id TEXT NOT NULL UNIQUE, -- 🔒 enforce one payment per booking

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_payments_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_payments_booking
        FOREIGN KEY (booking_id)
        REFERENCES bookings(id)
        ON DELETE CASCADE
);

-- Add FK after both tables exist (avoid circular dependency issue)
ALTER TABLE bookings
ADD CONSTRAINT fk_bookings_payment
FOREIGN KEY (payment_id)
REFERENCES payments(id)
ON DELETE SET NULL;

-- AUTO UPDATE TRIGGER

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Apply trigger to all tables
CREATE TRIGGER set_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER set_bookings_updated_at
BEFORE UPDATE ON bookings
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER set_payments_updated_at
BEFORE UPDATE ON payments
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- INDEXES

CREATE INDEX idx_bookings_user_id ON bookings(user_id);
CREATE INDEX idx_payments_user_id ON payments(user_id);
CREATE INDEX idx_payments_booking_id ON payments(booking_id);


-- +goose Down

DROP TRIGGER IF EXISTS set_payments_updated_at ON payments;
DROP TRIGGER IF EXISTS set_bookings_updated_at ON bookings;
DROP TRIGGER IF EXISTS set_users_updated_at ON users;

DROP FUNCTION IF EXISTS set_updated_at;

DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS bookings;
DROP TABLE IF EXISTS users;
