package configuration_center

import (
	"context"
)

type DrivenConfigurationCenter interface {
	DataUseType(ctx context.Context) (*DataUserTypeRes, error)
	GetChildrenDepartment(ctx context.Context, departmentId string) (*DepartmentObjectsList, error)
	GetUserRoles(ctx context.Context) ([]*UserRoleItem, error)
}
