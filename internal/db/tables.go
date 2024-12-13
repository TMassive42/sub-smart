package database

import "log"

const USERS_TABLE string = `CREATE TABLE
    IF NOT EXISTS USERS (
        user_id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
        username varchar(20) NOT NULL UNIQUE,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

const TOKENS = `
    CREATE TABLE IF NOT EXISTS TOKENS(
        id INTEGER PRIMARY KEY ,
        session_token VARCHAR(255) NOT NULL,
        csrf_token VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        expires_at TIMESTAMP NOT NULL,
        user_id VARCHAR(255),
        FOREIGN KEY (user_id) REFERENCES USERS(id)
    );
`

var statements = []string{USERS_TABLE, TOKENS}

func (m *UserModel) InitTables() {
	for _, statement := range statements {
		stmt, err := m.DB.Prepare(statement)
		if err != nil {
			log.Println(err)
			return
		}
		if _, err := stmt.Exec(); err != nil {
			log.Println(err.Error())
		}
		stmt.Close()
	}
}
