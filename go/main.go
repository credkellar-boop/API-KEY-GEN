package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

type ApiKeyGen struct {
	Prefix string
}

func NewApiKeyGen(prefix string) *ApiKeyGen {
	return &ApiKeyGen{Prefix: prefix}
}

func (g *ApiKeyGen) Generate() (string, string) {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	rawKey := g.Prefix + base64.RawURLEncoding.EncodeToString(bytes)
	
	hash := sha256.Sum256([]byte(rawKey))
	return rawKey, hex.EncodeToString(hash[:])
}
