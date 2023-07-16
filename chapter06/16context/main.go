package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//withCancel()
	//withTimeout()
	//withValue()
	withDeadline()
}

func withDeadline() {
	now := time.Now()
	newtime := now.Add(1 * time.Second)
	ctx, _ := context.WithDeadline(context.TODO(), newtime)
	go tv(ctx)
	go mobile(ctx)
	go game(ctx)

	time.Sleep(3 * time.Second)
}
func tv(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关电视")
			return
		default:

		}
		fmt.Println("看电视")
		time.Sleep(300 * time.Millisecond)
	}
}

// todo something special
func mobile(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关手机")
			return
		default:
		}
		fmt.Println("玩手机")
		time.Sleep(300 * time.Millisecond)
	}
}
func game(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关游戏机")
			return
		default:
		}
		fmt.Println("玩游戏机")
		time.Sleep(300 * time.Millisecond)
	}
}

func withValue() {
	ctx := context.WithValue(context.TODO(), 1, "钱包")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("WithValue:1", ctx.Value(1))
		fmt.Println("WithValue:2", ctx.Value(2))
		fmt.Println("WithValue:3", ctx.Value(3))
		fmt.Println("WithValue:4", ctx.Value(4))
	}(ctx)
	goTOPapa(ctx)
	time.Sleep(2 * time.Second)
}
func goTOPapa(ctx context.Context) {
	ctx = context.WithValue(ctx, 2, "充电宝")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goTOPapa:1", ctx.Value(1))
		fmt.Println("goTOPapa:2", ctx.Value(2))
		fmt.Println("goTOPapa:3", ctx.Value(3))
		fmt.Println("goTOPapa:4", ctx.Value(4))
	}(ctx)
	goTOMama(ctx)
}
func goTOMama(ctx context.Context) {
	ctx = context.WithValue(ctx, 3, "小夹克")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goTOMama:1", ctx.Value(1))
		fmt.Println("goTOMama:2", ctx.Value(2))
		fmt.Println("goTOMama:3", ctx.Value(3))
		fmt.Println("goTOMama:4", ctx.Value(4))
	}(ctx)
	goTOGrandma(ctx)
}
func goTOGrandma(ctx context.Context) {
	ctx = context.WithValue(ctx, 4, "大苹果")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goTOGrandma:1", ctx.Value(1))
		fmt.Println("goTOGrandma:2", ctx.Value(2))
		fmt.Println("goTOGrandma:3", ctx.Value(3))
		fmt.Println("goTOGrandma:4", ctx.Value(4))
	}(ctx)
	goToParty(ctx)
}
func goToParty(ctx context.Context) {
	fmt.Println("goTOGrandma:1", ctx.Value(1))
	fmt.Println("goTOGrandma:2", ctx.Value(2))
	fmt.Println("goTOGrandma:3", ctx.Value(3))
	fmt.Println("goTOGrandma:4", ctx.Value(4))

}

func withTimeout() {
	ctx, _ := context.WithTimeout(context.TODO(), 1*time.Second)
	fmt.Println("开始部署望远镜，发送信号")
	distributeMainFrame(ctx)
	distributeMainBody(ctx)
	distributeMainCover(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("任务超时")
	}
	time.Sleep(20 * time.Second) //等待20秒后收到任务取消的消息

}
func distributeMainFrame(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeMainFrame")
		return
	default:
	}
	fmt.Println("部署:distributeMainFrame")
}
func distributeMainBody(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeMainBody")
		return
	default:
	}
	fmt.Println("部署:distributeMainBody")
}

func distributeMainCover(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeMainCover")
		return
	default:
	}
	fmt.Println("部署:distributeMainCover")
}
func withCancel() {
	ctx := context.TODO()
	//context.Background()
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("做蛋挞，买材料")
	go buyFlour(ctx)
	go buyOil(ctx)
	go buyEgg(ctx)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("没电了取消购买所有材料")
	cancel() //当调用cancel后，所有由此上下文衍生出的context都会cancel
	time.Sleep(1 * time.Second)

}
func buyFlour(ctx context.Context) {
	fmt.Println("去买面")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done(): //todo 介绍一下struct{}
		fmt.Println("收到消息，不买面了")
		return
	default:

	}
	fmt.Println("去买面")

}
func buyOil(ctx context.Context) {
	fmt.Println("去买油")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买油了")
		return
	default:

	}
	fmt.Println("去买面油")

}
func buyEgg(ctx1 context.Context) {
	ctx, _ := context.WithCancel(ctx1)
	//defer cannel()//无论是否调用衍生出来得ctx的cannel，Done返回的Cannel都会关闭
	fmt.Println("去买蛋")
	//time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了")
		return
	default:
		fmt.Println("去买蛋")
		go buySEgg(ctx)
		go buyBEgg(ctx)
	}

}
func buySEgg(ctx context.Context) {
	fmt.Println("去买蛋buySEgg")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了buySEgg")
		return
	default:
	}
	fmt.Println("去买蛋buySEgg")
}
func buyBEgg(ctx context.Context) {
	fmt.Println("去买蛋buyBEgg")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了buyBEgg")
		return
	default:
	}
	fmt.Println("去买蛋buyBEgg")
}
