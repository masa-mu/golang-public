DROP TABLE IF EXISTS `message`;

CREATE TABLE message (
  ID int NOT NULL PRIMARY KEY,
  Message varchar(1024) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT message (ID,Message) VALUES (12345,"test");
