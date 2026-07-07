package main

type Entry struct {
	friendlyName string
	hostName     string
	desccription string
	ipAddress    string
	online       bool
}

// implement list.Item interface
func (e Entry) FilterValue() string {
	return e.friendlyName
}

func (e Entry) Title() string {
	return e.friendlyName
}

func (e Entry) Description() string {
	return e.desccription
}
