package main

import (
	"github.com/wazsmwazsm/QueryBuilder/builder"
	"log"
)

func main() {
	sb := builder.NewSQLBuilder()

	sql, err := sb.Table("`test`").
		Select("`test`.`name`", "`test`.`age`", "`test2`.`teacher`").
		JoinRaw("LEFT JOIN `test2` ON `test`.`class` = `test2`.`class` AND `test`.`num` = ?", 2333).
		Where("`test`.`age`", ">=", 18).
		GetQuerySQL()
	if err != nil {
		log.Fatal(err)
	}

	params := sb.GetQueryParams()

	log.Println(sql)    // SELECT `test`.`name`,`test`.`age`,`test2`.`teacher` FROM `test` LEFT JOIN `test2` ON `test`.`class` = `test2`.`class` AND `test`.`num` = ? WHERE `test`.`age` >= ?
	log.Println(params) // [2333 18]
}
