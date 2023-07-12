package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/12
  @desc: 策略模式

	例子： 游戏英雄拿着不同的武器去战斗，不同的武器就是不同的策略
**/

//武器策略（抽象的策略）
type WeaponStrategy interface {
	UseWeapon() //使用武器
}

//具体的策略
type Ak47 struct{}

func (a *Ak47) UseWeapon() {
	fmt.Println("使用ak47去战斗")
}

//具体的策略2
type Knife struct{}

func (k *Knife) UseWeapon() {
	fmt.Println("使用匕首去战斗")
}

//环境类(英雄)
type Hero struct {
	strategy WeaponStrategy //拥有一个抽象的策略
}

//设置策略
func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	//调用策略
	h.strategy.UseWeapon()
}

func main() {
	hero := new(Hero)
	//策略1
	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	//更换策略2
	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}
