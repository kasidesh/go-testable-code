func (*Requester) Get(url string) (string, error) {
	_, body, errs:= gorequest.New().Get(url).End()
	if len(errs) > 0 {
		return "", errs[0]
	}
	return body, nil
}


func NewRequester() *MemberService {
	return &MemberService()
}


// func (*ActorService) getActor(requester *Requester) ([*Actor, error]) {
func (*ActorService) getActor() ([*Actor, error]) {
	_, body, errs:= NewRequester().Get("https://swapi.co/api/people/1/").End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	actor := make([]*Actor, 0)
	err := json.Unmarshal([]byte(body), &actors)

	if err != nil {
		return nil, err
	}

	return actor, nil
}