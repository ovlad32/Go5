package types


type PlayerType struct {
	Name string
}
type PlayerListType []PlayerType


type BalanceType struct {
	Amount float64
	Currency CurrencyType
}
type Account struct {
	BalanceType
	Owner *PlayerType
	Name string
}

type CurrencyType struct {
	ISO string
	Name string
}
type CurrencyListType []CurrencyType

type CountryType struct {
	Name   string
	Currency CurrencyType
}
type CountryListType []CountryType

type CompanyType struct {
	Name string
	Industry IndustryType
	Country CountryType
}
type CompanyListType []CompanyType

type IndustryType struct{
	Name string
}
type IndustyListType []IndustryType


func Players() PlayerListType {

}