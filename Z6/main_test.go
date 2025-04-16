package main

import (
	"log"
	"net/url"
	"strings"
	"testing"

	"github.com/playwright-community/playwright-go"
)

func assertNotNil(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func beforeAll() playwright.BrowserContext {
	pw, err := playwright.Run()
	assertNotNil("could not start playright", err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertNotNil("could not launch Chromium: %w", err)

	context, err := browser.NewContext()
	assertNotNil("could not create context: %w", err)

	return context
}

var context = beforeAll()

func TestPageLoad(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)
}

func TestQueryGenerateCartId(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	url := page.URL()
	if strings.Contains(url, "?cartId=") == false {
		t.Errorf("cartId not found in URL")
	}
}

func TestProductsList(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	locator := page.Locator("[data-testid=\"product\"]")
	products, err := locator.All()

	if err != nil {
		t.Errorf("could not get products")
	}

	if len(products) == 0 {
		t.Errorf("no products found")
	}
}

func TestProductHasBuyButton(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	locator := page.Locator("[data-testid=\"product\"]")
	products, err := locator.All()

	if err != nil {
		t.Errorf("could not get products")
	}

	for _, product := range products {
		buttonLocator := product.Locator("[data-testid=\"buy-button\"]")
		button := buttonLocator.First()

		if button == nil {
			t.Errorf("buy button not found")
		}

		text, err := button.TextContent()

		if err != nil {
			t.Errorf("could not get button text")
		}

		if strings.Contains(text, "$") == false {
			t.Errorf("buy button does not contain dollar")
		}
	}
}

func TestProductBuyClickable(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	locator := page.Locator("[data-testid=\"product\"]")
	products, err := locator.All()

	if err != nil {
		t.Errorf("could not get products")
	}

	for _, product := range products {
		buttonLocator := product.Locator("[data-testid=\"buy-button\"]")
		button := buttonLocator.First()

		if button == nil {
			t.Errorf("buy button not found")
		}

		isDisabled, err := button.IsDisabled()

		if err != nil {
			t.Errorf("could not check button disabled")
		}

		if isDisabled {
			t.Errorf("buy button is disabled")
		}

		isVisible, err := button.IsVisible()

		if err != nil {
			t.Errorf("could not check button visible")
		}

		if !isVisible {
			t.Errorf("buy button is not visible")
		}
	}
}

func TestCartNotClickable(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	cartLocator := page.Locator("[data-testid=\"cart-link\"]")
	cart := cartLocator.First()

	if cart == nil {
		t.Errorf("cart link not found")
	}

	isVisible, err := cart.IsVisible()

	if err != nil {
		t.Errorf("could not check cart visible")
	}

	if isVisible {
		t.Errorf("cart is visible")
	}
}

func TestCartVisible(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	buttonLocator := page.Locator("[data-testid=\"buy-button\"]")
	button := buttonLocator.First()

	if button == nil {
		t.Errorf("buy button not found")
	}

	button.Click()

	page.WaitForTimeout(250)

	cartLocator := page.Locator("[data-testid=\"cart-link\"]")
	cart := cartLocator.First()

	if cart == nil {
		t.Errorf("cart link not found")
	}

	isVisible, err := cart.IsVisible()

	if err != nil {
		t.Errorf("could not check cart visible")
	}

	if !isVisible {
		t.Errorf("cart is not visible")
	}
}

func TestCartClickNavigation(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	buttonLocator := page.Locator("[data-testid=\"buy-button\"]")
	button := buttonLocator.First()

	if button == nil {
		t.Errorf("buy button not found")
	}

	button.Click()

	page.WaitForTimeout(250)

	cartLocator := page.Locator("[data-testid=\"cart-link\"]")
	cart := cartLocator.First()

	if cart == nil {
		t.Errorf("cart link not found")
	}

	cart.Click()

	page.WaitForTimeout(250)

	url := page.URL()
	subs := "/cart?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("could not wait for cart page, %s", page.URL())
	}
}

func TestNavigateHomeWhenEmptyCart(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	cartUrl, err := url.JoinPath(page.URL(), "/cart")
	assertNotNil("could not join cart url", err)

	page.Goto(cartUrl)

	if strings.Contains(page.URL(), "/cart") {
		t.Errorf("did not navigate home, %s", page.URL())
	}
}

func TestNavigateHomeWhenEmptyCheckout(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	cartUrl, err := url.JoinPath(page.URL(), "/checkout")
	assertNotNil("could not join cart url", err)

	page.Goto(cartUrl)

	if strings.Contains(page.URL(), "/checkout") {
		t.Errorf("did not navigate home, %s", page.URL())
	}
}

