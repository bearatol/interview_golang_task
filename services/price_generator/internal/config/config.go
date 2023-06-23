package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	AppAddr   string
	PDFConfig *PDFConfig
}

type PDFConfig struct {
	StorePath,
	FontPath,
	ImagePath,
	FontColorRGB,
	PageSize,
	ImageSize,
	ImageCoordinate,
	BarcodeCoordinate,
	TitleCoordinate,
	CostCoordinate string

	FontSize uint8
}

func NewConfig() (*Config, error) {
	var (
		exist  bool
		config = &Config{}
		env    string
	)

	if config.AppAddr, exist = os.LookupEnv("PRICE_GENERATOR_ADDR_LOCAL"); !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_ADDR_LOCAL] does not exist")
	}

	config.PDFConfig.StorePath, exist = os.LookupEnv("PRICE_GENERATOR_PDF_DIR_STORE")
	if !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_PDF_DIR_STORE] does not exist")
	}
	config.PDFConfig.FontPath, exist = os.LookupEnv("PRICE_GENERATOR_PDF_FONT_PATH")
	if !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_PDF_FONT_PATH] does not exist")
	}
	config.PDFConfig.ImagePath, exist = os.LookupEnv("PRICE_GENERATOR_PDF_IMAGE_PATH")
	if !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_PDF_IMAGE_PATH] does not exist")
	}
	config.PDFConfig.FontColorRGB, exist = os.LookupEnv("PRICE_GENERATOR_PDF_FONT_COLOR_RGB")
	if !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_PDF_FONT_COLOR_RGB] does not exist")
	}
	config.PDFConfig.PageSize, exist = os.LookupEnv("PRICE_GENERATOR_PDF_PAGE_SIZE")
	if !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_PDF_PAGE_SIZE] does not exist")
	}
	config.PDFConfig.ImageSize, exist = os.LookupEnv("PRICE_GENERATOR_PDF_IMAGE_SIZE")
	if !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_PDF_IMAGE_SIZE] does not exist")
	}
	config.PDFConfig.ImageCoordinate, exist = os.LookupEnv("PRICE_GENERATOR_PDF_IMAGE_COORDINATE")
	if !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_PDF_IMAGE_COORDINATE] does not exist")
	}
	config.PDFConfig.BarcodeCoordinate, exist = os.LookupEnv("PRICE_GENERATOR_PDF_BARCODE_COORDINATE")
	if !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_PDF_BARCODE_COORDINATE] does not exist")
	}
	config.PDFConfig.TitleCoordinate, exist = os.LookupEnv("PRICE_GENERATOR_PDF_TITLE_COORDINATE")
	if !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_PDF_TITLE_COORDINATE] does not exist")
	}
	config.PDFConfig.CostCoordinate, exist = os.LookupEnv("PRICE_GENERATOR_PDF_COST_COORDINATE")
	if !exist {
		return nil, errors.New("the env [PRICE_GENERATOR_PDF_COST_COORDINATE] does not exist")
	}
	{

		env, exist = os.LookupEnv("PRICE_GENERATOR_PDF_FONT_SIZE")
		if !exist {
			return nil, errors.New("the env [PRICE_GENERATOR_PDF_FONT_SIZE] does not exist")
		}
		fontSizeUint64, err := strconv.ParseUint(env, 10, 8)
		if err != nil {
			return nil, err
		}
		config.PDFConfig.FontSize = uint8(fontSizeUint64)
	}

	return config, nil
}
