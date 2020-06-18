package item

import "EasyBuy/commons"

func showItemService(page, rows int) (e *commons.Datagrid) {
	ts := selByPageDao(rows, page)
	if ts != nil {
		e = new(commons.Datagrid)
		e.Rows = ts //当前页显示的数据
		e.Total = selCount()
		return
	}
	return nil
}