func TestCartProductsVisible(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	buttonLocator := page.Locator("[data-testid=\"buy-button\"]")
	button := buttonLocator.First()

	if button == nil {
		t.Errorf("buy button not found")
	}

	button.Click()

	page.WaitForTimeout(250)

	cartLocator := page.Locator("[data-testid=\"cart-link\"]")
	cart := cartLocator.First()

	if cart == nil {
		t.Errorf("cart link not found")
	}

	cart.Click()

	page.WaitForTimeout(250)

	url := page.URL()
	subs := "/cart?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("could not wait for cart page, %s", page.URL())
	}

	productsLocator := page.Locator("[data-testid=\"product\"]")
	products, err := productsLocator.All()

	if err != nil {
		t.Errorf("could not get products: %s", err)
	}

	if len(products) == 0 {
		t.Errorf("no products found")
	}
}

func TestCartCheckoutClickable(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	buttonLocator := page.Locator("[data-testid=\"buy-button\"]")
	button := buttonLocator.First()

	if button == nil {
		t.Errorf("buy button not found")
	}

	button.Click()

	page.WaitForTimeout(250)

	cartLocator := page.Locator("[data-testid=\"cart-link\"]")
	cart := cartLocator.First()

	if cart == nil {
		t.Errorf("cart link not found")
	}

	cart.Click()

	page.WaitForTimeout(250)

	url := page.URL()
	subs := "/cart?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("could not wait for cart page, %s", page.URL())
	}

	productsLocator := page.Locator("[data-testid=\"product\"]")
	products, err := productsLocator.All()

	if err != nil {
		t.Errorf("could not get products: %s", err)
	}

	if len(products) == 0 {
		t.Errorf("no products found")
	}

	checkoutLocator := page.Locator("text=Checkout")
	checkout := checkoutLocator.First()

	if checkout == nil {
		t.Errorf("checkout button not found")
	}

	checkout.Click()

	page.WaitForTimeout(250)

	url = page.URL()
	subs = "/checkout?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("could not wait for checkout page, %s", page.URL())
	}
}

func getToCheckout(t *testing.T, page playwright.Page) {
	buttonLocator := page.Locator("[data-testid=\"buy-button\"]")
	button := buttonLocator.First()

	if button == nil {
		t.Errorf("buy button not found")
	}

	button.Click()

	page.WaitForTimeout(150)

	cartLocator := page.Locator("[data-testid=\"cart-link\"]")
	cart := cartLocator.First()

	if cart == nil {
		t.Errorf("cart link not found")
	}

	cart.Click()

	page.WaitForTimeout(150)

	url := page.URL()
	subs := "/cart?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("could not wait for cart page, %s", page.URL())
	}

	productsLocator := page.Locator("[data-testid=\"product\"]")
	products, err := productsLocator.All()

	if err != nil {
		t.Errorf("could not get products: %s", err)
	}

	if len(products) == 0 {
		t.Errorf("no products found")
	}

	checkoutLocator := page.Locator("text=Checkout")
	checkout := checkoutLocator.First()

	if checkout == nil {
		t.Errorf("checkout button not found")
	}

	checkout.Click()

	page.WaitForTimeout(250)

	url = page.URL()
	subs = "/checkout?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("could not wait for checkout page, %s", page.URL())
	}
}

func TestOnlyNameCheckout(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	getToCheckout(t, page)

	err = page.Locator("[name=\"name\"]").Fill("John Doe")
	assertNotNil("could not fill name input: %w", err)

	confirmLocator := page.Locator("text=Confirm")
	confirm := confirmLocator.First()

	if confirm == nil {
		t.Errorf("confirm button not found")
	}

	confirm.Click()

	page.WaitForTimeout(250)

	url := page.URL()
	subs := "/checkout?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("did not stay on checkout page, %s", page.URL())
	}
}

func TestOnlyEmailCheckout(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	getToCheckout(t, page)

	err = page.Locator("[name=\"email\"]").Fill("john.doe@example.com")
	assertNotNil("could not fill email input: %w", err)

	confirmLocator := page.Locator("text=Confirm")
	confirm := confirmLocator.First()

	if confirm == nil {
		t.Errorf("confirm button not found")
	}

	confirm.Click()

	page.WaitForTimeout(250)

	url := page.URL()
	subs := "/checkout?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("did not stay on checkout page, %s", page.URL())
	}
}

