/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package verifiable_test

import (
	"crypto/ed25519"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
)

// Private key generated by ed25519.GenerateKey(rand.Reader)
//nolint:gochecknoglobals
var privHolderKey = ed25519.PrivateKey{137, 207, 34, 115, 168, 80, 44, 145, 236, 206, 165, 152, 211, 195, 240, 205, 232, 37, 216, 101, 58, 133, 198, 107, 232, 119, 30, 80, 176, 137, 10, 251, 109, 144, 158, 93, 189, 195, 0, 24, 15, 29, 166, 185, 169, 69, 246, 25, 182, 179, 115, 54, 107, 4, 123, 30, 5, 88, 175, 94, 109, 15, 34, 113} //nolint:lll

func ExampleNewPresentation() {
	// A Holder sends to the Verifier a verifiable presentation in JWS form.
	vpJWS := "eyJhbGciOiJFZERTQSIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJkaWQ6ZXhhbXBsZTo0YTU3NTQ2OTczNDM2ZjZmNmM0YTRhNTc1NzMiLCJpc3MiOiJkaWQ6ZXhhbXBsZTplYmZlYjFmNzEyZWJjNmYxYzI3NmUxMmVjMjEiLCJqdGkiOiJ1cm46dXVpZDozOTc4MzQ0Zi04NTk2LTRjM2EtYTk3OC04ZmNhYmEzOTAzYzUiLCJ2cCI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL2V4YW1wbGVzL3YxIl0sInR5cGUiOlsiVmVyaWZpYWJsZVByZXNlbnRhdGlvbiIsIlVuaXZlcnNpdHlEZWdyZWVDcmVkZW50aWFsIl0sInZlcmlmaWFibGVDcmVkZW50aWFsIjpbeyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL2V4YW1wbGVzL3YxIl0sImNyZWRlbnRpYWxTY2hlbWEiOltdLCJjcmVkZW50aWFsU3ViamVjdCI6eyJkZWdyZWUiOnsidHlwZSI6IkJhY2hlbG9yRGVncmVlIiwidW5pdmVyc2l0eSI6Ik1JVCJ9LCJpZCI6ImRpZDpleGFtcGxlOmViZmViMWY3MTJlYmM2ZjFjMjc2ZTEyZWMyMSIsIm5hbWUiOiJKYXlkZW4gRG9lIiwic3BvdXNlIjoiZGlkOmV4YW1wbGU6YzI3NmUxMmVjMjFlYmZlYjFmNzEyZWJjNmYxIn0sImV4cGlyYXRpb25EYXRlIjoiMjAyMC0wMS0wMVQxOToyMzoyNFoiLCJpZCI6Imh0dHA6Ly9leGFtcGxlLmVkdS9jcmVkZW50aWFscy8xODcyIiwiaXNzdWFuY2VEYXRlIjoiMjAxMC0wMS0wMVQxOToyMzoyNFoiLCJpc3N1ZXIiOnsiaWQiOiJkaWQ6ZXhhbXBsZTo3NmUxMmVjNzEyZWJjNmYxYzIyMWViZmViMWYiLCJuYW1lIjoiRXhhbXBsZSBVbml2ZXJzaXR5In0sInJlZmVyZW5jZU51bWJlciI6ODMyOTQ4NDcsInR5cGUiOlsiVmVyaWZpYWJsZUNyZWRlbnRpYWwiLCJVbml2ZXJzaXR5RGVncmVlQ3JlZGVudGlhbCJdfV19fQ.F6-OoI27Q5d82ZR32uXxAXYH15QxhzrwCTZIJnwF7ABhXXUA73PmxxnFuXr-5keXpCzXPxtzHsRv0AKUWW_iAA" //nolint:lll

	// Holder received and decodes it.
	vp, err := verifiable.NewPresentation(
		[]byte(vpJWS),
		verifiable.WithPresPublicKeyFetcher(verifiable.SingleKey(privHolderKey.Public())))
	if err != nil {
		fmt.Println(fmt.Errorf("failed to decode VP JWS: %w", err))
	}

	// Marshal the VP to JSON to verify the result of decoding.
	vpBytes, err := json.Marshal(vp)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to marshal VP to JSON: %w", err))
	}

	fmt.Println(string(vpBytes))

	//nolint:lll
	// Output: {"@context":["https://www.w3.org/2018/credentials/v1","https://www.w3.org/2018/credentials/examples/v1"],"id":"urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c5","type":["VerifiablePresentation","UniversityDegreeCredential"],"verifiableCredential":[{"@context":["https://www.w3.org/2018/credentials/v1","https://www.w3.org/2018/credentials/examples/v1"],"credentialSchema":[],"credentialSubject":{"degree":{"type":"BachelorDegree","university":"MIT"},"id":"did:example:ebfeb1f712ebc6f1c276e12ec21","name":"Jayden Doe","spouse":"did:example:c276e12ec21ebfeb1f712ebc6f1"},"expirationDate":"2020-01-01T19:23:24Z","id":"http://example.edu/credentials/1872","issuanceDate":"2010-01-01T19:23:24Z","issuer":{"id":"did:example:76e12ec712ebc6f1c221ebfeb1f","name":"Example University"},"referenceNumber":83294847,"type":["VerifiableCredential","UniversityDegreeCredential"]}],"holder":"did:example:ebfeb1f712ebc6f1c276e12ec21"}
}

