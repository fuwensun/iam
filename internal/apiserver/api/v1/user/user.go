// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	srvv1 "github.com/marmotedu/iam/internal/apiserver/service/v1"
	"github.com/marmotedu/iam/internal/apiserver/store"
)

// UserHandler create a user handler used to handle request for user resource.
type UserHandler struct {
	// sfw 高层模块依赖低层模块的接口(抽象)
	srv   srvv1.Service // sfw Service interface
	store store.Factory // sfw Factory interface
}

// NewUserHandler creates a user handler.
func NewUserHandler(store store.Factory) *UserHandler {
	return &UserHandler{
		srv:   srvv1.NewService(store),
		store: store,
	}
}
