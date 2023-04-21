package abstract_factory

type Shopify struct {
	Url string
}

func (this *Shopify) Backet(response *AbsctractCommerce) {
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

func (this *Shopify) CheckStock(response *AbsctractCommerce) {
	response.ReturnCheckStock.sku = response.InputCheckStock.Sku
	response.ReturnCheckStock.quantity = 1
}
