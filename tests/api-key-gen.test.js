const test = require('node:test');
const assert = require('node:assert');
const ApiKeyGen = require('../ApiKeyGen'); // Adjust path as needed

test('ApiKeyGen', async (t) => {
    
    await t.test('applies the correct prefix', () => {
        const generator = new ApiKeyGen('test_node_');
        const { rawKey } = generator.generate();
        assert.strictEqual(rawKey.startsWith('test_node_'), true);
    });

    await t.test('generates unique keys and hashes', () => {
        const generator = new ApiKeyGen('sk_live_');
        const key1 = generator.generate();
        const key2 = generator.generate();
        
        assert.notStrictEqual(key1.rawKey, key2.rawKey);
        assert.notStrictEqual(key1.dbHash, key2.dbHash);
    });

    await t.test('creates a valid 64-character SHA-256 hash', () => {
        const generator = new ApiKeyGen();
        const { rawKey, dbHash } = generator.generate();
        
        assert.notStrictEqual(rawKey, dbHash);
        assert.strictEqual(dbHash.length, 64);
        assert.match(dbHash, /^[a-f0-9]+$/); // Ensures it is a valid hex string
    });
});
