package blogCfg

import (
	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"github.com/huandu/go-sqlbuilder"
)

func (data BlogCfgDict) ListBlogCfg(c *db.Core) {
	sql, args := sqlbuilder.NewSelectBuilder().Select("*").From("blog_cfg").Build()
	rows, err := c.DB.QueryxContext(*c.Ctx, sql, args...)
	srv.IsPanic(err)

	for rows.Next() {
		var row BlogCfgRow
		srv.IsPanic(rows.StructScan(&row))
		data[row.ID] = row.Cfg
	}
}

func (rows BlogCfgDictIDs) GetBlogCfg(c *db.Core) BlogCfgDict {
	ids := make([]interface{}, len(rows))
	for i:=0; i<len(rows); i++ {
		ids = append(ids, rows[i])
	}
	
	sb := sqlbuilder.NewSelectBuilder().Select("*").From("blog_cfg")
	sql, args := sb.Where(sb.In("id", ids...)).Build()
	rowsx, err := c.DB.QueryxContext(*c.Ctx, sql, args...)
	
	srv.IsPanic(err)
	
	data := make(BlogCfgDict)
	for rowsx.Next() {
		var row BlogCfgRow
		srv.IsPanic(rowsx.StructScan(&row))
		data[row.ID] = row.Cfg
	}

	return data
}

func (row *BlogCfgRow) ReplaceBlogCfg(c *db.Core) {
	sql, args := sqlbuilder.NewInsertBuilder().ReplaceInto("blog_cfg").Cols("id", "cfg").Values(row.ID, row.Cfg).Build()
	result := c.MustExec(sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(db.UpdateErr)}
}

func (row *BlogCfgRowId) RemoveBlogCfg(c *db.Core) {
	sb := sqlbuilder.NewDeleteBuilder().DeleteFrom("blog_cfg")
	sql, args := sb.Where(sb.E("id", row.ID)).Build()
	result := c.MustExec(sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(db.RemoveErr)}
}
