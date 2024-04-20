package main

import "fmt"

type gitUser struct {
	name string
}

func (g *gitUser) setUserName(s string) {
	g.name = s
}

func main() {

	// working with pointers

	g := &gitUser{
		name: "",
	}

	if g.name == "" {
		fmt.Printf("GitHub username is <empty>\n") // GitHub username is <empty>
		// fmt.Println(&g.name)                     // 0xc000028070
		g.setUserName("@nitinjilla")
		// fmt.Println(&g.name)                     // 0xc000028070
		fmt.Printf("GitHub username set to '%s'\n", g.name) // GitHub username set to '@nitinjilla'
	} else {
		fmt.Printf("GitHub username is %s\n", g.name)
	}

}
