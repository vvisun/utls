package algoutil

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/vvisun/utls/errutil"
)

func pemParse(data []byte, pemType string) ([]byte, int32) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errutil.Sys_UnKnown
	}
	if pemType != "" && block.Type != pemType {
		return nil, errutil.Sys_InvalidArg
	}
	return block.Bytes, errutil.Succ
}

func ParsePrivateKey(data []byte) (*rsa.PrivateKey, int32) {
	pemData, err := pemParse(data, "RSA PRIVATE KEY")
	if err != errutil.Succ {
		return nil, err
	}
	rlt, e := x509.ParsePKCS1PrivateKey(pemData)
	if e != nil {
		return nil, errutil.Sys_UnKnown
	}
	return rlt, errutil.Succ
}

func LoadPrivateKey(privKeyPath string) (*rsa.PrivateKey, int32) {
	certPEMBlock, err := os.ReadFile(privKeyPath)
	if err != nil {
		return nil, errutil.Sys_UnKnown
	}
	return ParsePrivateKey(certPEMBlock)
}
