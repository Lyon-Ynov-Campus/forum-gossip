CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    avatar TEXT,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    publication_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS likes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    UNIQUE(user_id, post_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    publication_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


SELECT * FROM users WHERE id= ?;
INSERT INTO users (email, username, password) values (?, ?, ?);
SELECT id, email, username, password FROM users WHERE email = ?;
SELECT posts.*, users.username FROM posts 
JOIN users ON posts.user_id = users.id 
WHERE posts.id = ?;

SELECT posts.*, users.username,
        COUNT(DISTINCT comments.id) AS nb_comments,
        COUNT(DISTINCT likes.id) AS nb_likes
FROM posts
JOIN users ON posts.user_id = users.id
LEFT JOIN comments ON posts.id = comments.post_id
LEFT JOIN likes ON posts.id = likes.post_id
WHERE users.username LIKE '%' || ? || '%'
GROUP BY posts.id
ORDER BY posts.publication_date DESC;

DELEETE FROM users WHERE id = ?;
//dans table users avatar = pdp a faire apres avec des images

