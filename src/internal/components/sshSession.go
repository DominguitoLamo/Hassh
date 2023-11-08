package components

import (
	"hassh/src/internal/svc"

	switchgo "github.com/DominguitoLamo/switchGo"
)

func InitSSHSession(ctx *svc.ServiceContext) {
	sshManager := switchgo.NewSessionManager()
	ctx.Components.SSHManager = sshManager
}