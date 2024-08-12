package lpapi

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDomainResults(t *testing.T){
	lpSiteId := "90412079"
	result, err := GetDomain(lpSiteId)
	if err != nil {
		t.Fatal("domain call returned error")
	}

	fmt.Printf("result is %v", len(result.BaseURIs))
	if len(result.BaseURIs) < 100 {
		t.Fatal("result is < 0")
	}

}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	msg := Hello()
    name := "hello"
    want := regexp.MustCompile(`\b`+name+`\b`)
	if !want.MatchString(msg){
		t.Fatal("hello test failed")
	}
    // msg, err := lpbot.Hello()
    // if !want.MatchString(msg) {
    //     t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    // }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
// func TestHelloEmpty(t *testing.T) {
//     msg, err := Hello("")
//     if msg != "" || err == nil {
//         t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
//     }
// }