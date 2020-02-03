func ICacher interface {
	Set(key string, value interface{}, expire time.Duration) error
}

func (c *Cacher) Set(key, value interface{}, expire time.Duration) error {
	client := redis.NewClient(&redis.Options{
		Addr: c.server,
		Password: ""
	})

	err := client.Set(key, value, expire).Err()
	if err != nil {
		return err
	}

	return nil
}

func (*MemberService) getMembers(requester IRequester, cacher ICacher) ([]*Member, error) {
	body, err := requester.Get("http://member-service/api/members")
	if err != nil {
		return nil, err
	}

	members := make([]*Member, 0)
	err = json.Unmarshal([]byte(body), &members)
	if err != nil {
		return nil, err
	}

    cacher.Set("member_cache", members, 2*time.Hour)
	return members, nil
}