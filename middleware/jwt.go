package middleware

import (
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
)

var (
	jwtKey = []byte(utils.JwtKey)
)

// 自定义一个cliams

type MyClaims struct {
	Username string
	jwt.StandardClaims
}

// 生成token

func NewToken(Username string) (string, int) {
	// 有效期
	expireTime := time.Now().Add(10 * time.Hour)
	// 声明一个Claims
	SetClaims := MyClaims{
		Username,
		jwt.StandardClaims{ExpiresAt: expireTime.Unix(), Issuer: "ginblog"},
	}
	// 注意是 HS256 不是ES
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaims.SignedString(jwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// 验证token

func CheckToken(token string) (*MyClaims, int) {
	//下面这个函数是官方文档中提供的函数，用来校验token
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	// 检验
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

//jwt中间件
// 定义一个gin的中间件

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 固定写法
		tokenHandler := c.Request.Header.Get("Authorization")
		var code int
		// not exist
		if tokenHandler == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		// check the token format
		checkToken := strings.SplitN(tokenHandler, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		// check the validity of the token
		key, valid := CheckToken(checkToken[1])
		if valid == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		// if the token timeout
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_TIME_OUT
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
