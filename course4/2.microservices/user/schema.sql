CREATE TABLE IF NOT EXISTS users (
                                           id SERIAL PRIMARY KEY,       -- Уникальный идентификатор пользователя
                                           email VARCHAR(255) UNIQUE NOT NULL,  -- Имя пользователя (уникальное)
                                           password VARCHAR(255) NOT NULL          -- Пароль пользователя (зашифрованный)
);