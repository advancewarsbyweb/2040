package uploadutils

type IS3Utils interface {
	GetObject()
	PutObject()
}

type S3Utils struct {
}

type MockS3Utils struct {
}

func (u *S3Utils) GetObject() {
}


func (u *
