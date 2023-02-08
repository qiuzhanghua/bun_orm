package main

import (
	"bun_orm/orm"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	ctx := context.Background()
	sqldb, err := sql.Open("mysql", "app:app@/app")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	defer db.Close()
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	db.RegisterModel((*orm.Author)(nil), (*orm.Book)(nil))
	//db.ResetModel(ctx, (*orm.Author)(nil), (*orm.Book)(nil))
	//fmt.Println(db.Dialect().Tables())
	if db.Dialect().Tables().ByName("authors") == nil {
		res, err := db.NewCreateTable().Model((*orm.Author)(nil)).Exec(ctx)
		fmt.Println(res, err)
	}
	if db.Dialect().Tables().ByName("books") == nil {
		res, err := db.NewCreateTable().Model((*orm.Book)(nil)).Exec(ctx)
		fmt.Println(res, err)
	}
	books := make([]orm.Book, 0)
	err = db.NewSelect().
		Model(&books).
		Relation("Author").
		Where("book.id = 1").
		Scan(ctx)

	if err != nil {
		panic(err)
	}

}
