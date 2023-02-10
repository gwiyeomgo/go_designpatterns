package main

func main() {
	stationManager := newStationManger()
	//승객기차
	passengerTrain := &PassengerTrain{
		mediator: stationManager, //스테이션 메니저 = 중재인
	}
	//화물열차
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
