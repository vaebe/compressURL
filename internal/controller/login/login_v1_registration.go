package login

import (
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"

	"compressURL/api/login/v1"
)

func (c *ControllerV1) Registration(ctx context.Context, req *v1.RegistrationReq) (res *v1.RegistrationRes, err error) {
	userinfo := entity.User{
		Email:       req.Email,
		WxId:        req.WxId,
		Password:    req.Password,
		NickName:    req.NickName,
		AccountType: "01",
		Role:        "01",
	}

	id, err := service.User().Create(ctx, userinfo)

	if err != nil {
		return nil, err
	}
	return &v1.RegistrationRes{Id: id}, nil
}
