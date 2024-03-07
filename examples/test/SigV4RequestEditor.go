package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/deepmap/oapi-codegen/v2/examples/test/petstore"
	"io"
	"net/http"
	"strings"
	"time"
)

func NewSigV4RequestEditor(config aws.Config) petstore.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		credentials, err := config.Credentials.Retrieve(ctx)
		if err != nil {
			return err
		}

		signer := v4.NewSigner()

		body, err := req.GetBody()
		if err != nil {
			return err
		}

		bodyBytes, err := io.ReadAll(body)
		if err != nil {
			return err
		}

		bodySha := sha256.Sum256(bodyBytes)

		shaHex := hex.EncodeToString(bodySha[:])

		err = signer.SignHTTP(ctx, credentials, req, shaHex, getServiceName(req), config.Region, time.Now())
		if err != nil {
			return err
		}

		return nil
	}
}

func getServiceName(req *http.Request) string {
	return req.URL.Host[0:strings.Index(req.URL.Host, ".")]
}
