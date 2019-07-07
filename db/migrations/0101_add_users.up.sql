CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT,
    identify varchar(255) NOT NULL,
    
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (id)
);