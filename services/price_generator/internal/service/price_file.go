package service

type PDF interface {
	Set(fileName, barcode, title string, cost int32) error
	Get(fileName string) ([]byte, error)
	Delete(fileName string) error
}

func (s *Service) SetFile(fileName, barcode, title string, cost int32) error {
	return s.pdf.Set(fileName, barcode, title, cost)
}
func (s *Service) GetFileByName(fileName string) ([]byte, error) {
	return s.pdf.Get(fileName)
}
func (s *Service) DeleteFileByName(fileName string) error {
	return s.pdf.Delete(fileName)
}
