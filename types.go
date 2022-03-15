package bulkvs2go

type BulkVS2GoClient struct {
	APIURL      string `json:"APIURL"`
	APIUsername string `json:"APIUsername"`
	APIPassword string `json:"APIPassword"`
}

type MessageSendResponse struct {
	RefID       string `json:"RefId"`
	From        string `json:"From"`
	MessageType string `json:"MessageType"`
	Results     []struct {
		To     string `json:"To"`
		Status string `json:"Status"`
	} `json:"Results"`
}

type MessageSendRequest struct {
	From      string   `json:"From"`
	To        []string `json:"To"`
	Message   string   `json:"Message"`
	MediaURLs []string `json:"MediaURLs"`
}

type GETPortTN300Response []struct {
	OrderID     string `json:"OrderId"`
	ReferenceID string `json:"ReferenceId"`
	LNPStatus   string `json:"LNP Status"`
	Rdd         string `json:"RDD"`
}
type ValidateAddressRequest struct {
	StreetNumber string `json:"Street Number"`
	StreetName   string `json:"Street Name"`
	Location     string `json:"Location"`
	City         string `json:"City"`
	State        string `json:"State"`
	Zip          string `json:"Zip"`
}
type CampaignsResponse []struct {
	Tcr            string `json:"Tcr"`
	Brand          string `json:"Brand"`
	Count          string `json:"Count"`
	ActivationDate string `json:"Activation Date"`
}
type Voice200Response []struct {
	CallStart       string `json:"callStart"`
	DurationSecs    string `json:"durationSecs"`
	CallSource      string `json:"callSource"`
	CallDestination string `json:"callDestination"`
	CallID          string `json:"callID"`
	PerMinute       string `json:"perMinute"`
	Amount          string `json:"amount"`
	CallType        string `json:"callType"`
	RemoteIP        string `json:"RemoteIP"`
	Tag             string `json:"tag"`
}
type Voice400Response struct {
	Status      string `json:"Status"`
	Code        string `json:"Code"`
	Description struct {
		ValidType []string `json:"Valid Type"`
	} `json:"Description"`
}
type Voice404Response struct {
	Status      string `json:"Status"`
	Code        string `json:"Code"`
	Description string `json:"Description"`
}
type ValidateAddressResponse struct {
	Status       string `json:"Status"`
	AddressID    string `json:"AddressID"`
	AddressLine1 string `json:"Address Line 1"`
	AddressLine2 string `json:"Address Line 2"`
	City         string `json:"City"`
	State        string `json:"State"`
	Zip          string `json:"Zip"`
}
type POSTE911RecordRequest struct {
	Tn         string   `json:"TN"`
	CallerName string   `json:"Caller Name"`
	AddressID  string   `json:"AddressID"`
	Sms        []string `json:"Sms"`
}
type POSTE911RecordResponse struct {
	Status           string   `json:"Status"`
	Tn               string   `json:"TN"`
	CallerName       string   `json:"Caller Name"`
	AddressLine1     string   `json:"Address Line 1"`
	AddressLine2     string   `json:"Address Line 2"`
	City             string   `json:"City"`
	State            string   `json:"State"`
	Zip              string   `json:"Zip"`
	Sms              []string `json:"Sms"`
	LastModification string   `json:"Last Modification"`
}
type GETE911RecordResponse []struct {
	Tn               string   `json:"TN"`
	CallerName       string   `json:"Caller Name"`
	AddressLine1     string   `json:"Address Line 1"`
	AddressLine2     string   `json:"Address Line 2"`
	City             string   `json:"City"`
	State            string   `json:"State"`
	Zip              string   `json:"Zip"`
	Sms              []string `json:"Sms"`
	LastModification string   `json:"Last Modification"`
}
type PUTPortTNResponse struct {
	OrderID     string `json:"OrderId"`
	CustOrderID string `json:"CustOrderId"`
	Rdd         string `json:"RDD"`
	Tn          string `json:"TN"`
	Code        string `json:"Code"`
	Description string `json:"Description"`
	TNGroups    struct {
		Verizon []string `json:"Verizon"`
		ATT     []string `json:"AT&T"`
	} `json:"TN Groups"`
}