func ExamplePresentation_JWTClaims() {
	// The Holder kept the presentation serialized to JSON in her personal verifiable credential wallet.
	vpStrFromWallet := `
{
  "@context": [
    "https://www.w3.org/2018/credentials/v1",
    "https://www.w3.org/2018/credentials/examples/v1"
  ],
  "id": "urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c5",
  "type": [
    "VerifiablePresentation",
    "UniversityDegreeCredential"
  ],
  "verifiableCredential": [
    {
      "@context": [
        "https://www.w3.org/2018/credentials/v1",
        "https://www.w3.org/2018/credentials/examples/v1"
      ],
      "credentialSchema": [],
      "credentialSubject": {
        "degree": {
          "type": "BachelorDegree",
          "university": "MIT"
        },
        "id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
        "name": "Jayden Doe",
        "spouse": "did:example:c276e12ec21ebfeb1f712ebc6f1"
      },
      "expirationDate": "2020-01-01T19:23:24Z",
      "id": "http://example.edu/credentials/1872",
      "issuanceDate": "2010-01-01T19:23:24Z",
      "issuer": {
        "id": "did:example:76e12ec712ebc6f1c221ebfeb1f",
        "name": "Example University"
      },
      "referenceNumber": 83294847,
      "type": [
        "VerifiableCredential",
        "UniversityDegreeCredential"
      ]
    }
  ],
  "holder": "did:example:ebfeb1f712ebc6f1c276e12ec21"
}
`

	// The Holder wants to send the presentation to the Verifier in JWS.
	vp, err := verifiable.NewPresentation([]byte(vpStrFromWallet), verifiable.WithPresSkippedEmbeddedProofCheck())
	if err != nil {
		fmt.Println(fmt.Errorf("failed to decode VP JSON: %w", err))
	}

	aud := []string{"did:example:4a57546973436f6f6c4a4a57573"}

	jws, err := vp.JWTClaims(aud, true).MarshalJWS(verifiable.EdDSA, privHolderKey, "")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to sign VP inside JWT: %w", err))
	}

	fmt.Println(jws)

	//nolint
	//Output: eyJhbGciOiJFZERTQSIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJkaWQ6ZXhhbXBsZTo0YTU3NTQ2OTczNDM2ZjZmNmM0YTRhNTc1NzMiLCJpc3MiOiJkaWQ6ZXhhbXBsZTplYmZlYjFmNzEyZWJjNmYxYzI3NmUxMmVjMjEiLCJqdGkiOiJ1cm46dXVpZDozOTc4MzQ0Zi04NTk2LTRjM2EtYTk3OC04ZmNhYmEzOTAzYzUiLCJ2cCI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL2V4YW1wbGVzL3YxIl0sInR5cGUiOlsiVmVyaWZpYWJsZVByZXNlbnRhdGlvbiIsIlVuaXZlcnNpdHlEZWdyZWVDcmVkZW50aWFsIl0sInZlcmlmaWFibGVDcmVkZW50aWFsIjpbeyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL2V4YW1wbGVzL3YxIl0sImNyZWRlbnRpYWxTY2hlbWEiOltdLCJjcmVkZW50aWFsU3ViamVjdCI6eyJkZWdyZWUiOnsidHlwZSI6IkJhY2hlbG9yRGVncmVlIiwidW5pdmVyc2l0eSI6Ik1JVCJ9LCJpZCI6ImRpZDpleGFtcGxlOmViZmViMWY3MTJlYmM2ZjFjMjc2ZTEyZWMyMSIsIm5hbWUiOiJKYXlkZW4gRG9lIiwic3BvdXNlIjoiZGlkOmV4YW1wbGU6YzI3NmUxMmVjMjFlYmZlYjFmNzEyZWJjNmYxIn0sImV4cGlyYXRpb25EYXRlIjoiMjAyMC0wMS0wMVQxOToyMzoyNFoiLCJpZCI6Imh0dHA6Ly9leGFtcGxlLmVkdS9jcmVkZW50aWFscy8xODcyIiwiaXNzdWFuY2VEYXRlIjoiMjAxMC0wMS0wMVQxOToyMzoyNFoiLCJpc3N1ZXIiOnsiaWQiOiJkaWQ6ZXhhbXBsZTo3NmUxMmVjNzEyZWJjNmYxYzIyMWViZmViMWYiLCJuYW1lIjoiRXhhbXBsZSBVbml2ZXJzaXR5In0sInJlZmVyZW5jZU51bWJlciI6OC4zMjk0ODQ3ZSswNywidHlwZSI6WyJWZXJpZmlhYmxlQ3JlZGVudGlhbCIsIlVuaXZlcnNpdHlEZWdyZWVDcmVkZW50aWFsIl19XX19.fYIKWhFN699O0GJl6DoYw0L_IcpR24GQREPT9G_0lIWGT02NFDuOFFuvydedujGd6twiNW9Drizm997Z7oYtDw
}

