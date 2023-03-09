package myrsa

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/utils"
	"github.com/lishimeng/go-log"
)

var (
	CONANT_CACHE_KEY = "cache_privateKey" // 存放redis键
	CONANT_VO_KEY    = "pubkey"           // 返回前台key值
)

// GetPubKey 将服务器中保存的密钥base64编码放在redis中
func GetPubKey(ctx iris.Context) {
	param := make(map[string]interface{})
	// 获取缓存实例
	cache := app.GetCache()

	var value string
	err := cache.Get(CONANT_CACHE_KEY, &value)
	if err != nil {
		log.Debug("privateKey doesn't exist, So we are gonna create one and put it into redis cache now!")
		// 创建，公钥、密钥
		privateKey, publicKey := utils.GenerateRSAKey()

		// 将密钥导出 PEM格式
		privPEM := utils.ExportPrivateKeyAsPEM(privateKey)
		// 将公钥导出 PEM格式
		pubPEM := utils.ExportPublicKeyAsPEM(publicKey)

		// 打印PEM
		log.Debug("%s", privPEM)
		log.Debug("%s", pubPEM)

		// 密钥best64编码，存入redis，这里只需要存放密钥即可，可以通过密钥获取公钥，前端只返回公钥(用于加密)
		privKeyEncodeString := utils.Base64EncodeString(privPEM)
		cache.Set(CONANT_CACHE_KEY, privKeyEncodeString)
	}
	// 取出 redis存放的 priKey
	cache.Get(CONANT_CACHE_KEY, &value)
	// 密钥
	log.Debug("Now we've got a privateKey from redis cache encoding by BASE64 %s", value)

	//best64解码
	privKeyDecodeString, err := utils.Base64DecodeString(value)
	if err != nil {
		param["code"] = 500
		param["error"] = err.Error()
		common.ResponseJSON(ctx, param)
		return
	}

	// 通过密钥获取公钥
	pubKey := utils.GetPublicKeyFromPriKey(privKeyDecodeString)
	// 将公钥转为 pem格式
	pubPEM := utils.ExportPublicKeyAsPEM(pubKey)
	// best64编码返回前台
	encodeString := utils.Base64EncodeString(pubPEM)

	// 封装响应
	param[CONANT_VO_KEY] = encodeString
	param["code"] = 200
	ctx.JSON(param)
}

// RefreshPubKey 刷新密钥、公钥
func RefreshPubKey(ctx iris.Context) {
	param := make(map[string]interface{})
	log.Debug("start refreshing pubkey...")
	// 获取缓存实例
	cache := app.GetCache()

	// 创建，公钥、密钥
	privateKey, _ := utils.GenerateRSAKey()
	// 将密钥到处 PEM格式
	privPEM := utils.ExportPrivateKeyAsPEM(privateKey)
	// best64编码，存入redis
	privKeyEncodeString := utils.Base64EncodeString(privPEM)
	// 放入缓存
	err := cache.Set(CONANT_CACHE_KEY, privKeyEncodeString)
	if err != nil {
		param["code"] = 500
		param["error"] = err.Error()
		ctx.JSON(param)
		return
	}

	param["code"] = 200
	ctx.JSON(param)
}
