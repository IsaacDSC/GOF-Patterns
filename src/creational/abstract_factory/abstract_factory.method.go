package abstract_factory

type TypesCommerce interface {
	Vtex | Shopify
}

type Commerce[T TypesCommerce] struct {
	Data T
}

type AbsctractCommerce struct {
	CommerceVtex    Commerce[Vtex]
	CommerceShopify Commerce[Shopify]
	InputCheckout   []struct {
		Sku      string
		Quantity int
	}
	ReturnCheckout struct {
		Id  string
		Url string
	}
	InputCheckStock struct {
		Sku string
	}
	ReturnCheckStock struct {
		sku      string
		quantity int
	}
}

func (_this *AbsctractCommerce) GetTypePlatform(enterpriseId string) string {
	return "vtex"
}

func (_this *AbsctractCommerce) Implements(enterpriseId string, action string) string {
	platform := _this.GetTypePlatform(enterpriseId)
	switch platform {
	case "vtex":
		commerce := Commerce[Vtex]{Data: Vtex{}}
		return _this.vtexImplements(action, &commerce)
	case "shopify":
		commerce := Commerce[Shopify]{Data: Shopify{}}
		return _this.shopifyImplements(action, &commerce)
	default:
		return "NOT-FOUND"
	}
}

func (_this *AbsctractCommerce) vtexImplements(action string, commerce *Commerce[Vtex]) string {
	switch action {
	case "checkout":
		commerce.Data.Backet(_this)
		return "SUCCESS"
	case "checkstock":
		commerce.Data.CheckStock(_this)
		return "SUCCESS"
	default:
		return "NOT-FOUND"
	}
}

func (_this *AbsctractCommerce) shopifyImplements(action string, commerce *Commerce[Shopify]) string {
	switch action {
	case "checkout":
		commerce.Data.Backet(_this)
		return "SUCCESS"
	case "checkstock":
		commerce.Data.CheckStock(_this)
		return "SUCCESS"
	default:
		return "NOT-FOUND"
	}
}
