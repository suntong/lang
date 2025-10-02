// the Template Method pattern in Go for processing CSV files with variations

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

// ============ INTERFACE DEFINITIONS ============

// CSVProcessor defines what every CSV processor must implement
type CSVProcessor interface {
	// Template method steps that must be implemented by concrete types
	ValidateHeaders(headers []string) error
	ParseRecord(record []string, headers []string) (Record, error)
	ProcessRecord(record Record) (Record, error)
	OutputRecord(record Record) string
	GetRequiredHeaders() []string
	GetFileType() string
}

// Record interface for all record types
type Record interface {
	GetID() string
	GetTimestamp() time.Time
	Validate() error
	String() string
}

// ============ BASE PROCESSOR WITH TEMPLATE METHOD ============

// BaseCSVProcessor provides the template method and common functionality
type BaseCSVProcessor struct {
	processor      CSVProcessor // Reference to concrete implementation
	records        []Record
	processedCount int
	errorCount     int
	startTime      time.Time
}

// Process is the TEMPLATE METHOD - defines the algorithm once
// This is the key pattern: one process definition for all variations
func (b *BaseCSVProcessor) Process(csvData string) error {
	fmt.Printf("\n📁 Processing %s File\n", b.processor.GetFileType())
	fmt.Println("=" + strings.Repeat("=", 50))

	b.startTime = time.Now()

	// Step 1: Read CSV
	reader := csv.NewReader(strings.NewReader(csvData))
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read headers: %w", err)
	}

	// Step 2: Validate headers
	fmt.Println("✓ Step 1: Validating headers...")
	if err := b.processor.ValidateHeaders(headers); err != nil {
		return fmt.Errorf("header validation failed: %w", err)
	}

	// Step 3: Parse records
	fmt.Println("✓ Step 2: Parsing records...")
	var rawRecords []Record
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading row: %v", err)
			b.errorCount++
			continue
		}

		record, err := b.processor.ParseRecord(row, headers)
		if err != nil {
			log.Printf("Error parsing record: %v", err)
			b.errorCount++
			continue
		}

		if err := record.Validate(); err != nil {
			log.Printf("Validation failed: %v", err)
			b.errorCount++
			continue
		}

		rawRecords = append(rawRecords, record)
	}

	// Step 4: Process records (business logic)
	fmt.Println("✓ Step 3: Processing records...")
	for _, record := range rawRecords {
		processed, err := b.processor.ProcessRecord(record)
		if err != nil {
			log.Printf("Error processing record %s: %v", record.GetID(), err)
			b.errorCount++
			continue
		}
		b.records = append(b.records, processed)
		b.processedCount++
	}

	// Step 5: Output results
	fmt.Println("✓ Step 4: Generating output...")
	b.outputResults()

	// Step 6: Generate summary
	b.printSummary()

	return nil
}

func (b *BaseCSVProcessor) outputResults() {
	fmt.Println("\n📊 Processed Records:")
	for i, record := range b.records {
		if i < 3 || i >= len(b.records)-1 { // Show first 3 and last record
			fmt.Println(b.processor.OutputRecord(record))
		} else if i == 3 {
			fmt.Println("   ... (more records) ...")
		}
	}
}

func (b *BaseCSVProcessor) printSummary() {
	duration := time.Since(b.startTime)
	fmt.Printf("\n📈 Processing Summary:\n")
	fmt.Printf("   Total Processed: %d\n", b.processedCount)
	fmt.Printf("   Errors: %d\n", b.errorCount)
	fmt.Printf("   Duration: %v\n", duration)
}

// ============ BASE RECORD WITH COMMON FIELDS ============

// BaseRecord contains fields common to all CSV types
type BaseRecord struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"`
}

func (b BaseRecord) GetID() string {
	return b.ID
}

func (b BaseRecord) GetTimestamp() time.Time {
	return b.Timestamp
}

func (b BaseRecord) String() string {
	return fmt.Sprintf("ID: %s, Time: %s, Status: %s",
		b.ID, b.Timestamp.Format("2006-01-02"), b.Status)
}

// ============ CUSTOMER CSV IMPLEMENTATION ============

type CustomerRecord struct {
	BaseRecord
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	LoyaltyTier string  `json:"loyalty_tier"`
	TotalSpent  float64 `json:"total_spent"`
}

func (c CustomerRecord) Validate() error {
	if c.Email == "" {
		return fmt.Errorf("email is required")
	}
	if c.TotalSpent < 0 {
		return fmt.Errorf("total spent cannot be negative")
	}
	return nil
}

type CustomerCSVProcessor struct {
	BaseCSVProcessor
}

