package yog

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func Test_init_taskfile(t *testing.T) {
	os.MkdirAll("./test/case1", os.ModePerm)
	os.MkdirAll("./test/case2", os.ModePerm)
	os.MkdirAll("./test/case3", os.ModePerm)
	os.MkdirAll("./test/case4", os.ModePerm)

	case1()
	case2()
	case3()
	case4()
}

func case1() {
	testString := "{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":223},\"distance\":{\"value\":2544}},{\"duration\":{\"value\":646},\"distance\":{\"value\":6473}},{\"duration\":{\"value\":718},\"distance\":{\"value\":8322}},{\"duration\":{\"value\":1116},\"distance\":{\"value\":14948}},{\"duration\":{\"value\":1325},\"distance\":{\"value\":20294}},{\"duration\":{\"value\":1218},\"distance\":{\"value\":16527}},{\"duration\":{\"value\":1722},\"distance\":{\"value\":25780}},{\"duration\":{\"value\":1594},\"distance\":{\"value\":25298}},{\"duration\":{\"value\":1385},\"distance\":{\"value\":22240}},{\"duration\":{\"value\":981},\"distance\":{\"value\":13268}}]},{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]},{\"elements\":[{\"duration\":{\"value\":956},\"distance\":{\"value\":16748}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":20718}},{\"duration\":{\"value\":1389},\"distance\":{\"value\":25208}},{\"duration\":{\"value\":1649},\"distance\":{\"value\":26520}},{\"duration\":{\"value\":1893},\"distance\":{\"value\":30988}},{\"duration\":{\"value\":1596},\"distance\":{\"value\":25942}},{\"duration\":{\"value\":2100},\"distance\":{\"value\":35194}},{\"duration\":{\"value\":1851},\"distance\":{\"value\":31655}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":26152}},{\"duration\":{\"value\":1193},\"distance\":{\"value\":17181}}]}]}"
	bytes, err := Encode(testString)
	if err != nil {
		return
	}

	if err := os.WriteFile("./test/case1/1", bytes, 0666); err != nil {
		log.Fatal(err)
	}

	var num []int
	for i := 0; i < 10; i++ {
		num = append(num, i)
	}
	var task TaskMeta
	task.Index = map[string]IndexItem{
		"1": {
			Origin:      num,
			Destination: num,
		},
	}

	marshal, err := json.Marshal(task)
	if err != nil {
		return
	}
	if err := os.WriteFile("./test/case1/"+META_FILE_NAME, marshal, 0666); err != nil {
		log.Fatal(err)
	}
}

func case2() {
	testString := "{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":223},\"distance\":{\"value\":2544}},{\"duration\":{\"value\":646},\"distance\":{\"value\":6473}},{\"duration\":{\"value\":718},\"distance\":{\"value\":8322}},{\"duration\":{\"value\":1116},\"distance\":{\"value\":14948}},{\"duration\":{\"value\":1325},\"distance\":{\"value\":20294}},{\"duration\":{\"value\":1218},\"distance\":{\"value\":16527}},{\"duration\":{\"value\":1722},\"distance\":{\"value\":25780}},{\"duration\":{\"value\":1594},\"distance\":{\"value\":25298}},{\"duration\":{\"value\":1385},\"distance\":{\"value\":22240}},{\"duration\":{\"value\":981},\"distance\":{\"value\":13268}}]},{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]},{\"elements\":[{\"duration\":{\"value\":956},\"distance\":{\"value\":16748}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":20718}},{\"duration\":{\"value\":1389},\"distance\":{\"value\":25208}},{\"duration\":{\"value\":1649},\"distance\":{\"value\":26520}},{\"duration\":{\"value\":1893},\"distance\":{\"value\":30988}},{\"duration\":{\"value\":1596},\"distance\":{\"value\":25942}},{\"duration\":{\"value\":2100},\"distance\":{\"value\":35194}},{\"duration\":{\"value\":1851},\"distance\":{\"value\":31655}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":26152}},{\"duration\":{\"value\":1193},\"distance\":{\"value\":17181}}]}]}"
	bytes, err := Encode(testString)
	if err != nil {
		return
	}

	if err := os.WriteFile("./test/case2/1", bytes, 0666); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./test/case2/2", bytes, 0666); err != nil {
		log.Fatal(err)
	}

	var num []int
	for i := 0; i < 10; i++ {
		num = append(num, i)
	}

	var num1 []int
	for i := 1; i < 20; i += 2 {
		num1 = append(num1, i)
	}

	var num2 []int
	for i := 0; i < 20; i += 2 {
		num2 = append(num2, i)
	}

	var task TaskMeta
	task.Index = map[string]IndexItem{
		"1": {
			Origin:      num,
			Destination: num1,
		},
		"2": {
			Origin:      num,
			Destination: num2,
		},
	}

	marshal, err := json.Marshal(task)
	if err != nil {
		return
	}
	if err := os.WriteFile("./test/case2/"+META_FILE_NAME, marshal, 0666); err != nil {
		log.Fatal(err)
	}
}

