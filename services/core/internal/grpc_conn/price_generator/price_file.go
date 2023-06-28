package price_generator

import (
	"context"
	"fmt"

	pb "github.com/bearatol/interview_golang_task/proto/price_generator"
	"github.com/bearatol/lg"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
)

type priceFile struct {
	*grpcConn
}

func NewPriceFile(g *grpcConn) *priceFile {
	return &priceFile{g}
}

func (f *priceFile) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, mapping.TimeoutConnect)
	defer cancel()

	_, err := f.client.Ping(ctx, &empty.Empty{})
	return err
}

func (f *priceFile) Create(ctx context.Context, fileName, barcode, title string, cost int32) error {
	ctx, cancel := context.WithTimeout(ctx, mapping.TimeoutConnect)
	defer cancel()

	if len(fileName) == 0 {
		return fmt.Errorf("file name is empty")
	}
	if len(barcode) == 0 {
		return fmt.Errorf("barcode is empty")
	}
	if len(title) == 0 {
		return fmt.Errorf("title is empty")
	}

	lg.Debugf("%+v/n", f.Ping(ctx))

	_, err := f.client.Set(ctx, &pb.PriceFileSetRequest{
		FileName: fileName,
		Barcode:  barcode,
		Title:    title,
		Cost:     cost,
	})

	lg.Debug(err)

	return err
}

func (f *priceFile) Get(ctx context.Context, fileName string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, mapping.TimeoutConnect)
	defer cancel()

	if len(fileName) == 0 {
		return nil, fmt.Errorf("file name is empty")
	}

	res, err := f.client.Get(ctx, &pb.PriceFileRequest{
		FileName: fileName,
	})
	if err != nil {
		return nil, err
	}
	if res != nil && len(res.File) != 0 {
		return res.File, nil
	}

	return nil, fmt.Errorf("file [%s] not found", fileName)
}

func (f *priceFile) Delete(ctx context.Context, fileName string) error {
	ctx, cancel := context.WithTimeout(ctx, mapping.TimeoutConnect)
	defer cancel()

	if len(fileName) == 0 {
		return fmt.Errorf("file name is empty")
	}

	_, err := f.client.Delete(ctx, &pb.PriceFileRequest{
		FileName: fileName,
	})

	return err
}
