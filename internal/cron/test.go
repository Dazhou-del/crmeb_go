package cron

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/server"
	"fmt"
	"time"
)

type Test struct {
	ctx    context.Context
	svcCtx *server.SvcContext
}

func NewTest(svcCtx *server.SvcContext) *Test {
	ctx := context.Background()

	return &Test{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Test) Exec() {
	fmt.Println(time.Now())
	fmt.Println("测试")
	first, err := l.svcCtx.Gen.User.WithContext(l.ctx).Where(l.svcCtx.Gen.User.UID.Eq(int64(1))).Debug().First()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("first", first)

	var user model.User
	err = l.svcCtx.Repo.UserRepository.QueryOne(l.ctx, "uid = ?", []interface{}{1}, &user)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("user", user)

}
