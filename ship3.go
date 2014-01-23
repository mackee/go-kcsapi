package kcsapi

type Ship3Data struct {
    ApiResult    int16       `json:"api_result"`
    ApiResultMsg string      `json:"api_result_msg"`
    ApiData      shipSection `json:"api_data"`
}

type shipSection struct {
    ShipData []shipData  `json:"api_ship_data"`
    DeckData []deckData  `json:"api_deck_data"`
    SlotData interface{} `json:"api_slot_data"`
}

type shipData struct {
    Id        int64   `json:"api_id"`
    Sotrno    int64   `json:"api_sortno"`
    ShipId    int64   `json:"api_ship_id"`
    Lv        int32   `json:"api_lv"`
    Exp       []int64 `json:"api_exp"`
    Nowhp     int32   `json:"api_nowhp"`
    Maxhp     int32   `json:"api_maxhp"`
    Leng      int32   `json:"api_leng"`
    Slot      []int32 `json:"api_slot"`
    Onslot    []int32 `json:"api_onslot"`
    Kyouka    []int32 `json:"api_kyouka"`
    Backs     int32   `json:"api_backs"`
    Fuel      int32   `json:"api_fuel"`
    Bull      int32   `json:"api_bull"`
    Slotnum   int32   `json:"api_slotnum"`
    NdockTime int64   `json:"api_ndock_time"`
    NdockItem []int32 `json:"api_ndock_item"`
    Srate     int32   `json:"api_srate"`
    Cond      int32   `json:"api_cond"`
    Karyoku   []int32 `json:"api_karyoku"`
    Raisou    []int32 `json:"api_raisou"`
    Taiku     []int32 `json:"api_taiku"`
    Soukou    []int32 `json:"api_soukou"`
    Kaihi     []int32 `json:"api_kaihi"`
    Taisen    []int32 `json:"api_taisen"`
    Sakuteki  []int32 `json:"api_sakuteki"`
    Lucky     []int32 `json:"api_lucky"`
    Locked    int32   `json:"api_locked"`
}
type deckData struct {
    MemberId int64   `json:"api_member_id"`
    Id       int16   `json:"api_id"`
    Name     string  `json:"api_name"`
    NameId   string  `json:"api_name_id"`
    Mission  []int64 `json:"api_mission"`
    Flagship string  `json:"api_flagship"`
    Ship     []int64 `json:"api_ship"`
}
