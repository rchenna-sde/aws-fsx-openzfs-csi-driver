/*
Copyright 2023 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloud

import (
	_ "unsafe" // Required for go:linkname

	"github.com/aws/aws-sdk-go-v2/service/fsx"
)

// Link to the unexported validateOpDeleteFileSystemInput function from the AWS SDK
//
//go:linkname validateOpDeleteFileSystemInput github.com/aws/aws-sdk-go-v2/service/fsx.validateOpDeleteFileSystemInput
func validateOpDeleteFileSystemInput(v *fsx.DeleteFileSystemInput) error

// Link to the unexported validateOpDeleteVolumeInput function from the AWS SDK
//
//go:linkname validateOpDeleteVolumeInput github.com/aws/aws-sdk-go-v2/service/fsx.validateOpDeleteVolumeInput
func validateOpDeleteVolumeInput(v *fsx.DeleteVolumeInput) error