type PUTPortTNRequest struct {
	ReferenceID    string   `json:"ReferenceId"`
	TNList         []string `json:"TN List"`
	Btn            string   `json:"BTN"`
	SubscriberType string   `json:"Subscriber Type"`
	AccountNumber  string   `json:"Account Number"`
	Pin            string   `json:"Pin"`
	Name           string   `json:"Name"`
	Contact        string   `json:"Contact"`
	StreetNumber   string   `json:"Street Number"`
	StreetName     string   `json:"Street Name"`
	City           string   `json:"City"`
	State          string   `json:"State"`
	Zip            string   `json:"Zip"`
	Rdd            string   `json:"RDD"`
	Time           string   `json:"Time"`
	PortoutPin     string   `json:"Portout Pin"`
	TrunkGroup     string   `json:"Trunk Group"`
	Lidb           string   `json:"Lidb"`
	Sms            bool     `json:"Sms"`
	Mms            bool     `json:"Mms"`
	SignLoa        bool     `json:"Sign Loa"`
	Notify         string   `json:"Notify"`
}

type GETPortTNResponse struct {
	OrderDetails struct {
		OrderID       string `json:"OrderId"`
		ReferenceID   string `json:"ReferenceId"`
		LosingCarrier string `json:"Losing Carrier"`
		Tier          string `json:"Tier"`
		TrunkGroup    string `json:"Trunk Group"`
		Btn           string `json:"BTN"`
	} `json:"Order Details"`
	EndUserInfo struct {
		Name          string `json:"Name"`
		Contact       string `json:"Contact"`
		Address       string `json:"Address"`
		City          string `json:"City"`
		State         string `json:"State"`
		Zip           string `json:"Zip"`
		AccountNumber string `json:"Account Number"`
		Pin           string `json:"Pin"`
	} `json:"End User Info"`
	TNList []struct {
		Tn        string `json:"TN"`
		LNPStatus string `json:"LNP Status"`
		Rdd       string `json:"RDD,omitempty"`
		Reason    string `json:"Reason,omitempty"`
	} `json:"TN List"`
}

type WebhooksResponse []struct {
	Webhook         string `json:"Webhook"`
	Description     string `json:"Description"`
	URL             string `json:"Url"`
	LastModifiction string `json:"Last Modifiction"`
}
type ValidatePortabilityResponse struct {
	Tn            string `json:"TN"`
	Portable      bool   `json:"Portable"`
	LosingCarrier string `json:"Losing Carrier"`
	Tier          string `json:"Tier"`
	RateCenter    string `json:"Rate Center"`
	State         string `json:"State"`
	PerMinuteRate string `json:"Per Minute Rate"`
	Mrc           string `json:"Mrc"`
	LNPFee        string `json:"LNP Fee"`
}

