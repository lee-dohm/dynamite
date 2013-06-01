package dynamite

import "testing"

type E struct {
    t *testing.T
}

type expectation struct {
    object *interface{}
    t *testing.T
}

func (e E) Expect(t *testing.T, object interface{}) expectation {
    return expectation{object: &object, t: t}
}

func (e expectation) To(matcher Matcher, params []interface{}) {
    p := make([]interface{}, 1)
    p[0] = e.object
    p = append(p, params...)

    if !matcher.Match(p) {
        e.t.Log(matcher.Message(p))
        e.t.FailNow()
    }
}

func (e expectation) ToNot(matcher Matcher, params []interface{}) {
    p := make([]interface{}, 1)
    p[0] = e.object
    p = append(p, params...)

    if matcher.Match(p) {
        e.t.Log(matcher.NegativeMessage(p))
        e.t.FailNow()
    }
}
