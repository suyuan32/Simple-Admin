package user

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/captcha"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/pkg/enum"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewLoginLogic(r *http.Request, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if ok := captcha.Store.Verify(req.CaptchaId, req.Captcha, true); ok {
		userInfo, err := l.svcCtx.CoreRpc.Login(l.ctx,
			&core.LoginReq{
				TenantAccount: req.TenantAccount,
				Username:      req.Username,
				Password:      req.Password,
			})
		if err != nil {
			return nil, err
		}

		token, err := GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, userInfo.Tid, userInfo.Uid, time.Now().Unix(),
			l.svcCtx.Config.Auth.AccessExpire, int64(userInfo.RoleId))
		if err != nil {
			return nil, err
		}

		// add token into database
		expiredAt := time.Now().Add(time.Second * 259200).Unix()
		_, err = l.svcCtx.CoreRpc.CreateOrUpdateToken(l.ctx, &core.TokenInfo{
			Id:        0,
			CreatedAt: 0,
			Uuid:      userInfo.Uid,
			Token:     token,
			Source:    "core",
			Status:    1,
			ExpiredAt: expiredAt,
		})

		if err != nil {
			return nil, err
		}

		resp = &types.LoginResp{
			BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.lang, i18n.Success)},
			Data: types.LoginInfo{
				TenantId: userInfo.Tid,
				UserId:   userInfo.Uid,
				Token:    token,
				Expire:   uint64(expiredAt),
				Role: types.RoleInfoSimple{
					Value:    userInfo.RoleValue,
					RoleName: userInfo.RoleName,
				},
			},
		}
		return resp, nil
	} else {
		return nil, errorx.NewCodeError(enum.InvalidArgument, "login.wrongCaptcha")
	}
}

// tacc is tenant account
func GetJwtToken(secretKey, tuid string, uuid string, iat, seconds, roleId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["tenantId"] = tuid
	claims["userId"] = uuid
	claims["roleId"] = roleId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
