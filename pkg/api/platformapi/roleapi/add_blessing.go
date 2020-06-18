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

package roleapi

import (
	"context"
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_infrastructure"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// AddBlessingParams is consumed by AddBlessing.
type AddBlessingParams struct {
	*api.API

	Blessing *models.Blessing
	RunnerID string
	ID       string
	Region   string
}

// Validate ensures the parameters are usable.
func (params AddBlessingParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid role add blessing params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Blessing == nil {
		merr = merr.Append(errors.New("blessing definition not specified and is required for this operation"))
	}

	if params.ID == "" {
		merr = merr.Append(errors.New("id not specified and is required for this operation"))
	}

	if params.RunnerID == "" {
		merr = merr.Append(errors.New("runner id not specified and is required for this operation"))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// AddBlessing adds a role blessing to a runner ID.
func AddBlessing(params AddBlessingParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.V1API.PlatformInfrastructure.AddBlueprinterBlessing(
			platform_infrastructure.NewAddBlueprinterBlessingParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithBlueprinterRoleID(params.ID).
				WithRunnerID(params.RunnerID).
				WithBody(params.Blessing),
			params.AuthWriter,
		),
	)
}
