package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"github.com/labstack/echo"
	"hottub/types"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateUserCosmosSettings(c echo.Context) error {
	reqSettings := new(types.CosmosUser)
	var user types.User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(types.ErrorParameterNotInteger.Status, types.ErrorParameterNotInteger)
	}

	if err = c.Bind(&reqSettings); err != nil {
		return c.JSON(types.ErrorCannotParseFields.Status, types.ErrorCannotParseFields)
	}

	enc, err := h.encryptMnemonic([]byte(reqSettings.Mnemonic))
	reqSettings.Mnemonic = string(enc)

	h.DB.First(&user, id)
	h.DB.Create(&user).Related(&reqSettings)

	return c.JSON(http.StatusOK, reqSettings)
}

func (h *Handler) GetCosmosSettings(c echo.Context) error {
	var user types.User
	var settings types.CosmosUser
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(types.ErrorParameterNotInteger.Status, types.ErrorParameterNotInteger)
	}
	h.DB.First(&user, id).Related(&settings)
	bytes, err := decrypt(settings.Mnemonic)
	settings.Mnemonic = string(bytes)
	return c.JSON(http.StatusOK, settings)
}

func (h *Handler) encryptMnemonic(plaintext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(Key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(Key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