func TestOnlyPhoneCheckout(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	getToCheckout(t, page)

	err = page.Locator("[name=\"phone\"]").Fill("1234567890")
	assertNotNil("could not fill phone input: %w", err)

	confirmLocator := page.Locator("text=Confirm")
	confirm := confirmLocator.First()

	if confirm == nil {
		t.Errorf("confirm button not found")
	}

	confirm.Click()

	page.WaitForTimeout(250)

	url := page.URL()
	subs := "/checkout?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("did not stay on checkout page, %s", page.URL())
	}
}

func TestOnlyAddressCheckout(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	getToCheckout(t, page)

	err = page.Locator("[name=\"address\"]").Fill("123 Main St")
	assertNotNil("could not fill address input: %w", err)

	confirmLocator := page.Locator("text=Confirm")
	confirm := confirmLocator.First()

	if confirm == nil {
		t.Errorf("confirm button not found")
	}

	confirm.Click()

	page.WaitForTimeout(250)

	url := page.URL()
	subs := "/checkout?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("did not stay on checkout page, %s", page.URL())
	}
}

func TestOnlyCityCheckout(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	getToCheckout(t, page)

	err = page.Locator("[name=\"city\"]").Fill("Anytown")
	assertNotNil("could not fill city input: %w", err)

	confirmLocator := page.Locator("text=Confirm")
	confirm := confirmLocator.First()

	if confirm == nil {
		t.Errorf("confirm button not found")
	}

	confirm.Click()

	page.WaitForTimeout(250)

	url := page.URL()
	subs := "/checkout?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("did not stay on checkout page, %s", page.URL())
	}
}

func TestOnlyZipCodeCheckout(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	getToCheckout(t, page)

	err = page.Locator("[name=\"zipCode\"]").Fill("12345")
	assertNotNil("could not fill zipcode input: %w", err)

	confirmLocator := page.Locator("text=Confirm")
	confirm := confirmLocator.First()

	if confirm == nil {
		t.Errorf("confirm button not found")
	}

	confirm.Click()

	page.WaitForTimeout(250)

	url := page.URL()
	subs := "/checkout?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("did not stay on checkout page, %s", page.URL())
	}
}

func TestOnlyStateCheckout(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	getToCheckout(t, page)

	err = page.Locator("[name=\"state\"]").Fill("CA")
	assertNotNil("could not fill state input: %w", err)

	confirmLocator := page.Locator("text=Confirm")
	confirm := confirmLocator.First()

	if confirm == nil {
		t.Errorf("confirm button not found")
	}

	confirm.Click()

	page.WaitForTimeout(250)

	url := page.URL()
	subs := "/checkout?cartId="

	if strings.Contains(url, subs) == false {
		t.Errorf("did not stay on checkout page, %s", page.URL())
	}
}

func TestFilledFormCheckout(t *testing.T) {
	page, err := context.NewPage()
	assertNotNil("could not create page: %w", err)

	defer page.Close()

	_, err = page.Goto("http://192.168.117.3:3000")
	assertNotNil("could not goto: %w", err)

	getToCheckout(t, page)

	// Fill all form fields
	err = page.Locator("[name=\"name\"]").Fill("John Doe")
	assertNotNil("could not fill name input: %w", err)

	err = page.Locator("[name=\"email\"]").Fill("john.doe@example.com")
	assertNotNil("could not fill email input: %w", err)

	err = page.Locator("[name=\"phone\"]").Fill("1234567890")
	assertNotNil("could not fill phone input: %w", err)

	err = page.Locator("[name=\"address\"]").Fill("123 Main St")
	assertNotNil("could not fill address input: %w", err)

	err = page.Locator("[name=\"city\"]").Fill("Anytown")
	assertNotNil("could not fill city input: %w", err)

	err = page.Locator("[name=\"zipCode\"]").Fill("12345")
	assertNotNil("could not fill zipCode input: %w", err)

	err = page.Locator("[name=\"state\"]").Fill("CA")
	assertNotNil("could not fill state input: %w", err)

	confirmLocator := page.Locator("text=Confirm")
	confirm := confirmLocator.First()

	if confirm == nil {
		t.Errorf("confirm button not found")
	}

	confirm.Click()

	page.WaitForTimeout(250)

	url := page.URL()
	subs := "/checkout?cartId="

	if strings.Contains(url, subs) == true {
		t.Errorf("did not navigate away from checkout page, %s", page.URL())
	}
}
