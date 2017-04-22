package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/flosch/pongo2"
)

func init() {
	os.Setenv("SPLIT_TEST", "value=with=equal=sings")
	os.Setenv("KUBERNETES_ADDRESS", "https://192.168.99.100")
	os.Setenv("FALSY", "")
}

func TestEnvVariables(t *testing.T) {
	tpl := []byte("{{ SPLIT_TEST }}: {{ KUBERNETES_ADDRESS }}")
	expected := fmt.Sprintf("%s: %s", os.Getenv("SPLIT_TEST"), os.Getenv("KUBERNETES_ADDRESS"))

	res, err := parse(tpl)

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if string(res) != expected {
		t.Errorf("bad expansion: expected '%s', got '%s'", expected, res)
	}
}

func TestEnvIf(t *testing.T) {
	tpl := []byte(`
{% if FALSY %}
I shouldn't appear
{% endif %}
I should!
`)
	expected := `

I should!
`

	res, err := parse(tpl)

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if string(res) != expected {
		t.Errorf("bad expansion: expected '%s', got '%s'", expected, res)
	}
}

func TestFilterBase64Encode(t *testing.T) {
	actualResult, err := filterBase64Encode(pongo2.AsValue("Hello"), pongo2.AsValue(""))

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	expectedResult := "SGVsbG8="

	if actualResult.String() != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, actualResult)
	}
}
