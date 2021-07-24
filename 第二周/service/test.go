package service

type Test struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (svc *Service) GetTest(param *Test) (*Test, error) {
	test, err := svc.dao.GetTest(param.ID)
	if err != nil {
		return nil, err
	}
	return &Test{
		ID:      test.ID,
		Title:   test.Title,
		Content: test.Content,
	}, nil
}
