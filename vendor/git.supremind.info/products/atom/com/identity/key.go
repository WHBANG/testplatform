package identity

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/pkg/errors"
)

func LoadRSAPrivateKey(filename string) (*rsa.PrivateKey, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read private key file failed")
	}
	block, _ := pem.Decode([]byte(file))
	if err != nil {
		return nil, errors.New("decode pem data failed")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "parse RSA key failed")
	}
	return key, nil
}

func LoadRSAPublicKey(filename string) (*rsa.PublicKey, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read public key file failed")
	}
	block, _ := pem.Decode([]byte(file))
	if err != nil {
		return nil, errors.New("decode pem data failed")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "parse RSA key failed")
	}
	pub, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, errors.Wrap(err, "not a valid rsa public key")
	}
	return pub, nil
}
