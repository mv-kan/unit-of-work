CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL
);
CREATE TABLE IF NOT EXISTS todos (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id uuid NOT NULL references users(id) ON DELETE CASCADE 
);
CREATE TABLE IF NOT EXISTS subs (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    todo_id uuid NOT NULL references todos(id) ON DELETE CASCADE
);