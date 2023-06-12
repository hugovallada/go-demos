package tests

type TestCase struct {
	Entry    []string
	Expected []string
}

func PrepareTestCase() []TestCase {
	return []TestCase{
		{
			Entry:    []string{"ola", "mundo", "hello", "ola"},
			Expected: []string{"ola", "mundo", "hello"},
		},
		{
			Entry:    []string{"ola", "ola", "ola", "ola", "ola"},
			Expected: []string{"ola"},
		},
	}
}
