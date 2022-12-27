package user

import (
	"context"
	"fmt"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *core.UUIDReq) (*core.UserInfoResp, error) {
	// 查询用户
	u, err := l.svcCtx.DB.User.Query().Where(user.UUIDEQ(in.Uuid)).First(l.ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	// 查询租户
	t, err := u.QueryTenant().All(l.ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("can't found user's tenant", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	// 查询角色
	roleName, err := l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d", u.RoleID))
	roleValue, err := l.svcCtx.Redis.Hget("roleData", fmt.Sprintf("%d_value", u.RoleID))
	if err != nil {
		return nil, err
	}

	// 组装数据
	tenants := make([]*core.TenantInfo, 0)
	for _, v := range t {
		tenants = append(tenants, &core.TenantInfo{
			TenantId:   v.UUID,
			TenantName: v.Name,
		})
	}

	return &core.UserInfoResp{
		Tenants:   tenants,
		Nickname:  u.Nickname,
		Avatar:    u.Avatar,
		RoleId:    u.RoleID,
		RoleName:  roleName,
		RoleValue: roleValue,
		Mobile:    u.Mobile,
		Email:     u.Email,
		Status:    uint64(u.Status),
		Id:        u.ID,
		Username:  u.Username,
		Uuid:      u.UUID,
		CreatedAt: u.CreatedAt.Unix(),
		UpdatedAt: u.UpdatedAt.Unix(),
	}, nil
}
