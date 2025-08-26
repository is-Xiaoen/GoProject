package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user"
	"gorm.io/gorm"
)

// CreateUser 创建用户
func (i *UserServiceImpl) CreateUser(
	ctx context.Context,
	req *user.CreateUserRequest) (
	*user.User, error) {
	// 1. 校验用户参数
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// 2. 生成一个User对象(ORM对象)
	ins := user.NewUser(req)

	if err := datasource.DBFromCtx(ctx).
		Create(ins).
		Error; err != nil {
		return nil, err
	}

	// 4. 返回结果
	return ins, nil
}

// DeleteUser 删除用户
func (i *UserServiceImpl) DeleteUser(
	ctx context.Context,
	req *user.DeleteUserRequest,
) (*user.User, error) {
	u, err := i.DescribeUser(ctx,
		user.NewDescribeUserRequestById(req.Id))
	if err != nil {
		return nil, err
	}

	return u, datasource.DBFromCtx(ctx).
		Where("id = ?", req.Id).
		Delete(&user.User{}).
		Error
}

// QueryUser 查询用户列表
func (i *UserServiceImpl) QueryUser(
	ctx context.Context,
	req *user.QueryUserRequest) (
	*types.Set[*user.User], error) {
	set := types.New[*user.User]()

	query := datasource.DBFromCtx(ctx).Model(&user.User{})

	// 查询总量
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.
		Order("created_at desc").
		Offset(int(req.ComputeOffset())).
		Limit(int(req.PageSize)).
		Find(&set.Items).
		Error
	if err != nil {
		return nil, err
	}

	return set, nil
}

// DescribeUser 查询用户详情
func (i *UserServiceImpl) DescribeUser(
	ctx context.Context,
	req *user.DescribeUserRequest) (
	*user.User, error) {

	query := datasource.DBFromCtx(ctx)

	// 1. 构造我们的查询条件
	switch req.DescribeBy {
	case user.DESCRIBE_BY_ID:
		query = query.Where("id = ?", req.DescribeValue)
	case user.DESCRIBE_BY_USERNAME:
		query = query.Where("user_name = ?", req.DescribeValue)
	}

	// SELECT * FROM `users` WHERE username = 'admin' ORDER BY `users`.`id` LIMIT 1
	ins := &user.User{}
	if err := query.First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("user %s not found", req.DescribeValue)
		}
		return nil, err
	}

	// 数据库里面存储的就是Hash
	ins.SetIsHashed()
	return ins, nil
}
