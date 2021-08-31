DROP IF EXISTS chat;
CREATE database chat;
use chat;

CREATE TABLE `user` (
    id int primary key auto_increment,
    username varchar(33) not null,
    password varchar(60) not null
);

CREATE TABLE `room` (
    id int primary key auto_increment,
    ownerID int not null,
    enterKey int not null,
    foreign key (ownerID) references user (id),
);

CREATE TABLE `message` (
  id int primary key auto_increment,
  owner int not null,
  roomID int not null,
  sendTime datetime not null,
  content message_text
);