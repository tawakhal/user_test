package repository

import (
	"testing"

	"github.com/rs/zerolog/log"
	cs "github.com/tawakhal/user_test/constanta"
	"github.com/tawakhal/user_test/extend/conf"
	"github.com/tawakhal/user_test/extend/logger"
)

func TestSetup(t *testing.T) {
	t.Fail()

	const (
		cfgFile = "../config/dev.yaml"
	)
	conf.Setup(cfgFile)

	logger.Setup()

	err := Setup()
	if err != nil {
		log.Err(err).Msg("Failed Setup Database")
		return
	}

	usr := &User{
		UserName: "olgi",
		Password: "1234",
		Email:    "olgi_tester@tester.com",
		Avatar:   "tester.com/avatar/test.png",
		Status:   cs.StatusActive,
	}
	id, err := usr.Insert()

	log.Info().Msgf("%d : %v", id, err)

}
