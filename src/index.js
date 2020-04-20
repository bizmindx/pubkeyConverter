const b58 = require('bs58check');

export const pubToXpub = pub => {

	let data = b58.decode(pub);
	data = data.slice(4);
	data = Buffer.concat([Buffer.from('0488b21e', 'hex'), data]);
	return b58.encode(data);
};

export const pubToYpub = pub => {
	let data = b58.decode(pub);
	data = data.slice(4);
	data = Buffer.concat([Buffer.from('049d7cb2', 'hex'), data]);
	return b58.encode(data);
};

export const pubToZpub = pub => {
	let data = b58.decode(pub);
	data = data.slice(4);
	data = Buffer.concat([Buffer.from('04b24746', 'hex'), data]);
	return b58.encode(data);
};
