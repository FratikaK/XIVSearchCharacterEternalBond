package main

import (
	"fmt"
	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone"
)

type character struct {
	firstName  string
	familyName string
	world      string
	id         uint32
}

func main() {
	s := godestone.NewScraper(bingode.New(), godestone.JA)

	c := character{
		firstName:  "Kamu",
		familyName: "Wraith",
		world:      "Anima",
	}

	characterOpts := godestone.CharacterOptions{
		Name:  c.firstName + " " + c.familyName,
		World: c.world,
	}

	for characters := range s.SearchCharacters(characterOpts) {
		if characters.Error != nil {
			fmt.Println("SearchCharacters err:", characters.Error)
		}
		c.id = characters.ID
	}
	eternal := isCharacterEternalBond(c.id)
	if eternal {
		fmt.Println(c.firstName, c.familyName+"さんは残念ながらエタバンしています...")
	} else {
		fmt.Println(c.firstName, c.familyName+"さんはおそらくエタバンしていないです！")
		fmt.Println("今がチャンスでしょう！")
	}
}

//キャラクターがエタバンしているかをbool値で返す
func isCharacterEternalBond(id uint32) bool {
	s := godestone.NewScraper(bingode.New(), godestone.JA)

	//エターナルチョコボを所持しているかで判断
	mounts, err := s.FetchCharacterMounts(id)
	if err != nil {
		fmt.Println("FetchCharacterMounts err:", err)
		return false
	}
	for _, mount := range mounts {
		if mount.ID == 41 {
			fmt.Println("エターナルチョコボを所持しています！")
			return true
		}
	}

	//エターナルリングを所持しているか
	c, err := s.FetchCharacter(id)
	if err != nil {
		fmt.Println("FetchCharacter err:", err)
		return false
	}
	
	ring1 := c.GearSet.Gear.Ring1.ID
	ring1Mirage := c.GearSet.Gear.Ring1.Mirage
	ring2 := c.GearSet.Gear.Ring2.ID
	ring2Mirage := c.GearSet.Gear.Ring2.Mirage

	if ring1 == 8575 || ring1Mirage == 8575 || ring2 == 8575 || ring2Mirage == 8575 {
		fmt.Println("エターナルリングを所持しています！")
		return true
	}
	return false
}
