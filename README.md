# API-KEY-GEN
<p align="center">
  <img src="1782718770759.png" alt="Profile Image" width="400"/>
</p>

API-KEY-GEN creates cryptographically secure, 256-bit API keys. Available in Node.js, Python, and Go, it features customizable prefixes (like 'sk_live_') and auto-generates SHA-256 hashes so you never have to store plaintext keys in your database.

## Why Use This?
Standard random number generators (like `Math.random`) are predictable and vulnerable to brute-force attacks. API-KEY-GEN uses cryptographically secure pseudorandom number generators (CSPRNG) combined with best practices for database storage.

## Features
* **256-bit Entropy:** Generates 32-byte, highly secure keys.
* **Secret-Scanning Safe:** Appends identifiable prefixes (e.g., `sk_live_`) so automated tools can flag leaked keys.
* **Dual Output:** Returns the raw key (to show your user) and a SHA-256 hash (to safely store in your DB).

## 🛠️ Architecture, Core Components & Profile Badges

These badges represent the extensive community, metadata, and architectural documentation present in your root directory.

<!-- Based on package.json, LICENSE, CODE_OF_CONDUCT.md, CONTRIBUTING.md, CONTRIBUTORS.md, CHANGELOG.md, CITATION.cff, CODEOWNERS, SECURITY.md, SECURITY.txt, and docs/ARCHITECTURE.md -->
![Version](https://img.shields.io/badge/Version-1.0.0-007EC6?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-success?style=for-the-badge)
![Architecture](https://img.shields.io/badge/Architecture-Documented-blueviolet?style=for-the-badge&logo=read-the-docs)
![Code of Conduct](https://img.shields.io/badge/Code_of_Conduct-Contributor_Covenant-EA4AAA?style=for-the-badge)
![Contributions](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=for-the-badge)
![Contributors](https://img.shields.io/badge/Contributors-Tracked-orange?style=for-the-badge&logo=github)
![Changelog](https://img.shields.io/badge/Changelog-Maintained-4B32C3?style=for-the-badge)
![Citation](https://img.shields.io/badge/Citation-CFF-lightgrey?style=for-the-badge)
![Codeowners](https://img.shields.io/badge/Codeowners-Enforced-181717?style=for-the-badge&logo=github)
![Security Policy](https://img.shields.io/badge/Security-Policy_Active-red?style=for-the-badge&logo=shield&logoColor=white)
![Security.txt](https://img.shields.io/badge/Security.txt-Configured-critical?style=for-the-badge)

---

## 📦 Core Programming Languages

These reflect the exact runtimes and languages implemented in your multi-language project.

<!-- Based on go/, node/, python/, .nvmrc, and .python-version -->
![Go](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Node.js](https://img.shields.io/badge/Node.js-20.10.0-339933?style=for-the-badge&logo=nodedotjs&logoColor=white)
![Python](https://img.shields.io/badge/Python-3.12-3776AB?style=for-the-badge&logo=python&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-ES6%2B-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)

---

## ⚙️ Core Systems (Cryptography & Testing)

Mapped directly to the native libraries and standard testing packages utilized across all three languages.

<!-- Based on crypto/sha256, crypto/rand, hashlib, secrets, node:crypto, node:test, unittest, and testing -->
![SHA-256](https://img.shields.io/badge/Hashing-SHA--256-D00000?style=for-the-badge&logo=letsencrypt&logoColor=white)
![CSPRNG](https://img.shields.io/badge/Entropy-CSPRNG-8A2BE2?style=for-the-badge)
![Encoding](https://img.shields.io/badge/Encoding-Base64%20%7C%20Hex-555555?style=for-the-badge)
![Zero Dependencies](https://img.shields.io/badge/Dependencies-Zero-2EA44F?style=for-the-badge)
![Test: Go](https://img.shields.io/badge/Testing-go%20test-00ADD8?style=for-the-badge&logo=go)
![Test: Node](https://img.shields.io/badge/Testing-node:test-339933?style=for-the-badge&logo=nodedotjs)
![Test: Python](https://img.shields.io/badge/Testing-unittest-3776AB?style=for-the-badge&logo=python)

---

## 🛠️ DevOps, Infrastructure & Build Tools

This covers the exhaustive list of dotfiles, package managers, and CI/CD configurations in your root tree.

<!-- Based on .github, Dockerfile, .dockerignore, Makefile, .editorconfig, .gitattributes, .gitignore, .nvmrc, .python-version, go.mod, package.json, pyproject.toml, MANIFEST.in -->
![GitHub Actions](https://img.shields.io/badge/CI-GitHub_Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white)
![Docker](https://img.shields.io/badge/Container-Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![GNU Make](https://img.shields.io/badge/Build-Makefile-000000?style=for-the-badge&logo=gnu&logoColor=white)
![Git](https://img.shields.io/badge/VCS-Git-F05032?style=for-the-badge&logo=git&logoColor=white)
![EditorConfig](https://img.shields.io/badge/Format-EditorConfig-FEFEFE?style=for-the-badge&logo=editorconfig&logoColor=black)
![NVM](https://img.shields.io/badge/Tool-NVM-4CAF50?style=for-the-badge)
![Pyenv](https://img.shields.io/badge/Tool-Pyenv-FFD43B?style=for-the-badge&logo=python&logoColor=blue)
![Go Modules](https://img.shields.io/badge/Package-Go_Modules-00ADD8?style=for-the-badge&logo=go)
![NPM](https://img.shields.io/badge/Package-NPM-CB3837?style=for-the-badge&logo=npm&logoColor=white)
![PyPI](https://img.shields.io/badge/Package-PyPA-3775A9?style=for-the-badge&logo=pypi&logoColor=white)