func ExampleCredential_Presentation() {
	// A Holder loads the credential from verifiable credential wallet in order to send to Verifier.
	// She embedded the credential into presentation and sends it to the Verifier in JWS form.
	vcStrFromWallet := `
{
  "@context": [
    "https://www.w3.org/2018/credentials/v1",
    "https://www.w3.org/2018/credentials/examples/v1"
  ],
  "credentialSchema": [],
  "credentialSubject": {
    "degree": {
      "type": "BachelorDegree",
      "university": "MIT"
    },
    "id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
    "name": "Jayden Doe",
    "spouse": "did:example:c276e12ec21ebfeb1f712ebc6f1"
  },
  "expirationDate": "2020-01-01T19:23:24Z",
  "id": "http://example.edu/credentials/1872",
  "issuanceDate": "2010-01-01T19:23:24Z",
  "issuer": {
    "id": "did:example:76e12ec712ebc6f1c221ebfeb1f",
    "name": "Example University"
  },
  "referenceNumber": 83294847,
  "type": [
    "VerifiableCredential",
    "UniversityDegreeCredential"
  ]
}
`

	vc, _, err := verifiable.NewCredential([]byte(vcStrFromWallet))
	if err != nil {
		fmt.Println(fmt.Errorf("failed to decode VC JSON: %w", err))
	}

	vp, err := vc.Presentation()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to build VP from VC: %w", err))
	}

	vp.ID = "urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c5"
	vp.Holder = "did:example:ebfeb1f712ebc6f1c276e12ec21"

	aud := []string{"did:example:4a57546973436f6f6c4a4a57573"}

	jws, err := vp.JWTClaims(aud, true).MarshalJWS(verifiable.EdDSA, privHolderKey, "")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to sign VP inside JWT: %w", err))
	}

	fmt.Println(jws)

	//nolint
	//Output: eyJhbGciOiJFZERTQSIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJkaWQ6ZXhhbXBsZTo0YTU3NTQ2OTczNDM2ZjZmNmM0YTRhNTc1NzMiLCJpc3MiOiJkaWQ6ZXhhbXBsZTplYmZlYjFmNzEyZWJjNmYxYzI3NmUxMmVjMjEiLCJqdGkiOiJ1cm46dXVpZDozOTc4MzQ0Zi04NTk2LTRjM2EtYTk3OC04ZmNhYmEzOTAzYzUiLCJ2cCI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL2V4YW1wbGVzL3YxIl0sInR5cGUiOlsiVmVyaWZpYWJsZVByZXNlbnRhdGlvbiJdLCJ2ZXJpZmlhYmxlQ3JlZGVudGlhbCI6W3siQGNvbnRleHQiOlsiaHR0cHM6Ly93d3cudzMub3JnLzIwMTgvY3JlZGVudGlhbHMvdjEiLCJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy9leGFtcGxlcy92MSJdLCJjcmVkZW50aWFsU2NoZW1hIjpbXSwiY3JlZGVudGlhbFN1YmplY3QiOnsiZGVncmVlIjp7InR5cGUiOiJCYWNoZWxvckRlZ3JlZSIsInVuaXZlcnNpdHkiOiJNSVQifSwiaWQiOiJkaWQ6ZXhhbXBsZTplYmZlYjFmNzEyZWJjNmYxYzI3NmUxMmVjMjEiLCJuYW1lIjoiSmF5ZGVuIERvZSIsInNwb3VzZSI6ImRpZDpleGFtcGxlOmMyNzZlMTJlYzIxZWJmZWIxZjcxMmViYzZmMSJ9LCJleHBpcmF0aW9uRGF0ZSI6IjIwMjAtMDEtMDFUMTk6MjM6MjRaIiwiaWQiOiJodHRwOi8vZXhhbXBsZS5lZHUvY3JlZGVudGlhbHMvMTg3MiIsImlzc3VhbmNlRGF0ZSI6IjIwMTAtMDEtMDFUMTk6MjM6MjRaIiwiaXNzdWVyIjp7ImlkIjoiZGlkOmV4YW1wbGU6NzZlMTJlYzcxMmViYzZmMWMyMjFlYmZlYjFmIiwibmFtZSI6IkV4YW1wbGUgVW5pdmVyc2l0eSJ9LCJyZWZlcmVuY2VOdW1iZXIiOjgzMjk0ODQ3LCJ0eXBlIjpbIlZlcmlmaWFibGVDcmVkZW50aWFsIiwiVW5pdmVyc2l0eURlZ3JlZUNyZWRlbnRpYWwiXX1dfX0.lak4J_q8LR0hiaC7f1QDcEd8BkDVl4c_vmqOPg8EFue-W27_mbkBJPIaoiOE-5UsxrSr-ApUFYvlxs24utq6Bw
}

