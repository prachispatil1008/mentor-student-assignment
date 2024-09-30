CREATE TABLE mentors (
    uid VARCHAR PRIMARY KEY,
    name VARCHAR,
    capacity INT
);

CREATE TABLE rules (
    id VARCHAR PRIMARY KEY,
    name VARCHAR,
    exp TEXT,
    groupid VARCHAR,
    priority INT
);

CREATE TABLE students (
    sid VARCHAR PRIMARY KEY,
    name VARCHAR,
    lang VARCHAR,
    standard VARCHAR,
    mentorid VARCHAR
);
