package user

import "EasyBuy/commons"

func LoginService(un, pwd string) (er commons.EgoResult) {
	u := SelByUnPwdDao(un, pwd)
	if u != nil {
		er.Status = 200
	} else {
		er.Status = 400
	}
	return
}
