package service

func (s *Service) AuthToken(token string) (uint, *ErrMsg) {
	uID, err := s.kp.ParseToken(token)
	if err != nil {
		return 0, EM_CommonFail_AuthFail.PutDesc(err.Error())
	}
	return uID, nil
}
