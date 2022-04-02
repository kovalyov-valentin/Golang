// Task. Начнем с создания прототипа, что генерирует 10 случайных билетов и отображает их в табличном виде.

// В таблице четыре столбца:
// 	Космическая станция (Spaceline), что предоставляет услуги;
// 	Продолжительность (Duration) в днях поездки на Марс в один конец;
// 	Покрывает ли цена поездку туда и обратно (Trip type);
// 	Цена (Price) в миллионах долларов.
// 	Для каждого билета случайным образом выбирается космическая станция: Space Adventures, SpaceX или Virgin Galactic.

// Датой отправления на каждом билете значится 13 Октября 2020 года. В этот день Марс будет на расстоянии 62 100 000 км от Земли.
	
// Скорость космического корабля будет выбрана случайным образом из диапазона от 16 до 30 км/ч. 
// Это определит продолжительность поездки на Марс, а также цену билета. 
// Более быстрые корабли намного дороже. 
// Цены на билеты варьируются от $36 до $50 миллионов. 
// Цена для поездки туда-обратно удваивается.

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// const secondsPerDay = 86400

// func main() {
// 	distance := 62100000
// 	spaceline := ""
// 	trip := ""

// 	fmt.Println("Spaceline        Days Trip type  Price")
// 	fmt.Println("=======================================")
// 	for ticket := 0; ticket < 10; ticket++ {
// 		switch rand.Intn(3) {
// 		case 0:
// 			spaceline = ("Space Adventures")
// 		case 1:
// 			spaceline = ("SpaceX")
// 		case 2:
// 			spaceline = ("Virgin Galactic")

// 		}
// 		speed := rand.Intn(15) + 16
// 		duration := distance / speed / secondsPerDay
// 		price := 20 + speed

// 		if rand.Intn(2) == 1 {
// 			trip = "Round-trip"
// 			price = price * 2
// 		} else {
// 			trip = "One-way"
// 		}
// 		fmt.Printf("%-16v %4v %-10v &%4v\n", spaceline, duration, trip, price)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// const VG = "Virgin Galactic"  // 1-й перевозчик
// const SX = "SpaceX"           // 2-й перевозчик
// const SA = "Space Adventures" // 3-й перевозчик
// const OW = "One-way"          // В один конец
// const RT = "Round-trip"       // С возвратом
// const distance = 62100000     // км до марса
// const hourPerDay = 24         // 24 часа в 1 сутках
// const secPerHour = 60 * 60    // секунд в 1 часе

// func main() {
// 	spaceLine := ""
// 	tripType := ""
// 	fmt.Printf("%-16v %4v %-10v %v\n", "SpaceLine", "Days", "Trip type", "Price")
// 	fmt.Println("======================================")
// 	for i := 0; i < 10; i++ { // Генерация 10-ти билетов
// 		speed := rand.Intn(15) + 16                          // скорость км/с в диапазоне от 16-30
// 		cost := speed + 20                                   // млн долларов, цена перелета
// 		days := (distance / speed) / secPerHour / hourPerDay // Продолжительность путешествия в днях
// 		switch rand.Intn(2) + 1 {
// 		case 1:
// 			tripType = OW
// 		case 2:
// 			tripType = RT
// 			cost *= 2 // Удвоим цену билета с возвратом
// 			days *= 2 // удвоим количество дней с возвратом
// 		}

// 		switch rand.Intn(3) + 1 {
// 		case 1:
// 			spaceLine = VG
// 		case 2:
// 			spaceLine = SX
// 		case 3:
// 			spaceLine = SA
// 		}
// 		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceLine, days, tripType, cost)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// var (
// 	spacelines, tripTypes         []string
// 	dur, cost, velocity, category int
// )

// const distance = 62100000 // km

// func main() {
// 	spacelines = []string{
// 		"Space Adventures",
// 		"SpaceX",
// 		"Virgin Galactic",
// 	}
// 	tripTypes = []string{
// 		"Round-trip",
// 		"One-way",
// 	}
// 	fmt.Printf("%-16v %4v %-14v %v\n", "Spaceline", "Days", "Trip type", "Price")
// 	fmt.Println("==========================================")
// 	for i := 0; i < 10; i++ {
// 		category = rand.Intn(14)
// 		velocity = category + 16
// 		cost = category + 36
// 		dur = distance / velocity / 86400
// 		selectSpaceline := rand.Intn(3)
// 		selectTripType := rand.Intn(2)
// 		if selectTripType == 0 {
// 			cost *= 2
// 		}
// 		fmt.Printf("%-16v %4v %-14v %v\n", spacelines[selectSpaceline], dur, tripTypes[selectTripType], cost)
// 	}
// }

package main

import (
	"fmt"
	"math/rand"
)

const secondPerDay = 86400

func main() {
	distance := 62100000

	spaceline := []string{
		"Space Adventures",
		"SpaceX",
		"Virgin Galactic",
	}
	tripType := []string{
		"Round-trip",
		"One-way",
	}

	fmt.Printf("%-16v %4v %-10v %5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("=======================================")

	for ticket := 0; ticket < 10; ticket++ {
		speed := rand.Intn(15) + 16
		duration := distance / speed / secondPerDay
		price := 20.0 + speed

		selectSpaceline := rand.Intn(3)
		selectTripType := rand.Intn(2)

		if selectTripType == 0 {
			price *= 2
		}

		// if rand.Intn(2) == 1 {
		// 	tripType = "Round-trip"
		// 	price *= 2
		// } else {
		// 	tripType = "One-way"
		// }
		fmt.Printf("%-16v %4v %-10v $%5v\n", spaceline[selectSpaceline], duration, tripType[selectTripType], price)
	}

}
