package util

import (
	"crypto/md5"
	"fmt"
	"math/big"
	"testing"
)

func TestBuffer(t *testing.T) {
	src := []byte("dfkjfdkjfn反对法dkjfndjnfkjdnfkdjn")
	m := md5.Sum(src)
	fmt.Println(string(m[:]))
	bi := &big.Int{}
	var z [16]byte
	z[15] = 9
	// bi.SetBytes(m[:])
	bi.SetBytes(z[:])
	fmt.Println(bi.Text(16))
	fmt.Println(len("b5d679e3a81be14c54376104e0f00692"))
	// d, _ := x509.ParseCertificate(nil)
	// x509.CreateCertificate()

}
