package basemodel

//{
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

type Country struct {
	Code string `json:"code"`
}

type City struct {
	Code string `json:"code"`
}

type Location struct {
	Country Country `json:"country"`
	City City `json:"city"`
}

type Context struct {
	Domain string `json:"domain"`
	Action string `json:"action"`
	Location string `json:"location"`
	Version string `json:"version"`
	BAP_ID string `json:"bap_id"`
	BAP_URI string `json:"bap_uri"`
	BPP_ID string `json:"bpp_id"`
	BPP_URI string `json:"bpp_uri"`
	TransactionID string `json:"transaction_id"`
	MessageID string `json:"message_id"`
	Timestamp string `json:"timestamp"`
}