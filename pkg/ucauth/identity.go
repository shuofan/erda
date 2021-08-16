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

	"github.com/erda-project/erda/pkg/http/httpclient"
)

func getIdentity(kratosPrivateAddr string, userID string) (*OryKratosIdentity, error) {
	var i OryKratosIdentity
	r, err := httpclient.New(httpclient.WithCompleteRedirect()).
		Get(kratosPrivateAddr).
		Path("/identities/" + userID).
		Do().JSON(&i)
	if err != nil {
		return nil, err
	}
	if !r.IsOK() {
		return nil, fmt.Errorf("bad session")
	}
	return &i, nil
}
