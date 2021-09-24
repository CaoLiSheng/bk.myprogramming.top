package blogCfg

import (
	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"github.com/huandu/go-sqlbuilder"
)

func (data *BlogCfgRows) ListBlogCfg(c *db.Core) {
	sql, args := sqlbuilder.NewSelectBuilder().Select("*").From("blog_cfg").Build()
	rows, err := c.DB.QueryxContext(*c.Ctx, sql, args...)
	srv.IsPanic(err)

	for rows.Next() {
		var row BlogCfgRow
		srv.IsPanic(rows.StructScan(&row))
		*data = append(*data, row)
	}
}

func (row *BlogCfgRowId) GetBlogCfg(c *db.Core) string {
	sb := sqlbuilder.NewSelectBuilder().Select("cfg").From("blog_cfg")
	sql, args := sb.Where(sb.E("id", row.ID)).Build()
	rowx := c.DB.QueryRowxContext(*c.Ctx, sql, args...)
	
	var res string
	srv.IsPanic(rowx.Scan(&res))

	return res
}

func (row *BlogCfgRow) ReplaceBlogCfg(c *db.Core) {
	sql, args := sqlbuilder.NewInsertBuilder().ReplaceInto("blog_cfg").Cols("id", "cfg").Values(row.ID, row.Cfg).Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(db.UpdateErr)}
}

func (row *BlogCfgRowId) RemoveBlogCfg(c *db.Core) {
	sb := sqlbuilder.NewDeleteBuilder().DeleteFrom("blog_cfg")
	sql, args := sb.Where(sb.E("id", row.ID)).Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(db.RemoveErr)}
}
