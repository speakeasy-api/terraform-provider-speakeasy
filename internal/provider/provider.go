// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"speakeasy/internal/sdk"
	"speakeasy/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &SpeakeasyProvider{}

type SpeakeasyProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// SpeakeasyProviderModel describes the provider data model.
type SpeakeasyProviderModel struct {
	ServerURL types.String `tfsdk:"server_url"`
	APIKey    types.String `tfsdk:"api_key"`
}

func (p *SpeakeasyProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "speakeasy"
	resp.Version = p.version
}

func (p *SpeakeasyProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://api.prod.speakeasyapi.dev)",
				Optional:            true,
				Required:            false,
			},
			"api_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *SpeakeasyProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data SpeakeasyProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://api.prod.speakeasyapi.dev"
	}

	apiKey := data.APIKey.ValueString()
	security := shared.Security{
		APIKey: apiKey,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *SpeakeasyProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAPIResource,
		NewAPIKeyResource,
		NewPluginResource,
		NewPortalResource,
		NewSDKConfigurationResource,
	}
}

func (p *SpeakeasyProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SpeakeasyProvider{
			version: version,
		}
	}
}
