package yagni

type ProductId string
type Product struct{}
type CountryEligibilityId string
type CountryEligibility struct{}

type FulfillmentPolicy int32

const (
	// Ship only if all items are in inventory
	FulfillmentPolicy_ONE_SHIPMENT_IMMEDIATELY = 0
	// Ship only if all items are in inventory or all items can be ship at a later time
	FulfillmentPolicy_ONE_SHIPMENT_QUEUED = 1
	// Ship if all items will become available, ship partial deliveries as they are available
	FulfillmentPolicy_MULTIPLE_SHIPMENTS_WHEN_AVAILABLE = 2
	// Ship if all items will become available, ship partial deliveries as they are available, accept multiple countries
	FulfillmentPolicy_MULTIPLE_SHIPMENTS_WHEN_AVAILABLE_ANY_COUNTRY = 3
	// Ship everything possible to ship as well as reserve incoming
	FulfillmentPolicy_BEST_EFFORT = 4
	// Ship everything possible to ship as well as reserve incoming from any country
	FulfillmentPolicy_BEST_EFFORT_MULTIPLE_COUNTRIES = 5
)

type ProductFulfillmentPolicy struct {
	EligibleCountries []*CountryEligibilityId
	Policy            FulfillmentPolicy
}
type ProductInventory struct{}
type InventoryReservation struct {
	Reserved       []*ProductInventoryReservation
	FutureReserved []*ProductInventoryReservation
	NotReserved    []*ProductInventoryReservation
}
type ProductInventoryReservation struct{}
type FulfillmentStatus struct{}

type Order struct {
	OrderLines []*OrderLine
}
type OrderLine struct{}

type ProductCatalogService interface {
	AddProduct(*Product) (ProductId, error)
	RemoveProduct(*Product) error
	UpdateProduct(*Product) error
	GetProduct(ProductId) (*Product, error)
	GetProductList() ([]*Product, error)
	GetProductListForCountry(*CountryEligibility) ([]*Product, error)

	GetCountryList() ([]*CountryEligibility, error)
	SetCountryEligible(ProductId, *CountryEligibility) error
	RemoveCountryEligible(ProductId, *CountryEligibility) error
}

type InventoryService interface {
	GetProductInventory(ProductId, CountryEligibilityId) (*ProductInventory, error)
	GetProductInventories(ProductId) (map[CountryEligibilityId]*ProductInventory, error)

	CheckProductInventory(*Order, *ProductFulfillmentPolicy) (*InventoryReservation, error)
	ReserveProductInventory(*Order, *ProductFulfillmentPolicy) (*InventoryReservation, error)
	CancelReservation(*ProductInventoryReservation) error
	FulfillReservation(*ProductInventoryReservation) (*FulfillmentStatus, error)
}

type OrderService interface {
	// ...
}

type PaymentService interface {
	// ...
}

type FulfillmentService interface {
	// ...
}
