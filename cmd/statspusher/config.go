package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/Sirupsen/logrus"
	"github.com/cerana/cerana/pkg/logrusx"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	flagSep     = regexp.MustCompile(`[\s._-]+`)
	specialCaps = regexp.MustCompile("(?i)^(url|cpu|ip|id)$")
)

type config struct {
	viper   *viper.Viper
	flagSet *pflag.FlagSet
}

// ConfigData defines the structure of the config data (e.g. in the config file)
type ConfigData struct {
	NodeDataURL     string `json:"nodeDataURL"`
	ClusterDataURL  string `json:"clusterDataURL"`
	LogLevel        string `json:"logLevel"`
	RequestTimeout  string `json:"requestTimeout"`
	DatasetInterval string `json:"datasetInterval"`
	BundleInterval  string `json:"bundleInterval"`
	NodeInterval    string `json:"nodeInterval"`
	DatasetDir      string `json:"datasetDir"`
}

func newConfig(flagSet *pflag.FlagSet, v *viper.Viper) *config {
	if flagSet == nil {
		flagSet = pflag.CommandLine
	}

	if v == nil {
		v = viper.New()
	}

	// Set normalization function before adding any flags
	flagSet.SetNormalizeFunc(canonicalFlagName)

	// Update Usage (--help output) to indicate flag
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		pflag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "Note: Flags can be used in either fooBar or foo[_-.]bar form.")
	}

	flagSet.StringP("configFile", "c", "", "path to config file")
	flagSet.StringP("nodeDataURL", "u", "", "url of coordinator for node information retrieval")
	flagSet.StringP("clusterDataURL", "e", "", "url of coordinator for the cluster information")
	flagSet.StringP("logLevel", "l", "warning", "log level: debug/info/warn/error/fatal/panic")
	flagSet.DurationP("requestTimeout", "r", 0, "default timeout for requests made")
	flagSet.DurationP("datasetInterval", "d", 0, "dataset heartbeat interval")
	flagSet.DurationP("bundleInterval", "b", 0, "bundle heartbeat interval")
	flagSet.DurationP("nodeInterval", "n", 0, "node heartbeat interval")
	flagSet.StringP("datasetDir", "a", "", "dataset directory")

	return &config{
		viper:   v,
		flagSet: flagSet,
	}
}

// canonicalFlagName translates flag names to camelCase using whitespace,
// periods, underscores, and dashes as word boundaries. All-caps words are
// preserved.
func canonicalFlagName(f *pflag.FlagSet, name string) pflag.NormalizedName {
	// Standardize separators to a single space and trim leading/trailing spaces
	name = strings.TrimSpace(flagSep.ReplaceAllString(name, " "))

	// Convert to title case (lower case with leading caps, preserved all caps)
	name = strings.Title(name)

	// Some words should always be all caps or all lower case (e.g. Interval)
	nameParts := strings.Split(name, " ")
	for i, part := range nameParts {
		caseFn := strings.ToUpper
		if i == 0 {
			caseFn = strings.ToLower
		}

		nameParts[i] = specialCaps.ReplaceAllStringFunc(part, caseFn)
	}

	// Split on space and examine the first part
	first := nameParts[0]
	if utf8.RuneCountInString(first) == 1 || first != strings.ToUpper(first) {
		// Lowercase the first letter if it is not an all-caps word
		r, n := utf8.DecodeRuneInString(first)
		nameParts[0] = string(unicode.ToLower(r)) + first[n:]
	}

	return pflag.NormalizedName(strings.Join(nameParts, ""))
}

func (c *config) loadConfig() error {
	if err := c.viper.BindPFlags(c.flagSet); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed to bind flags")
		return err
	}

	filePath := c.viper.GetString("configFile")
	if filePath == "" {
		return c.validate()
	}

	c.viper.SetConfigFile(filePath)
	if err := c.viper.ReadInConfig(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error":    err,
			"filePath": filePath,
		}).Error("failed to parse config file")
		return err
	}

	return c.validate()
}

func (c *config) nodeDataURL() *url.URL {
	// Error checking has been done during validation
	u, _ := url.ParseRequestURI(c.viper.GetString("nodeDataURL"))
	return u
}

func (c *config) clusterDataURL() *url.URL {
	// Error checking has been done during validation
	u, _ := url.ParseRequestURI(c.viper.GetString("clusterDataURL"))
	return u
}

func (c *config) requestTimeout() time.Duration {
	return c.viper.GetDuration("requestTimeout")
}

func (c *config) datasetInterval() time.Duration {
	return c.viper.GetDuration("datasetInterval")
}

func (c *config) nodeInterval() time.Duration {
	return c.viper.GetDuration("nodeInterval")
}

func (c *config) bundleInterval() time.Duration {
	return c.viper.GetDuration("bundleInterval")
}

func (c *config) datasetDir() string {
	return c.viper.GetString("datasetDir")
}

func (c *config) setupLogging() error {
	logLevel := c.viper.GetString("logLevel")
	if err := logrusx.SetLevel(logLevel); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"level": logLevel,
		}).Error("failed to set up logging")
		return err
	}
	return nil
}

func (c *config) validate() error {
	if c.datasetInterval() == 0 {
		return errors.New("dataset interval must be > 0")
	}
	if c.bundleInterval() == 0 {
		return errors.New("bundle interval must be > 0")
	}
	if c.nodeInterval() == 0 {
		return errors.New("node interval must be > 0")
	}
	if c.requestTimeout() == 0 {
		return errors.New("request timeout must be > 0")
	}

	if err := c.validateURL("nodeDataURL"); err != nil {
		return err
	}
	if err := c.validateURL("clusterDataURL"); err != nil {
		return err
	}
	if c.datasetDir() == "" {
		return errors.New("missing datasetDir")
	}

	return nil
}

func (c *config) validateURL(name string) error {
	u := c.viper.GetString(name)
	if u == "" {
		return errors.New("missing " + name)
	}
	if _, err := url.ParseRequestURI(u); err != nil {
		logrus.WithFields(logrus.Fields{
			name:    u,
			"error": err,
		}).Error("invalid config")
		return errors.New("invalid " + name)
	}
	return nil
}