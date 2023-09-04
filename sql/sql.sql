CREATE DATABASE IF NOT EXISTS api_go_lesson;
USE api_go_lesson;

    Drop TABLE IF EXISTS posts;
    Drop TABLE IF EXISTS followers;
    Drop TABLE IF EXISTS users;

Create TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers(
    userID int not null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    followerId int not null ,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    primary key (user_id, follower_id)
) ENGINE=INNODB;

Create TABLE posts(
    id int auto_increment primary key,
    title varchar(50) not null,
    content varchar(300) not null unique,

    authorId int not null,
    FOREIGN KEY (authorId)
    REFERENCES users(id)
    ON DELETE CASCADE,

    likes int default 0
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;