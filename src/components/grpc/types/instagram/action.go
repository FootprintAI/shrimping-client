package instagramtype

type ActionType string

const (
	LoginActionType        ActionType = "login"
	GetPostActionType                 = "post"
	GetProfileActionType              = "profile"
	GetTopSearchActionType            = "topsearch"
)
