package main
import(
  "sync"
)
// Set new 
type Set struct {
  m map[int]bool
  sync.RWMutex
}
// New new 
func New() *Set {
  return &Set{
    m: map[int]bool{},
  }
}
// Add add
func (s *Set) Add(item int) {
  s.Lock()
  defer s.Unlock()
  s.m[item] = true
}
// Remove Remove
func (s *Set) Remove(item int) {
  s.Lock()
  s.Unlock()
  delete(s.m, item)
}
// Has Has
func (s *Set) Has(item int) bool {
  s.RLock()
  defer s.RUnlock()
  _, ok := s.m[item]
  return ok
}
// Len Len
func (s *Set) Len() int {
  return len(s.List())
}
// Clear Clear
func (s *Set) Clear() {
  s.Lock()
  defer s.Unlock()
  s.m = map[int]bool{}
}
// IsEmpty IsEmpty
func (s *Set) IsEmpty() bool {
  if s.Len() == 0 {
    return true
  }
  return false
}
// List List
func (s *Set) List() []int {
  s.RLock()
  defer s.RUnlock()
  list := []int{}
  for item := range s.m {
    list = append(list, item)
  }
  return list
}
