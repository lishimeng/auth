package password

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/go-log"
	"golang.org/x/crypto/bcrypt"
	"hash"
)

func Generate(u model.AuthUser, plaintext string) (p string, err error) {

	bs, err := bcrypt.GenerateFromPassword([]byte(genPlainPassword(u, plaintext)), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	p = string(bs)
	return
}

func Compare(u model.AuthUser, plaintext string) (success bool) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(genPlainPassword(u, plaintext)))
	success = err == nil
	return
}

func genPlainPassword(u model.AuthUser, plaintext string) (s string) {
	ts := u.CreateTime.Format(common.DateFormatter)
	salt := fmt.Sprintf("%d.%d.%s_.%s", u.Id, u.Id, ts, plaintext)
	return digest(plaintext, salt, 10)
}

func digest(plaintext, salt string, loop int) (dig string) {
	dig = _digest(sha512.New(), []byte(plaintext), []byte(salt), loop)
	return
}
func _digest(dig hash.Hash, plaintext, salt []byte, loop int) (txt string) {
	var tmp = plaintext
	for i := 0; i < loop; i++ {
		dig.Write(tmp)
		dig.Write(salt)
		tmp = dig.Sum(nil)
	}
	return base64.StdEncoding.EncodeToString(tmp)
}

type PlaintextKey struct {
	Pri    string
	Pub    string
	OldPri string
	OldPub string
}

func GetPasswordKey() (cert PlaintextKey, err error) {
	err = app.GetCache().Get(_genPasswordCacheKey(), &cert)
	return
}

func _genPasswordCacheKey() (k string) {
	k = common.CacheSignInKey
	return
}

func RefreshPasswordKey() {
	var err error
	var rsaKey []byte
	var rsaPub []byte
	var keyDigest []byte
	var rsaOldKey string
	var rsaOldPub string

	var oldCert PlaintextKey

	// get old key
	err = app.GetCache().Get(_genPasswordCacheKey(), &oldCert)
	if err != nil {
		log.Info("old cache not exist")
	} else {
		rsaOldKey = oldCert.Pri
		rsaOldPub =  oldCert.Pub
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, common.CacheRsaBits)
	if err != nil {
		log.Info("gen rsa key failed")
		log.Info(err)
		return
	}
	priDer := x509.MarshalPKCS1PrivateKey(privateKey)
	rsaKey = genPem(priDer, "RSA PRIVATE KEY")

	pubDer := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	rsaPub = genPem(pubDer, "RSA PUBLIC KEY")

	digest := sha1.New()
	_, err = digest.Write(rsaPub)
	if err != nil {
		log.Info("gen rsa key digest failed")
		log.Info(err)
		return
	}

	keyDigest = digest.Sum(nil)

	var cacheKey = PlaintextKey{
		Pri:    string(rsaKey),
		Pub:    string(rsaPub),
		OldPri: rsaOldKey,
		OldPub: rsaOldPub,
	}
	err = app.GetCache().Set(string(keyDigest), cacheKey)
	if err != nil {
		log.Info("cache password key failed")
		log.Info(err)
	}
}

func genPem(der []byte, t string) []byte {

	contentPem := pem.EncodeToMemory(&pem.Block{
		Type:    t,
		Headers: nil,
		Bytes:   der,
	})
	return contentPem
}
