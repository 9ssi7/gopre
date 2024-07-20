package retry_test

import (
	"errors"
	"testing"

	"github.com/9ssi7/gopre/pkg/retry"
)

type retryFuncTestCase struct {
	name    string
	fn      retry.RetryFunc
	cnf     retry.Config
	logger  func(log string)
	wantErr error
}

func TestRun(t *testing.T) {
	secondTestCounter := 0
	retry.DefaultConfig.WaitTime = 1
	testConfig := retry.Config{}
	tests := []retryFuncTestCase{
		{
			name: "Success on first try",
			fn: func() error {
				return nil
			},
			cnf:     testConfig,
			wantErr: nil,
		},
		{
			name: "Success on second try",
			fn: func() error {
				if secondTestCounter == 0 {
					secondTestCounter++
					return errors.New("failed operation")
				}
				return nil
			},
			cnf:     testConfig,
			wantErr: nil,
		},
		{
			name: "Retries on failure (specified max)",
			fn: func() error {
				return errors.New("failed operation")
			},
			cnf:     testConfig,
			wantErr: errors.New("failed operation"),
		},
		{
			name: "Retries with logger",
			fn: func() error {
				return errors.New("failed operation")
			},
			logger: func(log string) {
				t.Log(log)
			},
			cnf:     testConfig,
			wantErr: errors.New("failed operation"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := retry.Run(tt.fn, retry.Config{
				Logger:     tt.logger,
				MaxRetries: tt.cnf.MaxRetries,
				WaitTime:   tt.cnf.WaitTime,
			})
			if tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
