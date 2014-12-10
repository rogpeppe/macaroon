package macaroon_test

import (
	gc "gopkg.in/check.v1"

	"gopkg.in/macaroon.v1"
)

type marshalSuite struct{}

var _ = gc.Suite(&marshalSuite{})

func (*marshalSuite) TestMarshalUnmarshalMacaroon(c *gc.C) {
	rootKey := []byte("secret")
	m := MustNew(rootKey, "some id", "a location")

	m.AddFirstPartyCaveat("a caveat")

	b, err := m.MarshalBinary()
	c.Assert(err, gc.IsNil)

	unmarshalledM := macaroon.Macaroon{}
	err = unmarshalledM.UnmarshalBinary(b)
	c.Assert(err, gc.IsNil)

	c.Assert(m.Location(), gc.Equals, unmarshalledM.Location())
	c.Assert(m.Id(), gc.Equals, unmarshalledM.Id())
	c.Assert(m.Signature(), gc.DeepEquals, unmarshalledM.Signature())
	c.Assert(m.Caveats(), gc.DeepEquals, unmarshalledM.Caveats())
}

func (*marshalSuite) TestMarshalUnmarshalMacaroons(c *gc.C) {
	rootKey := []byte("secret")
	m1 := MustNew(rootKey, "some id", "a location")
	m2 := MustNew(rootKey, "some other id", "another location")

	m1.AddFirstPartyCaveat("a caveat")
	m2.AddFirstPartyCaveat("another caveat")

	macaroons := macaroon.Macaroons{m1, m2}

	b, err := macaroons.MarshalBinary()
	c.Assert(err, gc.IsNil)

	unmarshalledMacs := macaroon.Macaroons{m1, m2}
	err = unmarshalledMacs.UnmarshalBinary(b)
	c.Assert(err, gc.IsNil)

	c.Assert(unmarshalledMacs, gc.HasLen, 2)
	for i := 0; i < 1; i++ {
		c.Assert(macaroons[i].Location(), gc.Equals, unmarshalledMacs[i].Location())
		c.Assert(macaroons[i].Id(), gc.Equals, unmarshalledMacs[i].Id())
		c.Assert(macaroons[i].Signature(), gc.DeepEquals, unmarshalledMacs[i].Signature())
		c.Assert(macaroons[i].Caveats(), gc.DeepEquals, unmarshalledMacs[i].Caveats())
	}
}
