package models

type UserAccountsRequest struct {
	Viewer struct {
		Digital struct {
			Services []struct {
				BackendServiceID string `graphql:"backendServiceId"`
				ServiceID        string `graphql:"serviceId"`
				Status           string `graphql:"status"`
				Type             string `graphql:"type"`
			} `graphql:"services"`
			Accounts []struct {
				ID        string `graphql:"id"`
				Type      string `graphql:"type"`
				AccountID string `graphql:"accountId"`
				Status    string `graphql:"status"`
			} `graphql:"accounts"`
			Features []struct {
				ID        string `graphql:"id"`
				Type      string `graphql:"type"`
				FeatureID string `graphql:"featureId"`
				Status    string `graphql:"status"`
			} `graphql:"features"`
			User struct {
				CustomerType string `graphql:"customerType"`
			} `graphql:"user"`
		} `graphql:"digital"`
	} `graphql:"viewer"`
}

func (u *UserAccountsRequest) OperationName() string {
	return "UserAccounts"
}

// {
//     "data": {
//         "viewer": {
//             "digital": {
//                 "services": [
//                     {
//                         "backendServiceId": "A-C8C88EA0-3566204",
//                         "serviceId": "da76e5c7-7627-4f21-ae77-8990e1c853dc",
//                         "status": "ACTIVE",
//                         "type": "HOT_WATER",
//                         "__typename": "DigitalService"
//                     },
//                     {
//                         "backendServiceId": "A-C8C88EA0-8185",
//                         "serviceId": "808e95e9-bd47-49dc-aeee-45b7b7773fa4",
//                         "status": "ACTIVE",
//                         "type": "CES_GAS_APPLIANCE",
//                         "__typename": "DigitalService"
//                     },
//                     {
//                         "backendServiceId": "A-DD5EBD8C-4269228",
//                         "serviceId": "3fb72a2f-f37a-4a0e-ba61-7c06f4d2667f",
//                         "status": "ACTIVE",
//                         "type": "CES_ELECTRICITY",
//                         "__typename": "DigitalService"
//                     }
//                 ],
//                 "accounts": [
//                     {
//                         "id": "eyJ0eXBlIjoiRGlnaXRhbEFjY291bnQiLCJpZCI6ImFiNjA5NDYwLWI1YWItNDRiMi1iYmU4LTEyMTVjZTEyNWI5NyJ9",
//                         "type": "KRAKEN",
//                         "accountId": "A-C8C88EA0",
//                         "status": "ACTIVE",
//                         "__typename": "DigitalAccount"
//                     },
//                     {
//                         "id": "eyJ0eXBlIjoiRGlnaXRhbEFjY291bnQiLCJpZCI6IjgxNGJiZjliLWE1MTQtNDhhYi04N2QwLTFlMWJlZTQxYzZiZCJ9",
//                         "type": "KRAKEN",
//                         "accountId": "A-DD5EBD8C",
//                         "status": "ACTIVE",
//                         "__typename": "DigitalAccount"
//                     }
//                 ],
//                 "features": [
//                     {
//                         "id": "eyJ0eXBlIjoiRmVhdHVyZSIsImlkIjoiODljY2JlNDgtZjFlNC00YWY0LWI4YzMtODgzNzgyNDEwNWI0IiwiZmVhdHVyZVR5cGUiOiJCSVpBIn0=",
//                         "type": "BIZA",
//                         "featureId": "89ccbe48-f1e4-4af4-b8c3-8837824105b4",
//                         "status": "ACTIVE",
//                         "__typename": "Feature"
//                     },
//                     {
//                         "id": "eyJ0eXBlIjoiRmVhdHVyZSIsImlkIjoiYjljMjU4ZWMtZjgzNy0xMWVjLThmM2ItMGE5NzQ1NmZiZTE2IiwiZmVhdHVyZVR5cGUiOiJCUkFaRSJ9",
//                         "type": "BRAZE",
//                         "featureId": "b9c258ec-f837-11ec-8f3b-0a97456fbe16",
//                         "status": "ACTIVE",
//                         "__typename": "Feature"
//                     },
//                     {
//                         "id": "eyJ0eXBlIjoiRmVhdHVyZSIsImlkIjoiNjFkZTUyYjAyMWU0ODA5MWEwYjQ3YTU2IiwiZmVhdHVyZVR5cGUiOiJJTlRFUkNPTSJ9",
//                         "type": "INTERCOM",
//                         "featureId": "61de52b021e48091a0b47a56",
//                         "status": "ACTIVE",
//                         "__typename": "Feature"
//                     }
//                 ],
//                 "user": {
//                     "customerType": "RESIDENTIAL",
//                     "__typename": "DigitalUser"
//                 },
//                 "__typename": "DigitalViewer"
//             },
//             "__typename": "Viewer"
//         }
//     }
// }
