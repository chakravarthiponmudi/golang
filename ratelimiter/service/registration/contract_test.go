package registration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
	suite.Equal(5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

// type Options struct {
// 	Logger        logging.Logger
// 	LockTimeout   time.Duration
// 	NumberOfRetry int
// 	RetryInterval time.Duration
// }

// type Option func(opt *Options)

// // SetLogger to an instance of logging.Logger, default to Noop
// func SetLogger(logger logging.Logger) Option {
// 	return func(opt *Options) {
// 		opt.Logger = logger
// 	}
// }

// func SetLockTimeout(d time.Duration) Option {
// 	return func(opt *Options) {
// 		opt.LockTimeout = d
// 	}
// }

// // SetNumberOfRetry sets max retries for acquire the table, default to 5 times
// func SetNumberOfRetry(t int) Option {
// 	return func(opt *Options) {
// 		opt.NumberOfRetry = t
// 	}
// }

// // SetRetryInterval sets sleep duration between each retry, default to 10 seconds
// func SetRetryInterval(d time.Duration) Option {
// 	return func(opt *Options) {
// 		opt.RetryInterval = d
// 	}
// }

// func SetupSuite() {
// 	pg := engine.NewPostgresEngine(library.SetupDSN())
// 	cleaner := dbcleaner.New(SetNumberOfRetry(10), SetLockTimeout(5*time.Second))
// 	Cleaner.setEngine(pg)
// }

// func TestGetContractByNameAndGroup(t *testing.T) {
// 	SetupSuite()
// }
