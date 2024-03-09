package main

import "fmt"

// Набросок деревьев и сегментации пользователей для реализации платформы ценообразования
// Деревья можно использовать as-is или переделать под себя, наработки даны просто как пример
func main() {
	//GetCategoriesTree().PrintTree(0)

	//GetLocationsTree().PrintTree(0)

	//fmt.Println(GetSegmentsByUserIDs([]int64{2000, 2600, 2700, 2888}))
	fmt.Println(GetSegmentsByUserIDs([]int64{2500}))
}
