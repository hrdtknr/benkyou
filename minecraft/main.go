package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

// 作りたい座標の左下を入力
// 片方にしか出力できない
const (
	X = 300// 70
	Y = 100// 75
	Z = 0// 100
)


func main(){
	readCSV()
	fmt.Println("-----")
	//compression()
	makeStairs()
}

/*
「csvを読み込んで、ドット絵を生成するfillコマンドを生成する」コードを作る

/fill 0 0 0 10 10 10 minecraft:stone replace minecraft:air
*/

func readCSV(){
	lines := useBufioScanner("src/minecraftDot.csv")


	height := len(lines)
	for i := 1; i < height; i++{ // index 0 はcsvの座標情報なので不要
		if len(lines[i]) == 0{ // ブロック情報がない行が出現したら、それ以降のデータは使わないので終了
			break
		}
		makeFillCommand(lines[i], height)
		fmt.Println()
	}
}

// 読み込んだcsv1行のデータを生成するfillコマンドを返す
func makeFillCommand(line string, height int){

	// カンマ区切りで分割
	// ex) 1,空気,空気,空気,空気,空気,空気,空気,... => 1 空気 空気　...
	var splitComma = strings.Split(line, ",")
	//	fmt.Println(splitComma[1])
	//	fmt.Println(len(splitComma))

	// 座標の始点
	var l int
	l, _ = strconv.Atoi(splitComma[0]) // y座標の始点に行数加えて調整するために必要
	adjustmentY := Y + height - l // 実際に作成するY座標に、作成する物の高さを加え、作成する行を引き、アイテムを配置する高さを決定する ex) 100 + 80 - 1 = 179
	adjustmentZ := 1 // 変数名がイケテナイ
	x := getItoa(X,0) // 1段つくるとき、xとy座標は固定
	y := getItoa(adjustmentY,0)


	// 先頭のブロック情報
	firstBlock := splitComma[1]
	var fillCommand []string

	var zStart, zEnd string


	for i := 2; i < len(splitComma); i++{
		// 違うブロックが出てきたら、ここまでをfill commandとして生成する
		if firstBlock != splitComma[i] {

			// fmt.Println("異なるブロック i = ", i)
			// fmt.Println("直前のブロック",firstBlock)
			// fmt.Println("現在のブロック",splitComma[i])

			// 同じブロックが連続する始点と終点
			zStart =getItoa(adjustmentZ,Z) // 実際に配置する座標を補正
			zEnd =getItoa(i-1,Z) // i に異なるブロックが出現しているので、その手前までが同じブロック


			// fmt.Println("zStart = ",zStart)
			// fmt.Println("zEnd = ",zEnd)
			// fmt.Println("adjustmentZ = ",adjustmentZ)

			// TODO:
			// 現状はz軸方向にのみブロックを置く実装なので、自由に調整できるようにしたい
			cmd := "/fill " + /*始点座標*/x + " " + y + " "+ zStart + " " + /*終点座標*/x + " " + y + " "+ zEnd + " " + getBlockData(firstBlock) + " replace air"
			fillCommand = append(fillCommand , cmd)

			// fmt.Println("----------------out fill command------------")
			// fmt.Println(cmd)
			// fmt.Println("----------------out fill command------------")

			// 現在のブロック種類を更新
			// fmt.Println("before:",firstBlock)
			firstBlock = splitComma[i]
			// fmt.Println("after:", firstBlock)


			// z座標を更新
			adjustmentZ = i // i に異なるブロックが出現しているので、次の始点は i
			//fmt.Println("次のz座標始点を更新")
			//fmt.Println("adjustmentZ = ",adjustmentZ)

		}
	}

	// 最後のコマンド生成
	cmd := "/fill " + /*始点座標*/x + " " + y + " "+ getItoa(adjustmentZ,Z) + " " + /*終点座標*/x + " " + y + " "+ getItoa(len(splitComma),Z-1) + " " + getBlockData(firstBlock) + " replace air"
	fillCommand = append(fillCommand , cmd)

	// test output
	for _,v := range fillCommand{
	fmt.Println(v)
	}

}

// TODO: コマンドで実行するなら、空気は空気でコマンド出力しない方がよい
// csvのブロックテキストをマイクラが読めるデータへ変換する
// 参考: https://minecraft-blog.net/?p=3746
func getBlockData(blockName string)string{
	switch blockName {
	case "空気", "白色のテラコッタ":
		return "concrete"
	case "桃色のコンクリート", "桃色の羊毛":
		return "pink_concrete"
		default:
			fmt.Println("変換未対応データあり")
			return blockName // 変換できないパターン
	}
}


func getItoa(i,j int)string{
	return strconv.Itoa(i+j)
}

// ファイル読み込み
func useBufioScanner(fileName string) []string{
    fp, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)
	lines := make([]string,0,100)
    for scanner.Scan() {
       lines = append(lines, scanner.Text())
    }

return lines
}

/*
// ズレた後の修正が大変なので、今回はこのままで
// 同ブロックを連続で配置するケースで、コマンドが分割されているものを圧縮する
// とりあえずz軸のみ対応
// ex)
// /fill 300 147 89 300 147 89 pink_concrete replace air
// /fill 300 147 90 300 147 93 pink_concrete replace air
// => /fill 300 147 89 300 147 93 pink_concrete replace air
func compression(){
	lines := useBufioScanner("result2.txt")
	for _,v := range lines {
		fmt.Println(v)
	}
}
*/

/*
/fill 170 13 -124 180 13 -124 pink_glazed_terracotta 
*/

func makeStairs(){
	// obelisk前階段　下
	//	y := 13
	//	for z := -124; z > -150; z--{
	//		fmt.Printf("/fill 170 %d %d 180 %d %d pink_glazed_terracotta\n", y,z,y,z)
	//		y++
	//	}

	y := 38
	for z := -147; z < -100; z++{
		fmt.Printf("/fill 181 %d %d 200 %d %d pink_glazed_terracotta\n", y,z,y,z)
		y++
	}


	y := 70
	for z := -147; z < -100; z++{
		fmt.Printf("/fill 181 %d %d 200 %d %d pink_glazed_terracotta\n", y,z,y,z)
		y--
	}
}