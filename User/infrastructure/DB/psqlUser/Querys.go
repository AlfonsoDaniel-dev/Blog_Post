package psqlUser

const SqlCreateUserQuery = `INSERT INTO users(id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
const SqlGetUser = `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`
const SqlGetUserPosts = `SELECT id, title, body, owner_id,  created_at, updated_at FROM posts WHERE owner_id = $1`
const SqlUpdateUserName = `UPDATE users SET name = $1 WHERE email = $2`
const SqlUpdateUserEmail = `UPDATE users SET email = $1 WHERE id = $2`
const SqlAdminGetAllUsers = `SELECT id, name, email, created_at, updated_at FROM users`
const SqlGetUserByName = `SELECT id, name, email, created_at, updated_at FROM users WHERE name = $1`
