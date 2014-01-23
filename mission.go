package kcsapi

type MissionStartData struct {
    ApiResult    int16       `json:"api_result"`
    ApiResultMsg string      `json:"api_result_msg"`
    ApiData      missionData `json:"api_data"`
}

type missionData struct {
    Complatetime    int64  `json:"api_complatetime"`
    ComplatetimeStr string `json:"api_complatetime_str"`
}
