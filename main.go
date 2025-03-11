package main

import "strings"

type ProductLineItem struct {
	Name string 
	ProcessedAt string
	Email string
	Item string
	PaymentStatus string
	PriceTotal float64
	PriceRefundTotal float64
	PriceTotalOutstanding float64
	TransactionType string
	TransactionProccessedAt string
	TransactionAmount float64
	FullfillmentStatus string
	FullfillmentProcessedAt string
	FullfillmentShipmentStatus string
	RefundAmount float64
}

type TransactionLineItem struct {
	id string
}

type FulfillmentLineItem struct {
	id string
}

func main() {
	//load initial csv file into memory
	data, err := readCSVFile("bigcommerce-orders-export.csv")
	if err!= nil {
		panic(err)
    }
	//parse initial csv into usuable format
	reader, err := parseCSV(data)
    if err!= nil {
        panic(err)
    }
	
	//remove header line
	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	//read records into array
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	//create csv writer to new file
	writer, file, err := createCSVWriter("matrixify-orders-import.csv")
	if err != nil {
        panic(err)
    }
	//defer file close to end of execution
	defer file.Close()

	//for order products
	for _, record := range records {
		//split products into unique line items
		products := strings.Split(record[36], "|")
		if len(products) > 1 {
			for _, product := range products {
				//seperate product values
				newProductLineItem := ProductLineItem{
					Email: record[0],
					Item: product,
				}
				writeCSVRecord(writer, []string{newProductLineItem.Email,newProductLineItem.Item})
			}
		} else {
			newProductLineItem := ProductLineItem{
				Email: record[0],
				Item: record[1],
			}
			writeCSVRecord(writer, []string{newProductLineItem.Email,newProductLineItem.Item})
		}
		//create transaction line
		newTransactionLineItem := TransactionLineItem {
			id: "",
		}
		writeCSVRecord(writer, []string{newTransactionLineItem.id})
		//create fulfillment line
		newFulfillmentLineItem := FulfillmentLineItem {
			id: "",
		}
		writeCSVRecord(writer, []string{newFulfillmentLineItem.id})

	}


	writer.Flush()
	if err := writer.Error(); err != nil {
        panic(err)
    }
}