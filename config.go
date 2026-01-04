package mdtable

import (
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

type Align int

const (
	Center Align = 0
	Left   Align = 1
	Right  Align = 2
)

type ColumnSortOption int

const (
	None       ColumnSortOption = 0
	Ascending  ColumnSortOption = 1
	Descending ColumnSortOption = 2
	Custom     ColumnSortOption = 3
)

var sortColumnsName = map[ColumnSortOption]string{
	None:       "None",
	Ascending:  "Ascending",
	Descending: "Descending",
	Custom:     "Custom",
}

func (cso ColumnSortOption) String() string {
	return sortColumnsName[cso]
}

type ColumnSortFunction func(a string, b string) int

type Config struct {
	// Align the rendered content for the Markdown table. 0 = Center, 1 = Left, 2 = Right
	Align Align

	// Caption of the table (as an HTML comment)
	Caption string

	// Should the markdown table be the compact version
	Compact bool

	// List of columns to be excluded from table construction
	ExcludedColumns []string

	// Indices of excluded columns (internal)
	excludedColumnsIndices []int

	// Indices of columns to convert to
	orderedColumnsIndices []int

	// Should the columns be sorted and how?
	SortColumns ColumnSortOption

	// Custom sort function
	SortFunction ColumnSortFunction

	// Log detailed diagnostic messages when running the program.
	VerboseLogging bool
}

// Validate the Config object passed as parameter.
// An error will be returned in case the configuration was invalid.
func ValidateConfig(cfg Config) error {

	cfgWarnings := []string{}

	if cfg.VerboseLogging {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Validating config 🤔")
	}

	if cfg.Align < Center || cfg.Align > Right {
		return errors.New("align value is out of range, please choose in range [0-2]")
	}

	if cfg.SortColumns < None || cfg.SortColumns > Custom {
		return errors.New("sort columns value is out of range, please choose in range [0-3]")
	}

	// custom sort but no custom sort function was provided, will affect sorting columns
	if cfg.SortColumns == Custom && cfg.SortFunction == nil {
		return errors.New("sort type is set to Custom but SortFunc was not set.")
	}

	// function passed in but not sort type is not custom
	if cfg.SortColumns != Custom && cfg.SortFunction != nil {
		cfgWarnings = append(cfgWarnings, fmt.Sprintf("Sort function only works when SortColumns is set to Custom. SortColumns received is %s, ignoring SortFunc.", cfg.SortColumns))
	}

	if len(cfgWarnings) > 0 {
		// config contains warnings, let's warn user but also continue the execution
		warnings := strings.Join(cfgWarnings, "\n")
		warningMsg := fmt.Sprintf("Config contains some warnings:\n%s", warnings)
		slog.Warn(warningMsg)
	} else if cfg.VerboseLogging {
		slog.Debug("Config is valid ✅")
	}

	return nil
}

// Populate orderColumnIndices in Config object
func populateColumnIndices(cfg Config, headerLine []string) Config {
	// get the new order of columns after sorted, compared to the original order of them.
	if cfg.SortColumns == None {
		for i := range len(headerLine) {
			cfg.orderedColumnsIndices = append(cfg.orderedColumnsIndices, i)
		}
	} else {
		cfg.orderedColumnsIndices = getIndicesAfterSorting(cfg, headerLine)
	}

	return cfg
}

// Get the indices of columns after sorted
func getIndicesAfterSorting(cfg Config, headerLine []string) []int {
	sortedColumns := make([]string, len(headerLine))
	copy(sortedColumns, headerLine)
	var columnsIndicesAfterSorting []int

	switch cfg.SortColumns {
	case Ascending:
		slices.SortFunc(sortedColumns, func(a, b string) int {
			return strings.Compare(strings.ToLower(a), strings.ToLower(b))
		})
	case Descending:
		slices.SortFunc(sortedColumns, func(a, b string) int {
			return strings.Compare(strings.ToLower(b), strings.ToLower(a))
		})
	case Custom:
		slices.SortFunc(sortedColumns, cfg.SortFunction)
	}

	for i := range sortedColumns {
		columnsIndicesAfterSorting = append(columnsIndicesAfterSorting, slices.Index(headerLine, sortedColumns[i]))
	}

	return columnsIndicesAfterSorting
}
