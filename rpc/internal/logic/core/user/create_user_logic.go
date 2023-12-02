package user

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *core.UserInfo) (*core.BaseUUIDResp, error) {
	result, err := l.svcCtx.DB.User.Create().
		SetNotNilUsername(in.Username).
		SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*in.Password))).
		SetNotNilNickname(in.Nickname).
		SetNotNilEmail(in.Email).
		SetNotNilMobile(in.Mobile).
		SetNotNilAvatar(in.Avatar).
		AddRoleIDs(in.RoleIds...).
		SetNotNilHomePath(in.HomePath).
		SetNotNilDescription(in.Description).
		SetNotNilDepartmentID(in.DepartmentId).
		AddPositionIDs(in.PositionIds...).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
