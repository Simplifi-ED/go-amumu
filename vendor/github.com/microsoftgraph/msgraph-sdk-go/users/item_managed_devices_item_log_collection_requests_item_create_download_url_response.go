package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponse 
// Deprecated: This class is obsolete. Use createDownloadUrlPostResponse instead.
type ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponse struct {
    ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlPostResponse
}
// NewItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponse instantiates a new ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponse and sets the default values.
func NewItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponse()(*ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponse) {
    m := &ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponse{
        ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlPostResponse: *NewItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlPostResponse(),
    }
    return m
}
// CreateItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponse(), nil
}
// ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponseable 
// Deprecated: This class is obsolete. Use createDownloadUrlPostResponse instead.
type ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponseable interface {
    ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}