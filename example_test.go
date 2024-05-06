package hubspot_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"bendingspoons.com/hubspot"
)

const (
	hubspotTestAccountApiKey = "eu1-f11a-f6c3-4999-abc2-493eb957f6cd"
)

type ExampleContact struct {
	email     string
	firstName string
	lastName  string
	phone     string
	zip       string
}

func ExampleContactServiceOp_Create() {
	cli, _ := hubspot.NewClient(hubspot.SetAPIKey("apikey"))

	example := &ExampleContact{
		email:     "hubspot@example.com",
		firstName: "Bryan",
		lastName:  "Cooper",
		phone:     "(877) 929-0687",
	}

	contact := &hubspot.Contact{
		Email:       hubspot.NewString(example.email),
		FirstName:   hubspot.NewString(example.firstName),
		LastName:    hubspot.NewString(example.lastName),
		MobilePhone: hubspot.NewString(example.phone),
		Website:     hubspot.NewString("example.com"),
		Zip:         nil,
	}

	res, err := cli.CRM.Contact.Create(contact)
	if err != nil {
		log.Fatal(err)
	}

	r, ok := res.Properties.(*hubspot.Contact)
	if !ok {
		log.Fatal("unable to type assertion")
	}

	// use properties
	_ = r

	fmt.Println(res)

	// // Output:
}

func ExampleContactServiceOp_Update() {
	cli, _ := hubspot.NewClient(hubspot.SetAPIKey("apikey"))

	example := &ExampleContact{
		email:     "hubspot@example.com",
		firstName: "Bryan",
		lastName:  "Cooper",
		phone:     "(877) 929-0687",
		zip:       "1000001",
	}

	contact := &hubspot.Contact{
		Email:       hubspot.NewString(example.email),
		FirstName:   hubspot.NewString(example.firstName),
		LastName:    hubspot.NewString(example.lastName),
		MobilePhone: hubspot.NewString(example.phone),
		Website:     hubspot.NewString("example.com"),
		Zip:         hubspot.NewString(example.zip),
	}

	res, err := cli.CRM.Contact.Update("contact001", contact)
	if err != nil {
		log.Fatal(err)
	}

	r, ok := res.Properties.(*hubspot.Contact)
	if !ok {
		log.Fatal("unable to type assertion")
	}

	// use properties
	_ = r

	fmt.Println(res)

	// // Output:
}

