package main

type database struct{}

func newDatabase() *database {
	return &database{}
}

func (db *database) Query(sql string, timeout int) (string, error) {
	return "", nil
}
