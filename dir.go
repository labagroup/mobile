package mobile

type DirInfo struct {
	Document  string
	Cache     string
	Temporary string
}

func NewDirInfo() *DirInfo {
	return new(DirInfo)
}
