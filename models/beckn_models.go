package models

import "github.com/pnaskardev/brpl-bpp/models/base_model"

// {
//   "context": {
//     "domain": "dent:ev-charging:0.1.0",
//     "action": "select",
//     "location": {
//       "country": {
//         "code": "IND"
//       },
//       "city": {
//         "code": "std:080"
//       }
//     },
//     "version": "1.1.0",
//     "bap_id": "example-bap.com",
//     "bap_uri": "https://api.example-bap.com/pilot/bap/energy/v1",
//     "bpp_id": "chargezone-energy-bpp.com",
//     "bpp_uri": "https://api.example-bpp.com/pilot/bpp/",
//     "transaction_id": "6743e9e2-4fb5-487c-92b7-13ba8018f176",
//     "message_id": "6743e9e2-4fb5-487c-92b7-13ba8018f176",
//     "timestamp": "2023-07-16T04:41:16Z"
//   },
//   "message": {
//     "order": {
//       "provider": {
//         "id": "chargezone.in"
//       },
//       "items": [
//         {
//           "id": "pe-charging-01",
//           "quantity": {
//             "selected": {
//               "measure": {
//                 "value": "4",
//                 "unit": "kWh"
//               }
//             }
//           },
//           "add_ons": [
//             {
//               "id": "pe-charging-01-addon-1"
//             }
//           ]
//         }
//       ],
//       "fulfillments": [
//         {
//           "id": "1"
//         }
//       ]
//     }
//   }
// }

type Message struct {
	Order Order `json:"order"`
}

type Order struct {
	Provider     Provider      `json:"provider"`
	Items        []Item        `json:"items"`
	Fulfillments []Fulfillment `json:"fulfillments"`
}

type Provider struct {
	ID string `json:"id"`
}

type Item struct {
	ID       string   `json:"id"`
	Quantity Quantity `json:"quantity"`
	AddOns   []AddOn  `json:"add_ons"`
}

type Quantity struct {
	Selected Selected `json:"selected"`
}

type Selected struct {
	Measure Measure `json:"measure"`
}

type Measure struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type AddOn struct {
	ID string `json:"id"`
}

type Fulfillment struct {
	ID string `json:"id"`
}

type SelectRequest struct {
	Context basemodel.Context `json:"context"`
	Messsage Message `json:"message"`
}