package user

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/tenant"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/utils"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *core.LoginReq) (*core.LoginResp, error) {
	// user, err := l.svcCtx.DB.User.Query().Where(user.UsernameEQ(in.Username)).First(l.ctx)
	user, err := l.svcCtx.DB.Tenant.Query().Where(tenant.AccountEQ(in.TenantAccount)).QueryUsers().First(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			logx.Errorw("user not found", logx.Field("username", in.Username))
			return nil, status.Error(codes.InvalidArgument, "login.userNotExist")
		}
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, status.Error(codes.Internal, i18n.DatabaseError)
	}
	tenants, err := user.QueryTenant().Where(tenant.AccountEQ(in.TenantAccount)).All(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			logx.Errorw("tentant not found", logx.Field("tenant", in.TenantAccount))
			return nil, status.Error(codes.InvalidArgument, "login.tenantNotExist")
		}
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, status.Error(codes.Internal, i18n.DatabaseError)
	}

	if ok := utils.BcryptCheck(in.Password, user.Password); !ok {
		logx.Errorw("wrong password", logx.Field("detail", in))
		return nil, status.Error(codes.InvalidArgument, "login.wrongUsernameOrPassword")
	}

	roleName, value, err := GetRoleInfo(user.RoleID, l.svcCtx.Redis, l.svcCtx.DB, l.ctx)
	if err != nil {
		return nil, err
	}

	logx.Infow("log in successfully", logx.Field("UUID", user.UUID))
	return &core.LoginResp{
		Tid:       tenants[0].UUID,
		Uid:       user.UUID,
		RoleValue: value,
		RoleName:  roleName,
		RoleId:    user.RoleID,
	}, nil
}

func GetRoleInfo(roleId uint64, rds *redis.Redis, db *ent.Client, ctx context.Context) (roleName, roleValue string, err error) {
	if s, err := rds.Hget("roleData", strconv.Itoa(int(roleId))); err != nil {
		roleData, err := db.Role.Query().All(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				logx.Error("fail to find any roles")
				return "", "", status.Error(codes.NotFound, i18n.TargetNotFound)
			}
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return "", "", status.Error(codes.NotFound, i18n.TargetNotFound)
		}

		for _, v := range roleData {
			err = rds.Hset("roleData", strconv.Itoa(int(v.ID)), v.Name)
			err = rds.Hset("roleData", fmt.Sprintf("%d_value", v.ID), v.Value)
			err = rds.Hset("roleData", fmt.Sprintf("%d_status", v.ID), strconv.Itoa(int(v.Status)))
			if err != nil {
				logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
				return "", "", statuserr.NewInternalError(i18n.RedisError)
			}
			if v.ID == roleId {
				roleName = v.Name
				roleValue = v.Value
			}
		}
	} else {
		roleName = s
		roleValue, err = rds.Hget("roleData", fmt.Sprintf("%d_value", roleId))
		if err != nil {
			logx.Error("fail to find the role data")
			return "", "", status.Error(codes.NotFound, i18n.TargetNotFound)
		}
	}
	return roleName, roleValue, nil
}
