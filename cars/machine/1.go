package machine

/*

База данных авто на продажу
Конфиг

1. Добавить в Конфиг год выпуска авто
2. Добавить а конфиг цену
3. Добавить поле продана или нет

Методы

1. Получение инфы о статусе продажи (продана или нет) GetStatusBuyCar() bool
2. Метод на покупку                                   BuyCar()bool
3. Метод на получение цены                            GetPrice() int
*/

type Config struct {
	Power              int
	DoorsCount         int
	MarkCar            string
	ModelCar           string
	YearManufactureCar int
	Price              int
	StatusBuyCar       bool
}

type infoCar struct {
	config Config
}

func (i *infoCar) GetStatusBuyCar() bool {
	return i.config.StatusBuyCar
}

func (i *infoCar) BuyCar() bool {
	if i.config.StatusBuyCar {
		return false
	} else {
		i.config.StatusBuyCar = true
	}
	return true
}

func (i *infoCar) GetPrice() int {
	return i.config.Price
}

func NewcInfoCar(cfg Config) *infoCar {
	car := infoCar{
		config: cfg,
	}
	return &car
}
