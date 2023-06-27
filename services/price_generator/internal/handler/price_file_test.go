package handler

import (
	"context"
	"testing"

	pb "github.com/bearatol/interview_golang_task/proto/price_generator"
	priceFileMock "github.com/bearatol/interview_golang_task/sevices/price_generator/internal/handler/mock_handler"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestPriceFileService(t *testing.T) {
	type mockBehavior func(
		r *priceFileMock.MockServicePriceFile,
		fileName,
		barcode,
		title string,
		cost int32,
	)

	tests := []struct {
		name,
		fileName,
		barcode,
		title string
		cost         int32
		mockBehavior mockBehavior
	}{
		{
			name:     "save file",
			fileName: "doc_test-barcode_test-date.pdf",
			barcode:  "test-barcode",
			title:    "test-title",
			cost:     10,
			mockBehavior: func(
				p *priceFileMock.MockServicePriceFile,
				fileName,
				barcode,
				title string,
				cost int32,
			) {
				p.EXPECT().SetFile(
					fileName,
					barcode,
					title,
					cost,
				).Return(nil)
			},
		},
		{
			name:     "get file",
			fileName: "doc_test-barcode_test-date.pdf",
			barcode:  "test-barcode",
			title:    "test-title",
			cost:     0,
			mockBehavior: func(
				p *priceFileMock.MockServicePriceFile,
				fileName,
				barcode,
				title string,
				cost int32,
			) {
				p.EXPECT().GetFileByName(fileName).Return([]byte{}, nil)
			},
		}, {
			name:     "delete file",
			fileName: "doc_test-barcode_test-date.pdf",
			barcode:  "test-barcode",
			title:    "test-title",
			cost:     0,
			mockBehavior: func(
				p *priceFileMock.MockServicePriceFile,
				fileName,
				barcode,
				title string,
				cost int32,
			) {
				p.EXPECT().DeleteFileByName(fileName).Return(nil)
			},
		},
	}

	for k, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			serv := priceFileMock.NewMockServicePriceFile(c)
			test.mockBehavior(
				serv,
				test.fileName,
				test.barcode,
				test.title,
				test.cost,
			)
			handl := &Handler{
				ctx:           context.TODO(),
				servPriceFile: serv,
			}

			var err error
			switch {
			case k == 0:
				_, err = handl.Set(context.TODO(), &pb.PriceFileSetRequest{
					FileName: test.fileName,
					Barcode:  test.barcode,
					Title:    test.title,
					Cost:     test.cost,
				})
			case k == 1:
				_, err = handl.Get(context.TODO(), &pb.PriceFileRequest{
					FileName: test.fileName,
				})
			case k == 2:
				_, err = handl.Delete(context.TODO(), &pb.PriceFileRequest{
					FileName: test.fileName,
				})
			}

			assert.Equal(t, err, nil)
		})
	}
}
