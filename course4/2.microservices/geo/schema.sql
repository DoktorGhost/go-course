CREATE TABLE IF NOT EXISTS search_history (
                                                  id SERIAL PRIMARY KEY,   -- Уникальный идентификатор
                                                  data TEXT NOT NULL       -- Строка с данными поиска
);

-- 2. Таблица для хранения адресов (схема geo)
CREATE TABLE IF NOT EXISTS address (
                                           id SERIAL PRIMARY KEY,   -- Уникальный идентификатор
                                           city VARCHAR(255) NOT NULL,    -- Город
                                           street VARCHAR(255) NOT NULL,  -- Улица
                                           house VARCHAR(50) NOT NULL,    -- Номер дома
                                           lat VARCHAR(50),               -- Широта
                                           lon VARCHAR(50)                -- Долгота
);

-- 3. Таблица для связи истории поиска с адресами (схема geo)
CREATE TABLE IF NOT EXISTS history_search_address (
                                                          id SERIAL PRIMARY KEY,      -- Уникальный идентификатор
                                                          id_search INT NOT NULL,     -- Идентификатор поиска из таблицы search_history
                                                          id_address INT NOT NULL,    -- Идентификатор адреса из таблицы address
                                                          FOREIGN KEY (id_search) REFERENCES search_history(id) ON DELETE CASCADE,  -- Указание схемы geo
                                                          FOREIGN KEY (id_address) REFERENCES address(id) ON DELETE CASCADE         -- Указание схемы geo
);

-- Расширение fuzzystrmatch (оно создается в схеме public)
CREATE EXTENSION IF NOT EXISTS fuzzystrmatch;