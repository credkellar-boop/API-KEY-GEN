const crypto = require('crypto');

class ApiKeyGen {
    constructor(prefix = 'sk_live_') {
        this.prefix = prefix;
    }

    generate() {
        const randomBytes = crypto.randomBytes(32).toString('base64url');
        const rawKey = `${this.prefix}${randomBytes}`;
        const dbHash = crypto.createHash('sha256').update(rawKey).digest('hex');
        return { rawKey, dbHash };
    }
}

module.exports = ApiKeyGen; // Essential for importing