func ExamplePresentation_SetCredentials() {
	// Holder wants to send 2 credentials to Verifier
	vp := &verifiable.Presentation{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1"},
		ID:     "urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c",
		Type:   []string{"VerifiablePresentation"},
		Holder: "did:example:ebfeb1f712ebc6f1c276e12ec21",
	}

	// The first VC is created on fly (or just decoded using NewCredential).
	vc := verifiable.Credential{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/examples/v1"},
		ID: "http://example.edu/credentials/1872",
		Types: []string{
			"VerifiableCredential",
			"UniversityDegreeCredential"},
		Subject: UniversityDegreeSubject{
			ID:     "did:example:ebfeb1f712ebc6f1c276e12ec21",
			Name:   "Jayden Doe",
			Spouse: "did:example:c276e12ec21ebfeb1f712ebc6f1",
			Degree: UniversityDegree{
				Type:       "BachelorDegree",
				University: "MIT",
			},
		},
		Issuer: verifiable.Issuer{
			ID:   "did:example:76e12ec712ebc6f1c221ebfeb1f",
			Name: "Example University",
		},
		Issued:  &issued,
		Expired: &expired,
		Schemas: []verifiable.TypedID{},
		CustomFields: map[string]interface{}{
			"referenceNumber": 83294847,
		},
	}

	// The second VC is provided in JWS form (e.g. kept in the wallet in that form).
	vcJWS := "eyJhbGciOiJFZERTQSIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Nzc5MDY2MDQsImlhdCI6MTI2MjM3MzgwNCwiaXNzIjoiZGlkOmV4YW1wbGU6NzZlMTJlYzcxMmViYzZmMWMyMjFlYmZlYjFmIiwianRpIjoiaHR0cDovL2V4YW1wbGUuZWR1L2NyZWRlbnRpYWxzLzE4NzIiLCJuYmYiOjEyNjIzNzM4MDQsInN1YiI6ImRpZDpleGFtcGxlOmViZmViMWY3MTJlYmM2ZjFjMjc2ZTEyZWMyMSIsInZjIjp7IkBjb250ZXh0IjpbImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL3YxIiwiaHR0cHM6Ly93d3cudzMub3JnLzIwMTgvY3JlZGVudGlhbHMvZXhhbXBsZXMvdjEiXSwiY3JlZGVudGlhbFNjaGVtYSI6W10sImNyZWRlbnRpYWxTdWJqZWN0Ijp7ImRlZ3JlZSI6eyJ0eXBlIjoiQmFjaGVsb3JEZWdyZWUiLCJ1bml2ZXJzaXR5IjoiTUlUIn0sImlkIjoiZGlkOmV4YW1wbGU6ZWJmZWIxZjcxMmViYzZmMWMyNzZlMTJlYzIxIiwibmFtZSI6IkpheWRlbiBEb2UiLCJzcG91c2UiOiJkaWQ6ZXhhbXBsZTpjMjc2ZTEyZWMyMWViZmViMWY3MTJlYmM2ZjEifSwiaXNzdWVyIjp7Im5hbWUiOiJFeGFtcGxlIFVuaXZlcnNpdHkifSwidHlwZSI6WyJWZXJpZmlhYmxlQ3JlZGVudGlhbCIsIlVuaXZlcnNpdHlEZWdyZWVDcmVkZW50aWFsIl19fQ.AHn2A2q5DL1heX3_izq_2yrsBDhoZ6BGGKhoRvhfMnMUuuOnBOdekdTg-dfUMJgipXRql_6WzBUIj4wTFehXCw" // nolint:lll

	err := vp.SetCredentials(vc, vcJWS)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to set credentials of VP: %w", err))
	}

	//Output:
}

