package regex

import (
	"fmt"
	"testing"
)

func TestR(t *testing.T) {

	var va = GetCookieAuthValue("BAILING_MANAGE_LOGIN_SUCCESS=%7B%22id%22%3A1974%2C%22userId%22%3A%22191668116%22%2C%22loginName%22%3A%22zhangliqun%22%2C%22userName%22%3A%22%E5%BC%A0%E7%AB%8B%E7%BE%A4%22%2C%22passWord%22%3A%228b5c1a5e07626bbbaba2e4d59e9942ed%22%2C%22status%22%3A0%2C%22roleId%22%3A%2244%22%2C%22lastLogin%22%3A%22Jan+17%2C+2022+8%3A58%3A13+AM%22%2C%22province%22%3A%22%22%2C%22city%22%3A%22%22%2C%22district%22%3A%22%22%2C%22role%22%3A%7B%22id%22%3A12%2C%22roleId%22%3A%2244%22%2C%22roleName%22%3A%22%E4%BA%92%E8%81%94%E7%BD%91%E5%8C%BB%E9%99%A2%E7%AE%A1%E7%90%86%E5%91%98%22%2C%22roleCode%22%3A%22hlwyygly%22%2C%22createTime%22%3A%22Dec+25%2C+2020+8%3A42%3A29+PM%22%2C%22updateTime%22%3A%22Dec+25%2C+2020+8%3A42%3A29+PM%22%7D%2C%22phoneNo%22%3A%2217853687363%22%2C%22hospitalCodes%22%3A%5B%5D%2C%22hospitalUser%22%3A%220%22%2C%22whetherDoctor%22%3A%221%22%2C%22createTime%22%3A%22Aug+26%2C+2021+9%3A58%3A19+AM%22%2C%22updateTime%22%3A%22Jan+17%2C+2022+8%3A50%3A55+AM%22%2C%22roleCode%22%3A%22hlwyygly%22%2C%22siteIds%22%3A%5B%5D%7D; %7B%22id%22%3A1974%2C%22userId%22%3A%22191668116%22%2C%22loginName%22%3A%22zhangliqun%22%2C%22userName%22%3A%22%E5%BC%A0%E7%AB%8B%E7%BE%A4%22%2C%22passWord%22%3A%228b5c1a5e07626bbbaba2e4d59e9942ed%22%2C%22status%22%3A0%2C%22roleId%22%3A%2244%22%2C%22lastLogin%22%3A%22Jan+17%2C+2022+8%3A58%3A13+AM%22%2C%22province%22%3A%22%22%2C%22city%22%3A%22%22%2C%22district%22%3A%22%22%2C%22role%22%3A%7B%22id%22%3A12%2C%22roleId%22%3A%2244%22%2C%22roleName%22%3A%22%E4%BA%92%E8%81%94%E7%BD%91%E5%8C%BB%E9%99%A2%E7%AE%A1%E7%90%86%E5%91%98%22%2C%22roleCode%22%3A%22hlwyygly%22%2C%22createTime%22%3A%22Dec+25%2C+2020+8%3A42%3A29+PM%22%2C%22updateTime%22%3A%22Dec+25%2C+2020+8%3A42%3A29+PM%22%7D%2C%22phoneNo%22%3A%2217853687363%22%2C%22hospitalCodes%22%3A%5B%5D%2C%22hospitalUser%22%3A%220%22%2C%22whetherDoctor%22%3A%221%22%2C%22createTime%22%3A%22Aug+26%2C+2021+9%3A58%3A19+AM%22%2C%22updateTime%22%3A%22Jan+17%2C+2022+8%3A50%3A55+AM%22%2C%22roleCode%22%3A%22hlwyygly%22%2C%22siteIds%22%3A%5B%5D%7D")
	fmt.Println(va)
}
