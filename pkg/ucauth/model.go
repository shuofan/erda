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

type OryKratosSession struct {
	ID       string            `json:"id"`
	Active   bool              `json:"active"`
	Identity OryKratosIdentity `json:"identity"`
}

type OryKratosIdentity struct {
	ID       USERID                  `json:"id"`
	SchemaID string                  `json:"schema_id"`
	State    string                  `json:"state"`
	Traits   OryKratosIdentityTraits `json:"traits"`
}

type OryKratosIdentityTraits struct {
	Email string                      `json:"email"`
	Name  OryKratosIdentityTraitsName `json:"name"`
	Nick  string                      `json:"nickName"`
	Phone string                      `json:"phone"`
}

type OryKratosIdentityTraitsName struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type OryKratosFlowResponse struct {
	ID string                  `json:"id"`
	UI OryKratosFlowResponseUI `json:"ui"`
}

type OryKratosFlowResponseUI struct {
	Action string `json:"action"`
}

type OryKratosRegistrationRequest struct {
	Traits   OryKratosIdentityTraits `json:"traits"`
	Password string                  `json:"password"`
	Method   string                  `json:"method"`
}

type OryKratosRegistrationResponse struct {
	Identity OryKratosIdentity `json:"identity"`
}

type OryKratosLoginRequest struct {
	Identify string `json:"password_identifier"`
	Password string `json:"password"`
	Method   string `json:"method"`
}

type OryKratosLoginResponse struct {
	SessionToken string `json:"session_token"`
}

type OryKratosUpdateIdentitiyRequest struct {
	State  string                  `json:"state"`
	Traits OryKratosIdentityTraits `json:"traits"`
}

var oryKratosStateMap = map[int]string{
	0: "active",
	1: "inactive",
}

func nameConversion(name OryKratosIdentityTraitsName) string {
	// TODO: eastern name vs western name
	return name.Last + name.First
}

func identityToUser(i OryKratosIdentity) User {
	return User{
		ID:    string(i.ID),
		Name:  nameConversion(i.Traits.Name),
		Nick:  i.Traits.Nick,
		Email: i.Traits.Email,
		Phone: i.Traits.Phone,
		State: i.State,
	}
}

func identityToUserInfo(i OryKratosIdentity) UserInfo {
	return userToUserInfo(identityToUser(i))
}

func userToUserInfo(u User) UserInfo {
	return UserInfo{
		ID:        USERID(u.ID),
		Email:     u.Email,
		Phone:     u.Phone,
		AvatarUrl: u.AvatarURL,
		UserName:  u.Name,
		NickName:  u.Nick,
		Enabled:   true,
	}
}

func userToUserInPaging(u User) userInPaging {
	return userInPaging{
		Id:       u.ID,
		Avatar:   u.AvatarURL,
		Username: u.Name,
		Nickname: u.Nick,
		Mobile:   u.Phone,
		Email:    u.Email,
		Enabled:  true,
		Locked:   u.State == "inactive",
		// TODO: LastLoginAt PwdExpireAt
	}
}
