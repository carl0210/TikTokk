namespace go tiktokk.video

struct VideoPublishReq{
    1: required i64 UserID
    2: required string title
    3: required binary data
}

struct VideoPublishResp{
    1: required i64 Code,
    2: optional string msg
}

service VideoService{
    VideoPublishResp VideoPublish(1: VideoPublishReq req)
}