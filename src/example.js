const pub = require('../dist/index');

let xpub = 'xpub6CUGRUonZSQ4TWtTMmzXdrXDtypWKiKrhko4egpiMZbpiaQL2jkwSB1icqYh2cfDfVxdx4df189oLKnC5fSwqPfgyP3hooxujYzAu3fDVmz';

const convToYpub = pub.pubToYpub(xpub);

const convToZpub = pub.pubToZpub(xpub);
console.log(convToZpub);
console.log(convToYpub);
