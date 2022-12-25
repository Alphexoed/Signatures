const crypto = require('crypto');

function generateDeviceId() {
    const identifier = crypto.randomBytes(20);
    const key = Buffer.from('AE49550458D8E7C51D566916B04888BFB8B3CA7D', 'hex');
    const hmac = crypto.createHmac('sha1', key);
    hmac.update(Buffer.concat([Buffer.from([0x52]), identifier]));
    return ('52' + identifier.toString('hex') + hmac.digest('hex')).toUpperCase();
}

console.log(generateDeviceId())