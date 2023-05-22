package phonenumber

var numberTests = []struct {
	description      string
	input            string
	expectErr        bool
	errorDescription string
	number           string
	areaCode         string
	formatted        string
	index            int
}{
	{
		description: "cleans the number",
		input:       "(223) 456-7890",
		number:      "2234567890",
		areaCode:    "223",
		formatted:   "(223) 456-7890",
		index:       1,
	},
	{
		description: "cleans numbers with dots",
		input:       "223.456.7890",
		number:      "2234567890",
		areaCode:    "223",
		formatted:   "(223) 456-7890",
		index:       2,
	},
	{
		description: "cleans numbers with multiple spaces",
		input:       "223 456   7890   ",
		number:      "2234567890",
		areaCode:    "223",
		formatted:   "(223) 456-7890",
		index:       3,
	},
	{
		description:      "invalid when 9 digits",
		input:            "123456789",
		expectErr:        true,
		errorDescription: "incorrect number of digits",
		index:            4,
	},
	{
		description:      "invalid when 11 digits does not start with a 1",
		input:            "22234567890",
		expectErr:        true,
		errorDescription: "11 digits must start with 1",
		index:            5,
	},
	{
		description: "valid when 11 digits and starting with 1",
		input:       "12234567890",
		number:      "2234567890",
		areaCode:    "223",
		formatted:   "(223) 456-7890",
		index:       6,
	},
	{
		description: "valid when 11 digits and starting with 1 even with punctuation",
		input:       "+1 (223) 456-7890",
		number:      "2234567890",
		areaCode:    "223",
		formatted:   "(223) 456-7890",
		index:       7,
	},
	{
		description:      "invalid when more than 11 digits",
		input:            "321234567890",
		expectErr:        true,
		errorDescription: "more than 11 digits",
		index:            8,
	},
	{
		description:      "invalid with letters",
		input:            "123-abc-7890",
		expectErr:        true,
		errorDescription: "letters not permitted",
		index:            9,
	},
	{
		description:      "invalid with punctuations",
		input:            "123-@:!-7890",
		expectErr:        true,
		errorDescription: "punctuations not permitted",
		index:            10,
	},
	{
		description:      "invalid if area code starts with 0",
		input:            "(023) 456-7890",
		expectErr:        true,
		errorDescription: "area code cannot start with zero",
		index:            11,
	},
	{
		description:      "invalid if area code starts with 1",
		input:            "(123) 456-7890",
		expectErr:        true,
		errorDescription: "area code cannot start with one",
		index:            12,
	},
	{
		description:      "invalid if exchange code starts with 0",
		input:            "(223) 056-7890",
		expectErr:        true,
		errorDescription: "exchange code cannot start with zero",
		index:            13,
	},
	{
		description:      "invalid if exchange code starts with 1",
		input:            "(223) 156-7890",
		expectErr:        true,
		errorDescription: "exchange code cannot start with one",
		index:            14,
	},
	{
		description:      "invalid if area code starts with 0 on valid 11-digit number",
		input:            "1 (023) 456-7890",
		expectErr:        true,
		errorDescription: "area code cannot start with zero",
		index:            15,
	},
	{
		description:      "invalid if area code starts with 1 on valid 11-digit number",
		input:            "1 (123) 456-7890",
		expectErr:        true,
		errorDescription: "area code cannot start with one",
		index:            16,
	},
	{
		description:      "invalid if exchange code starts with 0 on valid 11-digit number",
		input:            "1 (223) 056-7890",
		expectErr:        true,
		errorDescription: "exchange code cannot start with zero",
		index:            17,
	},
	{
		description:      "invalid if exchange code starts with 1 on valid 11-digit number",
		input:            "1 (223) 156-7890",
		expectErr:        true,
		errorDescription: "exchange code cannot start with one",
		index:            18,
	},
}
