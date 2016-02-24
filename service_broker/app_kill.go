package chaospeddler

//KillPercent - kills a given app by the given percentage
func (s *AppKill) KillPercent(binding ServiceBinding, percentKill int) (err error) {
	s.getAIInfo(binding.AppGUID)
	return
}

func (s *AppKill) getAIInfo(appGUID string) {

}
