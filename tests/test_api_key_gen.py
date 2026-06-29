import unittest
from api_key_gen import ApiKeyGen # Assuming your main file is named api_key_gen.py

class TestApiKeyGen(unittest.TestCase):
    
    def setUp(self):
        self.generator = ApiKeyGen("test_env_")

    def test_prefix_is_applied(self):
        result = self.generator.generate()
        self.assertTrue(result['raw_key'].startswith("test_env_"))

    def test_different_keys_generated(self):
        key1 = self.generator.generate()
        key2 = self.generator.generate()
        self.assertNotEqual(key1['raw_key'], key2['raw_key'])
        self.assertNotEqual(key1['db_hash'], key2['db_hash'])

    def test_hash_is_not_raw_key(self):
        result = self.generator.generate()
        self.assertNotEqual(result['raw_key'], result['db_hash'])
        self.assertEqual(len(result['db_hash']), 64) # SHA-256 hex digest is 64 chars

if __name__ == '__main__':
    unittest.main()
  
