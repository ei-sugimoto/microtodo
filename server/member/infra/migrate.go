package infra

func Migrate() {
	db := NewDB()
	defer db.Close()

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS members (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		panic(err)
	}

}
