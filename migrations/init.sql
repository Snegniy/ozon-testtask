CREATE TABLE IF NOT EXISTS links
		(
			shortlink  varchar(10)  PRIMARY KEY,
			baselink   text 	   NOT NULL UNIQUE
		);