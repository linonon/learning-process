package monster

import "testing"

func TestStore(t *testing.T) {

	monster := Monster{
		Name:  "紅孩兒",
		Age:   8,
		Skill: "fier",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.store() 錯誤")
	}
	t.Logf("Store() test correct")
}
func TestReStore(t *testing.T) {

	monster := &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("wrong, %v", res)
	}
	if monster.Name != "紅孩兒" {
		t.Fatalf("It's not 紅孩兒")
	}
	t.Logf("Restore() test correct")
}
