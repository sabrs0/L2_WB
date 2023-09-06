package main

import "fmt"

type PaymentService struct {
}

func (ps PaymentService) ProcessPayment(sum float64) {
	fmt.Printf("Paying %f $ is in process\n", sum)
}

type DeliveryService struct {
}

func (ds DeliveryService) StartDeliver(address string, days int) {
	fmt.Printf("Items will be dilivered to %s in %d days\n", address, days)
}

type NotificationService struct {
}

func (ns NotificationService) Notify(email, msg string) {
	fmt.Printf("To %s : %s\n", email, msg)
}

type OrderFacade struct {
	paymentService      PaymentService
	deliveryService     DeliveryService
	notificationService NotificationService
}

func NewOrderFacade() OrderFacade {
	return OrderFacade{
		paymentService:      PaymentService{},
		deliveryService:     DeliveryService{},
		notificationService: NotificationService{},
	}
}

func (of OrderFacade) ProcessOrder(sum float64, address string, days int, email, msg string) {
	of.paymentService.ProcessPayment(sum)
	of.deliveryService.StartDeliver(address, days)
	of.notificationService.Notify(email, msg)
}
func main() {
	orderFacade := NewOrderFacade()
	orderFacade.ProcessOrder(100, "Pushkin", 3, "user@gmail.com", "Order successfully processed")
}
