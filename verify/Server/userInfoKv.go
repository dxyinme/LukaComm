package Server

import (
	"fmt"
	"github.com/dxyinme/LukaComm/KvDao"
	"github.com/dxyinme/LukaComm/util/CoHash"
	"github.com/dxyinme/LukaComm/util/MD5"
	"github.com/dxyinme/LukaComm/verify"
	"time"
)

type UserInfoKv struct {
	PasswordDao             *KvDao.RedisDao
	LoginStatusDao          *KvDao.RedisDao
	UserLoginStatusLiveTime time.Duration
}

func (u *UserInfoKv) CheckUser(req *verify.LoginReq) error {
	password, err := u.PasswordDao.Get(req.User.Uid)
	if err != nil {
		return err
	}
	if string(password) != MD5.CalcMD5([]byte(req.User.Password)) {
		return fmt.Errorf("Wrong Password")
	}
	err = u.LoginStatusDao.Set(req.User.Uid, []byte(req.MessCode))
	defer func() {
		go func() {
			select {
			case <-time.After(u.UserLoginStatusLiveTime):
				messCodePre, _ := u.LoginStatusDao.Get(req.User.Uid)
				if string(messCodePre) == req.MessCode {
					_ = u.LoginStatusDao.Delete(req.User.Uid)
				}
				// 如果登录状态中的messCode仍然是当前的
				// 就删除登录状态中的在线状态. 否则不删除
			}
		}()
	}()
	return nil
}

func (u *UserInfoKv) SignUpUser(info *verify.UserInfo) error {
	info.Uid = CoHash.NewUID().Uid
	err := u.PasswordDao.Set(info.Uid, []byte(MD5.CalcMD5([]byte(info.Password))))
	return err
}
