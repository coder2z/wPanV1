package Authority

import (
	"github.com/gin-gonic/gin"
	"net/http"
	R "wPan/v1/Response"
)

type PolicyForm struct {
	Role   string `json:"role" form:"role" binding:"required"`
	Url    string `json:"url" form:"url" binding:"required"`
	Method string `json:"method" form:"method" binding:"required"`
}

func PolicyHandler(r *gin.Engine) *gin.Engine {

	//创建一个角色,并赋于权限
	r.POST("/api/v1/policy", func(c *gin.Context) {
		var Policy PolicyForm
		err := c.Bind(&Policy)
		if err != nil {
			R.Response(c, http.StatusUnprocessableEntity, R.MSG422, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if ok, _ := Enforcer.AddPolicy(Policy.Role, Policy.Url, Policy.Method); !ok {
			R.Error(c, R.POLICY_ADD_ERROR, nil)
			return
		} else {
			R.Ok(c, R.POLICY_ADD_OK, nil)
			return
		}
	})

	r.DELETE("/api/v1/policy", func(c *gin.Context) {
		var Policy PolicyForm
		err := c.Bind(&Policy)
		if err != nil {
			R.Response(c, http.StatusUnprocessableEntity, R.MSG422, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if ok, _ := Enforcer.RemovePolicy(Policy.Role, Policy.Url, Policy.Method); !ok {
			R.Error(c, R.POLICY_ERROR, nil)
			return
		} else {
			R.Ok(c, R.SUCCESSMSG, nil)
			return
		}
	})

	r.GET("/api/v1/policy", func(c *gin.Context) {
		list := Enforcer.GetPolicy()
		R.Ok(c, R.SUCCESSMSG, list)
	})

	r.POST("/api/v1/role", func(c *gin.Context) {
		_, _ = Enforcer.AddRoleForUser("admin_yzm", "admin")
		//s, _ := Enforcer.GetRolesForUser("admin")
		res, _ := Enforcer.GetUsersForRole("admin")
		R.Ok(c, R.SUCCESSMSG, res)
	})

	return r
}
