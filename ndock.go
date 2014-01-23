package kcsapi

type NdockData struct {
    ApiResult    int16          `json:"api_result"`
    ApiResultMsg string         `json:"api_result_msg"`
    ApiData      []ndockElement `json:"api_data"`
}

type ndockElement struct {
    MemberId        int64  `json:"api_member_id"`
    Id              int16  `json:"api_id"`
    State           int16  `json:"api_state"`
    ShipId          int64  `json:"api_ship_id"`
    CompleteTime    int64  `json:"api_complete_time"`
    CompleteTimeStr string `json:"api_complete_time_str"`
    Item1           int16  `json:"api_item1"`
    Item2           int16  `json:"api_item2"`
    Item3           int16  `json:"api_item3"`
    Item4           int16  `json:"api_item4"`
}
