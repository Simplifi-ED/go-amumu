package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChatsGetAllMessagesResponse 
// Deprecated: This class is obsolete. Use getAllMessagesGetResponse instead.
type ItemChatsGetAllMessagesResponse struct {
    ItemChatsGetAllMessagesGetResponse
}
// NewItemChatsGetAllMessagesResponse instantiates a new ItemChatsGetAllMessagesResponse and sets the default values.
func NewItemChatsGetAllMessagesResponse()(*ItemChatsGetAllMessagesResponse) {
    m := &ItemChatsGetAllMessagesResponse{
        ItemChatsGetAllMessagesGetResponse: *NewItemChatsGetAllMessagesGetResponse(),
    }
    return m
}
// CreateItemChatsGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsGetAllMessagesResponse(), nil
}
// ItemChatsGetAllMessagesResponseable 
// Deprecated: This class is obsolete. Use getAllMessagesGetResponse instead.
type ItemChatsGetAllMessagesResponseable interface {
    ItemChatsGetAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}