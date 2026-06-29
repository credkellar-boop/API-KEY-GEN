# API-KEY-GEN

API-KEY-GEN creates cryptographically secure, 256-bit API keys. Available in Node.js, Python, and Go, it features customizable prefixes (like 'sk_live_') and auto-generates SHA-256 hashes so you never have to store plaintext keys in your database.

## Why Use This?
Standard random number generators (like `Math.random`) are predictable and vulnerable to brute-force attacks. API-KEY-GEN uses cryptographically secure pseudorandom number generators (CSPRNG) combined with best practices for database storage.

## Features
* **256-bit Entropy:** Generates 32-byte, highly secure keys.
* **Secret-Scanning Safe:** Appends identifiable prefixes (e.g., `sk_live_`) so automated tools can flag leaked keys.
* **Dual Output:** Returns the raw key (to show your user) and a SHA-256 hash (to safely store in your DB).
