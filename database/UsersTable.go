package database

import (
	"log"
)

func CreateTableUser() (bool, error) {
	query := `CREATE TABLE users (
							id VARCHAR(255) PRIMARY KEY , 
							name VARCHAR(255) NOT NULL, 
							username VARCHAR(255) UNIQUE NOT NULL, 
							email VARCHAR(255) UNIQUE NOT NULL, 
							password VARCHAR(255) NOT NULL, 
							email_verified_at TIMESTAMP NULL,
							status INT DEFAULT 1, 
							created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
							updated_at TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
						)`

	if err := DB.Exec(query).Error; err != nil {
		log.Println("failed to execute query:", err)
		return false, err
	}

	log.Println("Table created successfully")
	return true, nil
}
