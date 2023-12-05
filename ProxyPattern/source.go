package ProxyPattern

import "fmt"

type IGiveGift interface {
	GiveDolls()
	GiveFollowers()
	GiveChocolate()
}

type SchoolGirl struct {
	name string
}

type PurSuit struct {
	SchoolGirl
	// name string
}

type Proxy struct {
	PurSuit
	// name string
}

func (p *PurSuit) GiveDolls() {
	fmt.Println(p.SchoolGirl.name + "送你洋娃娃")
}

func (p *PurSuit) GiveFollowers() {
	fmt.Println(p.SchoolGirl.name + "送你花")
}

func (p *PurSuit) GiveChocolate() {
	fmt.Println(p.SchoolGirl.name + "送你巧克力")
}

func (p *Proxy) GiveDolls() {
	p.PurSuit.GiveDolls()
}

func (p *Proxy) GiveFollowers() {
	p.PurSuit.GiveFollowers()
}

func (p *Proxy) GiveChocolate() {
	p.PurSuit.GiveChocolate()
}
