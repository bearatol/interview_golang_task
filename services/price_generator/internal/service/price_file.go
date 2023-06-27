package service

import "github.com/bearatol/lg"

type PDF interface {
	Set(fileName, barcode, title string, cost int32) error
	Get(fileName string) ([]byte, error)
	Delete(fileName string) error
}

func (s *Service) SetFile(fileName, barcode, title string, cost int32) error {
	err := s.pdf.Set(fileName, barcode, title, cost)
	if err != nil {
		lg.Error(err)
		return err
	}
	return nil
}
func (s *Service) GetFileByName(fileName string) ([]byte, error) {
	res, err := s.pdf.Get(fileName)
	if err != nil {
		lg.Error(err)
		return nil, err
	}
	return res, nil
}
func (s *Service) DeleteFileByName(fileName string) error {
	err := s.pdf.Delete(fileName)
	if err != nil {
		lg.Error(err)
		return err
	}
	return nil
}
