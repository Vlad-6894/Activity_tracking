CREATE SCHEMA school;

CREATE TABLE school.users (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1 CHECK(version>0),
    name VARCHAR(50) NOT NULL CHECK (char_length(name) BETWEEN 3 AND 50),
    password VARCHAR(20) NOT NULL CHECK (char_length(password) BETWEEN 8 AND 20),
    money BIGINT NOT NULL CHECK(money>=0),
    xp BIGINT NOT NULL CHECK(xp>0),
    level INTEGER NOT NULL CHECK(level>0),

    UNIQUE(name)
);

CREATE TABLE school.questions (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1 CHECK(version>0),
    name VARCHAR(100) NOT NULL CHECK(char_length(name) BETWEEN 1 AND 100),
    description VARCHAR(3000) NOT NULL CHECK(char_length(name) BETWEEN 10 AND 3000),
    module INTEGER NOT NULL REFERENCES school.modules(id),
    min_xp INTEGER NOT NULL CHECK(min_xp>0),
    max_xp INTEGER NOT NULL CHECK(min_xp>0)
);

CREATE TABLE school.tasks_hints (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1 CHECK(version>0),
    description VARCHAR(100) NOT NULL CHECK(char_length(name) BETWEEN 10 AND 100),
    task_id INTEGER NOT NULL REFERENCES school.tasks(id)
);

CREATE TABLE school.completed_questions (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1 CHECK(version>0),
    user_id INTEGER PRIMARY KEY REFERENCES school.users(id),
    question_id INTEGER NOT NULL REFERENCES school.questions(id)
);

CREATE TABLE school.tasks (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1 CHECK(version>0),
    name VARCHAR(100) NOT NULL CHECK(char_length(name) BETWEEN 1 AND 100),
    description VARCHAR(3000) NOT NULL CHECK(char_length(name) BETWEEN 10 AND 3000),
    module INTEGER NOT NULL REFERENCES school.modules (id),
    min_xp INTEGER NOT NULL CHECK(min_xp>0),
    max_xp INTEGER NOT NULL CHECK(min_xp>0)
);

CREATE TABLE school.tasks_hints (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1 CHECK(version>0),
    description VARCHAR(100) NOT NULL CHECK(char_length(name) BETWEEN 10 AND 100),
    task_id INTEGER NOT NULL REFERENCES school.tasks(id),
);

CREATE TABLE school.completed_tasks (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1 CHECK(version>0),
    user_id INTEGER REFERENCES school.users(id),
    task_id INTEGER NOT NULL REFERENCES school.tasks(id),
);

CREATE TABLE school.modules (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1 CHECK(version>0),
    name VARCHAR(50) NOT NULL CHECK(char_length(name) BETWEEN 1 AND 50),
    description VARCHAR(2000) NOT NULL CHECK(char_length(name) BETWEEN 10 AND 2000),
    module INTEGER NOT NULL REFERENCES school.modules (id),
    min_xp INTEGER NOT NULL CHECK(min_xp>0),
    max_xp INTEGER NOT NULL CHECK(max_xp>0)
);

CREATE TABLE school.Achievement (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1 CHECK(version>0),
    name VARCHAR(20) CHECK(char_length(name) BETWEEN 1 AND 20),
    description VARCHAR(100) CHECK(char_length(DESCRIPTION) BETWEEN 1 AND 100),
    user_id INTEGER REFERENCES school.users(id)
)