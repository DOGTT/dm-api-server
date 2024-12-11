package service

import "strings"

func (s *Service) AuthToken(token string) (uint, *ErrMsg) {
	tokenSpt := strings.Split(token, " ")
	if len(tokenSpt) != 2 {
		return 0, EM_CommonFail_AuthFail.PutDesc("invalid token")
	}
	tokenType := tokenSpt[0]
	tokenData := tokenSpt[1]
	if tokenType != "Bearer" {
		return 0, EM_CommonFail_AuthFail.PutDesc("invalid token type")
	}
	uID, err := s.kp.ParseToken(tokenData)
	if err != nil {
		return 0, EM_CommonFail_AuthFail.PutDesc(err.Error())
	}
	return uID, nil
}
