# yellowcard-endecrypt-library
3 Encryption and decryption algorithim 
```
npm i @yellowcard/yellowcard-encdecrypt-library
```
## usage
**Encrypt**
```
let ade = encrypt('hahahah');
console.log(ade);

```
>Result
```
{ blob1:
   'U2FsdGVkX1++cNLlgKn/E3rI2erYLZ/zarrcn8hBmSoxTw4ZKoFkKkB9y0mZvaaN0n7YpRYNTa2rwK3qchzAig==',
  key1: '9c5070da-f4c6-4b24-8c43-19e0140f4af8',
  blob2:
   'U2FsdGVkX18EoHQuzW4i7ynvh9NSzoIsPxYYlTii5Dx4bKCtzEI4Ue+vbFlr7riashRB6AI3MDDE8mevLHkumiUKKnN+y81YuX8crtcEQ3HkjQnMeBH5Vsi+CIyaxMwQvRk7YTr4JQJ8DW78bak/SCti+KzLSGTg8b3x6iBtdr6gHkl8Ixg4XIdJatQIop+IMpo+KOOjPu/kFZiacgL7jA==',
  key2: '717b62d9-5136-4d47-be12-a5c28c57b600' }
```

**Decrypt**

```

let dec = decrypt(ade.blob1, ade.key1, ade.blob2, ade.key2);
console.log(dec);

//hahahah
```