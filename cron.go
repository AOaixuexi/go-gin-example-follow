package main

// import (
//     "time"
//     "log"

//     "github.com/robfig/cron"

//     "github.com/AOaixuexi/go-gin-example-follow/models"
// )

// // 定时删除
// func main() {
//     log.Println("Starting...")

//     c := cron.New()
//     c.AddFunc("* * * * * *", func() {
//         log.Println("Run models.CleanAllTag...")
//         models.CleanAllTag()
//     })
//     c.AddFunc("* * * * * *", func() {
//         log.Println("Run models.CleanAllArticle...")
//         models.CleanAllArticle()
//     })

//     c.Start()

//     t1 := time.NewTimer(time.Second * 10)
//     for {
//         select {
//         case <-t1.C:
//             t1.Reset(time.Second * 10)
//         }
//     }
// }