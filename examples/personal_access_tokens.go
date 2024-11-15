//
// Copyright 2022, Ryan Glab <ryan.j.glab@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gitlab.com/gitlab-org/client-go"
)

func patRevokeExample() {
	git, err := gitlab.NewClient("glpat-123xyz")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := git.PersonalAccessTokens.RevokePersonalAccessToken(99999999)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response Code: %s", resp.Status)
}

func patListExampleWithUserFilter() {
	git, err := gitlab.NewClient("glpat-123xyz")
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListPersonalAccessTokensOptions{
		ListOptions: gitlab.ListOptions{Page: 1, PerPage: 10},
		UserID:      gitlab.Ptr(12345),
	}

	personalAccessTokens, _, err := git.PersonalAccessTokens.ListPersonalAccessTokens(opt)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(personalAccessTokens)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found personal access tokens: %s", data)
}

func patRotateExample() {
	git, err := gitlab.NewClient("glpat-123xyz")
	if err != nil {
		log.Fatal(err)
	}

	expiry := gitlab.ISOTime(time.Date(2023, time.August, 15, 0, 0, 0, 0, time.UTC))
	opts := &gitlab.RotatePersonalAccessTokenOptions{
		ExpiresAt: &expiry,
	}
	newPersonalAccessToken, _, err := git.PersonalAccessTokens.RotatePersonalAccessToken(12345, opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Your new token is %s\n", newPersonalAccessToken.Token)
}
