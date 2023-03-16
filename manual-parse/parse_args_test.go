package main

import "testing"

type testConfig struct {
	args []string
	err  error
	config
}

func TestParseArgs(t *testing.T) {
	tests := []testConfig{
		{
			args:   []string{"-h"},
			err:    nil,
			config: config{printUsage: true, numTimes: 0},
		},
		{
			args:   []string{"10"},
			err:    nil,
			config: config{printUsage: false, numTimes: 10},
		},
	}

	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("expected error to be: %v, got: %v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("expected nil error, got: %v\n", err)
		}
		if c.printUsage != tc.config.printUsage {
			t.Errorf("expected printUsage to be: %v, got: %v\n", tc.config.printUsage, c.printUsage)
		}
		if c.numTimes != tc.config.numTimes {
			t.Errorf("expected numTimes to be: %v, got: %v\n", tc.config.numTimes, c.numTimes)
		}
	}
}
