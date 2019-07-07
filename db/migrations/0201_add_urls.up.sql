CREATE TABLE urls (
    id int NOT NULL AUTO_INCREMENT,
	
	user_id int NULL DEFAULT NULL,
    url varchar(255) NOT NULL,
	short varchar(255) NOT NULL,
	hits int NULL DEFAULT NULL,
    
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (id)
);