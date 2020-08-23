package pubConverter

/*
This script uses version bytes as described in SLIP-132
https://github.com/satoshilabs/slips/blob/master/slip-0132.md
*/

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

var xpub = []byte("0488b21e")
var ypub = []byte("049d7cb2")
var zpub = []byte("04b24746")

func Xpub(pub string) string {
	data := base58.Decode(pub)
	// rune splice
	data = data[4:]
	// verfiyyy Hex and base64Encode
	db, err := decodeHex(xpub)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
	}

	f := base64Encode(db)
	fmt.Printf("%s", f)

	//bytes.join or bytes.buffer.writedString()
	pubX := append(f, data...)
	return base58.Encode(pubX)
}

func Ypub(pub string) string {
	data := base58.Decode(pub)
	// rune splice
	data = data[4:]
	// verfiyyy Hex and base64Encode
	db, err := decodeHex(ypub)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
	}

	f := base64Encode(db)
	fmt.Printf("%s", f)

	//bytes.join or bytes.buffer.writedString()
	pubY := append(f, data...)
	return base58.Encode(pubY)
}

func Zpub(pub string) string {
	data := base58.Decode(pub)
	// rune splice
	data = data[4:]
	// verfiyyy Hex and base64Encode
	db, err := decodeHex(zpub)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
	}

	f := base64Encode(db)
	fmt.Printf("%s", f)

	//bytes.join or bytes.buffer.writedString()
	pubZ := append(f, data...)
	return base58.Encode(pubZ)

}

func decodeHex(input []byte) ([]byte, error) {
	db := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(db, input)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func base64Encode(input []byte) []byte {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(eb, input)

	return eb
}
