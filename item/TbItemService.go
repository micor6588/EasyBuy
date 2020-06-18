package item

import "EasyBuy/commons"

func showItemService(page, rows int) (e *commons.Datagrid) {
	ts := selByPageDao(rows, page)
	if ts != nil {
		e = new(commons.Datagrid)
		e.Rows = ts
		return
	}
	return nil
}
