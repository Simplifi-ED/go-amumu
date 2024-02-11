package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartAreaFormat 
type WorkbookChartAreaFormat struct {
    Entity
}
// NewWorkbookChartAreaFormat instantiates a new workbookChartAreaFormat and sets the default values.
func NewWorkbookChartAreaFormat()(*WorkbookChartAreaFormat) {
    m := &WorkbookChartAreaFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartAreaFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAreaFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartAreaFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAreaFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartFillFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFill(val.(WorkbookChartFillable))
        }
        return nil
    }
    res["font"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartFontFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFont(val.(WorkbookChartFontable))
        }
        return nil
    }
    return res
}
// GetFill gets the fill property value. Represents the fill format of an object, which includes background formatting information. Read-only.
func (m *WorkbookChartAreaFormat) GetFill()(WorkbookChartFillable) {
    val, err := m.GetBackingStore().Get("fill")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChartFillable)
    }
    return nil
}
// GetFont gets the font property value. Represents the font attributes (font name, font size, color, etc.) for the current object. Read-only.
func (m *WorkbookChartAreaFormat) GetFont()(WorkbookChartFontable) {
    val, err := m.GetBackingStore().Get("font")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChartFontable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkbookChartAreaFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fill", m.GetFill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("font", m.GetFont())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFill sets the fill property value. Represents the fill format of an object, which includes background formatting information. Read-only.
func (m *WorkbookChartAreaFormat) SetFill(value WorkbookChartFillable)() {
    err := m.GetBackingStore().Set("fill", value)
    if err != nil {
        panic(err)
    }
}
// SetFont sets the font property value. Represents the font attributes (font name, font size, color, etc.) for the current object. Read-only.
func (m *WorkbookChartAreaFormat) SetFont(value WorkbookChartFontable)() {
    err := m.GetBackingStore().Set("font", value)
    if err != nil {
        panic(err)
    }
}
// WorkbookChartAreaFormatable 
type WorkbookChartAreaFormatable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFill()(WorkbookChartFillable)
    GetFont()(WorkbookChartFontable)
    SetFill(value WorkbookChartFillable)()
    SetFont(value WorkbookChartFontable)()
}
