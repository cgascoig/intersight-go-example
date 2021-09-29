package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cgascoig/intersight-go-example/pkg/intersight"
)

func main() {
	keyID := os.Getenv("IS_KEY_ID")
	keyFile := os.Getenv("IS_KEY_FILE")

	config := intersight.NewConfiguration()
	config.Debug = false

	client := intersight.NewAPIClient(config)

	authConfig := intersight.HttpSignatureAuth{
		KeyId:          keyID,
		PrivateKeyPath: keyFile,

		SigningScheme: intersight.HttpSigningSchemeRsaSha256,
		SignedHeaders: []string{
			intersight.HttpSignatureParameterRequestTarget, // The special (request-target) parameter expresses the HTTP request target.
			"Host",   // The Host request header specifies the domain name of the server, and optionally the TCP port number.
			"Date",   // The date and time at which the message was originated.
			"Digest", // A cryptographic digest of the request body.
		},
		SigningAlgorithm: intersight.HttpSigningAlgorithmRsaPKCS1v15,
	}

	authCtx, err := authConfig.ContextWithValue(context.Background())
	if err != nil {
		log.Fatal("Error creating authentication context")
	}

	res, _, err := client.NtpApi.GetNtpPolicyList(authCtx).Execute()

	for _, ntpPolicy := range res.NtpPolicyList.Results {
		fmt.Printf("Ntp Policy: %s\n", *ntpPolicy.Name)
	}
}
