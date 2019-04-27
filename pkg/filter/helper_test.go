package filter

import (
	"testing"
)

func TestCreatePredicate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		restrictions	[]string
		error		bool
		predicate	string
	}{{name: "no restrictions"}, {name: "invalid class restrictions", restrictions: []string{"this throws an error"}, error: true}, {name: "valid class restriction", restrictions: []string{"name in (Foo, Bar)"}, predicate: "name in (Bar,Foo)"}, {name: "valid class double restriction and wacky spacing", restrictions: []string{"name   in      (Foo,   Bar)", "name   notin   (Baz,   Barf)"}, predicate: "name in (Bar,Foo),name notin (Barf,Baz)"}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			predicate, err := CreatePredicate(tc.restrictions)
			if err != nil {
				if tc.error {
					return
				}
				t.Fatalf("Unexpected error from CreatePredicateForServiceClassesFromRestrictions: %v", err)
			}
			if predicate == nil {
				t.Fatalf("Failed to create predicate from restrictions: %+v", tc.restrictions)
			}
			if tc.restrictions == nil && !predicate.Empty() {
				t.Fatalf("Failed to create predicate an empty prediate from nil restrictions.")
			}
			ps := predicate.String()
			if ps != tc.predicate {
				t.Fatalf("Failed to create expected predicate, \n\texpected: \t%q,\n \tgot: \t\t%q", tc.predicate, ps)
			}
		})
	}
}
