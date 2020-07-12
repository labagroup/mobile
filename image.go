package mobile

import "github.com/gopub/types"

type Image types.Image

func NewImage() *Image {
	return new(Image)
}

type ImageList struct {
	List []*types.Image
}

func (l *ImageList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *ImageList) Get(index int) *types.Image {
	return l.List[index]
}
