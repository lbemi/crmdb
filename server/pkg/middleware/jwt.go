package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	"time"
)

func JWTAuth(rc *rctx.ReqCtx) error {

	permissions := rc.RequirePermission
	if !permissions.NeedToken {
		return nil
	}

	request := rc.Request.Request
	tokenStr := request.Header.Get("Authorization")

	if tokenStr == "" {
		return restfulx.TokenInvalid
	}

	token, claims, err := util.ParseToken(tokenStr)
	if err != nil || isInBlacklist(tokenStr) {
		return restfulx.TokenInvalid
	}
	rc.Set("id", claims.Id)
	rc.Set("token", token)

	if !permissions.NeedToken {
		return nil
	}

	enforcer := core.V1.Policy().GetEnforce()
	// 用户ID
	uid, isExit := rc.Get("id")
	if !isExit {
		return restfulx.TokenInvalid
	}

	p := request.URL.Path
	m := request.Method
	ok, err := enforcer.Enforce(uid, p, m)
	log.Logger.Infof("permission: %v -- %v --%v", uid, p, m)
	restfulx.ErrIsNil(err, "")
	if err != nil {
		return restfulx.ServerErr
	}
	if !ok {
		return restfulx.PermissionErr
	}

	return nil
}
func getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + util.MD5([]byte(tokenStr))
}

func JoinBlackList(token *jwt.Token) {
	nowUnix := time.Now().Unix()
	timer := time.Duration(token.Claims.(*util.CustomClaims).ExpiresAt-nowUnix) * time.Second
	core.V1.Redis().SetNX(getBlackListKey(token.Raw), nowUnix, timer)
}

func isInBlacklist(tokenStr string) bool {
	joinUnixStr, err := core.V1.Redis().Get(getBlackListKey(tokenStr)).Result()
	if err != nil || joinUnixStr == "" {
		return false
	}
	return true
}
