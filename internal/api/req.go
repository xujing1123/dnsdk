// Copyright 2024 dnsdk Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

type (
	DomainListReq struct {
		PageNumber string `url:"page_number,omitempty"` // 页码
		PageSize   string `url:"page_size,omitempty"`   // 每页数量
		Name       string `url:"name,omitempty"`        // 域名名称
	}
	TypeListReq   struct{}
	LineListReq   struct{}
	RecordListReq struct {
		Page      uint   `form:"page"`      // 页码 => 1
		Limit     uint   `form:"limit"`     // 每页数量 => 10
		DomainId  string `form:"domain_id"` // 域名Id => xxxxxxxxxxxx
		Type      string `form:"type"`      // 解析类型 => A
		Value     string `form:"value"`     // 记录值 => 1.1.1.1
		Remark    string `form:"remark"`    // 备注 => created by dnsdk
		Order     string `form:"order"`     // 排序 => type
		Direction string `form:"direction"` // 方向 => asc
	}
	RecordAddReq struct {
		DomainId string `json:"domain_id"` // 域名Id => xxxxxxxxxxxx
		Record   string `json:"record"`    // 主机记录 => www
		Name     string `json:"name"`      // 名称 => www.example.com
		Type     string `json:"type"`      // 类型 => A
		LineId   string `json:"line_id"`   // 线路id => 1
		Value    string `json:"value"`     // 记录值 => 1.1.1.1
		TTL      uint   `json:"ttl"`       // TTL => 60
		Remark   string `json:"remark"`    // 备注 => created by dnsdk
	}
	RecordUpdateReq struct {
		RecordId string `json:"record_id"` // 记录Id => xxxxxxxxxxxx
		DomainId string `json:"domain_id"` // 域名Id => xxxxxxxxxxxx
		Record   string `json:"record"`    // 主机记录 => www
		Name     string `json:"name"`      // 名称 => www.example.com
		Type     string `json:"type"`      // 类型 => A
		LineId   string `json:"line_id"`   // 线路id => 1
		Value    string `json:"value"`     // 记录值 => 1.1.1.1
		TTL      uint   `json:"ttl"`       // TTL => 60
		Remark   string `json:"remark"`    // 备注 => created by dnsdk
	}
	RecordDeleteReq struct {
		RecordId string `json:"record_id"` // 记录Id => xxxxxxxxxxxx
		DomainId string `json:"domain_id"` // 域名Id => xxxxxxxxxxxx
	}
	RecordEnableReq  struct{}
	RecordDisableReq struct{}
)