func NewCustomerProcessor() *CustomerCSVProcessor {
	p := &CustomerCSVProcessor{}
	p.BaseCSVProcessor.processor = p
	return p
}

func (c *CustomerCSVProcessor) GetFileType() string {
	return "Customer CSV"
}

func (c *CustomerCSVProcessor) GetRequiredHeaders() []string {
	return []string{"id", "timestamp", "status", "name", "email", "loyalty_tier", "total_spent"}
}

func (c *CustomerCSVProcessor) ValidateHeaders(headers []string) error {
	required := c.GetRequiredHeaders()
	if len(headers) < len(required) {
		return fmt.Errorf("expected at least %d columns, got %d", len(required), len(headers))
	}
	return nil
}

func (c *CustomerCSVProcessor) ParseRecord(row []string, headers []string) (Record, error) {
	if len(row) < 7 {
		return nil, fmt.Errorf("insufficient fields")
	}

	timestamp, _ := time.Parse("2006-01-02", row[1])
	totalSpent := 0.0
	fmt.Sscanf(row[6], "%f", &totalSpent)

	return &CustomerRecord{
		BaseRecord: BaseRecord{
			ID:        row[0],
			Timestamp: timestamp,
			Status:    row[2],
		},
		Name:        row[3],
		Email:       row[4],
		LoyaltyTier: row[5],
		TotalSpent:  totalSpent,
	}, nil
}

func (c *CustomerCSVProcessor) ProcessRecord(record Record) (Record, error) {
	customer := record.(*CustomerRecord)

	// Business logic: Upgrade loyalty tier based on spending
	if customer.TotalSpent > 10000 {
		customer.LoyaltyTier = "PLATINUM"
	} else if customer.TotalSpent > 5000 {
		customer.LoyaltyTier = "GOLD"
	} else if customer.TotalSpent > 1000 {
		customer.LoyaltyTier = "SILVER"
	}

	// Update status
	if customer.Status == "pending" {
		customer.Status = "processed"
	}

	return customer, nil
}

func (c *CustomerCSVProcessor) OutputRecord(record Record) string {
	customer := record.(*CustomerRecord)
	return fmt.Sprintf("   [CUSTOMER] %s | %s | Tier: %s | Spent: $%.2f",
		customer.Name, customer.Email, customer.LoyaltyTier, customer.TotalSpent)
}

// ============ PRODUCT CSV IMPLEMENTATION ============

