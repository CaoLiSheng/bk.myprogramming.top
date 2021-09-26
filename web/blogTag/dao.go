package blogTag

import (
	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"github.com/huandu/go-sqlbuilder"
)

func (data *BlogTagRows) ListBlogTags(c *db.Core) {
	sql, args := sqlbuilder.NewSelectBuilder().Select("*").From("blog_tag").Build()
	rows, err := c.DB.QueryxContext(*c.Ctx, sql, args...)

	srv.IsPanic(err)

	for rows.Next() {
		var row BlogTagRow
		srv.IsPanic(rows.StructScan(&row))
		*data = append(*data, row)
	}
}

func (row *AddBlogTagRow) AddBlogTag(c *db.Core) {
	sql, args := sqlbuilder.NewInsertBuilder().InsertInto("blog_tag").Cols("name", "alias").Values(row.Name, row.Alias).Build()
	result := c.MustExec(sql, args...)
	id, err := result.LastInsertId()

	srv.IsPanic(err)

	row.ID = id
}

func (row *BlogTagRowID) RemoveBlogTag(c *db.Core) {
	sb := sqlbuilder.NewDeleteBuilder().DeleteFrom("blog_tag")
	sql, args := sb.Where(sb.E("id", row.ID)).Build()
	result := c.MustExec(sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(db.RemoveErr)}
}

func (row *BlogTagRow) ModifyBlogTag(c *db.Core) {
	sb := sqlbuilder.NewUpdateBuilder().Update("blog_tag")
	sql, args := sb.Set(sb.Assign("name", row.Name), sb.Assign("alias", row.Alias)).Where(sb.E("id", row.ID)).Build()
	result := c.MustExec(sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(db.UpdateErr)}
}
