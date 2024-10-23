CREATE SCHEMA IF NOT EXISTS geo;
CREATE SCHEMA IF NOT EXISTS users;

-- 1. Таблица для хранения истории поиска (схема geo)
CREATE TABLE IF NOT EXISTS geo.search_history (
                                                  id SERIAL PRIMARY KEY,   -- Уникальный идентификатор
                                                  data TEXT NOT NULL       -- Строка с данными поиска
);

-- 2. Таблица для хранения адресов (схема geo)
CREATE TABLE IF NOT EXISTS geo.address (
                                           id SERIAL PRIMARY KEY,   -- Уникальный идентификатор
                                           city VARCHAR(255) NOT NULL,    -- Город
                                           street VARCHAR(255) NOT NULL,  -- Улица
                                           house VARCHAR(50) NOT NULL,    -- Номер дома
                                           lat VARCHAR(50),               -- Широта
                                           lon VARCHAR(50)                -- Долгота
);

-- 3. Таблица для связи истории поиска с адресами (схема geo)
CREATE TABLE IF NOT EXISTS geo.history_search_address (
                                                          id SERIAL PRIMARY KEY,      -- Уникальный идентификатор
                                                          id_search INT NOT NULL,     -- Идентификатор поиска из таблицы search_history
                                                          id_address INT NOT NULL,    -- Идентификатор адреса из таблицы address
                                                          FOREIGN KEY (id_search) REFERENCES geo.search_history(id) ON DELETE CASCADE,  -- Указание схемы geo
                                                          FOREIGN KEY (id_address) REFERENCES geo.address(id) ON DELETE CASCADE         -- Указание схемы geo
);

-- 4. Таблица для хранения пользователей (схема users)
CREATE TABLE IF NOT EXISTS users.users (
                                           id SERIAL PRIMARY KEY,       -- Уникальный идентификатор пользователя
                                           username VARCHAR(255) UNIQUE NOT NULL,  -- Имя пользователя (уникальное)
                                           password VARCHAR(255) NOT NULL          -- Пароль пользователя (зашифрованный)
);

-- Расширение fuzzystrmatch (оно создается в схеме public)
CREATE EXTENSION IF NOT EXISTS fuzzystrmatch;
