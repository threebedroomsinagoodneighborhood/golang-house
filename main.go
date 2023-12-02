package house

import "house/assets"

func main() {
	assets.CreateHouse()
}

// дз: сделать свой дом package furniture; import "папка/furniture"; type table struct;
// необходимо все размерные величины xyz, минимум 5 мебели и бытовой техники, состав семьи, функция на вид всего дома
// в мейне вызов только одной showHouse() и создания объектов
// gitignore для .idea
//	type House struct{
//		Tables [] furniture.Table
//
//func main() {
//	showHouse()
//}
