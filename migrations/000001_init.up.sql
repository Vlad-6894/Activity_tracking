CREATE SCHEMA app;

CREATE TABLE app.users (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL CHECK(char_length(full_name) BETWEEN 2 AND 100),
    age INTEGER NOT NULL CHECK(age>=0),
    google_refresh_token TEXT,
    steps_goal INTEGER CHECK(steps_goal>=0),
    rest_days INTEGER NOT NULL CHECK(rest_days>=0),
    streak INTEGER NOT NULL CHECK(streak>=0),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ

    CHECK(
        updated_at IS NULL 
        OR
        updated_at>=created_at
    )
);

CREATE TABLE app.activity (
    user_id INTEGER NOT NULL REFERENCES app.users(id) ON DELETE CASCADE,
    date TIMESTAMPTZ NOT NULL CHECK(date<=NOW()),
    steps INTEGER NOT NULL CHECK(steps>=0)
);

CREATE TABLE app.achievements (
    id SERIAL PRIMARY KEY,
    title VARCHAR(20) NOT NULL CHECK(char_length(title) BETWEEN 1 AND 20),
    description VARCHAR(1000)
);

CREATE TABLE app.achievement_user (
    user_id INTEGER NOT NULL REFERENCES app.users(id) ON DELETE CASCADE,
    achievement_id INTEGER NOT NULL REFERENCES app.achievements(id) ON DELETE CASCADE,
    count INTEGER NOT NULL CHECK(count>=0)
);

CREATE TABLE app.battlepass_rewards (
    min_streak INTEGER NOT NULL CHECK(min_streak>=0),
    reward VARCHAR(50) NOT NULL CHECK(char_length(reward) BETWEEN 1 AND 50)
);