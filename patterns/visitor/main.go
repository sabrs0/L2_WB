package main

import "fmt"

type Visitor interface {
	visitProfile(p *Profile)
}

type Profile struct {
	name       string
	subscribes []string
	followers  []string
	matches    []string
}

func (p *Profile) accept(v Visitor) {
	v.visitProfile(p)
}
func (p *Profile) addSubscribe(name string) {
	p.subscribes = append(p.subscribes, name)
}
func (p *Profile) addFollower(name string) {
	p.followers = append(p.followers, name)
}
func (p Profile) getFollowers() []string {
	return p.followers
}
func (p Profile) getSubsribes() []string {
	return p.subscribes
}
func (p *Profile) setMatches(matches []string) {
	p.matches = matches
}

type MatchedVisitor struct {
}

func (m MatchedVisitor) visitProfile(p *Profile) {
	ans := []string{}
	for _, f := range p.followers {
		for _, s := range p.subscribes {

			if s == f {
				ans = append(ans, s)
			}
		}
	}
	p.setMatches(ans)
}
func main() {
	matchedVisitor := MatchedVisitor{}
	profile := Profile{name: "Ivan"}
	profile.addSubscribe("Petr")
	profile.addSubscribe("Sidor")
	profile.addSubscribe("Oleg")

	profile.addFollower("Petr")
	profile.addFollower("Sidor")
	profile.addFollower("Igor")

	profile.accept(matchedVisitor)

	fmt.Println(profile.matches)

}
