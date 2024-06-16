package models

import "github.com/MaximoGCH/space-colony-game/game/common/animations"

const (
	Text = iota
	Title
)

type NotificationType int

type Notification struct {
	Text   string
	Type   NotificationType
	Height *animations.OneDimensionAnimation
	Remove bool
}

type NotificationSystem []*Notification

func (notificationSystem *NotificationSystem) Clear() {
	def := *notificationSystem
	for _, item := range def {
		item.Remove = true
	}
	*notificationSystem = def
}

func (notificationSystem *NotificationSystem) Add(nType NotificationType, text string) {
	def := *notificationSystem
	def = append(def, &Notification{
		Text:   text,
		Type:   nType,
		Height: animations.NewOneDimensionAnimation(0, 100),
		Remove: false,
	})
	*notificationSystem = def
}

func (notificationSystem *NotificationSystem) RemoveIndex(index int) {
	def := *notificationSystem
	*notificationSystem = append(def[:index], def[index+1:]...)
}
