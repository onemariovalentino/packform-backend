package usecases

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"packform-backend/src/app/orders/models"
	"packform-backend/src/app/orders/queries"
	"packform-backend/src/pkg/helper"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cast"
)

type (
	OrderUsecase struct {
		query queries.OrderQueryInterface
	}
)

func New(query queries.OrderQueryInterface) OrderUsecaseInterface {
	return &OrderUsecase{query: query}
}

func (u *OrderUsecase) GetOrderDetails(ctx context.Context, search string, startDate, endDate time.Time, page, perPage int) ([]*models.OrderDetails, error) {
	return u.query.GetOrderDetails(ctx, search, startDate, endDate, page, perPage)
}

func (u *OrderUsecase) FeedingDataFromCSV(ctx context.Context, destination string, files []string) error {
	wg := sync.WaitGroup{}

	successCh := make(chan string, len(files))
	errCh := make(chan error, len(files))

	numWorkers := 10 // use for distribute insert process

	for _, file := range files {
		wg.Add(1)
		go func(wg *sync.WaitGroup, file string, successCh chan<- string, errCh chan<- error) {
			defer wg.Done()

			csvReader, csvFile, err := helper.ReadCsvFile(file)
			if err != nil {
				errCh <- err
				return
			}
			defer csvFile.Close()

			switch destination {
			case "companies":
				workerChans := make([]chan []*models.Company, numWorkers)
				for i := range workerChans {
					workerChans[i] = make(chan []*models.Company, 1)
				}

				wg.Add(1)
				u.getCustomerCompanies(wg, workerChans, errCh, csvReader)

				for i := 0; i < numWorkers; i++ {
					wg.Add(1)
					go u.createCustomerCompanies(ctx, i, workerChans[i], wg, errCh)
				}

				successCh <- "success to insert companies"

			case "customers":
				workerChans := make([]chan []*models.Customer, numWorkers)
				for i := range workerChans {
					workerChans[i] = make(chan []*models.Customer, 1)
				}

				wg.Add(1)
				u.getCustomers(wg, workerChans, errCh, csvReader)

				for i := 0; i < numWorkers; i++ {
					wg.Add(1)
					go u.createCustomers(ctx, i, workerChans[i], wg, errCh)
				}

				successCh <- "success to insert customers"

			case "orders":
				workerChans := make([]chan []*models.Order, numWorkers)
				for i := range workerChans {
					workerChans[i] = make(chan []*models.Order, 1)
				}

				wg.Add(1)
				u.getOrders(wg, workerChans, errCh, csvReader)

				for i := 0; i < numWorkers; i++ {
					wg.Add(1)
					go u.createOrders(ctx, i, workerChans[i], wg, errCh)
				}

				successCh <- "success to insert orders"

			case "order_items":
				workerChans := make([]chan []*models.OrderItem, numWorkers)
				for i := range workerChans {
					workerChans[i] = make(chan []*models.OrderItem, 1)
				}

				wg.Add(1)
				u.getOrderItems(wg, workerChans, errCh, csvReader)

				for i := 0; i < numWorkers; i++ {
					wg.Add(1)
					go u.createOrderItems(ctx, i, workerChans[i], wg, errCh)
				}

				successCh <- "success to insert order items"

			case "order_item_deliveries":
				workerChans := make([]chan []*models.OrderItemDelivery, numWorkers)
				for i := range workerChans {
					workerChans[i] = make(chan []*models.OrderItemDelivery, 1)
				}

				wg.Add(1)
				u.getOrderItemDeliveries(wg, workerChans, errCh, csvReader)

				for i := 0; i < numWorkers; i++ {
					wg.Add(1)
					go u.createOrderItemDeliveries(ctx, i, workerChans[i], wg, errCh)
				}

				successCh <- "success to insert order item deliveries"

			default:
				errCh <- errors.New(`unknown command`)
				return
			}
		}(&wg, file, successCh, errCh)
	}
	wg.Wait()
	close(successCh)
	for v := range successCh {
		fmt.Println(v)
	}

	close(errCh)
	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *OrderUsecase) getCustomerCompanies(wg *sync.WaitGroup, workerChans []chan []*models.Company, errCh chan<- error, csv *csv.Reader) {
	defer wg.Done()

	_, err := csv.Read()
	if err != nil {
		errCh <- err
		return
	}

	companies := []*models.Company{}
	for {
		row, err := csv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			errCh <- err
			return
		}

		company := &models.Company{}
		company.CompanyID = cast.ToInt(row[0])
		company.CompanyName = row[1]
		companies = append(companies, company)
	}

	chunkSize := (len(companies) + len(workerChans) - 1) / len(workerChans)

	for i, workerChan := range workerChans {
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize
		if endIndex > len(companies) {
			endIndex = len(companies)
		}
		if startIndex > len(companies) {
			break
		}
		workerChan <- companies[startIndex:endIndex]
	}

	for _, ch := range workerChans {
		close(ch)
	}
}

