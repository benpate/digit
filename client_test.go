package digit

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestClient(t *testing.T) {

	spew.Dump(Lookup("http://localhost/@benpate"))
}
