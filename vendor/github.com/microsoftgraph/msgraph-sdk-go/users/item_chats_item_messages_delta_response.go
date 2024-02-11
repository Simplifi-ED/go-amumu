package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChatsItemMessagesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemChatsItemMessagesDeltaResponse struct {
    ItemChatsItemMessagesDeltaGetResponse
}
// NewItemChatsItemMessagesDeltaResponse instantiates a new ItemChatsItemMessagesDeltaResponse and sets the default values.
func NewItemChatsItemMessagesDeltaResponse()(*ItemChatsItemMessagesDeltaResponse) {
    m := &ItemChatsItemMessagesDeltaResponse{
        ItemChatsItemMessagesDeltaGetResponse: *NewItemChatsItemMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemChatsItemMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsItemMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemMessagesDeltaResponse(), nil
}
// ItemChatsItemMessagesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemChatsItemMessagesDeltaResponseable interface {
    ItemChatsItemMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