type QueryTNResponse []struct {
	Tn            string `json:"TN"`
	RateCenter    string `json:"Rate Center"`
	State         string `json:"State"`
	Tier          string `json:"Tier"`
	PerMinuteRate string `json:"Per Minute Rate"`
	Mrc           string `json:"Mrc"`
	Nrc           string `json:"Nrc"`
}
type OrderTNResponse []struct {
	Tn         string `json:"TN"`
	Status     string `json:"Status"`
	Lidb       string `json:"Lidb"`
	PortoutPin string `json:"Portout Pin"`
	Routing    struct {
		TrunkGroup  string `json:"Trunk Group"`
		CustomURI   string `json:"Custom URI"`
		CallForward string `json:"Call Forward"`
	} `json:"Routing"`
	Messaging struct {
		Sms bool `json:"Sms"`
		Mms bool `json:"Mms"`
	} `json:"Messaging"`
	TNDetails struct {
		RateCenter     string `json:"Rate Center"`
		State          string `json:"State"`
		Tier           int    `json:"Tier"`
		Cnam           bool   `json:"Cnam"`
		ActivationDate string `json:"Activation Date"`
	} `json:"TN Details"`
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

type OrderTNRequest struct {
	Tn         string `json:"TN"`
	Lidb       string `json:"Lidb"`
	PortoutPin string `json:"Portout Pin"`
	TrunkGroup string `json:"Trunk Group"`
	Sms        bool   `json:"Sms"`
	Mms        bool   `json:"Mms"`
}

type TrunkGroupsResponse []struct {
	TrunkGroup       string `json:"TrunkGroup"`
	Description      string `json:"Description"`
	PrimaryHost      string `json:"Primary Host"`
	SecondaryHost    string `json:"Secondary Host"`
	DigitDelivery    string `json:"Digit Delivery"`
	LastModification string `json:"Last Modification"`
}

type TNRecordRequest struct {
	Tn           string `json:"TN"`
	Lidb         string `json:"Lidb"`
	PortoutPin   string `json:"Portout Pin"`
	TrunkGroup   string `json:"Trunk Group"`
	CustomURI    string `json:"Custom URI"`
	CallForward  string `json:"Call Forward"`
	PSTNFailover string `json:"PSTN Failover"`
	Sms          bool   `json:"Sms"`
	Mms          bool   `json:"Mms"`
	Tcr          string `json:"Tcr"`
	Webhook      string `json:"Webhook"`
}

type GETTNRecordRequest struct {
	Number     string
	Status     string
	Tier       string
	TrunkGroup string
	PortoutPin string
	Lidb       string
	Class      string
	Sms        string
	Mms        string
	Webhook    string
	Tcr        string
	Limit      string
}

type TNRecordResponse []struct {
	Tn         string `json:"TN"`
	Status     string `json:"Status"`
	Lidb       string `json:"Lidb"`
	PortoutPin string `json:"Portout Pin"`
	Routing    struct {
		TrunkGroup   string `json:"Trunk Group"`
		CustomURI    string `json:"Custom URI"`
		CallForward  string `json:"Call Forward"`
		PSTNFailover string `json:"PSTN Failover"`
	} `json:"Routing"`
	Messaging struct {
		Class   string `json:"Class"`
		Sms     bool   `json:"Sms"`
		Mms     bool   `json:"Mms"`
		Webhook string `json:"Webhook"`
		Tcr     string `json:"Tcr"`
	} `json:"Messaging"`
	TNDetails struct {
		RateCenter     string `json:"Rate Center"`
		State          string `json:"State"`
		Tier           int    `json:"Tier"`
		Cnam           bool   `json:"Cnam"`
		ActivationDate string `json:"Activation Date"`
	} `json:"TN Details"`
}

type AccountDetailResponse struct {
	Username       string `json:"Username"`
	BalanceRelated struct {
		Balance             string `json:"Balance"`
		LowBalanceThreshold string `json:"Low Balance Threshold"`
	} `json:"Balance Related"`
	TnRelated struct {
		PortoutPin string `json:"Portout Pin"`
	} `json:"Tn Related"`
	ServiceStatus struct {
		Lrn                 string `json:"Lrn"`
		OutboundTermination string `json:"Outbound Termination"`
		InboundOrigination  string `json:"Inbound Origination"`
		E911                string `json:"E911"`
		EightYY             string `json:"8YY"`
		Sms                 string `json:"Sms"`
		Mms                 string `json:"Mms"`
		InvoicedBilling     string `json:"Invoiced Billing"`
	} `json:"Service Status"`
	MainContact struct {
		Name            bool `json:"Name"`
		Email           bool `json:"Email"`
		TelephoneNumber bool `json:"Telephone Number"`
	} `json:"Main Contact"`
	E911Contact struct {
		Name            bool `json:"Name"`
		Email           bool `json:"Email"`
		TelephoneNumber bool `json:"Telephone Number"`
	} `json:"E911 Contact"`
	LNPContact struct {
		Name            bool `json:"Name"`
		Email           bool `json:"Email"`
		TelephoneNumber bool `json:"Telephone Number"`
	} `json:"LNP Contact"`
	MaintenanceContact struct {
		Name            bool `json:"Name"`
		Email           bool `json:"Email"`
		TelephoneNumber bool `json:"Telephone Number"`
	} `json:"Maintenance Contact"`
	BillingContact struct {
		Name            bool `json:"Name"`
		Email           bool `json:"Email"`
		TelephoneNumber bool `json:"Telephone Number"`
	} `json:"Billing Contact"`
	Rates struct {
		E911 string `json:"E911"`
		Tn   string `json:"Tn"`
		Lnp  int    `json:"Lnp"`
		Cnam bool   `json:"Cnam"`
		Lidb bool   `json:"Lidb"`
	} `json:"Rates"`
}
