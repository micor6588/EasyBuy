package item

import (
	"EasyBuy/commons"
	"encoding/json"
	"net/http"
	"strconv"
)

func ItemHandler() {
	commons.Router.HandleFunc("/showItem", showItemController)
}

//显示商品信息
func showItemController(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))
	datagrid := showItemService(page, rows)
	b, _ := json.Marshal(datagrid)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}