func ExamplePresentation_MarshalJSON() {
	vp := &verifiable.Presentation{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1"},
		ID:     "urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c",
		Type:   []string{"VerifiablePresentation"},
		Holder: "did:example:ebfeb1f712ebc6f1c276e12ec21",
	}

	vc := verifiable.Credential{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/examples/v1"},
		ID: "http://example.edu/credentials/1872",
		Types: []string{
			"VerifiableCredential",
			"UniversityDegreeCredential"},
		Subject: UniversityDegreeSubject{
			ID:     "did:example:ebfeb1f712ebc6f1c276e12ec21",
			Name:   "Jayden Doe",
			Spouse: "did:example:c276e12ec21ebfeb1f712ebc6f1",
			Degree: UniversityDegree{
				Type:       "BachelorDegree",
				University: "MIT",
			},
		},
		Issuer: verifiable.Issuer{
			ID:   "did:example:76e12ec712ebc6f1c221ebfeb1f",
			Name: "Example University",
		},
		Issued:  &issued,
		Expired: &expired,
		Schemas: []verifiable.TypedID{},
		CustomFields: map[string]interface{}{
			"referenceNumber": 83294847,
		},
	}

	err := vp.SetCredentials(vc)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to set credentials of VP: %w", err))
	}

	// json.MarshalIndent() calls Presentation.MarshalJSON()
	vpJSON, err := json.MarshalIndent(vp, "", "\t")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to marshal VP to JSON: %w", err))
	}

	fmt.Println(string(vpJSON))

	// Output: {
	//	"@context": [
	//		"https://www.w3.org/2018/credentials/v1"
	//	],
	//	"id": "urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c",
	//	"type": [
	//		"VerifiablePresentation"
	//	],
	//	"verifiableCredential": [
	//		{
	//			"Context": [
	//				"https://www.w3.org/2018/credentials/v1",
	//				"https://www.w3.org/2018/credentials/examples/v1"
	//			],
	//			"CustomContext": null,
	//			"ID": "http://example.edu/credentials/1872",
	//			"Types": [
	//				"VerifiableCredential",
	//				"UniversityDegreeCredential"
	//			],
	//			"Subject": {
	//				"id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
	//				"name": "Jayden Doe",
	//				"spouse": "did:example:c276e12ec21ebfeb1f712ebc6f1",
	//				"degree": {
	//					"type": "BachelorDegree",
	//					"university": "MIT"
	//				}
	//			},
	//			"Issuer": {
	//				"ID": "did:example:76e12ec712ebc6f1c221ebfeb1f",
	//				"Name": "Example University"
	//			},
	//			"Issued": "2010-01-01T19:23:24Z",
	//			"Expired": "2020-01-01T19:23:24Z",
	//			"Proof": null,
	//			"Status": null,
	//			"Schemas": [],
	//			"Evidence": null,
	//			"TermsOfUse": null,
	//			"RefreshService": null,
	//			"CustomFields": {
	//				"referenceNumber": 83294847
	//			}
	//		}
	//	],
	//	"holder": "did:example:ebfeb1f712ebc6f1c276e12ec21"
	//}
}

