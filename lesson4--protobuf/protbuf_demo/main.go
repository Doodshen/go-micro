/*
 * @Author: Github Doodshen Github 2475169766@qq.com
 * @Date: 2024-03-02 14:04:50
 * @LastEditors: Github Doodshen Github 2475169766@qq.com
 * @LastEditTime: 2024-03-02 14:30:56
 * @FilePath: \go-micro\lesson4--protobuf\protbuf_demo\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"protobuf_demo/api"
)

// oneof字段演示
func oneofDemo() {
	//client
	req1 := &api.NoticeReaderRequest{
		Msg: "kongshen更新了",
		NoticeWay: &api.NoticeReaderRequest_Email{
			Email: "111111111", //这个位置只能给noticeway指定一个值
		},
	}

	//server
	req := req1
	switch v := req.NoticeWay.(type) { //类型断言
	case *api.NoticeReaderRequest_Email:
		noticeWithEmail(v) //发送邮件通知
	case *api.NoticeReaderRequest_Phone:
		noticeWithPhone(v)
	}
}

// 发通知相关的功能函数
func noticeWithEmail(in *api.NoticeReaderRequest_Email) {
	fmt.Printf("notice reader by email:%v", in.Email)
}
func noticeWithPhone(in *api.NoticeReaderRequest_Phone) {
	fmt.Printf("notice reader by email:%v", in.Phone)
}

func main() {
	oneofDemo()
}
