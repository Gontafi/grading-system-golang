
-- migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:6543/postgres?sslmode=disable" -verbose up

CREATE TABLE IF NOT EXISTS roles (
     id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY CHECK (id >= 0),
     name varchar(255) not null,
     status_id int DEFAULT 1
);

CREATE TABLE IF NOT EXISTS users (
    id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY CHECK (id >= 0),
    username varchar(255) unique not null,
    password_hash varchar(255) not null,
    name varchar(255) not null,
    created_at timestamptz,
    updated_at timestamptz
);

CREATE TABLE IF NOT EXISTS role_user (
     id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY CHECK (id >= 0),
     role_id int references roles(id) not null,
     user_id int references users(id) not null,
     CONSTRAINT uq_role_user UNIQUE (role_id, user_id)
);

CREATE TABLE IF NOT EXISTS lessons (
    id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY CHECK (id >= 0),
    name varchar(255)
);

CREATE TABLE IF NOT EXISTS student_lesson (
    id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY CHECK (id >= 0),
    student_id int references users(id) on delete cascade,
    lesson_id int references lessons(id) on delete cascade,
    CONSTRAINT uq_student_lesson_student_lesson UNIQUE (student_id, lesson_id)
);

CREATE TABLE IF NOT EXISTS marks (
    id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY CHECK (id >= 0),
    teacher_id int references users(id) on delete cascade,
    student_id int references users(id) on delete cascade,
    lesson_id int references lessons(id) on delete cascade,
    home_work_grade int not null check (home_work_grade >= 1 or home_work_grade < 6 ),
    attendance_grade int not null check (attendance_grade = 0 or attendance_grade = 1),
    date date not null,
    CONSTRAINT fk_student_lesson
        FOREIGN KEY (student_id, lesson_id)
        REFERENCES student_lesson(student_id, lesson_id)
);

CREATE INDEX idx_users_username ON users (username);
CREATE INDEX idx_marks_teacher_id ON marks (teacher_id);
CREATE INDEX idx_marks_student_id ON marks (student_id);
CREATE INDEX idx_marks_lesson_id ON marks (lesson_id);

INSERT INTO roles (name) VALUES ('teacher');
INSERT INTO roles (name) VALUES ('student');
INSERT INTO roles (name) VALUES ('admin');
INSERT INTO roles (name) VALUES ('guest');


--optional creates
--creating superuser admin
--username: admin
--password: admin
--TODO remove this queries
INSERT INTO users (username, password_hash, name, created_at, updated_at)
VALUES (
        'admin',
        '676a6462736a6b676466673133346b6a64736667626b6ad033e22ae348aeb5660fc2140aec35850c4da997',
        'Admin',
        NOW(),
        NOW()
    );
INSERT INTO role_user (role_id, user_id)
VALUES (3, (SELECT id FROM users WHERE username = 'admin'));
