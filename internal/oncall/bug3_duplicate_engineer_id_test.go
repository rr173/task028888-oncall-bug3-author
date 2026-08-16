package oncall

import "testing"

func TestBug3_RejectsDuplicateEngineerIDs(t *testing.T) {
	_, err := Build(Request{
		Roster: []string{"alice", "alice"},
		Start:  "2026-03-02",
		End:    "2026-03-03",
	})
	if err == nil {
		t.Fatal("expected duplicate engineer IDs to be rejected")
	}
}
