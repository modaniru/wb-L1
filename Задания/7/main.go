package main

import "sync"

/*
	Создаем структуру синхранизированной мапы, в ней у нас находятся мапа и RWMutex.
*/
type syncMap[K comparable, V any] struct{
	myMap map[K]V
	mutex sync.RWMutex
}
/*
	Конструктор
*/
func NewSyncMap[K comparable, V any]() *syncMap[K,V]{
	return &syncMap[K,V]{myMap: map[K]V{}, mutex: sync.RWMutex{}}
}

/*
	Установка ключа и значения в мапу
	Ставим блокировку на запись, в этот момент доступ к syncMap.myMap будет иметь только эта функция.
*/
func (s *syncMap[K, V]) Set(key K, value V){
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.myMap[key] = value
}

/*
	Получения значения по ключу
	Ставим блокировку на чтение. Они не начинают свое выполнение, пока существует любая блокировка на запись.
*/
func (s *syncMap[K, V]) Get(key K) (V, bool){
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	value, ok := s.myMap[key]
	return value, ok
}

/*
	Удаление по ключу
	Ставим блокировку на запись, в этот момент доступ к syncMap.myMap будет иметь только эта функция.
*/
func (s *syncMap[K, V]) Delete(key K){
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.myMap, key)
}
//TODO протестировать, запускать цикл горутин, рандомных операций над мапой??
func main(){
	mapa := NewSyncMap[int, string]()
	mapa.Set(2, "21")
}