// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package iam

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var _ resource.Resource = (*assignMachineUserResource)(nil)

func NewAssignMachineUserResource() resource.Resource {
	return &assignMachineUserResource{}
}

type assignMachineUserResource struct {
	client *cdp.Client
}

func (r *assignMachineUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = AssignMachineUserSchema
}

func (r *assignMachineUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_machine_user_role_assignment"
}

func (r *assignMachineUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data assignMachineUserResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := operations.NewAssignMachineUserRoleParamsWithContext(ctx).WithInput(
		&models.AssignMachineUserRoleRequest{
			MachineUserName: data.MachineUser.ValueStringPointer(),
			Role:            data.Role.ValueStringPointer(),
		})

	_, err := r.client.Iam.Operations.AssignMachineUserRole(request) // void method, does not have any return value
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "assign Machine User Role")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *assignMachineUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data assignMachineUserResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	params := operations.NewListMachineUserAssignedRolesParamsWithContext(ctx)
	params.WithInput(&models.ListMachineUserAssignedRolesRequest{
		MachineUserName: data.MachineUser.ValueStringPointer(),
	})

	machineUser, err := r.client.Iam.Operations.ListMachineUserAssignedRoles(params)
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "list Machine User Assigned Roles")
		return
	}

	hasAssignedRole := false
	for _, crn := range machineUser.Payload.RoleCrns {
		if crn == data.Role.ValueString() {
			resp.State.Set(ctx, &data)
			hasAssignedRole = true
			break
		}
	}

	if !hasAssignedRole {
		resp.Diagnostics.AddError("role", "Machine User does not have the specified role assigned")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *assignMachineUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not supported yet.")
}

func (r *assignMachineUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data assignMachineUserResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := operations.NewUnassignMachineUserRoleParamsWithContext(ctx).WithInput(
		&models.UnassignMachineUserRoleRequest{
			MachineUserName: data.MachineUser.ValueStringPointer(),
			Role:            data.Role.ValueStringPointer(),
		},
	)

	_, err := r.client.Iam.Operations.UnassignMachineUserRole(request) // void method, does not have any return value
	if err != nil {
		utils.AddIamDiagnosticsError(err, &resp.Diagnostics, "un-assign Machine User Role")
		return
	}
}
