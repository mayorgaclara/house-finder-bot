package main

import "time"

type Response struct {
	SiteID                 string `json:"site_id"`
	CountryDefaultTimeZone string `json:"country_default_time_zone"`
	Paging                 struct {
		Total          int `json:"total"`
		PrimaryResults int `json:"primary_results"`
		Offset         int `json:"offset"`
		Limit          int `json:"limit"`
	} `json:"paging"`
	Results []struct {
		ID     string `json:"id"`
		SiteID string `json:"site_id"`
		Title  string `json:"title"`
		Seller struct {
			ID               int      `json:"id"`
			Permalink        string   `json:"permalink"`
			RegistrationDate string   `json:"registration_date"`
			CarDealer        bool     `json:"car_dealer"`
			RealEstateAgency bool     `json:"real_estate_agency"`
			Tags             []string `json:"tags"`
			SellerReputation struct {
				PowerSellerStatus interface{} `json:"power_seller_status"`
				LevelID           interface{} `json:"level_id"`
				Metrics           struct {
					Cancellations struct {
						Period string `json:"period"`
						Rate   int    `json:"rate"`
						Value  int    `json:"value"`
					} `json:"cancellations"`
					Claims struct {
						Period string `json:"period"`
						Rate   int    `json:"rate"`
						Value  int    `json:"value"`
					} `json:"claims"`
					DelayedHandlingTime struct {
						Period string `json:"period"`
						Rate   int    `json:"rate"`
						Value  int    `json:"value"`
					} `json:"delayed_handling_time"`
					Sales struct {
						Period    string `json:"period"`
						Completed int    `json:"completed"`
					} `json:"sales"`
				} `json:"metrics"`
				Transactions struct {
					Canceled int    `json:"canceled"`
					Period   string `json:"period"`
					Total    int    `json:"total"`
					Ratings  struct {
						Negative int `json:"negative"`
						Neutral  int `json:"neutral"`
						Positive int `json:"positive"`
					} `json:"ratings"`
					Completed int `json:"completed"`
				} `json:"transactions"`
			} `json:"seller_reputation"`
		} `json:"seller,omitempty"`
		Price  int `json:"price"`
		Prices struct {
			ID     string `json:"id"`
			Prices []struct {
				ID            string      `json:"id"`
				Type          string      `json:"type"`
				Amount        int         `json:"amount"`
				RegularAmount interface{} `json:"regular_amount"`
				CurrencyID    string      `json:"currency_id"`
				LastUpdated   time.Time   `json:"last_updated"`
				Conditions    struct {
					ContextRestrictions []interface{} `json:"context_restrictions"`
					StartTime           interface{}   `json:"start_time"`
					EndTime             interface{}   `json:"end_time"`
					Eligible            bool          `json:"eligible"`
				} `json:"conditions"`
				ExchangeRateContext string `json:"exchange_rate_context"`
				Metadata            struct {
				} `json:"metadata"`
			} `json:"prices"`
			Presentation struct {
				DisplayCurrency string `json:"display_currency"`
			} `json:"presentation"`
			PaymentMethodPrices []interface{} `json:"payment_method_prices"`
			ReferencePrices     []interface{} `json:"reference_prices"`
			PurchaseDiscounts   []interface{} `json:"purchase_discounts"`
		} `json:"prices"`
		SalePrice          interface{} `json:"sale_price"`
		CurrencyID         string      `json:"currency_id"`
		AvailableQuantity  int         `json:"available_quantity"`
		SoldQuantity       int         `json:"sold_quantity"`
		BuyingMode         string      `json:"buying_mode"`
		ListingTypeID      string      `json:"listing_type_id"`
		StopTime           time.Time   `json:"stop_time"`
		Condition          string      `json:"condition"`
		Permalink          string      `json:"permalink"`
		Thumbnail          string      `json:"thumbnail"`
		ThumbnailID        string      `json:"thumbnail_id"`
		AcceptsMercadopago bool        `json:"accepts_mercadopago"`
		Installments       interface{} `json:"installments"`
		Address            struct {
			StateID   string `json:"state_id"`
			StateName string `json:"state_name"`
			CityID    string `json:"city_id"`
			CityName  string `json:"city_name"`
			AreaCode  string `json:"area_code"`
			Phone1    string `json:"phone1"`
		} `json:"address"`
		Promotions interface{} `json:"promotions"`
		Shipping   struct {
			FreeShipping bool          `json:"free_shipping"`
			Mode         string        `json:"mode"`
			Tags         []interface{} `json:"tags"`
			LogisticType interface{}   `json:"logistic_type"`
			StorePickUp  bool          `json:"store_pick_up"`
		} `json:"shipping"`
		SellerAddress struct {
			ID          string `json:"id"`
			Comment     string `json:"comment"`
			AddressLine string `json:"address_line"`
			ZipCode     string `json:"zip_code"`
			Country     struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"country"`
			State struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"state"`
			City struct {
				ID   interface{} `json:"id"`
				Name string      `json:"name"`
			} `json:"city"`
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
		} `json:"seller_address"`
		SellerContact struct {
			Contact   string `json:"contact"`
			OtherInfo string `json:"other_info"`
			AreaCode  string `json:"area_code"`
			Phone     string `json:"phone"`
			AreaCode2 string `json:"area_code2"`
			Phone2    string `json:"phone2"`
			Email     string `json:"email"`
			Webpage   string `json:"webpage"`
		} `json:"seller_contact"`
		Location struct {
			AddressLine     string      `json:"address_line"`
			ZipCode         string      `json:"zip_code"`
			Subneighborhood interface{} `json:"subneighborhood"`
			Neighborhood    struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"neighborhood"`
			City struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"city"`
			State struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"state"`
			Country struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"country"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
		Attributes []struct {
			Name             string      `json:"name"`
			ValueID          interface{} `json:"value_id"`
			ValueName        string      `json:"value_name"`
			AttributeGroupID string      `json:"attribute_group_id"`
			ID               string      `json:"id"`
			ValueStruct      interface{} `json:"value_struct"`
			Values           []struct {
				Source int         `json:"source"`
				ID     interface{} `json:"id"`
				Name   string      `json:"name"`
				Struct interface{} `json:"struct"`
			} `json:"values"`
			AttributeGroupName string `json:"attribute_group_name"`
			Source             int    `json:"source"`
		} `json:"attributes"`
		OriginalPrice    interface{} `json:"original_price"`
		CategoryID       string      `json:"category_id"`
		OfficialStoreID  interface{} `json:"official_store_id"`
		DomainID         string      `json:"domain_id"`
		CatalogProductID interface{} `json:"catalog_product_id"`
		Tags             []string    `json:"tags"`
		OrderBackend     int         `json:"order_backend"`
		UseThumbnailID   bool        `json:"use_thumbnail_id"`
		OfferScore       interface{} `json:"offer_score"`
		OfferShare       interface{} `json:"offer_share"`
		MatchScore       interface{} `json:"match_score"`
		WinnerItemID     interface{} `json:"winner_item_id"`
		Melicoin         interface{} `json:"melicoin"`
		Discounts        interface{} `json:"discounts"`
		Seller0          struct {
			ID               int      `json:"id"`
			Permalink        string   `json:"permalink"`
			RegistrationDate string   `json:"registration_date"`
			CarDealer        bool     `json:"car_dealer"`
			RealEstateAgency bool     `json:"real_estate_agency"`
			Tags             []string `json:"tags"`
			HomeImageURL     string   `json:"home_image_url"`
			SellerReputation struct {
				PowerSellerStatus interface{} `json:"power_seller_status"`
				LevelID           interface{} `json:"level_id"`
				Metrics           struct {
					Cancellations struct {
						Period string `json:"period"`
						Rate   int    `json:"rate"`
						Value  int    `json:"value"`
					} `json:"cancellations"`
					Claims struct {
						Period string `json:"period"`
						Rate   int    `json:"rate"`
						Value  int    `json:"value"`
					} `json:"claims"`
					DelayedHandlingTime struct {
						Period string `json:"period"`
						Rate   int    `json:"rate"`
						Value  int    `json:"value"`
					} `json:"delayed_handling_time"`
					Sales struct {
						Period    string `json:"period"`
						Completed int    `json:"completed"`
					} `json:"sales"`
				} `json:"metrics"`
				Transactions struct {
					Canceled int    `json:"canceled"`
					Period   string `json:"period"`
					Total    int    `json:"total"`
					Ratings  struct {
						Negative int `json:"negative"`
						Neutral  int `json:"neutral"`
						Positive int `json:"positive"`
					} `json:"ratings"`
					Completed int `json:"completed"`
				} `json:"transactions"`
			} `json:"seller_reputation"`
		} `json:"seller,omitempty"`
		Seller1 struct {
			ID               int      `json:"id"`
			Permalink        string   `json:"permalink"`
			RegistrationDate string   `json:"registration_date"`
			CarDealer        bool     `json:"car_dealer"`
			RealEstateAgency bool     `json:"real_estate_agency"`
			Tags             []string `json:"tags"`
			HomeImageURL     string   `json:"home_image_url"`
			SellerReputation struct {
				PowerSellerStatus interface{} `json:"power_seller_status"`
				LevelID           interface{} `json:"level_id"`
				Metrics           struct {
					Cancellations struct {
						Period string `json:"period"`
						Rate   int    `json:"rate"`
						Value  int    `json:"value"`
					} `json:"cancellations"`
					Claims struct {
						Period string `json:"period"`
						Rate   int    `json:"rate"`
						Value  int    `json:"value"`
					} `json:"claims"`
					DelayedHandlingTime struct {
						Period string `json:"period"`
						Rate   int    `json:"rate"`
						Value  int    `json:"value"`
					} `json:"delayed_handling_time"`
					Sales struct {
						Period    string `json:"period"`
						Completed int    `json:"completed"`
					} `json:"sales"`
				} `json:"metrics"`
				Transactions struct {
					Canceled int    `json:"canceled"`
					Period   string `json:"period"`
					Total    int    `json:"total"`
					Ratings  struct {
						Negative int `json:"negative"`
						Neutral  int `json:"neutral"`
						Positive int `json:"positive"`
					} `json:"ratings"`
					Completed int `json:"completed"`
				} `json:"transactions"`
			} `json:"seller_reputation"`
		} `json:"seller,omitempty"`
	} `json:"results"`
	Sort struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"sort"`
	AvailableSorts []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"available_sorts"`
	Filters []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Values []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"filters"`
	AvailableFilters []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Values []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Results int    `json:"results"`
		} `json:"values"`
	} `json:"available_filters"`
}
