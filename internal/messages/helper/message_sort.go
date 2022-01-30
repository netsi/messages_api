package helper

import "messages_api/internal/messages/model"

// ByCreationDateDesc implements sort.Interface.
type ByCreationDateDesc []model.Message

// Len returns the length of the slice
func (b ByCreationDateDesc) Len() int { return len(b) }

// Less reports whether the element with index i must sort before the element with index j.
func (b ByCreationDateDesc) Less(i, j int) bool { return b[i].CreationDate.After(b[j].CreationDate) }

// Swap swaps the elements with indexes i and j.
func (b ByCreationDateDesc) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
