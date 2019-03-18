package main

import (
	"fmt"
	"os/exec"
	"os"
	"bufio"
	"regexp"
)

func read_text(){
	filename := "tmp.txt"

	fp, err := os.Open(filename)
	if err != nil{
		fmt.Println("File not found.")
	}

	defer fp.Close()

	//r := regexp.MustCompile("ちょる")
	//ツイートの先頭かどうかの確認
	r2 := regexp.MustCompile("1107")
	//ツイートのユーザー名の確認
	g_name := regexp.MustCompile("s-now")
	c_you := regexp.MustCompile("youtube.com")
	youtube := regexp.MustCompile(".+(youtube.+)")

	//現在の状態を確認するための数字
	c_num := 0

	scanner := bufio.NewScanner(fp)
	for scanner.Scan(){
		//fmt.Println(scanner.Text())
		if c_num == 0{
			//冒頭に数字があるか確認することでツイートの先頭部分かどうかを確認する
			check := r2.MatchString(scanner.Text())
			if  check == true{
				//ツイートの先頭部分に指定名があるかどうかを確認(これがユーザー名だと仮定する,
				//[]内の文字に指定していないのでミスる可能性あり)
				check_name := g_name.MatchString(scanner.Text())
				if check_name == true{
					fmt.Println(scanner.Text())
					//ツイート取得中
					c_num = 1
				}
			}
		} else {
			//ツイート内容の取得
			head_check := r2.MatchString(scanner.Text())
			if head_check == false {
				//中にyoutubeリンクがあったら反応
				check_you := c_you.MatchString(scanner.Text())
				if check_you == true{
					//youtubeリンクを取り出す
					string_byte := []byte(scanner.Text())
					//group := youtube.FindSubmatch(scanner.Text())
					group := youtube.FindSubmatch(string_byte)
					youtube_url := string(group[1])
					fmt.Println(youtube_url)
				}
				//fmt.Println(scanner.Text())
			} else {
			//取得したツイートが先頭ツイートだった場合
			//中身の判定はできていない
				c_num = 0
			}
		}
	}

	return
}


func main(){
	err := exec.Command("sh","./line.sh").Run()

	if err != nil{
		fmt.Println("Command Exec Error.")
	}

	read_text()
}