func case3() {
	// 2000 * 2000
	// -rw-r--r--@ 1 xurui  staff    31M Aug  3 15:48 test.bin
	// -rw-r--r--@ 1 xurui  staff   227M Aug  3 15:48 test.json

	marshal, err := json.Marshal(genMatrixResult(2000, 2000))
	if err != nil {
		return
	}
	if err := os.WriteFile("./test/case3/test_01.json", marshal, 0666); err != nil {
		log.Fatal(err)
	}

	bytes, err := Encode(string(marshal))
	if err != nil {
		return
	}
	if err := os.WriteFile("./test/case3/test_01.bin", bytes, 0666); err != nil {
		log.Fatal(err)
	}

	// 20000 * 20000
	// -rw-r--r--@ 1 xurui  staff   3.0G Aug  8 10:06 test_02.bin
	// -rw-r--r--@ 1 xurui  staff    24G Aug  8 09:49 test_02.json
	//marshal, err = json.Marshal(genMatrixResult(20000, 20000))
	//if err != nil {
	//	return
	//}
	//if err := os.WriteFile("./test/case3/test_02.json", marshal, 0666); err != nil {
	//	log.Fatal(err)
	//}
	//
	//bytes, err = Encode(string(marshal))
	//if err != nil {
	//	return
	//}
	//if err := os.WriteFile("./test/case3/test_02.bin", bytes, 0666); err != nil {
	//	log.Fatal(err)
	//}
}

func case4() {
	testString := "{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":223},\"distance\":{\"value\":2544}},{\"duration\":{\"value\":646},\"distance\":{\"value\":6473}},{\"duration\":{\"value\":718},\"distance\":{\"value\":8322}},{\"duration\":{\"value\":1116},\"distance\":{\"value\":14948}},{\"duration\":{\"value\":1325},\"distance\":{\"value\":20294}},{\"duration\":{\"value\":1218},\"distance\":{\"value\":16527}},{\"duration\":{\"value\":1722},\"distance\":{\"value\":25780}},{\"duration\":{\"value\":1594},\"distance\":{\"value\":25298}},{\"duration\":{\"value\":1385},\"distance\":{\"value\":22240}},{\"duration\":{\"value\":981},\"distance\":{\"value\":13268}}]},{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]},{\"elements\":[{\"duration\":{\"value\":956},\"distance\":{\"value\":16748}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":20718}},{\"duration\":{\"value\":1389},\"distance\":{\"value\":25208}},{\"duration\":{\"value\":1649},\"distance\":{\"value\":26520}},{\"duration\":{\"value\":1893},\"distance\":{\"value\":30988}},{\"duration\":{\"value\":1596},\"distance\":{\"value\":25942}},{\"duration\":{\"value\":2100},\"distance\":{\"value\":35194}},{\"duration\":{\"value\":1851},\"distance\":{\"value\":31655}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":26152}},{\"duration\":{\"value\":1193},\"distance\":{\"value\":17181}}]}]}"
	bytes, err := Encode(testString)
	if err != nil {
		return
	}

	if err := os.WriteFile("./test/case4/1", bytes, 0666); err != nil {
		log.Fatal(err)
	}

	//if err := os.WriteFile("./test/case4/2", bytes, 0666); err != nil {
	//	log.Fatal(err)
	//}

	var num []int
	for i := 0; i < 10; i++ {
		num = append(num, i)
	}

	var num1 []int
	for i := 1; i < 20; i += 2 {
		num1 = append(num1, i)
	}

	var num2 []int
	for i := 0; i < 20; i += 2 {
		num2 = append(num2, i)
	}

	var task TaskMeta
	task.Index = map[string]IndexItem{
		"1": {
			Origin:      num,
			Destination: num1,
		},
		"2": {
			Origin:      num,
			Destination: num2,
		},
	}

	marshal, err := json.Marshal(task)
	if err != nil {
		return
	}
	if err := os.WriteFile("./test/case4/"+META_FILE_NAME, marshal, 0666); err != nil {
		log.Fatal(err)
	}
}

