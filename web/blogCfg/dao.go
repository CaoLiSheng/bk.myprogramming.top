package blogCfg

import (
	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"github.com/huandu/go-sqlbuilder"
)

func listBlogCfg(c *db.Core) []blogCfgRow {
	sql, args := sqlbuilder.NewSelectBuilder().Select("cfg").From("blog_cfg").Build()
	rows, err := c.DB.QueryxContext(*c.Ctx, sql, args...)
	srv.IsPanic(err)

	data := make([]blogCfgRow, 0)
	for rows.Next() {
		var row blogCfgRow
		srv.IsPanic(rows.StructScan(&row))
		data = append(data, row)
	}
	
	return data
}

func GetBlogCfg(c *db.Core, id string) string {
	sb := sqlbuilder.NewSelectBuilder().Select("cfg").From("blog_cfg")
	sql, args := sb.Where(sb.E("id", id)).Build()
	rowx := c.DB.QueryRowxContext(*c.Ctx, sql, args...)
	
	var res string
	srv.IsPanic(rowx.Scan(&res))

	return res
}
