const test = require('node:test');
const assert = require('node:assert');
const ApiKeyGen = require('../ApiKeyGen.js'); // Relative path

test('ApiKeyGen creates key', () => {
    const gen = new ApiKeyGen('sk_test_');
    const key = gen.generate();
    assert.strictEqual(key.rawKey.startsWith('sk_test_'), true);
});
