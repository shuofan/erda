// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package topology

import (
	"testing"

	"github.com/doug-martin/goqu/v9"
	"github.com/stretchr/testify/assert"
)

func Test_queryConditionsV2(t *testing.T) {
	type args struct {
		params  Vo
		tagInfo *TagInfo
	}
	tests := []struct {
		name    string
		args    args
		wantsql string
	}{
		{
			name: "without ApplicationName",
			args: args{
				params: Vo{
					TerminusKey: "123",
					StartTime:   1660192621035,
					EndTime:     1660196221035,
				},
				tagInfo: &TagInfo{},
			},
			wantsql: `SELECT * FROM "table" WHERE (("tenant_id" = '123') AND ("timestamp" >= fromUnixTimestamp64Milli(toInt64(1660192621035))) AND ("timestamp" <= fromUnixTimestamp64Milli(toInt64(1660196221035))) AND (tag_values[indexOf(tag_keys, 'component')] NOT IN ('registerCenter', 'configCenter', 'noticeCenter')) AND (tag_values[indexOf(tag_keys, 'target_addon_type')] NOT IN ('registerCenter', 'configCenter', 'noticeCenter')))`,
		},
		{
			name: "with ApplicationName",
			args: args{
				params: Vo{
					TerminusKey: "123",
					StartTime:   1660192621035,
					EndTime:     1660196221035,
				},
				tagInfo: &TagInfo{
					ApplicationName: "hello",
				},
			},
			wantsql: `SELECT * FROM "table" WHERE (("tenant_id" = '123') AND ("timestamp" >= fromUnixTimestamp64Milli(toInt64(1660192621035))) AND ("timestamp" <= fromUnixTimestamp64Milli(toInt64(1660196221035))) AND (tag_values[indexOf(tag_keys, 'component')] NOT IN ('registerCenter', 'configCenter', 'noticeCenter')) AND (tag_values[indexOf(tag_keys, 'target_addon_type')] NOT IN ('registerCenter', 'configCenter', 'noticeCenter')) AND ((tag_values[indexOf(tag_keys, 'application_name')] = 'hello') OR (tag_values[indexOf(tag_keys, 'target_application_name')] = 'hello') OR (tag_values[indexOf(tag_keys, 'source_application_name')] = 'hello')))`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sd := queryConditionsV2(goqu.From("table"), tt.args.params, tt.args.tagInfo)
			sqlstr, _, err := sd.ToSQL()
			if err != nil {
				assert.Fail(t, "must not error: %s", err)
				return
			}
			assert.Equalf(t, tt.wantsql, sqlstr, "queryConditionsV2(goqu.From(), %v, %v)", tt.args.params, tt.args.tagInfo)
		})
	}
}
