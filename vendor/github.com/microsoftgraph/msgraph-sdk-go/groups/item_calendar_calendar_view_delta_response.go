package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarCalendarViewDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarCalendarViewDeltaResponse struct {
    ItemCalendarCalendarViewDeltaGetResponse
}
// NewItemCalendarCalendarViewDeltaResponse instantiates a new ItemCalendarCalendarViewDeltaResponse and sets the default values.
func NewItemCalendarCalendarViewDeltaResponse()(*ItemCalendarCalendarViewDeltaResponse) {
    m := &ItemCalendarCalendarViewDeltaResponse{
        ItemCalendarCalendarViewDeltaGetResponse: *NewItemCalendarCalendarViewDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarCalendarViewDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarCalendarViewDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarCalendarViewDeltaResponse(), nil
}
// ItemCalendarCalendarViewDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarCalendarViewDeltaResponseable interface {
    ItemCalendarCalendarViewDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}