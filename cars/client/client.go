package client

/*
   - Добавить новый обьект (структуру) кошелек 
   . Кол-во денег 
      методы к кошельку:
	  - Оплатить          (k *koxrl)Pay(sum int)bool
	  - Проверить кол-во денег в кошельке  (k *koxrl)GetSumm()(int)


   Client: 
    - Метод на покупку авто 


*/
type Config struct{
	Car carse
}


type client struct{
	cfg Config
}


type carse interface{
	GetPrice() int
	BuyCar() bool
}


func NewClient(cfg Config)(*client){
		return &client{
			cfg: cfg,
		} 
}


func (c *client)PayCar(){
	priceBMW := c.config.Car.GetPrice()
	// проверяем хвататет ли денег 
	// если да то покупаем и радуемся 
	// 
}