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

func WriteToFile(key *string, t SecretType) error {
	if t == JWTAccess || t == JWTRefresh {
		return errors.New("JWT keys of type access or refresh cannot be saved")
	}
	strings.TrimSpace(*key)
	if *key == "" {
		return errors.New("the key string is empty, nothing to write")
	}
	path, err := genPathName(t)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write([]byte(*key)); err != nil {
		return err
	}

	return nil
}

func genPathName(t SecretType) (string, error) {
	path, err := filepath.Abs(defaultPath + string(t) + ".key")
	if err != nil {
		return "", errors.New("failed generating the absolute path to the secret key")
	}
	return path, nil
}
