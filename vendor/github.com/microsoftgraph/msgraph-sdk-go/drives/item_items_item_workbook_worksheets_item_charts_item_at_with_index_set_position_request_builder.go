package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder provides operations to call the setPosition method.
type ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilderInternal instantiates a new SetPositionRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/itemAt(index={index})/setPosition", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder instantiates a new SetPositionRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post positions the chart relative to cells on the worksheet.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chart-setposition?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation positions the chart relative to cells on the worksheet.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemAtWithIndexSetPositionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
