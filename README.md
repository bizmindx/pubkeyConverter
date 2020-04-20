# pubconverter
Bitcoin extended public key converter:- convert any extended public key to either XPUB, YPUB, ZPUB
```
npm i pubconverter
```
## usage

```
const pub = require('pubconverter');

let xpub = 'xpub6CUGRUonZSQ4TWtTMmzXdrXDtypWKiKrhko4egpiMZbpiaQL2jkwSB1icqYh2cfDfVxdx4df189oLKnC5fSwqPfgyP3hooxujYzAu3fDVmz';

const convToYpub = pub.pubToYpub(xpub); //ypub6XJXj9Uhi7wYJp5aC8n9qwcj4wxxGLKMcsKHS5ibjZyhmgDZHPvW4Efre3WH2XK9595ShYEDTnWMDcPkoMrxddMHqik8PinQ1H3pHbCYAtS

const convToZpub = pub.pubToZpub(xpub);
//zpub6r8o2p9croV2A7Gh2VZn42iEEv7QCxJrXyqWDUcV7aMapn2nY464gJKzfFTs2Ry4UnCFT1pmvSru6u1KX4GyRs2ti4SYydbtH17Tg8wL57f

const convToxpub = pub.pubToXpub(xpub);
//xpub6CUGRUonZSQ4TWtTMmzXdrXDtypWKiKrhko4egpiMZbpiaQL2jkwSB1icqYh2cfDfVxdx4df189oLKnC5fSwqPfgyP3hooxujYzAu3fDVmz


```


**Methods**

```
pubToXpub();
pubToZpub()
pubToYpub()
```