func ExampleContactServiceOp_Get() {
	cli, _ := hubspot.NewClient(hubspot.SetAPIKey("apikey"))

	res, err := cli.CRM.Contact.Get("contact001", &hubspot.Contact{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	r, ok := res.Properties.(*hubspot.Contact)
	if !ok {
		log.Fatal("unable to type assertion")
	}

	// use properties
	_ = r

	fmt.Println(res)

	// // Output:
}

func ExampleContactServiceOp_AssociateAnotherObj() {
	cli, _ := hubspot.NewClient(hubspot.SetAPIKey("apikey"))

	res, err := cli.CRM.Contact.AssociateAnotherObj("contact001", &hubspot.AssociationConfig{
		ToObject:   hubspot.ObjectTypeDeal,
		ToObjectID: "deal001",
		Type:       hubspot.AssociationTypeContactToDeal,
	})
	if err != nil {
		log.Fatal(err)
	}

	r, ok := res.Properties.(*hubspot.Contact)
	if !ok {
		log.Fatal("unable to type assertion")
	}

	// use properties
	_ = r

	fmt.Println(res)

	// // Output:
}

type ExampleDeal struct {
	amount  string
	name    string
	stage   string
	ownerID string
}

func ExampleDealServiceOp_Create_apikey() {
	cli, _ := hubspot.NewClient(hubspot.SetAPIKey("apikey"))

	example := &ExampleDeal{
		amount:  "1500.00",
		name:    "Custom data integrations",
		stage:   "presentation scheduled",
		ownerID: "910901",
	}

	deal := &hubspot.Deal{
		Amount:      hubspot.NewString(example.amount),
		DealName:    hubspot.NewString(example.name),
		DealStage:   hubspot.NewString(example.stage),
		DealOwnerID: hubspot.NewString(example.ownerID),
		PipeLine:    hubspot.NewString("default"),
	}

	res, err := cli.CRM.Deal.Create(deal)
	if err != nil {
		log.Fatal(err)
	}

	r, ok := res.Properties.(*hubspot.Deal)
	if !ok {
		log.Fatal("unable to type assertion")
	}

	// use properties
	_ = r

	fmt.Println(res)

	// // Output:
}

func ExampleDealServiceOp_Create_oauth() {
	cli, _ := hubspot.NewClient(hubspot.SetOAuth(&hubspot.OAuthConfig{
		GrantType:    hubspot.GrantTypeRefreshToken,
		ClientID:     "hubspot-client-id",
		ClientSecret: "hubspot-client-secret",
		RefreshToken: "hubspot-refresh-token",
	}))

	example := &ExampleDeal{
		amount:  "1500.00",
		name:    "Custom data integrations",
		stage:   "presentation scheduled",
		ownerID: "910901",
	}

	deal := &hubspot.Deal{
		Amount:      hubspot.NewString(example.amount),
		DealName:    hubspot.NewString(example.name),
		DealStage:   hubspot.NewString(example.stage),
		DealOwnerID: hubspot.NewString(example.ownerID),
		PipeLine:    hubspot.NewString("default"),
	}

	res, err := cli.CRM.Deal.Create(deal)
	if err != nil {
		log.Fatal(err)
	}

	r, ok := res.Properties.(*hubspot.Deal)
	if !ok {
		log.Fatal("unable to type assertion")
	}

	// use properties
	_ = r

	fmt.Println(res)

	// // Output:
}

type CustomDeal struct {
	hubspot.Deal
	CustomA string `json:"custom_a,omitempty"`
	CustomB string `json:"custom_b,omitempty"`
}

func ExampleDealServiceOp_Create_custom() {
	cli, _ := hubspot.NewClient(hubspot.SetAPIKey("apikey"))

	example := &ExampleDeal{
		amount:  "1500.00",
		name:    "Custom data integrations",
		stage:   "presentation scheduled",
		ownerID: "910901",
	}

	// Take advantage of structure embedding when using custom fields.
	deal := &CustomDeal{
		Deal: hubspot.Deal{
			Amount:      hubspot.NewString(example.amount),
			DealName:    hubspot.NewString(example.name),
			DealStage:   hubspot.NewString(example.stage),
			DealOwnerID: hubspot.NewString(example.ownerID),
			PipeLine:    hubspot.NewString("default"),
		},
		CustomA: "custom field A",
		CustomB: "custom field B",
	}

	res, err := cli.CRM.Deal.Create(deal)
	if err != nil {
		log.Fatal(err)
	}

	r, ok := res.Properties.(*CustomDeal)
	if !ok {
		log.Fatal("unable to type assertion")
	}

	// use custom struct
	_ = r

	// // Output:
}

func ExampleDealServiceOp_Update() {
	cli, _ := hubspot.NewClient(hubspot.SetAPIKey("apikey"))

	example := &ExampleDeal{
		amount:  "1500.00",
		name:    "Custom data integrations",
		stage:   "presentation scheduled",
		ownerID: "910901",
	}

	deal := &hubspot.Deal{
		Amount:      hubspot.NewString(example.amount),
		DealName:    hubspot.NewString(example.name),
		DealStage:   hubspot.NewString(example.stage),
		DealOwnerID: hubspot.NewString(example.ownerID),
		PipeLine:    hubspot.NewString("default"),
	}

	res, err := cli.CRM.Deal.Update("deal001", deal)
	if err != nil {
		log.Fatal(err)
	}

	r, ok := res.Properties.(*hubspot.Deal)
	if !ok {
		log.Fatal("unable to type assertion")
	}

	// use properties
	_ = r

	fmt.Println(res)

	// // Output:
}

func ExampleDealServiceOp_Get() {
	cli, _ := hubspot.NewClient(hubspot.SetAPIKey("apikey"))

	res, err := cli.CRM.Deal.Get("deal001", &hubspot.Deal{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	r, ok := res.Properties.(*hubspot.Deal)
	if !ok {
		log.Fatal("unable to type assertion")
	}

	// use properties
	_ = r

	fmt.Println(res)

	// // Output:
}

func ExampleDealServiceOp_Get_custom() {
	cli, _ := hubspot.NewClient(hubspot.SetAPIKey("apikey"))

	res, err := cli.CRM.Deal.Get("deal001", &CustomDeal{}, &hubspot.RequestQueryOption{
		CustomProperties: []string{
			"custom_a",
			"custom_b",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	r, ok := res.Properties.(*CustomDeal)
	if !ok {
		log.Fatal("unable to type assertion")
	}

	// use properties
	_ = r

	fmt.Println(res)

	// // Output:
}

func ExampleDealServiceOp_AssociateAnotherObj() {
	cli, _ := hubspot.NewClient(hubspot.SetAPIKey("apikey"))

	res, err := cli.CRM.Deal.AssociateAnotherObj("deal001", &hubspot.AssociationConfig{
		ToObject:   hubspot.ObjectTypeContact,
		ToObjectID: "contact001",
		Type:       hubspot.AssociationTypeDealToContact,
	})
	if err != nil {
		log.Fatal(err)
	}

	r, ok := res.Properties.(*hubspot.Deal)
	if !ok {
		log.Fatal("unable to type assertion")
	}

	// use properties
	_ = r

	fmt.Println(res)

	// // Output:
}

func TestExampleOwnerServiceOp_Get(t *testing.T) {

	t.Run("Test 1", func(t *testing.T) {
		cli, _ := hubspot.NewClient(hubspot.SetAPIKey(hubspotTestAccountApiKey))

		fmt.Println("Before Get")

		res, err := cli.CRM.Owner.Get("297111267", &hubspot.Owner{}, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Res after Get: %+v\n", reflect.TypeOf(res))

		switch res.(type) {
		case *hubspot.Owner:
			fmt.Printf("Matching type 1 found!\n")
		case hubspot.Owner:
			fmt.Printf("Matching type 2 found!\n")
		}

		r, ok := res.(*hubspot.Owner)
		if !ok {
			log.Fatal("unable to type assertion")
		}

		// use properties
		//_ = r

		fmt.Println(r)

		// // Output:

		t.Errorf("Response mismatch: %+v", r)
	})
}

func TestExampleOwnerServiceOp_GetAll(t *testing.T) {

	t.Run("Test 2", func(t *testing.T) {
		cli, _ := hubspot.NewClient(hubspot.SetAPIKey(hubspotTestAccountApiKey))

		fmt.Println("Before Get")

		res, err := cli.CRM.Owner.GetAll(&hubspot.Owner{}, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Res after Get: %+v\n", res)

		for _, results := range res.Results {
			r, ok := results.(*hubspot.Owner)
			if !ok {
				log.Fatal("unable to type assertion")
			}

			// use properties
			_ = r

			fmt.Println(r)

			// // Output:

			t.Errorf("Response mismatch: %+v", r)

		}
	})
}

func TestExampleOPipelineServiceOp_Get(t *testing.T) {

	t.Run("Test 1", func(t *testing.T) {
		cli, _ := hubspot.NewClient(hubspot.SetAPIKey(hubspotTestAccountApiKey))

		fmt.Println("Before Get")

		res, err := cli.CRM.Pipeline.Get("17542634", &hubspot.Pipeline{}, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Res after Get: %+v\n", reflect.TypeOf(res))

		switch res.(type) {
		case *hubspot.Pipeline:
			fmt.Printf("Matching type 1 found!\n")
		}

		r, ok := res.(*hubspot.Pipeline)
		if !ok {
			log.Fatal("unable to type assertion")
		}

		// use properties
		//_ = r

		fmt.Println(r)

		for i, stage := range *r.Stages {
			fmt.Printf("Stage %d: %+v\n", i, stage)
		}

		// // Output:

		t.Errorf("Response mismatch: %+v", r)
	})
}

func TestExamplePipelineServiceOp_GetAll(t *testing.T) {

	t.Run("Test 2", func(t *testing.T) {
		cli, _ := hubspot.NewClient(hubspot.SetAPIKey(hubspotTestAccountApiKey))

		fmt.Println("Before Get")

		res, err := cli.CRM.Owner.GetAll(&hubspot.Pipeline{}, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Res after Get: %+v\n", res)

		for i, results := range res.Results {
			r, ok := results.(*hubspot.Pipeline)
			if !ok {
				log.Fatal("unable to type assertion")
			}

			// use properties
			_ = r

			fmt.Printf("Result pipeline %d: %+v\n", i, r)

			// // Output:

			t.Errorf("Response mismatch: %+v", r)

		}
	})
}

func TestExampleCustomerServiceOp_Search(t *testing.T) {

	t.Run("Test 3", func(t *testing.T) {
		cli, _ := hubspot.NewClient(hubspot.SetAPIKey(hubspotTestAccountApiKey))

		fmt.Println("Before Search")

		filters := []hubspot.Filter{}
		filters = append(filters, hubspot.Filter{Value: "Company", PropertyName: "name", Operator: "EQ"})

		filterGroups := []hubspot.FilterGroup{}
		filterGroups = append(filterGroups, hubspot.FilterGroup{Filters: filters})

		res, err := cli.CRM.Company.Search(&hubspot.Company{}, &hubspot.RequestSearchOption{FilterGroups: filterGroups})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Res after Post: %+v\n", res)

		for i, result := range res.Results {

			fmt.Printf("Result pipeline %d: %+v\n", i, result)

			// // Output:

		}

		t.Errorf("Response mismatch: %+v", res)

	})
}
