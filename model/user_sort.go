package model

// ByUserID implements sort.Interface for []User based on the ID field.
type ByUserID []User

func (a ByUserID) Len() int           { return len(a) }
func (a ByUserID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUserID) Less(i, j int) bool { return a[i].ID < a[j].ID }
