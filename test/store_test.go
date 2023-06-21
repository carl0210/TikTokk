package test

import (
	"testing"
)

//func BenchmarkWithLen(b *testing.B) {
//	//初始化
//	tools.InitConfigTest()
//	tools.InitMysqlTest()
//	tools.InitStore()
//	tools.InitFeedlen()
//	ctx := context.Background()
//	b.ResetTimer()
//	//查询数据库信息
//	u, err := store.S.Users().GetByID(ctx, 3)
//	if err != nil {
//		b.Failed()
//	}
//	_, err = store.S.Videos().ListAllVideoByAuthorIDLen(ctx, u.UserId, int(u.WorkCount))
//	if err != nil {
//		b.Failed()
//	}
//
//}
//
//func BenchmarkWithoutLen(b *testing.B) {
//	//初始化
//	tools.InitConfigTest()
//	tools.InitMysqlTest()
//	tools.InitStore()
//	tools.InitFeedlen()
//	ctx := context.Background()
//	b.ResetTimer()
//	_, err := store.S.Videos().ListAllVideoByAuthorID(ctx, 3)
//	if err != nil {
//		b.Failed()
//	}
//}

func add1(s []int, l int) {
	for i := 0; i < l; i++ {
		s = append(s, i)
	}
}

func add2(s []int, l int) {
	for i := 0; i < l; i++ {
		//s = append(s, i)
		s[i] = i
	}
}

func BenchmarkWithLen(b *testing.B) {
	s := make([]int, 1000000)
	add2(s, 1000000)
}

func BenchmarkWithoutLen(b *testing.B) {
	var s []int
	add1(s, 1000000)
}
