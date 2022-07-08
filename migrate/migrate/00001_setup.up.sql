CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    fullname VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS todos (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255),
    username VARCHAR(255) NOT NULL  references users(username) ON DELETE CASCADE 
);
CREATE TABLE IF NOT EXISTS subs (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255),
    todo_id uuid NOT NULL references todos(id) ON DELETE CASCADE
);