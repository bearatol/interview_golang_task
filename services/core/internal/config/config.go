package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	DB *Postgres

	PriceGenAdd string
	AppAddr     string
}

type Postgres struct {
	PostgresUser,
	PostgresPass,
	PostgresDB,
	PostgresHost string
	PostgresPort uint64
}

func NewConfig(local bool) (*Config, error) {
	var (
		exist  bool
		env    string
		err    error
		config = &Config{}
	)
	if config.DB.PostgresUser, exist = os.LookupEnv("POSTGRES_USER"); !exist {
		return nil, errors.New("the env [POSTGRES_USER] does not exist")
	}
	if config.DB.PostgresPass, exist = os.LookupEnv("POSTGRES_PASSWORD"); !exist {
		return nil, errors.New("the env [POSTGRES_PASSWORD] does not exist")
	}
	if config.DB.PostgresDB, exist = os.LookupEnv("POSTGRES_DB"); !exist {
		return nil, errors.New("the env [POSTGRES_DB] does not exist")
	}

	if config.DB.PostgresHost, exist = os.LookupEnv("POSTGRES_HOST"); !exist {
		return nil, errors.New("the env [POSTGRES_HOST] does not exist")
	}
	{
		env, exist = os.LookupEnv("POSTGRES_PORT")
		if !exist {
			return nil, errors.New("the env [POSTGRES_PORT] does not exist")
		}
		config.DB.PostgresPort, err = strconv.ParseUint(env, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	if local {
		if config.DB.PostgresHost, exist = os.LookupEnv("POSTGRES_HOST_LOCAL"); !exist {
			return nil, errors.New("the env [POSTGRES_HOST_LOCAL] does not exist")
		}
		{
			env, exist = os.LookupEnv("POSTGRES_PORT_LOCAL")
			if !exist {
				return nil, errors.New("the env [POSTGRES_PORT_LOCAL] does not exist")
			}
			config.DB.PostgresPort, err = strconv.ParseUint(env, 10, 64)
			if err != nil {
				return nil, err
			}
		}
	}

	if config.PriceGenAdd, exist = os.LookupEnv("PRICE_GENERATOR_ADDR"); !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_ADDR] does not exist")
	}
	if local {
		if config.PriceGenAdd, exist = os.LookupEnv("PRICE_GENERATOR_ADDR_LOCAL"); !exist {
			return nil, errors.New("the env [PRICE_GENERATOR_ADDR_LOCAL] does not exist")
		}
	}

	if config.AppAddr, exist = os.LookupEnv("CORE_ADDR"); !exist {
		return nil, errors.New("the env [CORE_ADDR] does not exist")
	}

	return config, nil
}
