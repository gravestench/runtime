package desktop_notification

type SendsNotifications interface {
	Notify(title, message, appIcon string)
	Alert(title, message, appIcon string)
}