func (u *OrderUsecase) createCustomerCompanies(ctx context.Context, workerID int, inputChan <-chan []*models.Company, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	for {
		dataChunk, ok := <-inputChan
		if !ok {
			break
		}

		if len(dataChunk) > 0 {
			err := u.query.CreateCustomerCompanies(ctx, dataChunk)
			if err != nil {
				errCh <- err
				return
			}
		}
	}
}

func (u *OrderUsecase) getCustomers(wg *sync.WaitGroup, workerChans []chan []*models.Customer, errCh chan<- error, csv *csv.Reader) {
	defer wg.Done()

	_, err := csv.Read()
	if err != nil {
		errCh <- err
		return
	}

	customers := []*models.Customer{}
	for {
		row, err := csv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			errCh <- err
			return
		}

		customer := &models.Customer{}
		customer.UserID = row[0]
		customer.Login = row[1]
		customer.Password = row[2]
		customer.Name = row[3]
		customer.CompanyID = cast.ToInt(row[4])
		customer.CreditCards = row[5]

		customers = append(customers, customer)
	}

	chunkSize := (len(customers) + len(workerChans) - 1) / len(workerChans)

	for i, workerChan := range workerChans {
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize
		if endIndex > len(customers) {
			endIndex = len(customers)
		}
		if startIndex > len(customers) {
			break
		}
		workerChan <- customers[startIndex:endIndex]
	}

	for _, ch := range workerChans {
		close(ch)
	}
}

func (u *OrderUsecase) createCustomers(ctx context.Context, workerID int, inputChan <-chan []*models.Customer, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	for {
		dataChunk, ok := <-inputChan
		if !ok {
			break
		}

		if len(dataChunk) > 0 {
			err := u.query.CreateCustomers(ctx, dataChunk)
			if err != nil {
				errCh <- err
				return
			}
		}
	}
}

func (u *OrderUsecase) getOrders(wg *sync.WaitGroup, workerChans []chan []*models.Order, errCh chan<- error, csv *csv.Reader) {
	defer wg.Done()

	_, err := csv.Read()
	if err != nil {
		errCh <- err
		return
	}

	orders := []*models.Order{}
	for {
		row, err := csv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			errCh <- err
			return
		}

		tz, err := time.LoadLocation("Australia/Melbourne")
		if err != nil {
			errCh <- err
			return
		}
		createdAt, err := time.ParseInLocation("2006-01-02T15:04:05Z", row[1], tz)
		if err != nil {
			errCh <- err
			return
		}

		order := &models.Order{}
		order.ID = cast.ToInt(row[0])
		order.CreatedAt = createdAt.In(time.UTC)
		order.OrderName = row[2]
		order.CustomerID = row[3]

		orders = append(orders, order)
	}

	chunkSize := (len(orders) + len(workerChans) - 1) / len(workerChans)

	for i, workerChan := range workerChans {
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize
		if endIndex > len(orders) {
			endIndex = len(orders)
		}
		if startIndex > len(orders) {
			break
		}
		workerChan <- orders[startIndex:endIndex]
	}

	for _, ch := range workerChans {
		close(ch)
	}
}

