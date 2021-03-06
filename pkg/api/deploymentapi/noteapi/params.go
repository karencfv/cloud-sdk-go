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

package noteapi

import (
	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi/deputil"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

var (
	errEmptyNoteMessage = "note comment cannot be empty"
	errEmptyUserID      = "user id cannot be empty"
	errEmptyNoteID      = "note id cannot be empty"
)

// Params is used on Get and Update Notes
type Params struct {
	*api.API
	ID     string
	Region string
}

// Validate ensures that the parameters are usable by the consuming function.
func (params Params) Validate() error {
	var merr = multierror.NewPrefixed("deployment note")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if len(params.ID) != 32 {
		merr = merr.Append(deputil.NewInvalidDeploymentIDError(params.ID))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}
