//
// Copyright 2021, Sander van Harmelen
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
	"log"

	"gitlab.com/gitlab-org/client-go"
)

func pagination() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}

	for {
		// Get the first page with projects.
		ps, resp, err := git.Projects.ListProjects(opt)
		if err != nil {
			log.Fatal(err)
		}

		// List all the projects we've found so far.
		for _, p := range ps {
			log.Printf("Found project: %s", p.Name)
		}

		// Exit the loop when we've seen all pages.
		if resp.NextPage == 0 {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}
}

func keysetPagination() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			OrderBy:    "id",
			Pagination: "keyset",
			PerPage:    5,
			Sort:       "asc",
		},
		Owned: gitlab.Bool(true),
	}

	options := []gitlab.RequestOptionFunc{}

	for {
		// Get the first page with projects.
		ps, resp, err := git.Projects.ListProjects(opt, options...)
		if err != nil {
			log.Fatal(err)
		}

		// List all the projects we've found so far.
		for _, p := range ps {
			log.Printf("Found project: %s", p.Name)
		}

		// Exit the loop when we've seen all pages.
		if resp.NextLink == "" {
			break
		}

		// Set all query parameters in the next request to values in the
		// returned parameters. This could go along with any existing options.
		options = []gitlab.RequestOptionFunc{
			gitlab.WithKeysetPaginationParameters(resp.NextLink),
		}
	}
}
