package main
/*
중재자 패턴의 훌륭한 예는 철도역 교통 시스템입니다.
두 열차는 플랫폼의 가용성에 대해 서로 통신하지 않습니다.
중재자 역할을 하며 도착 하는 stationManager
열차 중 하나만 플랫폼을 사용할 수 있게 하고
나머지는 대기열에 보관합니다.
출발하는 열차는 역에 알려서 다음 열차가 도착할 수 있도록 합니다.
*/
type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}