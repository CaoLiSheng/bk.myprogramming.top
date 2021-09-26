package blogDir

import (
	"fmt"

	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"github.com/huandu/go-sqlbuilder"
)

func (data *BlogDirRows) ListBlogDir(c *db.Core) {
	sql, args := sqlbuilder.NewSelectBuilder().Select("*").From("blog_dir").OrderBy("pid").Asc().Build()
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
	result := c.MustExec(sql, args...)
	id, err := result.LastInsertId()

	srv.IsPanic(err)

	row.ID = id
}

func (row *BlogDirRowId) RemoveBlogDir(c *db.Core) {
	sb := sqlbuilder.NewDeleteBuilder().DeleteFrom("blog_dir")
	sb.Where(sb.E("id", row.ID))
	sql, args := sb.Build()
	result := c.MustExec(sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(db.RemoveErr)}
}

func (row *BlogDirRow) ModifyBlogDir(c *db.Core) {
	sb := sqlbuilder.NewUpdateBuilder().Update("blog_dir")
	sb.Set(sb.Assign("pid", row.PID), sb.Assign("name", row.Name), sb.Assign("alias", row.Alias)).Where(sb.E("id", row.ID))
	sql, args := sb.Build()
	result := c.MustExec(sql, args...)
	ra, err := result.RowsAffected()

	srv.IsPanic(err)
	if ra <= 0 {panic(db.UpdateErr)}
}

func (row *BlogDirRowId) GetBlogDirDict(c *db.Core) map[int64]BlogDirRow {
	sb1 := sqlbuilder.NewSelectBuilder()
	sb1.Select("id", "@pathnodes:=IF( pid=0, ',0,', CONCAT( IF( LOCATE( CONCAT('|',pid,':'), @pathall ) > 0, SUBSTRING_INDEX( SUBSTRING_INDEX( @pathall, CONCAT('|',pid,':'), -1), '|', 1 ), @pathnodes ) ,pid, ',' ) ) paths", "@pathall:=CONCAT( @pathall, '|', id, ':', @pathnodes, '|' )")
	sb1.From("blog_dir").SQL(", (SELECT @pathnodes:='', @pathall:='') AS a").OrderBy("pid", "id")
	
	sb2 := sqlbuilder.NewSelectBuilder()
	sb2.Select(fmt.Sprintf("CONCAT(paths, '%d,') paths", row.ID)).From(sb2.BuilderAs(sb1, "b")).Where(sb2.E("b.id", row.ID))

	sb := sqlbuilder.NewSelectBuilder()
	sql, args := sb.Select("d.*").From(sb.As("blog_dir", "d"), sb.BuilderAs(sb2, "c")).Where(sb.G("INSTR(c.paths, CONCAT(',',d.id,','))", 0)).Build()

	rows, err := c.DB.QueryxContext(*c.Ctx, sql, args...)

	srv.IsPanic(err)

	data := make(map[int64]BlogDirRow)
	for rows.Next() {
		var datum BlogDirRow
		srv.IsPanic(rows.StructScan(&datum))
		data[datum.ID] = datum
	}

	return data
}

func (row *BlogDirRowId) GetBlogDirPath(c *db.Core) string {
	data := row.GetBlogDirDict(c)

	datum := data[row.ID]
	res := "/" + datum.Name
	for datum.PID != 0 {
		datum = data[datum.PID]
		res = "/" + datum.Name + res
	}

	return res
}
