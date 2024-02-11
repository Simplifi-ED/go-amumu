package notification

import (
	"go-send/graphhelper"
)

type Observer interface {
	Update(Message, *graphhelper.GraphHelper)
}
