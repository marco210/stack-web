-- Run mysql docker
docker run --name mysql -d     -p 3306:3306     -e MYSQL_ROOT_PASSWORD=password     --restart unless-stopped    mysql:8


CREATE database book_example;

use book_example;

CREATE TABLE books (
    ID int,
    Title varchar(255),
    Author varchar(255)    
);

Insert into books(ID,Title,Author) values(1,"Sherlock Holmes","Conan Doyle");


docker build -t web-stack .
docker run --name myapp -p 9000:9000 --restart unless-stopped  web-stack

docker tag web-stack marco210/web-stack:latest
docker push marco210/web-stack:latest


SELECT user,authentication_string,plugin,host FROM mysql.user;
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'Current-Root-Password';
FLUSH PRIVILEGES;

INSERT INTO mysql.user (Host, User, Password) VALUES ('%', 'root', password('YOURPASSWORD'));
GRANT ALL ON *.* TO 'root'@'%' WITH GRANT OPTION;

CREATE USER 'user'@'%' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;