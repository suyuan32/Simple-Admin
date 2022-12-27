package oauth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
	r      *http.Request
}

func NewOauthCallbackLogic(r *http.Request, svcCtx *svc.ServiceContext) *OauthCallbackLogic {
	return &OauthCallbackLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
		r:      r,
	}
}

func (l *OauthCallbackLogic) OauthCallback() (resp *types.CallbackResp, err error) {
	result, err := l.svcCtx.CoreRpc.OauthCallback(l.ctx, &core.CallbackReq{
		State: l.r.FormValue("state"),
		Code:  l.r.FormValue("code"),
	})

	if err != nil {
		return nil, err
	}

	token, err := user.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, result.Tid, result.Uid, time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire, int64(result.RoleId))

	// add token into database
	expiredAt := time.Now().Add(time.Second * 259200).Unix()
	_, err = l.svcCtx.CoreRpc.CreateOrUpdateToken(l.ctx, &core.TokenInfo{
		Uuid:      result.Uid,
		Token:     token,
		Source:    strings.Split(l.r.FormValue("state"), "-")[1],
		Status:    1,
		ExpiredAt: expiredAt,
	})

	if err != nil {
		return nil, err
	}

	return &types.CallbackResp{
		UserId: result.Uid,
		Role: types.RoleInfoSimple{
			RoleName: result.RoleName,
			Value:    result.RoleValue,
		},
		Token:  token,
		Expire: uint64(time.Now().Add(time.Second * 259200).Unix()),
	}, nil
}
