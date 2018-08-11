package main

import "testing"

func TestLoadPrivateKey(t *testing.T) {
	keyPath := "test/nokey.pem"
	key, err := loadPrivateKey(keyPath)
	if key != nil {
		t.Errorf("loadPrivateKey() returned a result - it should be returning nil due to a missing file")
	}
	if err == nil {
		t.Errorf("loadPrivateKey() returned no error - it should be failing due to a missing file")
	}

	keyPath = "test/key.pem"
	key, err = loadPrivateKey(keyPath)
	if key == nil {
		t.Errorf("loadPrivateKey() returned no key")
	}
	if err != nil {
		t.Errorf("loadPrivateKey() returned error - %s", err)
	}
}

func TestGenerateToken(t *testing.T) {
	secretPrivateKey, _ = loadPrivateKey("./test/key.pem")

	token, err := generateToken("42a")
	if token == "" {
		t.Errorf("Blank token returned")
	}
	if err != nil {
		t.Errorf("generateToken() - %s", err)
	}
}
