/*
Copyright 2021 The KodeRover Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gitlab

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

// ListBranches lists branches by projectID <- urlEncode(namespace/projectName)
func (c *Client) ListBranches(owner, repo, key string, opts *ListOptions) ([]*gitlab.Branch, error) {
	if opts.Page == 0 {
		opts.Page = 1
	}
	if opts.PerPage == 0 {
		opts.PerPage = 100
	}

	got := 0
	limit := 100
	req := opts.PerPage
	opts.Page = 1
	opts.PerPage = limit
	res := []*gitlab.Branch{}

	for got < req {
		branches, err := wrap(paginated(func(o *gitlab.ListOptions) ([]interface{}, *gitlab.Response, error) {
			bs, r, err := c.Branches.ListBranches(generateProjectName(owner, repo), &gitlab.ListBranchesOptions{ListOptions: *o, Search: &key})
			var res []interface{}
			for _, b := range bs {
				res = append(res, b)
			}
			return res, r, err
		}, opts))

		if err != nil {
			return nil, err
		}

		bs, ok := branches.([]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to convert branches to []interface{}")
		}
		for _, b := range bs {
			res = append(res, b.(*gitlab.Branch))
		}

		got += len(bs)

		if len(bs) < limit {
			break
		}
		if !opts.MatchBranches {
			break
		}

		opts.Page++
	}

	return res, nil
}
