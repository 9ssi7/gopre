package assets

import "embed"

//go:embed mail/*
var emailTemplate embed.FS

func EmbedMailTemplate() embed.FS {
	return emailTemplate
}

type templates struct {
	AuthRegistered string
	AuthVerify     string
	AuthInvited    string
}

var Templates = templates{
	AuthRegistered: "auth/registered",
	AuthVerify:     "auth/verify",
	AuthInvited:    "auth/invited",
}
