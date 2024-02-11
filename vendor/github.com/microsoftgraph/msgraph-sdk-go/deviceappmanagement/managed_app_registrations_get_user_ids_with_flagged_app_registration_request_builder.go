package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder provides operations to call the getUserIdsWithFlaggedAppRegistration method.
type ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetQueryParameters invoke function getUserIdsWithFlaggedAppRegistration
type ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetQueryParameters
}
// NewManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal instantiates a new GetUserIdsWithFlaggedAppRegistrationRequestBuilder and sets the default values.
func NewManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    m := &ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/managedAppRegistrations/getUserIdsWithFlaggedAppRegistration(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder instantiates a new GetUserIdsWithFlaggedAppRegistrationRequestBuilder and sets the default values.
func NewManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getUserIdsWithFlaggedAppRegistration
// Deprecated: This method is obsolete. Use GetAsGetUserIdsWithFlaggedAppRegistrationGetResponse instead.
func (m *ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration)(ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationResponseable), nil
}
// GetAsGetUserIdsWithFlaggedAppRegistrationGetResponse invoke function getUserIdsWithFlaggedAppRegistration
func (m *ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder) GetAsGetUserIdsWithFlaggedAppRegistrationGetResponse(ctx context.Context, requestConfiguration *ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration)(ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationGetResponseable), nil
}
// ToGetRequestInformation invoke function getUserIdsWithFlaggedAppRegistration
func (m *ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder) WithUrl(rawUrl string)(*ManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    return NewManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
