package mobile

import (
	"encoding/json"

	"github.com/gopub/types"
)

type Video types.Video

func NewVideo() *Video {
	return new(Video)
}

func (v *Video) GetImage() *Image {
	return (*Image)(v.Image)
}

type Audio types.Audio

func NewAudio() *Audio {
	return new(Audio)
}

type File types.File

func NewFile() *File {
	return new(File)
}

type WebPage types.WebPage

func NewWebPage() *WebPage {
	return new(WebPage)
}

type Any types.Any

func (a *Any) TypeName() string {
	return (*types.Any)(a).TypeName()
}

func (a *Any) SetImage(i *Image) {
	(*types.Any)(a).SetValue((*types.Image)(i))
}

func (a *Any) SetAudio(au *Audio) {
	(*types.Any)(a).SetValue((*types.Audio)(au))
}

func (a *Any) SetVideo(v *Video) {
	(*types.Any)(a).SetValue((*types.Video)(v))
}

func (a *Any) SetFile(f *File) {
	(*types.Any)(a).SetValue((*types.File)(f))
}

func (a *Any) SetWebPage(wp *WebPage) {
	(*types.Any)(a).SetValue((*types.WebPage)(wp))
}

func (a *Any) Image() *Image {
	return (*Image)((*types.Any)(a).Value().(*types.Image))
}

func (a *Any) Video() *Video {
	return (*Video)((*types.Any)(a).Value().(*types.Video))
}

func (a *Any) Audio() *Audio {
	return (*Audio)((*types.Any)(a).Value().(*types.Audio))
}

func (a *Any) File() *File {
	return (*File)((*types.Any)(a).Value().(*types.File))
}

func (a *Any) WebPage() *WebPage {
	return (*WebPage)((*types.Any)(a).Value().(*types.WebPage))
}

func NewAnyObj() *Any {
	return new(Any)
}

type AnyList struct {
	List []*types.Any
}

func NewAnyListObj() *AnyList {
	return new(AnyList)
}

func NewAnyList(list []*types.Any) *AnyList {
	return &AnyList{List: list}
}

func (a *AnyList) Size() int {
	if a == nil {
		return 0
	}
	return len(a.List)
}

func (a *AnyList) Get(index int) *Any {
	if a == nil {
		return nil
	}
	return (*Any)(a.List[index])
}

func (a *AnyList) Append(v *Any) {
	a.List = append(a.List, (*types.Any)(v))
}

func (a *AnyList) Prepend(v *Any) {
	a.List = append([]*types.Any{(*types.Any)(v)}, a.List...)
}

func (a *AnyList) Insert(i int, v *Any) {
	if len(a.List) <= i {
		a.List = append(a.List, (*types.Any)(v))
	} else {
		l := a.List[i:]
		l = append([]*types.Any{(*types.Any)(v)}, l...)
		a.List = append(a.List[0:i], l...)
	}
}

func (a *AnyList) RemoveAt(index int) {
	a.List = append(a.List[0:index], a.List[index+1:]...)
}

func (a *AnyList) Remove(v *Any) {
	i := a.IndexOf(v)
	if i >= 0 {
		a.RemoveAt(i)
	}
}

func (a *AnyList) IndexOf(v *Any) int {
	for i, m := range a.List {
		if (*Any)(m) == v {
			return i
		}
	}
	return -1
}

func (a *AnyList) FirstImage() *Image {
	for _, m := range a.List {
		if img, ok := m.Value().(*types.Image); ok {
			return (*Image)(img)
		}
	}
	return nil
}

func (a *AnyList) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &a.List)
}

func (a *AnyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.List)
}
