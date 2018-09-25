// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package credentials_test

import (
	"cloud.google.com/go/iam/credentials/apiv1"
	"golang.org/x/net/context"
	credentialspb "google.golang.org/genproto/googleapis/iam/credentials/v1"
)

func ExampleNewIamCredentialsClient() {
	ctx := context.Background()
	c, err := credentials.NewIamCredentialsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleIamCredentialsClient_GenerateAccessToken() {
	ctx := context.Background()
	c, err := credentials.NewIamCredentialsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &credentialspb.GenerateAccessTokenRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GenerateAccessToken(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamCredentialsClient_GenerateIdToken() {
	ctx := context.Background()
	c, err := credentials.NewIamCredentialsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &credentialspb.GenerateIdTokenRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GenerateIdToken(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamCredentialsClient_SignBlob() {
	ctx := context.Background()
	c, err := credentials.NewIamCredentialsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &credentialspb.SignBlobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SignBlob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamCredentialsClient_SignJwt() {
	ctx := context.Background()
	c, err := credentials.NewIamCredentialsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &credentialspb.SignJwtRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SignJwt(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
