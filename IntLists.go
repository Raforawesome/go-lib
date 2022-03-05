package lists

// IntList INTEGER LISTS
type IntList struct {
	Components []int
	Length     int
}

func (arr *IntList) Append(toAdd int) {
	newLength := arr.Length + 1
	newList := make([]int, newLength, newLength)

	for i, v := range arr.Components {
		newList[i] = v
	}
	newList[newLength-1] = toAdd

	arr.Components = newList
	arr.Length = newLength
}

func (arr *IntList) Entries() []int {
	return arr.Components
}

func (arr *IntList) SliceAt(index int) bool {
	arr.Components = append(arr.Components[:index], arr.Components[index+1:]...)
	return false
}

func (arr *IntList) SliceElement(element int) bool {
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

func (arr *IntList) IndexOf(element int) int {
	for i, v := range arr.Components {
		if v == element {
			return i
		}
	}
	return -1
}

// NewIntList exporting constructor functions
func NewIntList() IntList {
	list := IntList{}
	list.Components = make([]int, 0, 0)
	list.Length = 0
	return list
}
