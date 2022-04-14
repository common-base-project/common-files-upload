/*
@Author : zggong
*/
package routers

// import (
// 	"fmt"
// 	"unicorn-files/handler/auth"

// 	"github.com/gin-gonic/gin"
// 	"github.com/spf13/viper"
// )

// // 用户组
// func GroupRouter(g *gin.Engine) {
// 	authRouterGroup := fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/group")
// 	authGroup := g.Group(authRouterGroup)
// 	{
// 		authGroup.POST("", auth.CreateGroupHandler)
// 		authGroup.PUT("/:id", auth.UpdateGroupHandler)
// 		authGroup.DELETE("/:id", auth.DeleteGroupHandler)
// 		authGroup.GET("", auth.GroupListHandler)
// 		authGroup.GET("/:id", auth.GroupDetailHandler)
// 	}
// }

// // 用户
// func UserRouter(g *gin.Engine) {
// 	authRouterUser := fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/user")
// 	authUser := g.Group(authRouterUser)
// 	{
// 		authUser.POST("", auth.CreateUserHandler)
// 		authUser.PUT("/:id", auth.UpdateUserHandler)
// 		authUser.DELETE("/:id", auth.DeleteUserHandler)
// 		authUser.GET("", auth.UserListHandler)
// 		authUser.GET("/:id", auth.UserDetailHandler)
// 		authUser.GET("/:id/dept-user", auth.DeptUserListHandler)
// 	}
// }

// // 用户和组关联
// func GroupUserRouter(g *gin.Engine) {
// 	authRouterGroupUser := fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/groupuser")
// 	authGroupUser := g.Group(authRouterGroupUser)
// 	{
// 		authGroupUser.GET("/:id", auth.UserGroupListHandler)
// 		authGroupUser.POST("/:id", auth.BindUserHandler)
// 	}
// }

// // 部门
// func DepartRouter(g *gin.Engine) {
// 	authRouterDepart := fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/depart")
// 	authDepart := g.Group(authRouterDepart)
// 	{
// 		authDepart.POST("", auth.CreateDepartHandler)
// 		authDepart.PUT("/:id", auth.UpdateDepartHandler)
// 		authDepart.DELETE("/:id", auth.DeleteDepartHandler)
// 		authDepart.GET("", auth.DepartListHandler)
// 		authDepart.GET("/:id", auth.DepartDetailHandler)
// 	}
// }
