# Architecture Overview

API-KEY-GEN is a modular utility designed for cross-language compatibility.

- **Generator Logic:** Uses language-native cryptographically secure random number generators (CSPRNG).
- **Storage Strategy:** Keys are generated as raw tokens (with prefixes) and salted/hashed using SHA-256 for secure DB persistence.
- **Project Structure:** Code is isolated by language directory (`python/`, `node/`, `go/`) to manage dependencies independently.
