package main

//https://refactoring.guru/design-patterns/bridge
import "fmt"

// 추상화
type Remote interface {
	volumeUp()
	volumeDown()
	mute()
}

// 정제된 추상화

type RemoteControl struct {
	device Device
}

func (a *RemoteControl) volumeUp() {
	a.device.setVolume(a.device.getVolume() + 10)
}
func (a *RemoteControl) volumeDown() {
	a.device.setVolume(a.device.getVolume() - 10)
}
func (a *RemoteControl) mute() { //무음
	a.device.setVolume(0)
}

type AdvancedRemoteControl struct {
	device Device
}

func (a *AdvancedRemoteControl) volumeUp() {
	a.device.setVolume(a.device.getVolume() + 10)
}
func (a *AdvancedRemoteControl) volumeDown() {
	a.device.setVolume(a.device.getVolume() - 10)
}
func (a *AdvancedRemoteControl) mute() { //무음
	a.device.setVolume(0)
}

// printer.go: 구현
type Device interface {
	getVolume() int
	setVolume(percent int)
}

// epson.go: 구상 구현
type TV struct {
	volume int
}

func (t *TV) getVolume() int {
	fmt.Println(t.volume)
	return t.volume
}
func (t *TV) setVolume(percent int) {
	fmt.Println("TV 볼륨 변경")
	t.volume = percent
}

type Radio struct {
	volume int
}

func (r *Radio) getVolume() int {
	fmt.Println(r.volume)
	return r.volume
}
func (r *Radio) setVolume(percent int) {
	fmt.Println("Radio 볼륨 설정")
	r.volume = percent
}

func main() {

	tv := &TV{volume: 0}
	remote := &RemoteControl{device: tv}
	remote.volumeUp()
	remote.volumeUp()
	remote.volumeDown()
	fmt.Println("tv 의 현재 볼륨")
	tv.getVolume()
	fmt.Println()

	radio := &Radio{volume: 0}
	advancedRemote := &AdvancedRemoteControl{device: radio}
	advancedRemote.volumeUp()
	advancedRemote.volumeUp()
	advancedRemote.mute()
	fmt.Println("radio 의 현재 볼륨")
	radio.getVolume()
	fmt.Println()
}
