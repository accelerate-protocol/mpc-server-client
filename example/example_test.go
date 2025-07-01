package example

import (
	"context"
	"encoding/hex"
	"net/http"
	"testing"

	client "github.com/accelerate-protocal/mpc-server-client"
	"github.com/stretchr/testify/assert"
)

const (
	mpcServerAddress = "http://127.0.0.1:8082"
)

func TestSign(t *testing.T) {
	c, err := client.NewClientWithResponses(mpcServerAddress)
	assert.Nil(t, err)

	// generate key
	resp, err := c.PostApiV1CustodialAccountGenerateKeyWithResponse(context.Background(), client.PostApiV1CustodialAccountGenerateKeyJSONRequestBody{
		KeyType: client.KeyTypeEd25519,
	})
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())

	t.Logf("generate key: %s, token: %s", *resp.JSON200.Data.PublicKey, *resp.JSON200.Data.Token)

	// sign message
	rawMsg := []byte("hello world")
	msg := hex.EncodeToString(rawMsg)
	res, err := c.PostApiV1CustodialAccountSignWithResponse(context.Background(), client.PostApiV1CustodialAccountSignJSONRequestBody{
		PublicKey: *resp.JSON200.Data.PublicKey,
		Token:     *resp.JSON200.Data.Token,
		Msg:       msg,
	})
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode())
	t.Logf("sign message: %s", *res.JSON200.Data.Signature)

	// // verify signature
	// publicKey := &manager.Ed25519PublicKey{}
	// err = publicKey.Decode(*resp.JSON200.Data.PublicKey)
	// assert.Nil(t, err)
	// verifiedResult := publicKey.Verify(rawMsg, *res.JSON200.Data.Signature)
	// assert.True(t, verifiedResult)
}
