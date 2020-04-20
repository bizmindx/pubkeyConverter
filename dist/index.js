"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.pubToZpub = exports.pubToYpub = exports.pubToXpub = void 0;

var b58 = require('bs58check');

var check = function check(pub) {
  if (!pub) {
    throw new Error('Key is required');
  }

  var key = b58.decode(pub);

  if (!key) {
    throw new Error('Invalid extended public key');
  }
};

var pubToXpub = function pubToXpub(pub) {
  check(pub);
  var data = b58.decode(pub);
  data = data.slice(4);
  data = Buffer.concat([Buffer.from('0488b21e', 'hex'), data]);
  return b58.encode(data);
};

exports.pubToXpub = pubToXpub;

var pubToYpub = function pubToYpub(pub) {
  check(pub);
  var data = b58.decode(pub);
  data = data.slice(4);
  data = Buffer.concat([Buffer.from('049d7cb2', 'hex'), data]);
  return b58.encode(data);
};

exports.pubToYpub = pubToYpub;

var pubToZpub = function pubToZpub(pub) {
  check(pub);
  var data = b58.decode(pub);
  data = data.slice(4);
  data = Buffer.concat([Buffer.from('04b24746', 'hex'), data]);
  return b58.encode(data);
};

exports.pubToZpub = pubToZpub;
