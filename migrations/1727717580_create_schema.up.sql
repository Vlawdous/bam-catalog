CREATE TABLE IF NOT EXISTS Book (
                                    id UUID PRIMARY KEY,
                                    title VARCHAR(255) NOT NULL,
                                    published_year INT,
                                    isbn VARCHAR(20) UNIQUE NOT NULL,
                                    price FLOAT NOT NULL,
                                    description TEXT,
                                    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS Genre (
                                     id UUID PRIMARY KEY,
                                     genre_name VARCHAR(100) NOT NULL,
                                     description TEXT,
                                     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS BookGenre (
                                         id UUID PRIMARY KEY,
                                         book_id UUID REFERENCES Book(id) ON DELETE CASCADE,
                                         genre_id UUID REFERENCES Genre(id) ON DELETE CASCADE,
                                         deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS Author (
                                      id UUID PRIMARY KEY,
                                      name VARCHAR(100) NOT NULL,
                                      surname VARCHAR(100) NOT NULL,
                                      birthdate DATE,
                                      bio TEXT,
                                      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS BookAuthor (
                                          id UUID PRIMARY KEY,
                                          book_id UUID REFERENCES Book(id) ON DELETE CASCADE,
                                          author_id UUID REFERENCES Author(id) ON DELETE CASCADE,
                                          deleted BOOLEAN NOT NULL DEFAULT FALSE
);