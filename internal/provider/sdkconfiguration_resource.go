// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"speakeasy/internal/sdk"
	"speakeasy/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SDKConfigurationResource{}
var _ resource.ResourceWithImportState = &SDKConfigurationResource{}

func NewSDKConfigurationResource() resource.Resource {
	return &SDKConfigurationResource{}
}

// SDKConfigurationResource defines the resource implementation.
type SDKConfigurationResource struct {
	client *sdk.SDK
}

// SDKConfigurationResourceModel describes the resource data model.
type SDKConfigurationResourceModel struct {
	APIKey        types.String            `tfsdk:"api_key"`
	APIID         types.String            `tfsdk:"api_id"`
	Configuration GenerationConfiguration `tfsdk:"configuration"`
	VersionID     types.String            `tfsdk:"version_id"`
	WorkspaceID   types.String            `tfsdk:"workspace_id"`
}

func (r *SDKConfigurationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sdk_configuration"
}

func (r *SDKConfigurationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SDKConfiguration Resource",

		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Required: true,
			},
			"api_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"api_id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"is_monorepo": schema.BoolAttribute{
						Computed: true,
						Optional: true,
					},
					"repo_setup": schema.StringAttribute{
						Computed: true,
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"provided",
								"managed",
							),
						},
					},
					"repo_url": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"schema_auth": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"schema_url": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"version_id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"workspace_id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			"version_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required:    true,
				Description: `Workspace's ID`,
			},
		},
	}
}

func (r *SDKConfigurationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SDKConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SDKConfigurationResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	generationConfigurationRequest := data.ToCreateSDKType()
	apiID := data.APIID.ValueString()
	versionID := data.VersionID.ValueString()
	workspaceID := data.WorkspaceID.ValueString()
	request := operations.UpsertGenerationConfigurationRequest{
		GenerationConfigurationRequest: generationConfigurationRequest,
		APIID:                          apiID,
		VersionID:                      versionID,
		WorkspaceID:                    workspaceID,
	}
	res, err := r.client.SDKs.UpsertGenerationConfiguration(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SDKConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SDKConfigurationResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SDKConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SDKConfigurationResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	generationConfigurationRequest := data.ToUpdateSDKType()
	apiID := data.APIID.ValueString()
	versionID := data.VersionID.ValueString()
	workspaceID := data.WorkspaceID.ValueString()
	request := operations.UpsertGenerationConfigurationRequest{
		GenerationConfigurationRequest: generationConfigurationRequest,
		APIID:                          apiID,
		VersionID:                      versionID,
		WorkspaceID:                    workspaceID,
	}
	res, err := r.client.SDKs.UpsertGenerationConfiguration(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SDKConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SDKConfigurationResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiID := data.APIID.ValueString()
	versionID := data.VersionID.ValueString()
	workspaceID := data.WorkspaceID.ValueString()
	request := operations.DeleteGenerationConfigurationRequest{
		APIID:       apiID,
		VersionID:   versionID,
		WorkspaceID: workspaceID,
	}
	res, err := r.client.SDKs.DeleteGenerationConfiguration(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SDKConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