func genMatrixResult(oc, dc int) MatrixData {
	var matrixData MatrixData
	matrixData.Rows = make([]MatrixRow, oc)
	i := 0
	for o := 0; o < oc; o++ {
		matrixData.Rows[o] = MatrixRow{make([]MatrixElement, dc)}
		for d := 0; d < dc; d++ {
			matrixData.Rows[o].Elements[d] = MatrixElement{
				Duration: Value{rand.Uint32()},
				Distance: Value{rand.Uint32()},
			}
			i++
		}
	}
	return matrixData
}

func TestSeDecode(t *testing.T) {

	testcases := []string{
		"{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":223},\"distance\":{\"value\":2544}},{\"duration\":{\"value\":646},\"distance\":{\"value\":6473}},{\"duration\":{\"value\":718},\"distance\":{\"value\":8322}},{\"duration\":{\"value\":1116},\"distance\":{\"value\":14948}},{\"duration\":{\"value\":1325},\"distance\":{\"value\":20294}},{\"duration\":{\"value\":1218},\"distance\":{\"value\":16527}},{\"duration\":{\"value\":1722},\"distance\":{\"value\":25780}},{\"duration\":{\"value\":1594},\"distance\":{\"value\":25298}},{\"duration\":{\"value\":1385},\"distance\":{\"value\":22240}},{\"duration\":{\"value\":981},\"distance\":{\"value\":13268}}]},{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]},{\"elements\":[{\"duration\":{\"value\":956},\"distance\":{\"value\":16748}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":20718}},{\"duration\":{\"value\":1389},\"distance\":{\"value\":25208}},{\"duration\":{\"value\":1649},\"distance\":{\"value\":26520}},{\"duration\":{\"value\":1893},\"distance\":{\"value\":30988}},{\"duration\":{\"value\":1596},\"distance\":{\"value\":25942}},{\"duration\":{\"value\":2100},\"distance\":{\"value\":35194}},{\"duration\":{\"value\":1851},\"distance\":{\"value\":31655}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":26152}},{\"duration\":{\"value\":1193},\"distance\":{\"value\":17181}}]}]}",
		"{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]},{\"elements\":[{\"duration\":{\"value\":956},\"distance\":{\"value\":16748}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":20718}},{\"duration\":{\"value\":1389},\"distance\":{\"value\":25208}},{\"duration\":{\"value\":1649},\"distance\":{\"value\":26520}},{\"duration\":{\"value\":1893},\"distance\":{\"value\":30988}},{\"duration\":{\"value\":1596},\"distance\":{\"value\":25942}},{\"duration\":{\"value\":2100},\"distance\":{\"value\":35194}},{\"duration\":{\"value\":1851},\"distance\":{\"value\":31655}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":26152}},{\"duration\":{\"value\":1193},\"distance\":{\"value\":17181}}]}]}",
		"{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]}]}",
	}

	for _, testString := range testcases {

		bytes, err := Encode(testString)
		if err != nil {
			return
		}

		matrixResponse, err := Decode(bytes)
		if err != nil {
			return
		}

		var m1 MatrixData
		err = json.Unmarshal([]byte(testString), &m1)
		if err != nil {
			return
		}

		var m2 MatrixData
		err = json.Unmarshal([]byte(matrixResponse), &m2)
		if err != nil {
			return
		}

		assert.Equal(t, m1, m2)
	}
}

