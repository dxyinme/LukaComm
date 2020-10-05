package Server

import (
	"context"
	"github.com/dxyinme/LukaComm/util/Const"
	"github.com/dxyinme/LukaComm/verify"
)

type Server struct {
	userInfoKv *UserInfoKv
}

func (s *Server) Login(ctx context.Context, req *verify.LoginReq) (*verify.LoginRsp, error) {
	var result string
	err := s.userInfoKv.CheckUser(req.User)
	if err == nil {
		result = Const.LoginSuccess
	} else {
		result = err.Error()
	}
	return &verify.LoginRsp{
		MessCode: req.MessCode,
		LoginResult: result,
	},nil
}

func (s *Server) SignUp(ctx context.Context, req *verify.SignUpReq) (*verify.SignUpRsp, error) {
	err := s.userInfoKv.SignUpUser(req.SignUser)
	if err != nil {
		return &verify.SignUpRsp{
			SignUpResult:  err.Error(),
			SignUserReply: nil,
		}, nil
	} else {
		return &verify.SignUpRsp{
			SignUpResult:  Const.SignUpSuccess,
			SignUserReply: req.SignUser,
		}, nil
	}
}