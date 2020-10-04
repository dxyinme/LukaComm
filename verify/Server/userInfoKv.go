package Server

import (
	"fmt"
	"github.com/dxyinme/LukaComm/KvDao"
	"github.com/dxyinme/LukaComm/util/CoHash"
	"github.com/dxyinme/LukaComm/util/MD5"
	"github.com/dxyinme/LukaComm/verify"
)

type UserInfoKv struct {
	dao *KvDao.RedisDao
}

func (u *UserInfoKv) Initial(newDao *KvDao.RedisDao) {
	u.dao = newDao
}

func (u *UserInfoKv) CheckUser(info *verify.UserInfo) error {
	password,err := u.dao.Get(info.Uid)
	if err != nil {
		return err
	}
	if string(password) != MD5.CalcMD5([]byte(info.Password)) {
		return fmt.Errorf("Wrong Password")
	}
	return nil
}

func (u *UserInfoKv) SignUpUser(info *verify.UserInfo) error {
	info.Uid = CoHash.NewUID().Uid
	err := u.dao.Set(info.Uid, []byte(MD5.CalcMD5([]byte(info.Password))))
	return err
}