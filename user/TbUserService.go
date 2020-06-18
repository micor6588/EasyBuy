package user

import "EasyBuy/commons"

// LoginService 实现用户登录相关逻辑
func LoginService(un, pwd string) (er commons.EgoResult) {
	u := SelByUnPwdDao(un, pwd)
	if u != nil {
		// 如果u!=nil,即成功
		er.Status = 200
	} else {
		er.Status = 400
	}
	return
}
