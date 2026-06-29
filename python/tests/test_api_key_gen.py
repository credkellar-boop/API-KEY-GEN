import sys
import os
import unittest

# Fix import path to find the file in the parent directory
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from api_key_gen import ApiKeyGen

class TestApiKeyGen(unittest.TestCase):
    def test_prefix(self):
        gen = ApiKeyGen("test_")
        self.assertTrue(gen.generate()['raw_key'].startswith("test_"))

if __name__ == '__main__':
    unittest.main()
  
