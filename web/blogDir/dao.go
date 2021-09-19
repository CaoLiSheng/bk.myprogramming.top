package blogDir

import (
	"errors"

	"bk.myprogramming.top/db"
	"github.com/huandu/go-sqlbuilder"
)

func ListBlogDir(c *db.Core) []blogDirRow {
	sql, args := sqlbuilder.NewSelectBuilder().Select("*").From("blog_dir").Build()
	rows, err := c.DB.QueryxContext(*c.Ctx, sql, args...)

	if err != nil {panic(err)}

	data := make([]blogDirRow, 0)
	for rows.Next() {
		var row blogDirRow
		err := rows.StructScan(&row)
		if err != nil {
			panic(err)
		}
		data = append(data, row)
	}

	return data
}

func addBlogDir(c *db.Core, row *addBlogDirRow) {
	sql, args := sqlbuilder.NewInsertBuilder().Cols("pid", "name", "alias").InsertInto("blog_dir").Values(row.PID, row.Name, row.Alias).Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	id, err := result.LastInsertId()

	if err != nil {panic(err)}

	row.ID = id
}

func removeBlogDir(c *db.Core, row *removeBlogDirRow) {
	sb := sqlbuilder.NewDeleteBuilder().DeleteFrom("blog_dir")
	sb.Where(sb.E("id", row.ID))
	sql, args := sb.Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	ra, err := result.RowsAffected()

	if err != nil {panic(err)}
	if ra <= 0 {panic(errors.New("要删除的数据不存在！"))}
}

func modifyBlogDir(c *db.Core, row *blogDirRow) {
	sb := sqlbuilder.NewUpdateBuilder().Update("blog_dir")
	sb.Set(sb.Assign("pid", row.PID), sb.Assign("name", row.Name), sb.Assign("alias", row.Alias)).Where(sb.E("id", row.ID))
	sql, args := sb.Build()
	result := c.DB.MustExecContext(*c.Ctx, sql, args...)
	ra, err := result.RowsAffected()

	if err != nil {panic(err)}
	if ra <= 0 {panic(errors.New("要修改的数据不存在！"))}
}