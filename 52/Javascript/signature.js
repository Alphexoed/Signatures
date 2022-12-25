const crypto = require('crypto');

function generateSignature(data) {
    const mac = crypto.createHmac('sha1', Buffer.from('EAB4F1B9E3340CD1631EDE3B587CC3EBEDF1AFA9', 'hex'));
    mac.update(data);
    return Buffer.concat([Buffer.from('R'), mac.digest()]).toString('base64');
}

console.log(generateSignature('teste'))
