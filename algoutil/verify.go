// Package algoutil contain some scaffold algo
package algoutil

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	mathrand "math/rand"
	"runtime"
	"strconv"
	"time"

	"github.com/vvisun/utls/errutil"
	"github.com/vvisun/utls/randutil"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	//"triple/modules/security"
	"github.com/vvisun/utls/crypto"
)

// MD5String md5 digest in string
func MD5String(plain string) string {
	cipher := MD5([]byte(plain))
	return hex.EncodeToString(cipher)
}

// MD5 md5 digest
func MD5(plain []byte) []byte {
	md5Ctx := md5.New()
	md5Ctx.Write(plain)
	cipher := md5Ctx.Sum(nil)
	return cipher[:]
}

// CallSite the caller's file & line
func CallSite() interface{} {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	return string(file + ":" + strconv.FormatInt(int64(line), 10))
}

// GenRSAKey gen a rsa key pair, the bit size is 512
func GenRSAKey() (privateKey, publicKey string, err error) {
	//public gen the private key
	privKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		return "", "", err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privKey)
	privateKey = base64.StdEncoding.EncodeToString(derStream)

	//gen the public key
	pubKey := &privKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", "", err
	}
	publicKey = base64.StdEncoding.EncodeToString(derPkix)
	return privateKey, publicKey, nil
}

// RSAEncrypt encrypt data by rsa
func RSAEncrypt(plain []byte, pubKey string) ([]byte, int32) {
	buf, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return nil, errutil.Sys_UnKnown
	}
	p, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return nil, errutil.Sys_UnKnown
	}
	if pub, ok := p.(*rsa.PublicKey); ok {
		bs, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plain)
		if err != nil {
			return nil, errutil.Sys_UnKnown
		} else {
			return bs, errutil.Succ
		}
	}
	return nil, errutil.Sys_UnKnown
}

// RsaDecrypt decrypt data by rsa
func RSADecrypt(cipher []byte, privKey string) ([]byte, int32) {
	if cipher == nil {
		return nil, errutil.Sys_UnKnown
	}
	buf, err := base64.StdEncoding.DecodeString(privKey)
	if err != nil {
		return nil, errutil.Sys_UnKnown
	}
	priv, err := x509.ParsePKCS1PrivateKey(buf)
	if err != nil {
		return nil, errutil.Sys_UnKnown
	}
	bs, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cipher) //RSA解密算法
	if err != nil {
		return nil, errutil.Sys_UnKnown
	}
	return bs, errutil.Succ
}

// Sign with database appsecret string(base64 encode)
func Sign(plain []byte, privKey string) (string, int32) {
	if plain == nil {
		return "", errutil.Sys_UnKnown
	}
	buf, err := base64.StdEncoding.DecodeString(privKey)
	if err != nil {
		return "", errutil.Sys_UnKnown
	}
	priv, err := x509.ParsePKCS1PrivateKey(buf)
	if err != nil {
		return "", errutil.Sys_UnKnown
	}
	return crypto.Sign(priv, plain)
}

// Verify with database appkey string(base64 encode)
func Verify(pubKey string, data []byte, sign string) int32 {

	buf, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return errutil.Sys_UnKnown
	}
	p, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return errutil.Sys_UnKnown
	}
	if pub, ok := p.(*rsa.PublicKey); ok {
		return crypto.Verify(pub, data, sign)
	}
	return errutil.Sys_InvalidArg
}

func VerifyRSAWithMD5(pubKey string, data []byte, sign string) int32 {
	buf, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return errutil.Sys_UnKnown
	}
	p, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return errutil.Sys_UnKnown
	}
	if pub, ok := p.(*rsa.PublicKey); ok {
		return crypto.VerifyRSAWithMD5(pub, data, sign)
	}
	return errutil.Sys_UnKnown
}

// 生成随机字符串
func RandStr(strlen int) string {
	randutil.RandSeed()
	data := make([]byte, strlen)
	var num int
	for i := 0; i < strlen; i++ {
		num = mathrand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = mathrand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}

func Utf8ToGBK(utf8str string) string {
	result, _, _ := transform.String(simplifiedchinese.GBK.NewEncoder(), utf8str)
	return result
}

// TimeRange adjust the time range.
func TimeRange(start, end int64) (int64, int64) {
	if start < 0 {
		start = 0
	}
	if end < 0 || end > time.Now().Unix() {
		end = time.Now().Unix()
	}

	if start > end {
		start, end = end, start
	}

	return start, end
}
