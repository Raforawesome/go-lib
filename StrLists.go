package lists

// StrList STRING LISTS
type StrList struct {
	Components []string
	Length     int
}

func (arr *StrList) Append(toAdd string) {
	newLength := arr.Length + 1
	newList := make([]string, newLength, newLength)

	for i, v := range arr.Components {
		newList[i] = v
	}
	newList[newLength-1] = toAdd

	arr.Components = newList
	arr.Length = newLength
}

func (arr *StrList) Entries() []string {
	return arr.Components
}

func (arr *StrList) SliceAt(index int) bool {
	arr.Components = append(arr.Components[:index], arr.Components[index+1:]...)
	return false
}

func (arr *StrList) SliceElement(element string) bool {
	indexAt := -1
	for i, v := range arr.Components {
		if element == v {
			indexAt = i
		}
	}

	if indexAt > -1 {
		arr.SliceAt(indexAt)
		return true
	} else {
		return false
	}
}

func (arr *StrList) IndexOf(element string) int {
	for i, v := range arr.Components {
		if v == element {
			return i
		}
	}
	return -1
}

// NewStrList exporting constructor functions
func NewStrList() StrList {
	list := StrList{}
	list.Components = make([]string, 0, 0)
	list.Length = 0
	return list
}
