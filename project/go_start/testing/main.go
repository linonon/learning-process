package main

import ms "gostart/testing/testcase02"

func main() {
	monster := ms.Monster{
		Name:  "linonon",
		Age:   25,
		Skill: "swimming",
	}
	monster.Store()
	monster.ReStore()
}
