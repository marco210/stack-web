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


docker build -t web-stack
docker run -p 9000:9000 -tid web-stack