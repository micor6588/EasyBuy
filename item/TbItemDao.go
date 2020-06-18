package item

import (
	"EasyBuy/commons"
	"database/sql"
	"fmt"
)

/*
rows:每页显示的条数
page:当前第几页
*/
func selByPageDao(rows, page int) []TbItem {
	//第一个表示:从哪条开始查询,0算起  第二个:查询几个
	r, err := commons.Dql("select * from tb_item limit ?,?", rows*(page-1), rows)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	ts := make([]TbItem, 0)
	for r.Next() {
		var t TbItem
		var s sql.NullString
		//如果直接使用t.Barcode由于数据库中列为Null导致填充错误
		r.Scan(&t.Id, &t.Title, &t.SellPoint, &t.Price, &t.Num, &s, &t.Image, &t.Cid, &t.Status, &t.Created, &t.Updated)
		t.Barcode = s.String
		ts = append(ts, t)
	}
	commons.CloseConn()
	return ts
}

//查询总条数
/*
如果返回值为<0表示查询失败
*/
func selCount() (count int) {
	rows, err := commons.Dql("select count(*) from tb_item")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	rows.Next()
	rows.Scan(&count)
	commons.CloseConn()
	return
}
