// package main

// import (
// 	"fmt"
// 	"io/ioutil"

// 	"gopkg.in/yaml.v2"
// )

// type Mysql struct {
// 	Host     string `yaml:"host"`
// 	Name     string `yaml:"name"`
// 	Username string `yaml:"username"`
// 	Password string `yaml:"password"`
// }

// type Redis struct {
// 	Host string `yaml:"host"`
// 	Port string `yaml:"port"`
// }

// type Config struct {
// 	Mysql
// 	Redis
// }

// func main() {
// 	config, err := ioutil.ReadFile("config.yaml")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	aa := &Config{}
// 	if err = yaml.Unmarshal(config, &aa); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(aa.Mysql.Host)
// }

package main

func asteroidCollision(asteroids []int) []int {
	dd := func(n int) int {
		if n <= 0 {
			return -n
		}
		return n
	}
	if len(asteroids) <= 1 {
		return asteroids
	}
	arr := []int{asteroids[0]}
	for i := 1; i < len(asteroids); i++ {

		if arr[len(arr)-1] >= 0 && asteroids[i] >= 0 || arr[len(arr)-1] <= 0 && asteroids[i] <= 0 { // 方向相同
			arr = append(arr, asteroids[i])
		} else {
			if dd(arr[len(arr)-1]) == dd(asteroids[i]) {
				arr = arr[:len(arr)-1]
			} else if dd(arr[len(arr)-1]) < dd(asteroids[i]) {
				for 
			}
		}
	}
	return arr
}

func asteroidCollision(asteroids []int) (st []int) {
    for _, aster := range asteroids {
        alive := true
        for alive && aster < 0 && len(st) > 0 && st[len(st)-1] > 0 {
            alive = st[len(st)-1] < -aster // aster 是否存在
            if st[len(st)-1] <= -aster {   // 栈顶行星爆炸
                st = st[:len(st)-1]
            }
        }
        if alive {
            st = append(st, aster)
        }
    }
    return
}

