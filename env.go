package env_utils

import (
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"fmt"
)

func GetEnv(envKey string) (val string) {
	val = os.Getenv(envKey)
	if val == "" {
		err := errors.Errorf("Env %s does not exists", envKey)
		log.Fatal(err.Error())
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return val
}

func GetEnvInt(envKey string) (val int64) {
	valString := GetEnv(envKey)
	val, err := strconv.ParseInt(valString, 10, 64)
	if err != nil {
		err = errors.Wrapf(err, "error parsing string %s to int for env %s ", valString, envKey)
		log.Fatal(err.Error())
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return val
}

func GetEnvInt16(envKey string) (val int16) {
	valString := GetEnv(envKey)
	valInt64, err := strconv.ParseInt(valString, 10, 64)
	if err != nil {
		err = errors.Wrapf(err, "error parsing string %s to int for env %s ", valString, envKey)
		log.Fatal(err.Error())
		fmt.Println(err.Error())
		os.Exit(1)
	}
	val = int16(valInt64)
	return val
}

func GetEnvInt8(envKey string) (val int8) {
	valString := GetEnv(envKey)
	valInt64, err := strconv.ParseInt(valString, 10, 64)
	if err != nil {
		err = errors.Wrapf(err, "error parsing string %s to int for env %s ", valString, envKey)
		log.Fatal(err.Error())
		fmt.Println(err.Error())
		os.Exit(1)
	}
	val = int8(valInt64)
	return val
}