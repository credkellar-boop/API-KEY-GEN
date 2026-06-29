import secrets
import hashlib

class ApiKeyGen:
    def __init__(self, prefix="sk_live_"):
        self.prefix = prefix

    def generate(self):
        random_bytes = secrets.token_urlsafe(32)
        raw_key = f"{self.prefix}{random_bytes}"
        key_hash = hashlib.sha256(raw_key.encode()).hexdigest()
        return {"raw_key": raw_key, "db_hash": key_hash}
      
