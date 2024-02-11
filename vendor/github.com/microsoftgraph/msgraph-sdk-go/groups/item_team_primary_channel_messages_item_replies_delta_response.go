package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamPrimaryChannelMessagesItemRepliesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamPrimaryChannelMessagesItemRepliesDeltaResponse struct {
    ItemTeamPrimaryChannelMessagesItemRepliesDeltaGetResponse
}
// NewItemTeamPrimaryChannelMessagesItemRepliesDeltaResponse instantiates a new ItemTeamPrimaryChannelMessagesItemRepliesDeltaResponse and sets the default values.
func NewItemTeamPrimaryChannelMessagesItemRepliesDeltaResponse()(*ItemTeamPrimaryChannelMessagesItemRepliesDeltaResponse) {
    m := &ItemTeamPrimaryChannelMessagesItemRepliesDeltaResponse{
        ItemTeamPrimaryChannelMessagesItemRepliesDeltaGetResponse: *NewItemTeamPrimaryChannelMessagesItemRepliesDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamPrimaryChannelMessagesItemRepliesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamPrimaryChannelMessagesItemRepliesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPrimaryChannelMessagesItemRepliesDeltaResponse(), nil
}
// ItemTeamPrimaryChannelMessagesItemRepliesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamPrimaryChannelMessagesItemRepliesDeltaResponseable interface {
    ItemTeamPrimaryChannelMessagesItemRepliesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
