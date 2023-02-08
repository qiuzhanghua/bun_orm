package orm

type Book struct {
	ID       int64 `bun:",pk,autoincrement"`
	AuthorID int64
	Author   Author `bun:"rel:belongs-to,join:author_id=id"`
}

type Author struct {
	ID int64 `bun:",pk,autoincrement"`
}
