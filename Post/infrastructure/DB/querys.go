package psqlPost

const SqlCreatePost = `INSERT INTO posts (id, title, body, owner_email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING title`
const SqlFindPostByName = `SELECT id, title, body, owner_email, created_at, updated_at FROM posts WHERE title = $1`
const SlqGetPostsByEmail = `SELECT id, title, body, owner_email, created_at, updated_at FROM posts WHERE owner_email = $1`