func TestRead(t *testing.T) {

	y := New("", "./test/case1", "")
	err := y.ReloadMeta()
	if err != nil {
		assert.Equalf(t, err, nil, "yog reload meta failed %v", err)
	}

	type args struct {
		o int
		d int
	}
	tests := []struct {
		name         string
		args         args
		wantDuration uint32
		wantDistance uint32
		wantErr      bool
	}{
		{
			"test01",
			args{
				0, 0,
			},
			223, 2544,
			false,
		}, {
			"test02",
			args{
				0, 1,
			},
			646, 6473,
			false,
		}, {
			"test03",
			args{
				1, 0,
			},
			699, 6981,
			false,
		}, {
			"test04",
			args{
				1, 1,
			},
			436, 2684,
			false,
		}, {
			"test04",
			args{
				9, 9,
			},
			1193, 17181,
			false,
		},
	}

	for _, tt := range tests {
		gotDuration, gotDistance, err := y.Read(tt.args.o, tt.args.d)
		assert.Equalf(t, err != nil, tt.wantErr, "%v Error(%v, %v)", tt.name, tt.wantErr, err)
		assert.Equalf(t, tt.wantDuration, gotDuration, "%v Read(%v, %v)", tt.name, gotDuration, tt.wantDuration)
		assert.Equalf(t, tt.wantDistance, gotDistance, "%v Read(%v, %v)", tt.name, gotDistance, tt.wantDistance)
	}
}

func TestYog_ReadChunk(t *testing.T) {
	yogArray := []*Yog{
		New("", "./test/case1", ""),
		New("", "./test/case2", ""),
		New("", "./test/case4", ""),
	}

	for i := range yogArray {
		err := yogArray[i].ReloadMeta()
		if err != nil {
			assert.Equalf(t, err, nil, "yog reload meta failed %v", err)
		}
	}

	tests := []struct {
		name         string
		y            *Yog
		chunkSize    int
		wantDuration []uint32
		wantDistance []uint32
		wantErr      bool
	}{
		{

			"test01",
			yogArray[1],
			11,
			[]uint32{223, 223, 646, 646, 718, 718, 1116, 1116, 1325, 1325, 1218},
			[]uint32{2544, 2544, 6473, 6473, 8322, 8322, 14948, 14948, 20294, 20294, 16527},
			false,
		},
		{
			"test02",
			yogArray[2],
			11,
			[]uint32{0, 223, 0, 646, 0, 718, 0, 1116, 0, 1325, 0},
			[]uint32{0, 2544, 0, 6473, 0, 8322, 0, 14948, 0, 20294, 0},
			false,
		},
	}

	for _, tt := range tests {
		gotDuration, gotDistance, err := tt.y.ReadChunk(tt.chunkSize)
		assert.Equalf(t, err != nil, tt.wantErr, "%v Error(%v, %v)", tt.name, tt.wantErr, err)
		assert.Equalf(t, tt.wantDuration, gotDuration, "%v Read(%v, %v)", tt.name, gotDuration, tt.wantDuration)
		assert.Equalf(t, tt.wantDistance, gotDistance, "%v Read(%v, %v)", tt.name, gotDistance, tt.wantDistance)
	}
}

func TestYog_DiffTest(t *testing.T) {
	yogArray := []*Yog{
		New("", "./test/case1", ""),
		New("", "./test/case2", ""),
		New("", "./test/case4", ""),
	}

	for _, y := range yogArray {
		err := y.ReloadMeta()
		if err != nil {
			assert.Equal(t, err, nil)
			return
		}

		for o := 0; o < y.taskMeta.MatrixInfo.OriginCount; o++ {
			for d := 0; d < y.taskMeta.MatrixInfo.DestinationCount; d++ {
				durations, distances, err := y.ReadChunk(1)
				assert.Equal(t, err, nil)
				duration, distance, err := y.Read(o, d)
				assert.Equal(t, err, nil)
				assert.Equal(t, strconv.Itoa(int(durations[0])), strconv.Itoa(int(duration)))
				assert.Equal(t, strconv.Itoa(int(distances[0])), strconv.Itoa(int(distance)))
			}
		}
	}
}

