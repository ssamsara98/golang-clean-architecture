package bootstrap

import (
	"github.com/ssamsara98/golang-clean-architecture/src/api/controllers"
	"github.com/ssamsara98/golang-clean-architecture/src/api/middlewares"
	"github.com/ssamsara98/golang-clean-architecture/src/api/routes"
	"github.com/ssamsara98/golang-clean-architecture/src/api/services"
	"github.com/ssamsara98/golang-clean-architecture/src/helpers"
	"github.com/ssamsara98/golang-clean-architecture/src/infrastructure"
	"github.com/ssamsara98/golang-clean-architecture/src/lib"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	lib.Module,
	infrastructure.Module,
	helpers.Module,
	services.Module,
	controllers.Module,
	middlewares.Module,
	routes.Module,
)