type ProductRecord struct {
	BaseRecord
	SKU          string  `json:"sku"`
	Name         string  `json:"name"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	WarningLevel int     `json:"warning_level"`
}

func (p ProductRecord) Validate() error {
	if p.SKU == "" {
		return fmt.Errorf("SKU is required")
	}
	if p.Price < 0 {
		return fmt.Errorf("price cannot be negative")
	}
	if p.Stock < 0 {
		return fmt.Errorf("stock cannot be negative")
	}
	return nil
}

type ProductCSVProcessor struct {
	BaseCSVProcessor
}

func NewProductProcessor() *ProductCSVProcessor {
	p := &ProductCSVProcessor{}
	p.BaseCSVProcessor.processor = p
	return p
}

func (p *ProductCSVProcessor) GetFileType() string {
	return "Product Inventory CSV"
}

func (p *ProductCSVProcessor) GetRequiredHeaders() []string {
	return []string{"id", "timestamp", "status", "sku", "name", "category", "price", "stock"}
}

func (p *ProductCSVProcessor) ValidateHeaders(headers []string) error {
	required := p.GetRequiredHeaders()
	if len(headers) < len(required) {
		return fmt.Errorf("expected at least %d columns, got %d", len(required), len(headers))
	}
	return nil
}

func (p *ProductCSVProcessor) ParseRecord(row []string, headers []string) (Record, error) {
	if len(row) < 8 {
		return nil, fmt.Errorf("insufficient fields")
	}

	timestamp, _ := time.Parse("2006-01-02", row[1])
	price := 0.0
	stock := 0
	fmt.Sscanf(row[6], "%f", &price)
	fmt.Sscanf(row[7], "%d", &stock)

	return &ProductRecord{
		BaseRecord: BaseRecord{
			ID:        row[0],
			Timestamp: timestamp,
			Status:    row[2],
		},
		SKU:          row[3],
		Name:         row[4],
		Category:     row[5],
		Price:        price,
		Stock:        stock,
		WarningLevel: 10, // Default warning level
	}, nil
}

func (p *ProductCSVProcessor) ProcessRecord(record Record) (Record, error) {
	product := record.(*ProductRecord)

	// Business logic: Set status based on stock levels
	if product.Stock == 0 {
		product.Status = "out_of_stock"
	} else if product.Stock < product.WarningLevel {
		product.Status = "low_stock"
	} else {
		product.Status = "in_stock"
	}

	// Apply category-based pricing rules
	if product.Category == "premium" {
		product.Price = product.Price * 1.2 // 20% markup
	}

	return product, nil
}

func (p *ProductCSVProcessor) OutputRecord(record Record) string {
	product := record.(*ProductRecord)
	stockStatus := "✓"
	if product.Status == "out_of_stock" {
		stockStatus = "✗"
	} else if product.Status == "low_stock" {
		stockStatus = "⚠"
	}

	return fmt.Sprintf("   [PRODUCT] %s %s | %s | $%.2f | Stock: %d %s",
		product.SKU, product.Name, product.Category, product.Price, product.Stock, stockStatus)
}

// ============ ORDER CSV IMPLEMENTATION ============

type OrderRecord struct {
	BaseRecord
	OrderNumber  string    `json:"order_number"`
	CustomerID   string    `json:"customer_id"`
	TotalAmount  float64   `json:"total_amount"`
	PaymentType  string    `json:"payment_type"`
	ShippingDate time.Time `json:"shipping_date"`
	Priority     string    `json:"priority"`
}

func (o OrderRecord) Validate() error {
	if o.OrderNumber == "" {
		return fmt.Errorf("order number is required")
	}
	if o.TotalAmount <= 0 {
		return fmt.Errorf("total amount must be positive")
	}
	return nil
}

type OrderCSVProcessor struct {
	BaseCSVProcessor
}

func NewOrderProcessor() *OrderCSVProcessor {
	p := &OrderCSVProcessor{}
	p.BaseCSVProcessor.processor = p
	return p
}

func (o *OrderCSVProcessor) GetFileType() string {
	return "Order Processing CSV"
}

func (o *OrderCSVProcessor) GetRequiredHeaders() []string {
	return []string{"id", "timestamp", "status", "order_number", "customer_id", "total_amount", "payment_type"}
}

func (o *OrderCSVProcessor) ValidateHeaders(headers []string) error {
	required := o.GetRequiredHeaders()
	if len(headers) < len(required) {
		return fmt.Errorf("expected at least %d columns, got %d", len(required), len(headers))
	}
	return nil
}

func (o *OrderCSVProcessor) ParseRecord(row []string, headers []string) (Record, error) {
	if len(row) < 7 {
		return nil, fmt.Errorf("insufficient fields")
	}

	timestamp, _ := time.Parse("2006-01-02", row[1])
	totalAmount := 0.0
	fmt.Sscanf(row[5], "%f", &totalAmount)

	return &OrderRecord{
		BaseRecord: BaseRecord{
			ID:        row[0],
			Timestamp: timestamp,
			Status:    row[2],
		},
		OrderNumber: row[3],
		CustomerID:  row[4],
		TotalAmount: totalAmount,
		PaymentType: row[6],
		Priority:    "normal",
	}, nil
}

func (o *OrderCSVProcessor) ProcessRecord(record Record) (Record, error) {
	order := record.(*OrderRecord)

	// Business logic: Set priority based on amount
	if order.TotalAmount > 1000 {
		order.Priority = "high"
	} else if order.TotalAmount > 500 {
		order.Priority = "medium"
	}

	// Set shipping date based on priority
	baseDate := order.Timestamp
	switch order.Priority {
	case "high":
		order.ShippingDate = baseDate.AddDate(0, 0, 1) // Next day
	case "medium":
		order.ShippingDate = baseDate.AddDate(0, 0, 3) // 3 days
	default:
		order.ShippingDate = baseDate.AddDate(0, 0, 5) // 5 days
	}

	// Update status
	if order.Status == "pending" {
		order.Status = "confirmed"
	}

	return order, nil
}

func (o *OrderCSVProcessor) OutputRecord(record Record) string {
	order := record.(*OrderRecord)
	priorityIcon := ""
	switch order.Priority {
	case "high":
		priorityIcon = "🔴"
	case "medium":
		priorityIcon = "🟡"
	default:
		priorityIcon = "🟢"
	}

	return fmt.Sprintf("   [ORDER] #%s %s | Customer: %s | $%.2f | Ships: %s",
		order.OrderNumber, priorityIcon, order.CustomerID, order.TotalAmount,
		order.ShippingDate.Format("Jan 02"))
}

// ============ MAIN DEMONSTRATION ============

func main() {
	fmt.Println("🏭 CSV Processing Factory - Template Method Pattern")
	fmt.Println("=" + strings.Repeat("=", 60))
	fmt.Println("Demonstrating: One process definition, multiple behaviors")

	// Sample CSV data for each type
	customerCSV := `id,timestamp,status,name,email,loyalty_tier,total_spent
C001,2024-01-15,pending,John Doe,john@email.com,SILVER,2500.00
C002,2024-01-16,pending,Jane Smith,jane@email.com,GOLD,7500.00
C003,2024-01-17,pending,Bob Wilson,bob@email.com,BRONZE,500.00
C004,2024-01-18,pending,Alice Brown,alice@email.com,SILVER,12000.00`

	productCSV := `id,timestamp,status,sku,name,category,price,stock
P001,2024-01-15,active,SKU-001,Laptop Pro,premium,999.99,5
P002,2024-01-16,active,SKU-002,Mouse Wireless,standard,29.99,2
P003,2024-01-17,active,SKU-003,Keyboard Mechanical,premium,149.99,0
P004,2024-01-18,active,SKU-004,Monitor 4K,premium,599.99,15`

	orderCSV := `id,timestamp,status,order_number,customer_id,total_amount,payment_type
O001,2024-01-15,pending,ORD-2024-001,C001,250.00,credit_card
O002,2024-01-16,pending,ORD-2024-002,C002,750.00,paypal
O003,2024-01-17,pending,ORD-2024-003,C003,1500.00,credit_card
O004,2024-01-18,pending,ORD-2024-004,C004,50.00,debit_card`

	// Process different CSV types using the SAME process flow
	processors := []struct {
		processor CSVProcessor
		data      string
	}{
		{NewCustomerProcessor(), customerCSV},
		{NewProductProcessor(), productCSV},
		{NewOrderProcessor(), orderCSV},
	}

	for _, p := range processors {
		// Each processor follows the SAME process defined in BaseCSVProcessor
		// but with different behavior for each step
		if baseProcessor, ok := p.processor.(*CustomerCSVProcessor); ok {
			baseProcessor.Process(p.data)
		} else if baseProcessor, ok := p.processor.(*ProductCSVProcessor); ok {
			baseProcessor.Process(p.data)
		} else if baseProcessor, ok := p.processor.(*OrderCSVProcessor); ok {
			baseProcessor.Process(p.data)
		}
	}

	// Demonstrate polymorphic processor selection
	fmt.Println("\n🔄 Dynamic Processor Selection")
	fmt.Println("=" + strings.Repeat("=", 50))

	// Factory function for dynamic processor selection
	getProcessor := func(csvType string) CSVProcessor {
		switch csvType {
		case "customer":
			return NewCustomerProcessor()
		case "product":
			return NewProductProcessor()
		case "order":
			return NewOrderProcessor()
		default:
			return nil
		}
	}

	// Simulate dynamic file processing
	fileTypes := []string{"customer", "product", "order"}
	for _, fileType := range fileTypes {
		processor := getProcessor(fileType)
		if processor != nil {
			fmt.Printf("✓ Processor ready for %s files\n", fileType)
		}
	}

	fmt.Println("\n✅ All CSV files processed using the same base process!")
}

/*

## Key Design Patterns Demonstrated:

### 1. **Template Method Pattern**

The `Process()` method in `BaseCSVProcessor` defines the algorithm once:

```go
1. Read CSV
2. Validate headers
3. Parse records
4. Process records (business logic)
5. Output results
6. Generate summary
```

### 2. **Composition & Embedding**

- `BaseRecord` is embedded in all record types
- `BaseCSVProcessor` is embedded in all processor types

### 3. **Polymorphism Through Interfaces**

- All processors implement `CSVProcessor` interface
- All records implement `Record` interface

### 4. **Hook Methods**

Each concrete processor implements:

- `ValidateHeaders()` - Different validation rules
- `ParseRecord()` - Different field parsing
- `ProcessRecord()` - Different business logic
- `OutputRecord()` - Different output formats

## Benefits of This Design:

1. **DRY (Don't Repeat Yourself)**
   - Process flow defined once
   - Common fields handled in base types

2. **Open/Closed Principle**
   - Easy to add new CSV types without modifying existing code
   - Just implement the interface

3. **Single Responsibility**
   - Each processor handles one CSV type
   - Base processor handles the workflow

4. **Flexibility**
   - Each CSV type can have unique fields
   - Business logic is customizable per type
   - Output format is specific to each type

5. **Maintainability**
   - Changes to the process flow happen in one place
   - Type-specific logic is isolated

This pattern is perfect for scenarios where you have:

- Similar data structures with variations
- A common processing pipeline
- Different business rules for each variation
- Need for extensibility without modifying core logic

*/
