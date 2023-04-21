package abstract_factory

type Vtex struct {
	Url string
}

func (this *Vtex) Backet(response *AbsctractCommerce) {
	this.Url = "url.test.com"
	for index := range response.InputCheckout {
		if index == 0 {
			this.Url += "?sku=" + response.InputCheckout[index].Sku + "&quantity=" + response.InputCheckout[index].Sku
		} else {
			this.Url += "&sku=" + response.InputCheckout[index].Sku + "&quantity=" + response.InputCheckout[index].Sku
		}
	}
	response.ReturnCheckout.Id = "1233213"
	response.ReturnCheckout.Url = this.Url
}

func (this *Vtex) CheckStock(response *AbsctractCommerce) {
	response.ReturnCheckStock.sku = response.InputCheckStock.Sku
	response.ReturnCheckStock.quantity = 1
}
