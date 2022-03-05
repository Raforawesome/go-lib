package lists

// FloatList STRING LISTS
type FloatList struct {
	Components []float64
	Length     int
}

func (arr *FloatList) Append(toAdd float64) {
	newLength := arr.Length + 1
	newList := make([]float64, newLength, newLength)

	for i, v := range arr.Components {
		newList[i] = v
	}
	newList[newLength-1] = toAdd

	arr.Components = newList
	arr.Length = newLength
}

func (arr *FloatList) Entries() []float64 {
	return arr.Components
}

func (arr *FloatList) SliceAt(index int) bool {
	arr.Components = append(arr.Components[:index], arr.Components[index+1:]...)
	return false
}

func (arr *FloatList) SliceElement(element float64) bool {
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

func (arr *FloatList) IndexOf(element float64) int {
	for i, v := range arr.Components {
		if v == element {
			return i
		}
	}
	return -1
}

// NewFloat64List exporting constructor functions
func NewFloat64List() FloatList {
	list := FloatList{}
	list.Components = make([]float64, 0, 0)
	list.Length = 0
	return list
}
