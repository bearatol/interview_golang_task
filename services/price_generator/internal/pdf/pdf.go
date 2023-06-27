package pdf

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/bearatol/interview_golang_task/sevices/price_generator/internal/config"
	"github.com/signintech/gopdf"
)

type PDF struct {
	DirStore  string
	FontFile  io.Reader
	ImageFile image.Image

	FontSize,
	FontColorR,
	FontColorG,
	FontColorB uint8

	PageSize  *gopdf.Rect
	ImageSize *gopdf.Rect
	ImageX,
	ImageY,
	BarcodeX,
	BarcodeY,
	TitleX,
	TitleY,
	CostX,
	CostY float64
}

func NewPDF(conf *config.Config) (p *PDF, err error) {
	p = &PDF{}
	path, err := os.Getwd()
	if err != nil {
		return
	}

	p.DirStore = fmt.Sprintf("%s/%s", path, conf.PDFConfig.StorePath)
	if _, err = os.Stat(p.DirStore); os.IsNotExist(err) {
		err = fmt.Errorf("%s doesn't exist", p.DirStore)
		return
	} else if err != nil {
		return
	}
	p.FontFile, err = p.getFont(fmt.Sprintf("%s/%s", path, conf.PDFConfig.FontPath))
	if err != nil {
		return
	}
	p.ImageFile, err = p.getImg(fmt.Sprintf("%s/%s", path, conf.PDFConfig.ImagePath))
	if err != nil {
		return
	}

	p.FontSize = conf.PDFConfig.FontSize
	p.FontColorR, p.FontColorG, p.FontColorB, err = p.getRGB(conf.PDFConfig.FontColorRGB)
	if err != nil {
		return
	}

	w, h, err := p.getSplitedPositionAndSize(conf.PDFConfig.PageSize)
	if err != nil {
		return
	}
	p.PageSize = &gopdf.Rect{W: w, H: h}

	w, h, err = p.getSplitedPositionAndSize(conf.PDFConfig.ImageSize)
	if err != nil {
		return
	}
	p.ImageSize = &gopdf.Rect{W: w, H: h}

	p.ImageX, p.ImageY, err = p.getSplitedPositionAndSize(conf.PDFConfig.ImageCoordinate)
	if err != nil {
		return
	}
	p.BarcodeX, p.BarcodeY, err = p.getSplitedPositionAndSize(conf.PDFConfig.BarcodeCoordinate)
	if err != nil {
		return
	}
	p.TitleX, p.TitleY, err = p.getSplitedPositionAndSize(conf.PDFConfig.TitleCoordinate)
	if err != nil {
		return
	}
	p.CostX, p.CostY, err = p.getSplitedPositionAndSize(conf.PDFConfig.CostCoordinate)

	return
}

func (p *PDF) getSplitedPositionAndSize(xY string) (x, y float64, err error) {
	xYList := strings.Split(xY, "x")
	if len(xYList) < 2 {
		err = fmt.Errorf("X and Y are invalid")
		return
	}
	x, err = strconv.ParseFloat(xYList[0], 64)
	if err != nil {
		return
	}
	y, err = strconv.ParseFloat(xYList[1], 64)
	if err != nil {
		return
	}
	return
}

func (p *PDF) getFont(ttfPath string) (io.Reader, error) {
	if _, err := os.Stat(ttfPath); os.IsNotExist(err) {
		return nil, err
	}
	data, err := ioutil.ReadFile(ttfPath)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func (p *PDF) getImg(imgPath string) (image.Image, error) {
	if _, err := os.Stat(imgPath); os.IsNotExist(err) {
		return nil, err
	}
	data, err := os.ReadFile(imgPath)
	if err != nil {
		return nil, err
	}
	return png.Decode(bytes.NewReader(data))
}

func (p *PDF) getRGB(rgbColor string) (r, g, b uint8, err error) {
	rgbList := strings.Split(rgbColor, ".")
	if len(rgbList) < 3 {
		err = fmt.Errorf("rgb is invalid")
		return
	}
	rUint64, err := strconv.ParseUint(rgbList[0], 10, 8)
	if err != nil {
		err = fmt.Errorf("red color is invalid: %s", err)
		return
	}
	gUint64, err := strconv.ParseUint(rgbList[1], 10, 8)
	if err != nil {
		err = fmt.Errorf("green color is invalid: %s", err)
		return
	}
	bUint64, err := strconv.ParseUint(rgbList[2], 10, 8)
	if err != nil {
		err = fmt.Errorf("blue color is invalid: %s", err)
		return
	}
	return uint8(rUint64), uint8(gUint64), uint8(bUint64), nil
}

func (p *PDF) createPDF(barcode, title string, cost int32) ([]byte, error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *p.PageSize})
	defer pdf.Close()

	pdf.AddPage()
	if err := pdf.AddTTFFontByReader("font", p.FontFile); err != nil {
		return nil, err
	}
	err := pdf.ImageFrom(p.ImageFile, p.ImageX, p.BarcodeY, p.ImageSize)
	if err != nil {
		return nil, err
	}
	err = pdf.SetFont("font", "", p.FontSize)
	if err != nil {
		return nil, err
	}
	pdf.SetTextColor(p.FontColorR, p.FontColorG, p.FontColorB)

	pdf.SetXY(p.BarcodeX, p.BarcodeY)
	if err := pdf.Cell(nil, barcode); err != nil {
		return nil, err
	}

	pdf.SetXY(p.TitleX, p.TitleY)
	if err := pdf.Cell(nil, title); err != nil {
		return nil, err
	}

	pdf.SetXY(p.CostX, p.CostY)
	if err := pdf.Cell(nil, fmt.Sprintf("%d", cost)); err != nil {
		return nil, err
	}

	return pdf.GetBytesPdf(), nil
}

func (p *PDF) Set(fileName, barcode, title string, cost int32) error {
	pdf, err := p.createPDF(barcode, title, cost)
	if err != nil {
		return err
	}

	return os.WriteFile(p.DirStore+"/"+fileName, pdf, 0644)
}

func (p *PDF) Get(fileName string) ([]byte, error) {
	return os.ReadFile(p.DirStore + "/" + fileName)
}

func (p *PDF) Delete(fileName string) error {
	return os.Remove(p.DirStore + "/" + fileName)
}
