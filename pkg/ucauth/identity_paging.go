// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package ucauth

import (
	"fmt"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/pkg/http/httpclient"
	"github.com/erda-project/erda/pkg/strutil"
)

func getIdentityPage(kratosPrivateAddr string, page, perPage int) ([]*OryKratosIdentity, error) {
	var i []*OryKratosIdentity
	r, err := httpclient.New(httpclient.WithCompleteRedirect()).
		Get(kratosPrivateAddr).
		Path("/identities").
		Param("page", fmt.Sprintf("%d", page)).
		Param("per_page", fmt.Sprintf("%d", perPage)).
		Do().JSON(&i)
	if err != nil {
		return nil, err
	}
	if !r.IsOK() {
		return nil, fmt.Errorf("bad session")
	}

	return i, nil
}

func getUserList(kratosPrivateAddr string, req *apistructs.UserPagingRequest) ([]User, int, error) {
	var identities []*OryKratosIdentity
	cnt := 0
	p := 1
	size := 1000
	for {
		ul, err := getIdentityPage(kratosPrivateAddr, p, size)
		if err != nil {
			return nil, 0, err
		}
		for _, u := range ul {
			if strutil.ContainsOrEmpty(nameConversion(u.Traits.Name), req.Name) && strutil.ContainsOrEmpty(u.Traits.Nick, req.Nick) &&
				strutil.ContainsOrEmpty(u.Traits.Email, req.Email) && strutil.ContainsOrEmpty(u.Traits.Phone, req.Phone) &&
				strutil.ContainsOrEmpty(u.State, oryKratosStateMap[*req.Locked]) {
				identities = append(identities, u)
				cnt++
			}
		}
		if cnt >= 10 {
			break
		}
		p++
		if p > 100 {
			break
		}
	}

	var users []User
	for _, u := range paginate(identities, req.PageNo, req.PageSize) {
		users = append(users, identityToUser(*u))
	}

	return users, len(identities), nil
}

func paginate(i []*OryKratosIdentity, pageNo int, pageSize int) []*OryKratosIdentity {
	start := (pageNo - 1) * pageSize
	if start > len(i) {
		return nil
	}
	end := start + pageSize
	if end > len(i) {
		return i[start:]
	}
	return i[start:end]
}
