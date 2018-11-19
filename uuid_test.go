package uuid

import (
	"testing"

	"github.com/google/uuid"
)

func TestEncode(t *testing.T) {
	var tests = []struct {
		raw     string
		encoded string
	}{
		{
			raw:     "32312b5f-22fa-4843-b46d-abe50b0ab8bc",
			encoded: "gea1sz3n9jrr8pdpix1osniazo",
		},
		{
			raw:     "35b8bf87-e6fd-4990-bd72-6f134752a04f",
			encoded: "gshm9b9g9ir3bxm1phjwqwiyjh",
		},
		{
			raw:     "570acacb-dea3-47d7-92f3-0601d5538d8f",
			encoded: "khfci166wpd7xrzuyay7kwhpth",
		},
	}
	for _, test := range tests {
		t.Run(test.raw, func(t *testing.T) {
			parsed := uuid.Must(uuid.Parse(test.raw))
			output := Encode(parsed)
			if output != test.encoded {
				t.Errorf("got %q, want %q", output, test.encoded)
			}
			decoded, ok := Decode(test.encoded)
			if !ok {
				t.Errorf("Decoding failed")
			}
			if decoded.String() != test.raw {
				t.Errorf("Decoded to %q", decoded.String())
			}
		})
	}
}

func TestDecode(t *testing.T) {
	var tests = []struct {
		input   string
		decoded string
		ok      bool
	}{
		{
			input:   "khfci166wpd7xrzuyay7kwhpth",
			decoded: "570acacb-dea3-47d7-92f3-0601d5538d8f",
			ok:      true,
		},
		{
			input:   "xxx",
			decoded: "00000000-0000-0000-0000-000000000000",
			ok:      false,
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			uuid, ok := Decode(test.input)
			if uuid.String() != test.decoded || ok != test.ok {
				t.Errorf("got (%s, %t)", uuid.String(), ok)
			}
		})
	}
}