func Test_findSubset(t *testing.T) {
	type args struct {
		min   int
		max   int
		array []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			"test01",
			args{2, 4, []int{0, 1, 2, 3, 4, 5}},
			2, 4,
		},
		{
			"test01",
			args{0, 4, []int{2, 3, 4, 5, 7, 8, 9}},
			0, 2,
		},
		{
			"test02",
			args{10, 11, []int{2, 3, 4, 5, 7, 8, 9}},
			-1, -1,
		},
		{
			"test03",
			args{5, 11, []int{2, 3, 4, 5, 7, 8, 9}},
			3, 6,
		},
		{
			"test04",
			args{0, 3, []int{4, 5, 7, 8, 9}},
			-1, -1,
		},
		{
			"test05",
			args{0, 11, []int{2, 4, 6, 8, 10, 12}},
			0, 4,
		},
	}
	for _, tt := range tests {
		got, got1 := findSubset(tt.args.min, tt.args.max, tt.args.array)
		assert.Equalf(t, tt.want, got, "findSubset(%v, %v, %v)", tt.name, tt.want, got)
		assert.Equalf(t, tt.want1, got1, "findSubset(%v, %v, %v)", tt.name, tt.want1, got1)
	}
}

//func Test_yog(t *testing.T) {
//
//	for i := 0; i < 5; i++ {
//		o_size, d_size := rand.Intn(2000)+1, rand.Intn(2000)+1
//
//		o_count := rand.Intn((o_size-1)/10) + 1
//		d_count := rand.Intn((d_size-1)/10) + 1
//
//		oChunkArray := split(o_size, o_count)
//		dChunkArray := split(d_size, d_count)
//
//		var num []int
//		for i := 0; i < 10; i++ {
//			num = append(num, i)
//		}
//		var task TaskMeta
//		task.Index = map[string]IndexItem{}
//
//		path := fmt.Sprintf("./test/random/case_%v", i)
//		os.MkdirAll(path, os.ModePerm)
//
//		chunkname := 0
//		oOffset := 0
//		dOffset := 0
//		for _, ol := range oChunkArray {
//			for _, dl := range dChunkArray {
//				task.Index[strconv.Itoa(chunkname)] = IndexItem{}
//				chunkname++
//				dOffset += dl
//			}
//			oOffset += ol
//		}
//
//		chunkname = 0
//		for _, ol := range oChunkArray {
//			for _, dl := range dChunkArray {
//				marshal, err := json.Marshal(genMatrixResult(ol, dl))
//				bytes, err := Encode(string(marshal))
//				if err != nil {
//					return
//				}
//
//				if err := os.WriteFile(path+"/"+strconv.Itoa(chunkname), bytes, 0666); err != nil {
//					log.Fatal(err)
//				}
//
//				chunkname++
//			}
//		}
//
//		marshal, err := json.Marshal(task)
//		if err != nil {
//			return
//		}
//		if err := os.WriteFile(path+"/"+META_FILE_NAME, marshal, 0666); err != nil {
//			log.Fatal(err)
//		}
//	}
//}
//
//func split(num, n int) []int {
//	// 计算每份的基准值
//	min := num / n / 5 * 4
//	base := (num)/n - min
//
//	// 生成n-1个随机数
//	randomNums := make([]int, n-1)
//	for i := 0; i < n-1; i++ {
//		randomNums[i] = rand.Intn(base) + min
//	}
//
//	sum := 0
//	for _, v := range randomNums {
//		sum += v
//	}
//
//	// 计算剩余的数字
//	remaining := num - sum
//
//	// 将剩余的数字作为第n份的值
//	randomNums = append(randomNums, remaining)
//	return randomNums
//}
