func (*ActorService) getActor() ([*Actor, error]) {
	// _, body, errs:= gorequest.New().Get("https://swapi.co/api/people/1/").End()
	resp, body, errs := gorequest.New().Get("https://swapi.co/api/people/1").End()
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