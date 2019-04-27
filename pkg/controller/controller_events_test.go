package controller

import (
	"fmt"
	"strings"
)

func checkEventCounts(actual, expected []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(actual) != len(expected) {
		return fmt.Errorf("Checking event count: %s", expectedGot(len(expected), len(actual)))
	}
	return nil
}
func checkEvents(actual, expected []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := checkEventCounts(actual, expected); err != nil {
		return err
	}
	for i, actualEvt := range actual {
		if expectedEvt := expected[i]; actualEvt != expectedEvt {
			return fmt.Errorf("event %d: %s", i, expectedGot(expectedEvt, actualEvt))
		}
	}
	return nil
}
func checkEventPrefixes(actual, expected []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := checkEventCounts(actual, expected); err != nil {
		return err
	}
	for i, e := range expected {
		a := actual[i]
		if !strings.HasPrefix(a, e) {
			return fmt.Errorf("received unexpected event prefix:\n %s", expectedGot(e, a))
		}
	}
	return nil
}
func checkEventContains(actual, expected string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !strings.Contains(actual, expected) {
		return fmt.Errorf("received unexpected event (contains):\n %s", expectedGot(expected, actual))
	}
	return nil
}
func expectedGot(a ...interface{}) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("\nexpected:\n\t '%v',\ngot:\n\t '%v'", a...)
}
