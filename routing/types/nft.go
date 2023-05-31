package types

type PledgeNgtReq struct {
	NftId    string `json:"nft_id"`
	Duration string `json:"duration"`
	Hash     string `json:"hash"`
	Chain    string `json:"chain"`
}
type PledgeNgtResp struct {
	Code int64 `json:"code"`
}