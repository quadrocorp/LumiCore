package keychain

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	GPGMaster        SecretType = "GPGMaster"
	GPGMasterRefresh SecretType = "GPGMasterRefresh"
	JWTMaster        SecretType = "JWTMaster"
	JWTAccess        SecretType = "JWTAccess"
	JWTRefresh       SecretType = "JWTRefresh"
)

const (
	defaultPath string = "./configs/secrets/"
)

func GenerateJWTSecret(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be positive")
	}
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to read random bytes: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

func WriteToFile(key string, t SecretType) error {
	if t == JWTAccess || t == JWTRefresh {
		return errors.New("JWT keys of type access or refresh cannot be saved")
	}
	key = strings.TrimSpace(key)
	if key == "" {
		return errors.New("the key string is empty, nothing to write")
	}
	path, err := genPathName(t)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("failed to create secret directory %q: %w", dir, err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to open secret file %q: %w", path, err)
	}
	defer file.Close()

	_, err = file.Write([]byte(key))
	return err
}

func genPathName(t SecretType) (string, error) {
	path, err := filepath.Abs(defaultPath + string(t) + ".key")
	if err != nil {
		return "", errors.New("failed generating the absolute path to the secret key")
	}
	return path, nil
}
