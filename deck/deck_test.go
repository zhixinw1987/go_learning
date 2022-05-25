package deck
import "testing"

//https://appliedgo.net/testmain/

func TestNewDeck(t *testing.T){
	d:= newDeck()

	//formatted string, %v is placeholder for the following value
	if len(d) != 24 {
		t.Errorf("Expected deck length of 24 but got %v", len(d))
	}

	if d[0] != "Ace of Cluba" {
		t.Errorf("Expected first card in deck as Ace of Club but got %v", d[0])
	}

	if d[len(d)-1] != "King of Spade" {
		t.Errorf("Expected last card in deck as King of Spade but got %v", d[len(d)-1])
	}
}