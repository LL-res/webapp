package models

type Post struct {
	ID       int64  `json:"id" db:"post_id"`
	Title    string `json:"title" db:"title"`
	Content  string `json:"content" db:"content"`
	AuthorID int64  `json:"author_id" db:"author_id"`
}
