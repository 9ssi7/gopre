package config

type roles struct {
	Admin      string
	AdminSuper string
	Contact    contactRoles
}

type contactRoles struct {
	List     string
	MarkSeen string
	Super    string
}

var Roles = roles{
	Admin:      "admin",
	AdminSuper: "admin_super",
	Contact: contactRoles{
		List:     "contact_list",
		MarkSeen: "contact_mark_seen",
		Super:    "contact_super",
	},
}
