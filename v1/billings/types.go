package billings

import v1 "github.com/almeidazs/go-abacate-types/v1"

type Billing = v1.APICharge

type ListBillingsData = v1.RESTGetListBillingsData

type CreateBillingBody = v1.RESTPostCreateNewChargeBody

type CreateBillingData = v1.RESTPostCreateNewChargeData
