CREATE TABLE
    books (
        id STRING (36) NOT NULL DEFAULT (GENERATE_UUID ()),
        title STRING (MAX) NOT NULL,
        author STRING (36) NOT NULL,
        price FLOAT64 NOT NULL,
        createdAt TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP())
    ) PRIMARY KEY (id);