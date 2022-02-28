package state

import "testing"

func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn hs ---\\ \\__/ _________________/---]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
func TestPutRevInBoat(t *testing.T) {
	wanted := "[kylling korn hs ---\\ \\rev__/ _________________/---]"
	state := PutRevInBoat()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
