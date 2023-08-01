package controllers

import "github.com/google/wire"

var Set = wire.NewSet(
	NewUserController,
	NewBlogController,
)
