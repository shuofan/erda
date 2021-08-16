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
	"net/http"
	"strings"

	"github.com/erda-project/erda/pkg/http/httpclient"
)

func whoami(kratosPublicAddr string, sessionID string) (UserInfo, error) {
	var s OryKratosSession
	r, err := httpclient.New(httpclient.WithCompleteRedirect()).
		Get(kratosPublicAddr).
		Cookie(&http.Cookie{
			Name:  "ory_kratos_session",
			Value: sessionID,
		}).
		Path("/sessions/whoami").
		Do().JSON(&s)
	if err != nil {
		return UserInfo{}, err
	}
	if !r.IsOK() {
		return UserInfo{}, fmt.Errorf("bad session")
	}
	return identityToUserInfo(s.Identity), nil
}

func getUserByID(kratosPrivateAddr string, userID string) (*User, error) {
	i, err := getIdentity(kratosPrivateAddr, userID)
	if err != nil {
		return nil, err
	}
	u := identityToUser(*i)
	return &u, nil
}

func getUserByIDs(kratosPrivateAddr string, userIDs []string) ([]User, error) {
	var users []User
	for _, id := range userIDs {
		u, err := getUserByID(kratosPrivateAddr, id)
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}
	return users, nil
}

func getUserPage(kratosPrivateAddr string, page, perPage int) ([]User, error) {
	i, err := getIdentityPage(kratosPrivateAddr, page, perPage)
	if err != nil {
		return nil, err
	}
	var users []User
	for _, u := range i {
		users = append(users, identityToUser(*u))
	}
	return users, nil
}

func getUserByKey(kratosPrivateAddr string, key string) ([]User, error) {
	p := 1
	size := 1000
	cnt := 0
	var users []User
	for {
		ul, err := getUserPage(kratosPrivateAddr, p, size)
		if err != nil {
			return nil, err
		}
		for _, u := range ul {
			if strings.Contains(u.Name, key) || strings.Contains(u.Nick, key) || strings.Contains(u.Email, key) {
				users = append(users, u)
				cnt++
			}
		}
		if cnt >= 10 {
			return users, nil
		}
		p++
		if p > 100 {
			return users, nil
		}
	}
}

func CreateUser(kratosPublicAddr string, req OryKratosRegistrationRequest) error {
	var rsp OryKratosFlowResponse
	r, err := httpclient.New(httpclient.WithCompleteRedirect()).
		Get(kratosPublicAddr).
		Path("/self-service/registration/api").
		Do().JSON(&rsp)
	if err != nil {
		return err
	}
	if !r.IsOK() {
		return fmt.Errorf("bad session")
	}

	var register OryKratosRegistrationResponse
	r, err = httpclient.New(httpclient.WithCompleteRedirect()).
		Post(kratosPublicAddr).
		Path("/self-service/registration").
		Param("flow", rsp.ID).
		JSONBody(req).
		Do().JSON(&register)
	if err != nil {
		return err
	}
	if !r.IsOK() {
		return fmt.Errorf("bad session")
	}
	return nil
}
