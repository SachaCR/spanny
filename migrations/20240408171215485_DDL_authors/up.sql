CREATE TABLE
    authors (
        id STRING (36) NOT NULL DEFAULT (GENERATE_UUID ()),
        fullname STRING (MAX) NOT NULL,
        editor STRING (36) NOT NULL,
        createdAt TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP())
    ) PRIMARY KEY (id);