package jwt

import (
	"fmt"
	proxy "github.com/dgrijalva/jwt-go"
	"time"
)

type BaseToken struct {
	UID  string `json:"uid"`     // id
	Type int32 `json:"type"` // 登录方式
}

type TokenReq struct {
	BaseToken
	Audience  string
	Subject   string
	Expire time.Duration
}

type Claims struct {
	BaseToken
	proxy.StandardClaims
}

type Handler struct {
	key    []byte        // key
	expire time.Duration // 有效时间
	issuer string
}

func New(key []byte, issuer string, expire time.Duration) Handler {

	return Handler{key: key, issuer: issuer, expire: expire}
}

func (h *Handler) GenToken(t TokenReq) (claims *Claims, expire time.Duration, signedToken string, success bool) {

	claims = &Claims{
		BaseToken:      t.BaseToken,
		StandardClaims: proxy.StandardClaims{
			Issuer: h.issuer,
		},
	}
	if len(t.Audience) > 0 {
		claims.StandardClaims.Audience = t.Audience
	}
	if len(t.Subject) > 0 {
		claims.StandardClaims.Subject = t.Subject
	}
	if t.Expire > 0 {
		expire = t.Expire
	} else {
		expire = h.expire
	}
	claims.ExpiresAt = time.Now().Add(expire).Unix()
	signedToken, success = h.CreateToken(claims)
	return
}

func (h *Handler) VerifyToken(signedToken string) (claims *Claims, success bool) {

	claims, success = h.ValidateToken(signedToken)
	if success {
		success = claims.VerifyIssuer(h.issuer, false)
	}
	return
}

// CreateToken create token
func (h *Handler) CreateToken(claims *Claims) (signedToken string, success bool) {
	token := proxy.NewWithClaims(proxy.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(h.key)
	if err != nil {
		return
	}

	success = true
	return
}

// ValidateToken validate token
func (h *Handler) ValidateToken(signedToken string) (claims *Claims, success bool) {
	token, err := proxy.ParseWithClaims(signedToken, &Claims{},
		func(token *proxy.Token) (interface{}, error) {
			if _, ok := token.Method.(*proxy.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return h.key, nil
		})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		success = true
		return
	}

	return
}
