package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponse 
// Deprecated: This class is obsolete. Use imageWithWidthWithHeightWithFittingModeGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponse struct {
    ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponse()(*ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponse{
        ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeGetResponse: *NewItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponse(), nil
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponseable 
// Deprecated: This class is obsolete. Use imageWithWidthWithHeightWithFittingModeGetResponse instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemChartsItemWithNameImageWithWidthWithHeightWithFittingModeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}