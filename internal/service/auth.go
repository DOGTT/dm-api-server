package service

import (
	"strings"

	"github.com/DOGTT/dm-api-server/internal/utils"
)

func (s *Service) AuthToken(token string) (*utils.TokenClaims, *ErrMsg) {
	tokenSpt := strings.Split(token, " ")
	if len(tokenSpt) != 2 {
		return nil, EM_CommonFail_AuthFail.PutDesc("invalid token")
	}
	tokenType := tokenSpt[0]
	tokenData := tokenSpt[1]
	if tokenType != "Bearer" {
		return nil, EM_CommonFail_AuthFail.PutDesc("invalid token type")
	}
	tc, err := s.kp.ParseToken(tokenData)
	if err != nil {
		return nil, EM_CommonFail_AuthFail.PutDesc(err.Error())
	}
	return &tc, nil
}
