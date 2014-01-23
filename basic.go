package kcsapi

type BasicData struct {
    ApiResult    int16   `json:"api_result"`
    ApiResultMsg string  `json:"api_result_msg"`
    ApiData      apiData `json:"api_data"`
}

type apiData struct {
    MemberId         string  `json:"api_member_id"`
    Nickname         string  `json:"api_nickname"`
    NicknameId       string  `json:"api_nickname_id"`
    ActiveFlag       int16   `json:"api_active_flag"`
    Starttime        int64   `json:"api_starttime"`
    Level            int32   `json:"api_level"`
    Rank             int32   `json:"api_rank"`
    Experience       int64   `json:"api_experience"`
    Fleetname        string  `json:"api_fleetname"`
    Comment          string  `json:"api_comment"`
    CommentId        int64   `json:"api_comment_id"`
    MacChara         int32   `json:"api_max_chara"`
    MaxSlotitem      int32   `json:"api_max_slotitem"`
    MaxKagu          int32   `json:"api_max_kagu"`
    Playtime         int32   `json:"api_playtime"`
    Tutorial         int16   `json:"api_tutorial"`
    Furniture        []int32 `json:"api_furniture"`
    CountDeck        int32   `json:"api_count_deck"`
    CountKdock       int32   `json:"api_count_kdock"`
    CountNdock       int32   `json:"api_count_ndock"`
    Fcoin            int64   `json:"api_fcoin"`
    StWin            int64   `json:"api_st_win"`
    StLose           int64   `json:"api_st_lose"`
    MsCount          int64   `json:"api_ms_count"`
    MsSuccess        int64   `json:"api_ms_success"`
    PtWin            int64   `json:"api_pt_win"`
    PtLose           int64   `json:"api_pt_lose"`
    PtChallenged     int64   `json:"api_pt_challenged"`
    PtChallengedWin  int64   `json:"api_pt_challenged_win"`
    Firstflag        int16   `json:"api_firstflag"`
    TutorialProgress int16   `json:"api_tutorial_progress"`
    Pvp              []int32 `json:"api_pvp"`
}
