-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL,
                                     email VARCHAR(255) NOT NULL UNIQUE,
                                     password VARCHAR(255) NOT NULL,
                                     is_admin BOOLEAN DEFAULT false
);

-- Таблица хакатонов
CREATE TABLE IF NOT EXISTS hackathons (
                                          id SERIAL PRIMARY KEY,
                                          name VARCHAR(255) NOT NULL,
                                          description TEXT,
                                          start_date TIMESTAMP,
                                          end_date TIMESTAMP,
                                          registration_deadline TIMESTAMP,
                                          registration_link VARCHAR(255),
                                          format VARCHAR(255)
);

-- Таблица завершенных хакатонов
CREATE TABLE IF NOT EXISTS completed_hackathons (
                                                    id SERIAL PRIMARY KEY,
                                                    hackathon_id INT,
                                                    hackathon_name VARCHAR(255),
                                                    team_id INT,
                                                    team_name VARCHAR(255),
                                                    solution_description TEXT,
                                                    solution_link VARCHAR(255)
);

-- Таблица команд
CREATE TABLE IF NOT EXISTS teams (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL UNIQUE
);

-- Таблица участников команд
CREATE TABLE IF NOT EXISTS team_members (
                                            team_id INT,
                                            user_id INT,
                                            is_leader BOOLEAN DEFAULT false
);

-- Таблица хакатонов, в которых участвует команда
CREATE TABLE IF NOT EXISTS team_hackathons (
                                               team_id INT,
                                               hackathon_id INT
);

-- Таблица приглашенных пользователей в команду
CREATE TABLE IF NOT EXISTS invited_users (
                                             team_id INT references teams(id),
                                             user_id INT references users(id)
);

-- Таблица пользователей, которые хотят присоединиться к команде
CREATE TABLE IF NOT EXISTS users_want_team (
    user_id INT references users(id)
);

-- Таблица информации о пользователях
CREATE TABLE IF NOT EXISTS user_info (
                                         user_id INT references users(id),
                                         name VARCHAR(255),
                                         description TEXT,
                                         skills TEXT,
                                         github VARCHAR(255),
                                         team_id INT,
                                         registration_date TIMESTAMP
);

-- Таблица запросов на присоединение к команде
CREATE TABLE IF NOT EXISTS team_requests (
                                             team_id INT references teams(id),
                                             user_id INT references users(id)
);