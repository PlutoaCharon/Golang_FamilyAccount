package utils

import (
	"fmt"
)

type FamilyAccount struct {
	key     string  // 保存接受用户输入的选项
	loop    bool    // 控制是否退出for
	balance float64 // 定义账户的余额
	money   float64 // 每次收支的金额
	note    string  // 每次收支的说明
	flag    bool    // 记录是否有收支的行为
	details string  // 对收支进行拼接处理
}

//构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {

	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说明",
	}

}

// 显示明细
func (this *FamilyAccount) showDetails() {
	fmt.Println("----------------------当前收支明细记录----------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细,来一笔吧!")
	}
}

// 登记记录
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money // 修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 登记支出
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)

	if this.money > this.balance {
		fmt.Println("余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 退出系统
func (this *FamilyAccount) exit() {
	fmt.Println("确认退出吗? 确定退出请输入y 否则输入n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" {
			this.loop = false
			break
		} else if choice == "n" {
			break
		}
		fmt.Println("输入有误,请重新输入 y/n")
	}
}

// 显示主菜单
func (this *FamilyAccount) MainMenu() {

	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项..")
		}

		if !this.loop {
			break
		}
	}

}
