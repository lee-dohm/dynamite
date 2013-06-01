package dynamite

import "fmt"

// Matcher is the interface that represents a set of methods that can check if an object matches a certain condition.
//
// Match performs the actual check to see if the object matches the condition. Message returns the message when the match
// fails the positive case. NegativeMessage returns the message when the match fails the negative case.
type Matcher interface {
	Match(params []interface{}) bool
    Message(params []interface{}) string
    NegativeMessage(params []interface{}) string
}

type beNilMatcher struct{}

var BeNil = beNilMatcher{}

func (matcher beNilMatcher) Match(params []interface{}) bool {
    value := matcher.params(params)

    return value == nil
}

func (matcher beNilMatcher) Message(params []interface{}) string {
    value := matcher.params(params)

    return fmt.Sprint("Expected %v to be nil", value)
}

func (matcher beNilMatcher) NegativeMessage(params []interface{}) string {
    value := matcher.params(params)

    return fmt.Sprint("Expected %v to not be nil", value)
}

func (matcher beNilMatcher) params(params []interface{}) interface{} {
    return params[0]
}
