package admin

import (
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
)

func SetupAdminGroup(rg *gin.RouterGroup, appName string) {
	gAdmin := rg.Group("admin")
	gAdmin.Use(
		web.HandlerAdapter(middleware.Authn), middleware.AppScope(appName),
	)
	setupDashboard(gAdmin)

	setupAppGroup(gAdmin)

	setupTokenGroup(gAdmin)

	setupStaffGroup(gAdmin)

	setupGroupGroup(gAdmin)

	setupGroupshipGroup(gAdmin)

	setupPermissionGroup(gAdmin)

	setupPolicyGroup(gAdmin)

	setupDeptGroup(gAdmin)

	setupLogGroup(gAdmin)
}
