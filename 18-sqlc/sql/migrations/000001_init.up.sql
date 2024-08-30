CREATE TABLE categories (
    id varchar(36) PRIMARY KEY NOT NULL,
    name text NOT NULL,
    description text
);

CREATE TABLE courses (
    id varchar(36) PRIMARY KEY NOT NULL,
    category_id varchar(36) NOT NULL,
    name text NOT NULL,
    price decimal(10,2) NOT NULL,
    description text,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
