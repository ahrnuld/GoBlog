A little GO experiment project

To run MySQL:
```bash
docker run --name go-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=dbname -d mysql:latest
```
To run Adminer: 
```bash
docker run --link go-mysql:db -p 8080:8080 adminer  
```
To run the application
```bash
go run .\main.go
```

Test data query:
```SQL
CREATE TABLE `post` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `posted_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `content` varchar(8000) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `post` (`id`, `title`, `posted_at`, `content`) VALUES
(1,	'test',	'2022-12-29 10:35:13',	'test content'),
(2,	'This is a test 2',	'2022-12-29 19:54:52',	'<p>asdasdasd</p>\r\n<p>&nbsp;</p>\r\n<h4>test heading</h4>'),
(3,	'Another test 2',	'2022-12-29 19:11:49',	'<p>This time with some HTML</p>'),
(4,	'A test',	'2022-12-29 19:50:24',	'<p>With some <strong>HTML </strong>this <em>time</em></p>\r\n<blockquote>\r\n<p>And a quote</p>\r\n</blockquote>\r\n<h1>And a heading</h1>\r\n<p>&nbsp;</p>\r\n<p>some more</p>');
```