func (u *OrderUsecase) createOrders(ctx context.Context, workerID int, inputChan <-chan []*models.Order, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	for {
		dataChunk, ok := <-inputChan
		if !ok {
			break
		}

		if len(dataChunk) > 0 {
			err := u.query.CreateOrders(ctx, dataChunk)
			if err != nil {
				errCh <- err
				return
			}
		}
	}
}

func (u *OrderUsecase) getOrderItems(wg *sync.WaitGroup, workerChans []chan []*models.OrderItem, errCh chan<- error, csv *csv.Reader) {
	defer wg.Done()

	_, err := csv.Read()
	if err != nil {
		errCh <- err
		return
	}

	orderItems := []*models.OrderItem{}
	for {
		row, err := csv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			errCh <- err
			return
		}

		orderItem := &models.OrderItem{}
		orderItem.ID = cast.ToInt(row[0])
		orderItem.OrderID = cast.ToInt(row[1])
		orderItem.PricePerUnit = cast.ToFloat64(row[2])
		orderItem.Qty = cast.ToInt(row[3])
		orderItem.Product = strings.Title(strings.ToLower(row[4]))

		orderItems = append(orderItems, orderItem)
	}

	chunkSize := (len(orderItems) + len(workerChans) - 1) / len(workerChans)

	for i, workerChan := range workerChans {
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize
		if endIndex > len(orderItems) {
			endIndex = len(orderItems)
		}
		if startIndex > len(orderItems) {
			break
		}
		workerChan <- orderItems[startIndex:endIndex]
	}

	for _, ch := range workerChans {
		close(ch)
	}
}

func (u *OrderUsecase) createOrderItems(ctx context.Context, workerID int, inputChan <-chan []*models.OrderItem, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	for {
		dataChunk, ok := <-inputChan
		if !ok {
			break
		}

		if len(dataChunk) > 0 {
			err := u.query.CreateOrderItems(ctx, dataChunk)
			if err != nil {
				errCh <- err
				return
			}
		}
	}
}

func (u *OrderUsecase) getOrderItemDeliveries(wg *sync.WaitGroup, workerChans []chan []*models.OrderItemDelivery, errCh chan<- error, csv *csv.Reader) {
	defer wg.Done()

	_, err := csv.Read()
	if err != nil {
		errCh <- err
		return
	}

	orderItemDeliveries := []*models.OrderItemDelivery{}
	for {
		row, err := csv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			errCh <- err
			return
		}

		orderItemDelivery := &models.OrderItemDelivery{}
		orderItemDelivery.ID = cast.ToInt(row[0])
		orderItemDelivery.OrderItemID = cast.ToInt(row[1])
		orderItemDelivery.DeliveredQty = cast.ToInt(row[2])

		orderItemDeliveries = append(orderItemDeliveries, orderItemDelivery)
	}

	chunkSize := (len(orderItemDeliveries) + len(workerChans) - 1) / len(workerChans)

	for i, workerChan := range workerChans {
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize
		if endIndex > len(orderItemDeliveries) {
			endIndex = len(orderItemDeliveries)
		}
		if startIndex > len(orderItemDeliveries) {
			break
		}
		workerChan <- orderItemDeliveries[startIndex:endIndex]
	}

	for _, ch := range workerChans {
		close(ch)
	}
}

func (u *OrderUsecase) createOrderItemDeliveries(ctx context.Context, workerID int, inputChan <-chan []*models.OrderItemDelivery, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	for {
		dataChunk, ok := <-inputChan
		if !ok {
			break
		}

		if len(dataChunk) > 0 {
			err := u.query.CreateOrderItemDeliveries(ctx, dataChunk)
			if err != nil {
				errCh <- err
				return
			}
		}
	}
}
