CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL
);

SELECT * FROM users WHERE id= $1;
insert into users (email, username, password) values ($1, $2, $3);
SELECT id, email, username, password FROM users WHERE email = $1;
select id, email, username, password from users;



/*
//db, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
// Autheticate user
// Create Admin User
// SELECT auth_user_add('admin2', 'admin2', 1);

// Change password for user
SELECT auth_user_change('user', 'userpassword', 0);

// Delete user
SELECT user_delete('user');

conn, err := db.Conn(context.Background())
// if err != nil { ... }
defer conn.Close()
err = conn.Raw(func (driverConn any) error {
	sqliteConn := driverConn.(*sqlite3.SQLiteConn)
	// ... use sqliteConn
})
// if err != nil { ... }
*/