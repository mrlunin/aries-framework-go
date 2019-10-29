/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package authcrypt

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/btcsuite/btcutil/base58"
	chacha "golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/poly1305"

	"github.com/hyperledger/aries-framework-go/pkg/internal/cryptoutil"
	"github.com/hyperledger/aries-framework-go/pkg/kms"
)

// Encrypt will encode the payload argument
// Using the protocol defined by Aries RFC 0019
func (c *Crypter) Encrypt(payload, sender []byte, recipientPubKeys [][]byte) ([]byte, error) {
	var err error

	if len(recipientPubKeys) == 0 {
		return nil, errors.New("empty recipients keys, must have at least one recipient")
	}

	nonce := make([]byte, chacha.NonceSize)
	_, err = c.randSource.Read(nonce)
	if err != nil {
		return nil, err
	}

	// cek (content encryption key) is a symmetric key, for chacha20, a symmetric cipher
	cek := &[chacha.KeySize]byte{}
	_, err = c.randSource.Read(cek[:])
	if err != nil {
		return nil, err
	}

	var recipients []recipient

	recKeys := make([]string, len(recipientPubKeys))
	for i, key := range recipientPubKeys {
		recKeys[i] = base58.Encode(key)
	}

	recipients, err = c.buildRecipients(cek, sender, recipientPubKeys)
	if err != nil {
		return nil, err
	}

	p := protected{
		Enc:        "chacha20poly1305_ietf",
		Typ:        "JWM/1.0",
		Alg:        "Authcrypt",
		Recipients: recipients,
	}

	return c.buildEnvelope(nonce, payload, cek[:], &p)
}

func (c *Crypter) buildEnvelope(nonce, payload, cek []byte, protected *protected) ([]byte, error) {
	protectedBytes, err := json.Marshal(protected)
	if err != nil {
		return nil, err
	}

	protectedB64 := base64.URLEncoding.EncodeToString(protectedBytes)

	chachaCipher, err := chacha.New(cek)
	if err != nil {
		return nil, err
	}

	// 	Additional data is b64encode(jsonencode(protected))
	symPld := chachaCipher.Seal(nil, nonce, payload, []byte(protectedB64))

	// symPld has a length of len(pld) + poly1305.TagSize
	// fetch the tag from the tail
	tag := symPld[len(symPld)-poly1305.TagSize:]
	// fetch the cipherText from the head (0:up to the trailing tag)
	cipherText := symPld[0 : len(symPld)-poly1305.TagSize]

	env := legacyEnvelope{
		Protected:  protectedB64,
		IV:         base64.URLEncoding.EncodeToString(nonce),
		CipherText: base64.URLEncoding.EncodeToString(cipherText),
		Tag:        base64.URLEncoding.EncodeToString(tag),
	}
	out, err := json.Marshal(env)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Crypter) buildRecipients(cek *[chacha.KeySize]byte, senderKey []byte, recPubKeys [][]byte) ([]recipient, error) { // nolint: lll
	var encodedRecipients = make([]recipient, len(recPubKeys))

	for i, recKey := range recPubKeys {
		recipient, err := c.buildRecipient(cek, senderKey, recKey)
		if err != nil {
			return nil, err
		}

		encodedRecipients[i] = *recipient
	}

	return encodedRecipients, nil
}

// buildRecipient encodes the necessary data for the recipient to decrypt the message
// 	encrypting the CEK and sender Pub key
func (c *Crypter) buildRecipient(cek *[chacha.KeySize]byte, senderKey, recKey []byte) (*recipient, error) { // nolint: lll
	var nonce [24]byte

	_, err := c.randSource.Read(nonce[:])
	if err != nil {
		return nil, err
	}

	// We need the private key to be converted and keypair to be persisted
	senderEncKey, err := c.kms.ConvertToEncryptionKey(senderKey)
	if err != nil {
		return nil, err
	}

	recEncKey, err := cryptoutil.PublicEd25519toCurve25519(recKey)
	if err != nil {
		return nil, err
	}

	box, err := kms.NewCryptoBox(c.kms)
	if err != nil {
		return nil, err
	}

	encCEK, err := box.Easy(cek[:], nonce[:], recEncKey, senderEncKey)
	if err != nil {
		return nil, err
	}

	// assumption: senderKey is ed25519
	encSender, err := box.Seal([]byte(base58.Encode(senderKey)), recEncKey, c.randSource)
	if err != nil {
		return nil, err
	}

	return &recipient{
		EncryptedKey: base64.URLEncoding.EncodeToString(encCEK),
		Header: recipientHeader{
			KID:    base58.Encode(recKey), // recKey is the Ed25519 recipient pk in b58 encoding
			Sender: base64.URLEncoding.EncodeToString(encSender),
			IV:     base64.URLEncoding.EncodeToString(nonce[:]),
		},
	}, nil
}
