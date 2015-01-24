package main

import (
	"reflect"
	"testing"
)

func TestResourceListParsing(t *testing.T) {
	normalCase := `http://www.google.com
https://github.com/gophergala/blinky_light_display
forecast.io/#/f/41.9452,-87.7068`

	normalCaseExpected := []string{"http://www.google.com", "https://github.com/gophergala/blinky_light_display", "forecast.io/#/f/41.9452,-87.7068"}

	if !reflect.DeepEqual(normalCaseExpected, parseResourceList(normalCase)) {
		t.Errorf("Normal case list parsing failed. Got %s", parseResourceList(normalCase))
	}

	extraNewLinesCase := `http://www.google.com

https://github.com/gophergala/blinky_light_display

forecast.io/#/f/41.9452,-87.7068`

	if !reflect.DeepEqual(normalCaseExpected, parseResourceList(extraNewLinesCase)) {
		t.Errorf("Extra new lines case list parsing failed. Got %s", parseResourceList(extraNewLinesCase))
	}

}

func TestUrlsStringification(t *testing.T) {
	someUrls := []string{"http://www.google.com", "https://github.com/gophergala/blinky_light_display", "forecast.io/#/f/41.9452,-87.7068"}

	expectedString := `http://www.google.com
https://github.com/gophergala/blinky_light_display
forecast.io/#/f/41.9452,-87.7068
`

	if !reflect.DeepEqual(makeUrlsString(someUrls), expectedString) {
		t.Errorf("Url Stringification failed. Got:\n%s\n Expected:\n%s", makeUrlsString(someUrls), expectedString)
	}
}
