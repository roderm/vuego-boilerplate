package ricebox

import rice "github.com/GeertJohan/go.rice"

func Static() *rice.Box {
	return rice.MustFindBox("../static")
}
