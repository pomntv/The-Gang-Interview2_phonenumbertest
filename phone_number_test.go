package phonenumber

import (
	"fmt"
	"testing"
)

func TestNumber(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := Number(test.input)
		fmt.Printf("###> TestNumber ### ")
		if actual != "" {
			fmt.Printf("%v: Number: %v\n", test.index, actual)
		} else {
			fmt.Printf("%v: !!! actualErr: %v\n", test.index, actualErr)
		}
		if !test.expectErr {
			if actualErr != nil {
				var _ error = actualErr
				t.Errorf("FAIL: %s\nNumber(%q): expected no error, but error is: %s", test.description, test.input, actualErr)
				fmt.Printf("FAIL: %s\nNumber(%q): expected no error, but error is: %s \n", test.description, test.input, actualErr)
			}
			if actual != test.number {
				t.Errorf("FAIL: %s\nNumber(%q): expected [%s], actual: [%s]", test.description, test.input, test.number, actual)
				fmt.Printf("FAIL: %s\nNumber(%q): expected [%s], actual: [%s] \n", test.description, test.input, test.number, actual)
			}
		} else if actualErr == nil {
			t.Errorf("FAIL: %s\nNumber(%q): expected an error, but error is nil", test.description, test.input)
			fmt.Printf("FAIL: %s\nNumber(%q): expected an error, but error is nil \n", test.description, test.input)
		}
	}

}

func BenchmarkNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numberTests {
			Number(test.input)
		}
	}
}

func TestAreaCode(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := AreaCode(test.input)
		fmt.Printf("####> TestAreaCode ### ")
		if actual != "" {
			fmt.Printf("%v: AreaCode: %v\n", test.index, actual)
		} else {
			fmt.Printf("%v: !!! actualErr: %v\n", test.index, actualErr)
		}
		if !test.expectErr {
			if actual != test.areaCode {
				t.Errorf("FAIL: %s\nAreaCode(%q): expected [%s], actual: [%s]", test.description, test.input, test.areaCode, actual)
				fmt.Printf("FAIL: %s\nAreaCode(%q): expected [%s], actual: [%s] \n", test.description, test.input, test.areaCode, actual)
			}
			if actualErr != nil {
				var _ error = actualErr
				t.Errorf("FAIL: %s\nAreaCode(%q): expected no error, but error is: %s", test.description, test.input, actualErr)
				fmt.Printf("FAIL: %s\nAreaCode(%q): expected no error, but error is: %s \n", test.description, test.input, actualErr)
			}
		} else if actualErr == nil {
			t.Errorf("FAIL: %s\nAreaCode(%q): expected an error, but error is nil", test.description, test.input)
			fmt.Printf("FAIL: %s\nAreaCode(%q): expected an error, but error is nil \n", test.description, test.input)
		}
	}
}

func BenchmarkAreaCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numberTests {
			AreaCode(test.input)
		}
	}
}

func TestFormat(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := Format(test.input)
		fmt.Printf("#####> TestFormat ### ")
		if actual != "" {
			fmt.Printf("%v: Format: %v\n", test.index, actual)
		} else {
			fmt.Printf("%v: !!! actualErr: %v\n", test.index, actualErr)
		}
		if !test.expectErr {
			if actualErr != nil {
				var _ error = actualErr
				t.Errorf("FAIL: %s\nFormat(%q): expected no error, but error is: %s", test.description, test.input, actualErr)
				fmt.Printf("FAIL: %s\nFormat(%q): expected no error, but error is: %s \n", test.description, test.input, actualErr)
			}
			if actual != test.formatted {
				t.Errorf("FAIL: %s\nFormat(%q): expected [%s], actual: [%s]", test.description, test.input, test.formatted, actual)
				fmt.Printf("FAIL: %s\nFormat(%q): expected [%s], actual: [%s] \n", test.description, test.input, test.formatted, actual)
			}
		} else if actualErr == nil {
			t.Errorf("FAIL: %s\nFormat(%q): expected an error, but error is nil", test.description, test.input)
			fmt.Printf("FAIL: %s\nFormat(%q): expected an error, but error is nil \n", test.description, test.input)

		}
	}
}

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range numberTests {
			Format(test.input)
		}
	}
}
