package XrayAPI

import (
	"github.com/crossfw/Air-Universe/pkg/structures"
)

func (xrayCtl *XrayController) AddUsers(users *[]structures.UserInfo) (err error) {
	for _, u := range *users {
		switch u.Protocol {
		case "vmess":
			err = addVmessUser(*xrayCtl.HsClient, &u)
		case "trojan":
			err = addTrojanUser(*xrayCtl.HsClient, &u)
			//case "ss":
			//	err = addSSUser(*xrayCtl.HsClient, &u)
		}

		if err != nil {
			return err
		}
	}
	return
}

func (xrayCtl *XrayController) RemoveUsers(users *[]structures.UserInfo) (err error) {
	for _, u := range *users {
		err := removeUser(*xrayCtl.HsClient, &u)
		if err != nil {
			return err
		}
	}
	return
}

func (xrayCtl *XrayController) QueryUsersTraffic(users *[]structures.UserInfo) (usersTraffic *[]structures.UserTraffic, err error) {
	usersTraffic = new([]structures.UserTraffic)
	var ut structures.UserTraffic

	for _, u := range *users {
		ut.Id = u.Id
		ut.Up, err = queryUserTraffic(*xrayCtl.SsClient, u.Tag, "up")
		ut.Down, err = queryUserTraffic(*xrayCtl.SsClient, u.Tag, "down")
		// when a user used this node, post traffic data
		if ut.Up+ut.Down > 0 {
			*usersTraffic = append(*usersTraffic, ut)
		}
		if err != nil {
			return
		}
	}
	return
}
