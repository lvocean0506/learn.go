package main

import (
	"fmt"
	"learn.go/chapter02/015.faterate/calc"
)

func main() {
	for {
		mainFatRateBody()
		if cont := whetherContinue(); !cont {
			break
		}
	}
	fmt.Println("修复bug")

	fmt.Println("开发到了100%")

}

func mainFatRateBody() {
	//录入各项
	weight, tall, age, _, sex := getMaterialsFromInput()

	//计算体脂率

	fateRate := calcFateRate(weight, tall, age, sex)
	if fateRate <= 0 {
		panic("fat rate is not allowed to be negative")
	}
	//给出建议
	getHealthinessSuggestions(sex, age, fateRate)
}

func getHealthinessSuggestions(sex string, age int, fatRate float64) {
	//给出建议
	if sex == "男" {
		//  编写男性的体脂率与体脂状态表
		if age >= 18 && age <= 39 {
			// 对该年龄段的体脂率进行if、else判断
			if fatRate <= 0.1 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏瘦")
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				// 得出结论，并进行建议
				fmt.Println("目前是：标准")
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				// 得出结论，并进行建议
				fmt.Println("目前是：肥胖")
			} else {
				// 得出结论，并进行建议
				fmt.Println("目前是：非常肥胖")
			}
		} else if age >= 40 && age <= 59 {
			// 对该年龄段的体脂率进行if、else判断
			if fatRate <= 0.11 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏瘦")
			} else if fatRate > 0.11 && fatRate <= 0.17 {
				// 得出结论，并进行建议
				fmt.Println("目前是：标准")
			} else if fatRate > 0.17 && fatRate <= 0.22 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.22 && fatRate <= 0.27 {
				// 得出结论，并进行建议
				fmt.Println("目前是：肥胖")
			} else {
				// 得出结论，并进行建议
				fmt.Println("目前是：非常肥胖")
			}
		} else if age >= 60 {
			// 对该年龄段的体脂率进行if、else判断
			if fatRate <= 0.13 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏瘦")
			} else if fatRate > 0.13 && fatRate <= 0.19 {
				// 得出结论，并进行建议
				fmt.Println("目前是：标准")
			} else if fatRate > 0.19 && fatRate <= 0.24 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.24 && fatRate <= 0.29 {
				// 得出结论，并进行建议
				fmt.Println("目前是：肥胖")
			} else {
				// 得出结论，并进行建议
				fmt.Println("目前是：非常肥胖")
			}
		} else {
			fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
		}
	} else {
		// 编写女性的体脂率与体脂状态表
		if age >= 18 && age <= 39 {
			// 对该年龄段的体脂率进行if、else判断
			if fatRate <= 0.2 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏瘦")
			} else if fatRate > 0.2 && fatRate <= 0.27 {
				// 得出结论，并进行建议
				fmt.Println("目前是：标准")
			} else if fatRate > 0.27 && fatRate <= 0.34 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.34 && fatRate <= 0.39 {
				// 得出结论，并进行建议
				fmt.Println("目前是：肥胖")
			} else {
				// 得出结论，并进行建议
				fmt.Println("目前是：非常肥胖")
			}
		} else if age >= 40 && age <= 59 {
			// 对该年龄段的体脂率进行if、else判断
			if fatRate <= 0.21 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏瘦")
			} else if fatRate > 0.21 && fatRate <= 0.28 {
				// 得出结论，并进行建议
				fmt.Println("目前是：标准")
			} else if fatRate > 0.28 && fatRate <= 0.35 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.35 && fatRate <= 0.40 {
				// 得出结论，并进行建议
				fmt.Println("目前是：肥胖")
			} else {
				// 得出结论，并进行建议
				fmt.Println("目前是：非常肥胖")
			}
		} else if age >= 60 {
			// 对该年龄段的体脂率进行if、else判断
			if fatRate <= 0.22 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏瘦")
			} else if fatRate > 0.22 && fatRate <= 0.29 {
				// 得出结论，并进行建议
				fmt.Println("目前是：标准")
			} else if fatRate > 0.29 && fatRate <= 0.36 {
				// 得出结论，并进行建议
				fmt.Println("目前是：偏胖")
			} else if fatRate > 0.36 && fatRate <= 0.41 {
				// 得出结论，并进行建议
				fmt.Println("目前是：肥胖")
			} else {
				// 得出结论，并进行建议
				fmt.Println("目前是：非常肥胖")
			}
		} else {
			fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
		}
	}
}

func calcFateRate(weight float64, tall float64, age int, sex string) float64 {
	//计算体脂率

	bmi := calc.CalcBMI(tall, weight)
	fateRate := calc.ClacFateRate(bmi, age, sex)
	fmt.Println("体脂率是", fateRate)
	return fateRate
}

func getMaterialsFromInput() (float64, float64, int, int, string) {
	//录入各项
	var name string
	fmt.Print("姓名: ")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重（千克）：") //辅助信息
	fmt.Scanln(&weight)  //录入命令

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	var sexWeight int
	sex := "男"
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	return weight, tall, age, sexWeight, sex
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Print("是否录入下一个（y/n）?")
	fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		return false
	}
	return true
}
