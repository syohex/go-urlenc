package urlenc_test

import (
	"net/url"
	"testing"

	"github.com/lestrrat/go-urlenc"
	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Bar   string    `urlenc:"bar"`
	Baz   int       `urlenc:"baz"`
	Qux   []string  `urlenc:"qux"`
	Corge []float64 `urlenc:"corge"`
}

func TestUnmarshal(t *testing.T) {
	const src = `bar=one&baz=2&qux=three&qux=4&corge=1.41421356237&corge=2.2360679775`

	var foo Foo
	if !assert.NoError(t, urlenc.Unmarshal([]byte(src), &foo), "Unmarshal should succeed") {
		return
	}

	if !assert.Equal(t, foo.Bar, "one", "Bar is 'one'") {
		return
	}
	if !assert.Equal(t, foo.Baz, 2, "Baz is '2'") {
		return
	}
	if !assert.Equal(t, foo.Qux, []string{"three", "4"}, "Qux is 'three, 4'") {
		return
	}
	if !assert.Equal(t, foo.Corge, []float64{1.41421356237, 2.2360679775}, "Corge is '1.41421356237, 2.2360679775'") {
		return
	}
}

func TestMarshal(t *testing.T) {
	const src = `bar=one&baz=2&qux=three&qux=4&corge=1.41421356237&corge=2.2360679775`

	foo := Foo{
		Bar:   "one",
		Baz:   2,
		Qux:   []string{"three", "4"},
		Corge: []float64{1.41421356237, 2.2360679775},
	}
	buf, err := urlenc.Marshal(foo)
	if !assert.NoError(t, err, "Marshal should succeed") {
		return
	}

	produced, err := url.ParseQuery(string(buf))
	if !assert.NoError(t, err, "ParseQuery should succeed") {
		return
	}
	expected, err := url.ParseQuery(src)
	if !assert.NoError(t, err, "ParseQuery should succeed") {
		return
	}

	if !assert.Equal(t, produced, expected, "Marshal produces the same result") {
		return
	}
}
