CREATE TABLE IF NOT EXISTS todos (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS subs (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255),
    todo_id uuid references todos(id)
);