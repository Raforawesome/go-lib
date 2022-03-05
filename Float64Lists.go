package lists

// Float64List STRING LISTS
type Float64List struct {
	Components []float64
	Length     int
}

func (arr *Float64List) Append(toAdd float64) {
	newLength := arr.Length + 1
	newList := make([]float64, newLength, newLength)

	for i, v := range arr.Components {
		newList[i] = v
	}
	newList[newLength-1] = toAdd

	arr.Components = newList
	arr.Length = newLength
}

func (arr *Float64List) Entries() []float64 {
	return arr.Components
}

func (arr *Float64List) SliceAt(index int) bool {
	arr.Components = append(arr.Components[:index], arr.Components[index+1:]...)
	return false
}

func (arr *Float64List) SliceElement(element float64) bool {
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

func (arr *Float64List) IndexOf(element float64) int {
	for i, v := range arr.Components {
		if v == element {
			return i
		}
	}
	return -1
}

// NewFloat64List exporting constructor functions
func NewFloat64List() Float64List {
	list := Float64List{}
	list.Components = make([]float64, 0, 0)
	list.Length = 0
	return list
}
