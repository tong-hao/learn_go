package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
)

var pubKey string = "-----BEGIN RSA PUBLIC KEY-----\nMIGJAoGBANxHSR4gyaZX7uet7fGzCwqhUcvTYTQpakPDihLkW+e4Ib4kBCd84Ldb\ndI7cziiOk1e6NjMEqnsjs6hOZ0tTPXrE7eKHMR9vwIJW08O0pyGw275DSQLVbP5k\nmlWs0W/pGfnO+sh3apaeLHF86qzkFmS6Q6pjYGTw3jCQy6bOP0F3AgMBAAE=\n-----END RSA PUBLIC KEY-----"
var prvKey string = "-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQDcR0keIMmmV+7nre3xswsKoVHL02E0KWpDw4oS5FvnuCG+JAQn\nfOC3W3SO3M4ojpNXujYzBKp7I7OoTmdLUz16xO3ihzEfb8CCVtPDtKchsNu+Q0kC\n1Wz+ZJpVrNFv6Rn5zvrId2qWnixxfOqs5BZkukOqY2Bk8N4wkMumzj9BdwIDAQAB\nAoGAV/9uWUfV5sr4GLul88mH3q5FY/zEtzbYScvi69soT/CCIlh3BGNhzj4N2Uii\nXxdeC0zLfCQgCuNfURxJLZFKXHQ3CNO7JYt+Lcmj2KBkkS0hS0LBuZ7uqPYQ5hN+\nZBj8oVREkZFhdv23F4Iwvm/6PwcBkEkJ3rjH3jcUIl6bonECQQD52p3IXJxdpuWq\nFX4rI5eqnfDyh4MAclOXf4D9t9++J8YgawGfvR0keGnMPWtEt66gpPx0++vjvDCk\nguznmrqpAkEA4bJtDZBCHlDJeE6VcPqNa5swTGoeYw1OJZGJC/KuAbvwsXd61AKy\nV81D3Jjzbg/rY17B2mKPBqxrFN6XKXTPHwJAby5sJduoLTh2XHBB+5pUBDVSIepR\nTiKRtgmj8cMfyjNSw9w6FcYGsNLwaVUvZZ3DFHM2cCwmNOnqT1p/ZzSAOQJAXENb\nb0sErG3sHmFJmBjkzRNwyBwtdeKPiq4W7ypy1cSlnXaxYJAFpf0Ee96OzPR3DnVD\nG+pke57qW2qvMsRMkwJBANWIWbF3oSEtA8YnzOSiiHWtvAiTF/ablQQ8iybm2lhy\nqp4ks45ROPNhaY3/GaIIiSYXaSIzqO71/Tn4em+N1k8=\n-----END RSA PRIVATE KEY-----"


func RsaEncrypt2(data []byte) (string,error) {
	log.Println("001")
	bts := []byte(pubKey)
	block, _ := pem.Decode(bts)
	if block == nil {
		return "", errors.New("block is null")
	}
	log.Println("002")
	pubInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println("003")
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pubInterface, data)
	if err != nil {
		log.Println(err)
		return "",err
	}
	log.Println("004")
	ret := b64.StdEncoding.EncodeToString(ciphertext)
	return ret,nil
}

// 公钥加密
func RsaEncrypt(data []byte) (string, error) {
	//解密pem格式的公钥
	bts := []byte(pubKey)
	block, _ := pem.Decode(bts)
	if block == nil {
		return "", errors.New("public key is error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pubInterface, data)
	if err != nil {
		return "", err
	}
	//return string(ciphertext), nil

	ret := b64.StdEncoding.EncodeToString(ciphertext)
	return ret, nil

}

// 私钥解密
func RsaDecrypt(ciphertext []byte, keyBytes []byte) (string, error) {
	// base64 decode
	b64Decode, err := b64.StdEncoding.DecodeString(string(ciphertext))

	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return "", errors.New("private key is error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, []byte(b64Decode))
	if err != nil {
		return "", err
	}
	return string(data), err

}



func RsaDecrypt2(ciphertext []byte) (string, error) {
	b64Decode, err := b64.StdEncoding.DecodeString(string(ciphertext))

	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		return "", errors.New("private key is error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return "", err
	}
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, []byte(b64Decode))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(data), err
}

func main() {
	data := "您好"
	fmt.Println(pubKey)
	fmt.Println(prvKey)

	encodeStr, _ := RsaEncrypt([]byte(data))
	fmt.Println("公钥加密后的数据：", encodeStr)
	decodeStr, _ := RsaDecrypt([]byte(encodeStr), []byte(prvKey))
	fmt.Println("私钥解密后的数据：", decodeStr)

	encodeStr2, _ := RsaEncrypt2([]byte(data))
	fmt.Println("公钥加密后的数据：", encodeStr2)
	//decodeStr, _ := RsaDecrypt2([]byte(encodeStr))
	//fmt.Println("私钥解密后的数据：", decodeStr)

}
