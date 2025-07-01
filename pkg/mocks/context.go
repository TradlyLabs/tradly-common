package mocks

import (
	"context"

	"github.com/golang/mock/gomock"
)

const mOCK_CONTROLLER_KEY = "gmock.Controller"

func WithController(ctx context.Context, ctrl *gomock.Controller) context.Context {
	return context.WithValue(ctx, mOCK_CONTROLLER_KEY, ctrl)
}

func GetController(ctx context.Context) *gomock.Controller {
	i := ctx.Value(mOCK_CONTROLLER_KEY)
	ctrl, ok := i.(*gomock.Controller)
	if !ok {
		panic("controller is not exists")
	}
	return ctrl
}
