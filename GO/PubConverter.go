Package PubConverter

/*
This script uses version bytes as described in SLIP-132
https://github.com/satoshilabs/slips/blob/master/slip-0132.md
*/

import (
	   "github.com/btcsuite/btcutil/base58"
       "strings"
       "encoding/hex"
       "encoding/base64"
       "fmt"
)

var pubMap := map[string]string{"xpub":"0488b21e", "ypub": "049d7cb2", "zpub": "04b24746"}




func Xpub(pub []byte) []byte {
	data := base58.Decode(pub)
	// rune splice
    data := data[4:]
   // verfiyyy Hex and base64Encode
   db, err := decodeHex(pubmap["xpub"])
   if err != nil {
       fmt.Printf("failed to decode hex: %s", err)
       return
   }

   f := base64Encode(db)
   fmt.Printf("%s", f)

//bytes.join or bytes.buffer.writedString()
xpub,err := bytes.join(f,data)
if err != nil {
    return nil, err
}

return base58.encode(xpub)
}

	
func Ypub(pub []byte) []byte {
	data := base58.Decode(pub)
	// rune splice
    data := data[4:]
   // verfiyyy Hex and  base64Encode
   db, err := decodeHex(pubmap["ypub"])
   if err != nil {
       fmt.Printf("failed to decode hex: %s", err)
       return
   }

   f := base64Encode(db)
   fmt.Printf("%s", f)

//bytes.join or bytes.buffer.writedString()
ypub:= bytes.join(f,data)
return base58.encode(zpub)  
}
	
func Zpub(pub []byte) []byte {
	data := base58.Decode(pub)
	// rune splice
    data := data[4:]
   // verfiyyy Hex and  base64Encode
   db, err := decodeHex(pubmap["zpub"])
   if err != nil {
       fmt.Printf("failed to decode hex: %s", err)
       return
   }

   f := base64Encode(db)
   fmt.Printf("%s", f)

//bytes.join or bytes.buffer.writedString()
zpub:= bytes.join(f,data)
return base58.encode(zpub)
}



func decodeHex(input []byte) ([]byte, error) {
   db := make([]byte, hex.DecodedLen(len(input)))
   _, err := hex.Decode(db, input)
   if err != nil {
       return nil, err
   }
   return db, nil
}

func base64Encode(input []byte) ([]byte) {
   eb := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
   base64.StdEncoding.Encode(eb, input)

   return eb
}
 