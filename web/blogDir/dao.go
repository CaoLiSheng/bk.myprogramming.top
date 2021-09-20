package blogDir

import (
	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"github.com/huandu/go-sqlbuilder"
)

func (data *BlogDirRows) ListBlogDir(c *db.Core) {
	sql, args := sqlbuilder.NewSelectBuilder().Select("*").From("blog_dir").Build()
	rows, err := c.DB.QueryxContext(*c.Ctx, sql, args...)

	srv.IsPanic(err)

	for rows.Next() {
		var row BlogDirRow
		srv.IsPanic(rows.StructScan(&row))
		*data = append(*data, row)
	}
}

func (row *AddBlogDirRow) AddBlogDir(c *db.Core) {
	sql, args := sqlbuilder.NewInsertBuilder().InsertInto("blog_dir").Cols("pid", "name", "alias").Values(row.PID, row.Name, row.Alias).Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	id, err := result.LastInsertId()

	srv.IsPanic(err)

	row.ID = id
}

func (row *BlogDirRowId) RemoveBlogDir(c *db.Core) {
	sb := sqlbuilder.NewDeleteBuilder().DeleteFrom("blog_dir")
	sb.Where(sb.E("id", row.ID))
	sql, args := sb.Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic("要删除的数据不存在！")}
}

func (row *BlogDirRow) ModifyBlogDir(c *db.Core) {
	sb := sqlbuilder.NewUpdateBuilder().Update("blog_dir")
	sb.Set(sb.Assign("pid", row.PID), sb.Assign("name", row.Name), sb.Assign("alias", row.Alias)).Where(sb.E("id", row.ID))
	sql, args := sb.Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic("要修改的数据不存在！")}
}

func (row *BlogDirRowId) GetBlogDir(c *db.Core) *BlogDirRow {
	sb := sqlbuilder.NewSelectBuilder().Select("*").From("blog_dir")
	sb.Where(sb.E("id", row.ID))
	sql, args := sb.Build()
	rowx := c.DB.QueryRowxContext(*c.Ctx, sql, args...)

	res := new(BlogDirRow)
	srv.IsPanic(rowx.StructScan(res))

	return res
}
