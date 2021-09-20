package blogDir

import (
	"errors"

	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"github.com/huandu/go-sqlbuilder"
)

func ListBlogDir(c *db.Core) []blogDirRow {
	sql, args := sqlbuilder.NewSelectBuilder().Select("*").From("blog_dir").Build()
	rows, err := c.DB.QueryxContext(*c.Ctx, sql, args...)

	srv.IsPanic(err)

	data := make([]blogDirRow, 0)
	for rows.Next() {
		var row blogDirRow
		srv.IsPanic(rows.StructScan(&row))
		data = append(data, row)
	}

	return data
}

func addBlogDir(c *db.Core, row *addBlogDirRow) {
	sql, args := sqlbuilder.NewInsertBuilder().Cols("pid", "name", "alias").InsertInto("blog_dir").Values(row.PID, row.Name, row.Alias).Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	id, err := result.LastInsertId()

	srv.IsPanic(err)

	row.ID = id
}

func removeBlogDir(c *db.Core, row *blogDirRowId) {
	sb := sqlbuilder.NewDeleteBuilder().DeleteFrom("blog_dir")
	sb.Where(sb.E("id", row.ID))
	sql, args := sb.Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(errors.New("要删除的数据不存在！"))}
}

func modifyBlogDir(c *db.Core, row *blogDirRow) {
	sb := sqlbuilder.NewUpdateBuilder().Update("blog_dir")
	sb.Set(sb.Assign("pid", row.PID), sb.Assign("name", row.Name), sb.Assign("alias", row.Alias)).Where(sb.E("id", row.ID))
	sql, args := sb.Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(errors.New("要修改的数据不存在！"))}
}

func getBlogDir(c *db.Core, row *blogDirRowId) *blogDirRow {
	sb := sqlbuilder.NewSelectBuilder().Select("*").From("blog_dir")
	sb.Where(sb.E("id", row.ID))
	sql, args := sb.Build()
	rowx := c.DB.QueryRowxContext(*c.Ctx, sql, args...)

	res := new(blogDirRow)
	srv.IsPanic(rowx.StructScan(res))

	return res
}
