use UnityWebPlayer;

create table games(
	id INT auto_increment NOT NULL,
  title VARCHAR(128) NOT NULL,
  auther VARCHAR(128) NOT NULL,
  created_at DATA,
  deleted_at DATA,

  PRIMARY KEY (id)
);

INSERT INTO games (title, auther, created_at) 
VALUES ("test1", "user1", CURDATE()),
       ("test2", "user2", CURDATE());


