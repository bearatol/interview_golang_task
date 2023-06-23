package price_generator

import (
	"context"
	"fmt"

	pb "github.com/bearatol/interview_golang_task/proto/price_generator"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
)

type priceFile struct {
	*grpcConn
}

func NewPriceFile(g *grpcConn) *priceFile {
	return &priceFile{g}
}

func (f *priceFile) Create(ctx context.Context, barcode, title string, cost int32) error {
	ctx, cancel := context.WithTimeout(ctx, mapping.TimeoutConnect)
	defer cancel()

	if len(barcode) == 0 {
		return fmt.Errorf("barcode is empty")
	}

	_, err := f.client.Set(ctx, &pb.PriceFileSetRequest{
		Barcode: barcode,
		Title:   title,
		Cost:    cost,
	})

	return err
}

func (f *priceFile) Get(ctx context.Context, barcodes []string) ([]*mapping.PriceFile, error) {
	ctx, cancel := context.WithTimeout(ctx, mapping.TimeoutConnect)
	defer cancel()

	if len(barcodes) == 0 {
		return nil, fmt.Errorf("barcodes is empty")
	}

	res, err := f.client.Get(ctx, &pb.PriceFilesRequest{
		Barcodes: barcodes,
	})
	if err != nil {
		return nil, err
	}
	if res != nil && len(res.Files) != 0 {
		fileList := make([]*mapping.PriceFile, 0, len(res.Files))
		for _, file := range res.Files {
			fileList = append(fileList, &mapping.PriceFile{
				Name:    file.Name,
				Content: file.Content,
			})
		}
		return fileList, nil
	}

	return nil, nil
}

func (f *priceFile) Delete(ctx context.Context, barcodes []string) error {
	ctx, cancel := context.WithTimeout(ctx, mapping.TimeoutConnect)
	defer cancel()

	if len(barcodes) == 0 {
		return fmt.Errorf("barcodes is empty")
	}

	_, err := f.client.Delete(ctx, &pb.PriceFilesRequest{
		Barcodes: barcodes,
	})

	return err
}
