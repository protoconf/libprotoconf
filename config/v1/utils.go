package v1

type ByPrio []*LibprotoconfConfig_Loadable

// Len is i the sort.Interface
func (b ByPrio) Len() int {
	return len(b)
}

// Less is in the sort.Interface
func (b ByPrio) Less(i, j int) bool {
	// Higher prio comes first
	if b[i].GetPriority() == b[j].GetPriority() {
		return b[i].GetPath() > b[j].GetPath()
	}
	return b[i].GetPriority() > b[j].GetPriority()
}

// Swap is in the sort.Interface
func (b ByPrio) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
