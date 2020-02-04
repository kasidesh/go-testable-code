package main

import "fmt"

func main() {
	svc := NewMemberService()
	requster := NewRequester()
	members, _ := svc.getMembers(requster)
	fmt.Printf("%v", members)
}
