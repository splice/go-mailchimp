package status

const (
	// Subscribed - This address is on the list and ready to receive email. You can only send campaigns to ‘subscribed’ addresses.
	Subscribed = "subscribed"
	// Unsubscribed - This address is on the list and ready to receive email. You can only send campaigns to ‘subscribed’ addresses.
	Unsubscribed = "unsubscribed"
	// Pending - This address used to be on the list but isn’t anymore.
	Pending = "pending"
	// Cleaned - This address requested to be added with double-opt-in but hasn’t confirmed their subscription yet.
	Cleaned = "cleaned"
)
