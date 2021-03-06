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

package depresourceapi

import (
	"github.com/pkg/errors"

	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// StopParams is consumed by Stop.
type StopParams struct {
	Params

	All bool
}

// StopInstancesParams is consumed by StopInstances.
type StopInstancesParams struct {
	StopParams
	IgnoreMissing *bool
	InstanceIDs   []string
}

// Validate ensures the parameters are usable by StopInstances.
func (params *StopInstancesParams) Validate() error {
	var merr = multierror.NewPrefixed("deployment stop")
	if len(params.InstanceIDs) == 0 {
		merr = merr.Append(errors.New("at least 1 instance ID must be provided"))
	}

	merr = merr.Append(params.StopParams.Validate())

	return merr.ErrorOrNil()
}

// Stop stops all instances belonging to a deployment resource kind.
func Stop(params StopParams) (models.DeploymentResourceCommandResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.V1API.Deployments.StopDeploymentResourceInstancesAll(
		deployments.NewStopDeploymentResourceInstancesAllParams().
			WithDeploymentID(params.DeploymentID).
			WithResourceKind(params.Kind).
			WithRefID(params.RefID),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Unwrap(err)
	}

	return res.Payload, nil
}

// StopInstances stops defined instances belonging to a deployment resource.
func StopInstances(params StopInstancesParams) (models.DeploymentResourceCommandResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.V1API.Deployments.StopDeploymentResourceInstances(
		deployments.NewStopDeploymentResourceInstancesParams().
			WithDeploymentID(params.DeploymentID).
			WithResourceKind(params.Kind).
			WithIgnoreMissing(params.IgnoreMissing).
			WithInstanceIds(params.InstanceIDs).
			WithRefID(params.RefID),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Unwrap(err)
	}

	return res.Payload, nil
}

// StopAllOrSpecified stops all or defined instances belonging to a deployment resource.
func StopAllOrSpecified(params StopInstancesParams) (models.DeploymentResourceCommandResponse, error) {
	if params.All {
		res, err := Stop(params.StopParams)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	res, err := StopInstances(params)
	if err != nil {
		return nil, err
	}
	return res, nil
}
