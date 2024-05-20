package psqlUser

const SqlCreatePost = `INSERT INTO posts (id, title, body, owner_email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING title`
const SqlFindPostByName = `SELECT id, title, body, owner_email, created_at, updated_at FROM posts WHERE title = $1`
const SqlGetUserPosts = `SELECT id, title, body, owner_email,  created_at, updated_at FROM posts WHERE owner_email = $1`
const SqlGetPostsById = `SELECT id, title, body, owner_email, created_at, updated_at FROM posts WHERE id = $1`
const SqlUpdatePostTitle = `UPDATE posts SET title = $1, updated_at = $2 WHERE owner_email = $2`
const SqlUpdatePostBody = `UPDATE posts SET body = $1, updated_at = $2 WHERE owner_email = $2`
const SqlFindPostId = `SELECT id From posts WHERE owner_email = $1 AND title = $2`
const SqlDeletePost = `DELETE id, title, body, owner_email, created_at, updated_at FROM posts WHERE title = $1 AND owner_email = $2`
