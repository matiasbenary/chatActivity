CREATE TABLE IF NOT EXISTS message (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		value VARCHAR(255) NOT NULL,
		send_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		user_id VARCHAR(255) NOT NULL ,
		room_id VARCHAR(255) NOT NULL 
	);