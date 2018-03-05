package env_utils

import (
"os"
"strconv"
"github.com/pkg/errors"
	"log"
)

func GetEnv(envKey string) (val string) {
	val = os.Getenv(envKey)
	if val == "" {
		err := errors.Errorf("Env %s does not exists", envKey)
		log.Fatal(err.Error())
		os.Exit(1)
	}
	return val
}

func GetEnvInt(envKey string) (val int) {
	valString := GetEnv(envKey)
	valInt64, err := strconv.ParseInt(valString, 10, 64)
	if err != nil {
		err = errors.Wrapf(err,"error parsing string %s to int for env %s ", valString, envKey)
		log.Fatal(err.Error())
		os.Exit(1)
	}
	val = int(valInt64)
	return val
}
