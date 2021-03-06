// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package deploymentapi

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestResync(t *testing.T) {
	type args struct {
		params ResyncParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Fails due to parameter validation (Cluster ID)",
			args: args{},
			wantErr: multierror.NewPrefixed("deployment resync",
				errors.New("deployment params: api reference is required for the operation"),
				errors.New(`deployment params: id "" is invalid`),
			),
		},
		{
			name: "Fails due to parameter validation (API)",
			args: args{params: ResyncParams{
				ID: "d324608c97154bdba2dff97511d40368",
			}},
			wantErr: multierror.NewPrefixed("deployment resync",
				errors.New("deployment params: api reference is required for the operation"),
			),
		},
		{
			name: "Fails due to unknown API response",
			args: args{params: ResyncParams{
				ID: "2c221bd86b7f48959a59ee3128d5c5e8",
				API: api.NewMock(mock.Response{Response: http.Response{
					StatusCode: http.StatusForbidden,
					Body:       mock.NewStringBody(`{"error": "some forbidden error"}`),
				}}),
			}},
			wantErr: errors.New(`{"error": "some forbidden error"}`),
		},
		{
			name: "Fails due to API error",
			args: args{params: ResyncParams{
				ID: "2c221bd86b7f48959a59ee3128d5c5e8",
				API: api.NewMock(mock.Response{
					Error: errors.New("error with API"),
				}),
			}},
			wantErr: &url.Error{
				Op:  "Post",
				URL: "https://mock.elastic.co/api/v1/deployments/2c221bd86b7f48959a59ee3128d5c5e8/_resync",
				Err: errors.New("error with API"),
			},
		},
		{
			name: "Succeeds to resynchronize a deployment without errors",
			args: args{params: ResyncParams{
				ID: "d324608c97154bdba2dff97511d40368",
				API: api.NewMock(mock.Response{Response: http.Response{
					StatusCode: http.StatusOK,
					Body:       mock.NewStringBody(`{}`),
				}}),
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Resync(tt.args.params); !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Resync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResyncAll(t *testing.T) {
	type args struct {
		params ResyncAllParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		want    *models.IndexSynchronizationResults
	}{
		{
			name:    "Fails due to parameter validation (API)",
			args:    args{params: ResyncAllParams{}},
			wantErr: errors.New("api reference is required for the operation"),
		},
		{
			name: "Fails due to unknown API response",
			args: args{params: ResyncAllParams{
				API: api.NewMock(mock.Response{Response: http.Response{
					StatusCode: http.StatusForbidden,
					Body:       mock.NewStringBody(`{"error": "some forbidden error"}`),
				}}),
			}},
			wantErr: errors.New(`{"error": "some forbidden error"}`),
		},
		{
			name: "Fails due to API error",
			args: args{params: ResyncAllParams{
				API: api.NewMock(mock.Response{
					Error: errors.New("error with API"),
				}),
			}},
			wantErr: &url.Error{
				Op:  "Post",
				URL: "https://mock.elastic.co/api/v1/deployments/_resync",
				Err: errors.New("error with API"),
			},
		},
		{
			name: "Succeeds to resynchronize all deployments without errors",
			args: args{params: ResyncAllParams{
				API: api.NewMock(mock.Response{Response: http.Response{
					StatusCode: http.StatusOK,
					Body:       mock.NewStringBody(`{}`),
				}}),
			}},
			want: &models.IndexSynchronizationResults{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResyncAll(tt.args.params)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("ResyncAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResyncAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
