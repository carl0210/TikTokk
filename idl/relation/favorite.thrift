namespace go tiktokk.relation.favorite

struct FavoriteActionReq{
    1: i64 VideoID
    2: i64 UserID
    3: i8 ActionType
}

struct FavoriteActionResp{
    1: i64 StatusCode
    2: i64 StatusMsg
}

service FavoriteService{
    FavoriteActionResp FavoriteAction(1: FavoriteActionReq req)
}
