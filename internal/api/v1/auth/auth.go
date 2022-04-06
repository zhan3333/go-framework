package auth

import (
	"net/http"

	"go-framework/internal/domain/user"
	"go-framework/pkg/auth"
	"go-framework/pkg/lgo"
)

// Register godoc
// @Summary 注册新用户
// @Description 注册新用户
// @Produce  json
// @Param user body RegisterReq true "注册信息"
// @Success 200 {string} default
// @Router /api/v1/auth/register [post]
func Register(c *lgo.CustomContext) error {
	var (
		req RegisterReq
		err error
	)
	if err := c.Bind(&req); err != nil {
		return c.BadRequest(err.Error())
	}
	if isUsed, err := c.NewUser().IsEmailUsed(req.Email); err != nil {
		return err
	} else if isUsed {
		return lgo.NewHTTPError(http.StatusBadRequest, "邮箱已被使用")
	}
	// 调用领域
	params := user.StoreUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if _, err = c.NewUser().Store(params); err != nil {
		// 处理错误
		return err
	}
	return c.OK()
}

// Login godoc
// @Summary 登录
// @Description 登录
// @Produce  json
// @Param user body LoginReq true "登录"
// @Success 200 {object} LoginResp
// @Router /api/v1/auth/login [post]
func Login(c *lgo.CustomContext) error {
	var (
		req LoginReq
	)
	if err := c.Bind(&req); err != nil {
		return c.BadRequest(err.Error())
	}

	u, err := c.NewUser().FirstUserByEmail(req.Email)
	if err != nil {
		return err
	}
	if u == nil {
		return c.BadRequest("用户不存在")
	}
	if err = auth.Compare(u.Password, req.Password); err != nil {
		return c.BadRequest("密码不正确")
	}
	if token, err := c.JWT.Create(uint64(u.ID)); err != nil {
		return err
	} else {
		return c.OK(LoginResp{
			AccessToken: token.Token,
			Type:        token.Type,
			ExpiresAt:   token.ExpiresAt,
		})
	}
}

type MeResp struct {
	Name  string `json:"name"`
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}

// Me godoc
// @Summary 获取当前登录用户信息
// @Description 获取当前登录用户信息
// @Produce  json
// @Param Authorization header string true "accessToken"
// @Success 200 {object} MeResp
// @Router /api/v1/me [get]
func Me(c *lgo.CustomContext) error {
	u, err := c.NewUser().First(c.UserID)
	if err != nil {
		return err
	}
	if u == nil {
		return c.Unauthorized("用户不存在")
	}
	return c.OK(MeResp{
		Name:  u.Name,
		ID:    uint64(u.ID),
		Email: u.Email,
	})
}
