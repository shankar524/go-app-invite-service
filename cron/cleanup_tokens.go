package cron

import (
	"github.com/robfig/cron"
	"github.com/shankar524/go-app-invite-service/api/service"
)

type Cron struct {
	Job *cron.Cron
}

const TOKEN_CLEANUP_AFTER_DAYS = 7

func NewCleanupTokens(service service.ITokenService) Cron {
	c := cron.New()
	c.AddFunc("0 0 * * *", func() {
		service.InvalidateToken(TOKEN_CLEANUP_AFTER_DAYS)
	})
	return Cron{c}
}
