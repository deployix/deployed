package v1

type PromotionsConfig struct {
	// Pause Promotions: Bool flag to indicate if promotion should be paused
	PausePromotion     bool               `yaml:"pausePromotion"`
	NotificationConfig NotificationConfig `yaml:"notificationConfig"`
}