//nolint:gocyclo
func ExamplePresentation_MarshalledCredentials() {
	vp := &verifiable.Presentation{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1"},
		ID:     "urn:uuid:3978344f-8596-4c3a-a978-8fcaba3903c",
		Type:   []string{"VerifiablePresentation"},
		Holder: "did:example:ebfeb1f712ebc6f1c276e12ec21",
	}

	vc := verifiable.Credential{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/examples/v1"},
		ID: "http://example.edu/credentials/1872",
		Types: []string{
			"VerifiableCredential",
			"UniversityDegreeCredential"},
		Subject: UniversityDegreeSubject{
			ID:     "did:example:ebfeb1f712ebc6f1c276e12ec21",
			Name:   "Jayden Doe",
			Spouse: "did:example:c276e12ec21ebfeb1f712ebc6f1",
			Degree: UniversityDegree{
				Type:       "BachelorDegree",
				University: "MIT",
			},
		},
		Issuer: verifiable.Issuer{
			ID:   "did:example:76e12ec712ebc6f1c221ebfeb1f",
			Name: "Example University",
		},
		Issued:  &issued,
		Expired: &expired,
		Schemas: []verifiable.TypedID{},
		CustomFields: map[string]interface{}{
			"referenceNumber": 83294847,
		},
	}

	// Put JWS form of VC into VP.
	jwtClaims, err := vc.JWTClaims(true)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to set credentials of VP: %w", err))
	}

	vcJWS, err := jwtClaims.MarshalJWS(verifiable.EdDSA, privIssuerKey, "i-kid")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to sign VC JWT: %w", err))
	}

	err = vp.SetCredentials(vcJWS)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to set credentials of VP: %w", err))
	}

	// Marshal VP to JWS as well.
	vpJWS, err := vp.JWTClaims(nil, true).MarshalJWS(verifiable.EdDSA, privHolderKey, "h-kid")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to sign VP inside JWT: %w", err))
	}

	// Decode VP from JWS.
	// Note that VC-s inside will be decoded as well. If they are JWS, their signature is verified
	// and thus we need to make sure the public key fetcher can access the
	vp, err = verifiable.NewPresentation(
		[]byte(vpJWS),
		verifiable.WithPresPublicKeyFetcher(func(issuerID, keyID string) (interface{}, error) {
			switch issuerID {
			case "did:example:76e12ec712ebc6f1c221ebfeb1f":
				return privIssuerKey.Public(), nil
			case "did:example:ebfeb1f712ebc6f1c276e12ec21":
				return privHolderKey.Public(), nil
			default:
				return nil, fmt.Errorf("unexpected key: %s", keyID)
			}
		}))
	if err != nil {
		fmt.Println(fmt.Errorf("failed to decode VP JWS: %w", err))
	}

	// Get credentials in binary form.
	vpCreds, err := vp.MarshalledCredentials()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to get marshalled credentials from decoded presentation: %w", err))
	}

	if len(vpCreds) != 1 {
		fmt.Println("Expected 1 credential inside presentation")
	}

	// Decoded credential. Note that no public key fetcher is passed as the VC was already decoded (and proof verified)
	// when VP was decoded.
	vcDecoded, _, err := verifiable.NewCredential(vpCreds[0])
	if err != nil {
		fmt.Println(fmt.Errorf("failed to decode VC: %w", err))
	}

	vcDecodedJSON, err := json.MarshalIndent(vcDecoded, "", "\t")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to marshal VP to JSON: %w", err))
	}

	fmt.Println(string(vcDecodedJSON))

	// Output: {
	//	"@context": [
	//		"https://www.w3.org/2018/credentials/v1",
	//		"https://www.w3.org/2018/credentials/examples/v1"
	//	],
	//	"credentialSchema": [],
	//	"credentialSubject": {
	//		"degree": {
	//			"type": "BachelorDegree",
	//			"university": "MIT"
	//		},
	//		"id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
	//		"name": "Jayden Doe",
	//		"spouse": "did:example:c276e12ec21ebfeb1f712ebc6f1"
	//	},
	//	"expirationDate": "2020-01-01T19:23:24Z",
	//	"id": "http://example.edu/credentials/1872",
	//	"issuanceDate": "2010-01-01T19:23:24Z",
	//	"issuer": {
	//		"id": "did:example:76e12ec712ebc6f1c221ebfeb1f",
	//		"name": "Example University"
	//	},
	//	"referenceNumber": 83294847,
	//	"type": [
	//		"VerifiableCredential",
	//		"UniversityDegreeCredential"
	//	]
	//}
}