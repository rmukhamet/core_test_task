CREATE TABLE IF NOT EXISTS retailer(
    id uuid PRIMARY KEY,
    name VARCHAR(255),
    address_city VARCHAR(255),
    address_street VARCHAR(255),
    address_house VARCHAR(255),
    owner_first_name VARCHAR(255),
    owner_last_name VARCHAR(255),
    open_time  timestamp,
    close_time  timestamp,
    version integer,
    actor VARCHAR(255),
    created_at timestamp,
    updated_at timestamp NULL,
    deleted_at timestamp NULL
);