package department

import (
	"context"
	"fmt"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dbfunc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type UpdateDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDepartmentLogic {
	return &UpdateDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDepartmentLogic) UpdateDepartment(in *core.DepartmentInfo) (*core.BaseResp, error) {

	// 获取部门的祖先标识
	cd := dbfunc.NewGenAncestorsLogic(l.ctx, l.svcCtx)
	ancestors := cd.Department_ancestors(in.Id)
	fmt.Printf("========%s\n", *ancestors)

	err := l.svcCtx.DB.Department.UpdateOneID(*in.Id).
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilSort(in.Sort).
		SetNotNilName(in.Name).
		SetNotNilAncestors(ancestors).
		SetNotNilLeader(in.Leader).
		SetNotNilPhone(in.Phone).
		SetNotNilEmail(in.Email).
		SetNotNilRemark(in.Remark).
		SetNotNilParentID(in.ParentId).
		Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
