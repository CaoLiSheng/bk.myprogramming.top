package blogTag

import (
	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"github.com/huandu/go-sqlbuilder"
)

func (data *BlogTags) ListBlogTags(c *db.Core) {
	sql, args := sqlbuilder.NewSelectBuilder().Select("*").From("blog_tag").Build()
	rows, err := c.DB.QueryxContext(*c.Ctx, sql, args...)

	srv.IsPanic(err)

	for rows.Next() {
		var row BlogTag
		srv.IsPanic(rows.StructScan(&row))
		*data = append(*data, row)
	}
}

func (row *AddBlogTag) AddBlogTag(c *db.Core) {
	sql, args := sqlbuilder.NewInsertBuilder().InsertInto("blog_tag").Cols("name").Values(row.Name).Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	id, err := result.LastInsertId()

	srv.IsPanic(err)

	row.ID = id
}

func (row *BlogTagID) RemoveBlogTag(c *db.Core) {
	sb := sqlbuilder.NewDeleteBuilder().DeleteFrom("blog_tag")
	sql, args := sb.Where(sb.E("id", row.ID)).Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(db.RemoveErr)}
}

func (row *BlogTag) ModifyBlogTag(c *db.Core) {
	sb := sqlbuilder.NewUpdateBuilder().Update("blog_tag")
	sql, args := sb.Set(sb.Assign("name", row.Name)).Where(sb.E("id", row.ID)).Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(db.UpdateErr)}
}
