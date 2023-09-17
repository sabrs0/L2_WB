package pattern

import "fmt"

type Route [][2]float64
type Drone struct {
	flyStrategy FlyStrategy
}

func (drone Drone) Fly() {
	drone.flyStrategy.Fly()
}

type FlyStrategy interface {
	Fly()
}

type FlyByFeelStrategy struct {
}

func (strategy FlyByFeelStrategy) Fly() {
	fmt.Println("....Flying by feel....")
}

type FlyByRoute struct {
}

func (strategy FlyByRoute) Fly() {
	fmt.Println("<<<<Flying By Known Route>>>>")
}

type DroneManager struct {
	drone    Drone
	places   map[string]Route
	curRoute Route
}

func (manager *DroneManager) SetDroneStrategy(place string) {
	curRoute, isKnown := manager.places[place]
	if !isKnown {
		manager.drone.flyStrategy = FlyByFeelStrategy{}
	} else {
		manager.drone.flyStrategy = FlyByRoute{}
		manager.curRoute = curRoute
	}
}

func strategyPattern() {
	places := []string{"Pushkin Street", "Lermontov Street", "Tolstoi Street"}
	routePlaces := make(map[string]Route)
	routePlaces["Pushkin Street"] = Route{{55.5555555, 66.6666666}, {56.0, 66.6666665}, {55.656565656, 66.777777}}
	droneManager := DroneManager{
		drone:  Drone{},
		places: routePlaces,
	}
	for _, v := range places {
		droneManager.SetDroneStrategy(v)
		droneManager.drone.Fly()
	}
}
