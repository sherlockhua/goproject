package config


type Notifyer interface {
	Callback(*Config)
}
