package yog

import (
	"encoding/json"
	"errors"
	"fmt"
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
	os.MkdirAll("./test/case5", os.ModePerm)
	os.MkdirAll("./test/case6", os.ModePerm)

	case1()
	case2()
	case3()
	case4()
	case5()
	case6()
}

func TestSeDecode(t *testing.T) {

	testcases := []string{
		"{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":223},\"distance\":{\"value\":2544}},{\"duration\":{\"value\":646},\"distance\":{\"value\":6473}},{\"duration\":{\"value\":718},\"distance\":{\"value\":8322}},{\"duration\":{\"value\":1116},\"distance\":{\"value\":14948}},{\"duration\":{\"value\":1325},\"distance\":{\"value\":20294}},{\"duration\":{\"value\":1218},\"distance\":{\"value\":16527}},{\"duration\":{\"value\":1722},\"distance\":{\"value\":25780}},{\"duration\":{\"value\":1594},\"distance\":{\"value\":25298}},{\"duration\":{\"value\":1385},\"distance\":{\"value\":22240}},{\"duration\":{\"value\":981},\"distance\":{\"value\":13268}}]},{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]},{\"elements\":[{\"duration\":{\"value\":956},\"distance\":{\"value\":16748}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":20718}},{\"duration\":{\"value\":1389},\"distance\":{\"value\":25208}},{\"duration\":{\"value\":1649},\"distance\":{\"value\":26520}},{\"duration\":{\"value\":1893},\"distance\":{\"value\":30988}},{\"duration\":{\"value\":1596},\"distance\":{\"value\":25942}},{\"duration\":{\"value\":2100},\"distance\":{\"value\":35194}},{\"duration\":{\"value\":1851},\"distance\":{\"value\":31655}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":26152}},{\"duration\":{\"value\":1193},\"distance\":{\"value\":17181}}]}]}",
		"{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]},{\"elements\":[{\"duration\":{\"value\":956},\"distance\":{\"value\":16748}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":20718}},{\"duration\":{\"value\":1389},\"distance\":{\"value\":25208}},{\"duration\":{\"value\":1649},\"distance\":{\"value\":26520}},{\"duration\":{\"value\":1893},\"distance\":{\"value\":30988}},{\"duration\":{\"value\":1596},\"distance\":{\"value\":25942}},{\"duration\":{\"value\":2100},\"distance\":{\"value\":35194}},{\"duration\":{\"value\":1851},\"distance\":{\"value\":31655}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":26152}},{\"duration\":{\"value\":1193},\"distance\":{\"value\":17181}}]}]}",
		"{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]}]}",
	}

	for _, testString := range testcases {

		bytes, err := encode(testString)
		if err != nil {
			return
		}

		matrixResponse, err := decode(bytes)
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
		wantDuration int32
		wantDistance int32
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
		wantDuration []int32
		wantDistance []int32
		wantErr      bool
	}{
		{

			"test01",
			yogArray[1],
			11,
			[]int32{223, 223, 646, 646, 718, 718, 1116, 1116, 1325, 1325, 1218},
			[]int32{2544, 2544, 6473, 6473, 8322, 8322, 14948, 14948, 20294, 20294, 16527},
			false,
		},
		{
			"test02",
			yogArray[2],
			11,
			[]int32{0, 223, 0, 646, 0, 718, 0, 1116, 0, 1325, 0},
			[]int32{0, 2544, 0, 6473, 0, 8322, 0, 14948, 0, 20294, 0},
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
		New("", "./test/case5", ""),
		New("", "./test/case6", ""),
	}

	for _, y := range yogArray {
		err := y.ReloadMeta()
		if err != nil {
			assert.Equal(t, err, nil)
			return
		}

		if y.path == "./test/case5/" {
			fmt.Println("sss")
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

// encode MatrixInfo Result -> binary i = o * d.length + d
func encode(data string) ([]byte, error) {

	var serializer Int32BinarySerializer

	var resp MatrixData
	err := json.Unmarshal([]byte(data), &resp)
	if err != nil {
		return nil, err
	}

	header, err := serializer.encode(int32(len(resp.Rows)), int32(len(resp.Rows[0].Elements)))
	if err != nil {
		return nil, err
	}

	// 将 MatrixData.Rows 转化成 binary
	res := make([]byte, 0)

	// add header
	res = append(res, header...)

	// source index
	for _, row := range resp.Rows {
		// destination index
		for _, element := range row.Elements {
			chunk, err := serializer.encode(element.Duration.Value, element.Distance.Value)
			if err != nil {
				return nil, err
			}
			res = append(res, chunk...)
		}
	}
	return res, nil
}

// decode binary -> MatrixInfo Result
func decode(bin []byte) (string, error) {

	var serializer Int32BinarySerializer

	// check data
	if len(bin)%8 > 0 {
		return "", errors.New("illegal binary data format")
	}

	//read header
	sourceLength, destinationLength, err := serializer.decode(bin[0:8])
	if err != nil {
		return "", errors.New("binary header decode failed")
	}
	bin = bin[8:]

	// resize result
	var m MatrixData
	m.Rows = make([]MatrixRow, sourceLength)
	for i := range m.Rows {
		m.Rows[i].Elements = make([]MatrixElement, destinationLength)
	}

	// decode
	for i := 0; i < len(bin); i = i + 8 {
		duration, distance, err := serializer.decode(bin[i : i+8])
		if err != nil {
			return "", errors.New("binary decode failed")
		}

		n := int32(i / 8)
		var d = n % destinationLength
		var s = (n - d) / destinationLength
		//fmt.Println(fmt.Sprintf("s=%v, d=%v", s, d))

		m.Rows[s].Elements[d] = MatrixElement{
			Distance: Value{distance},
			Duration: Value{duration},
		}
	}

	marshal, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}

func genMatrixResult(oc, dc int) MatrixData {
	var matrixData MatrixData
	matrixData.Rows = make([]MatrixRow, oc)
	i := 0
	for o := 0; o < oc; o++ {
		matrixData.Rows[o] = MatrixRow{make([]MatrixElement, dc)}
		for d := 0; d < dc; d++ {
			matrixData.Rows[o].Elements[d] = MatrixElement{
				Duration: Value{rand.Int31()},
				Distance: Value{rand.Int31()},
			}
			i++
		}
	}
	return matrixData
}

func case1() {
	testString := "{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":223},\"distance\":{\"value\":2544}},{\"duration\":{\"value\":646},\"distance\":{\"value\":6473}},{\"duration\":{\"value\":718},\"distance\":{\"value\":8322}},{\"duration\":{\"value\":1116},\"distance\":{\"value\":14948}},{\"duration\":{\"value\":1325},\"distance\":{\"value\":20294}},{\"duration\":{\"value\":1218},\"distance\":{\"value\":16527}},{\"duration\":{\"value\":1722},\"distance\":{\"value\":25780}},{\"duration\":{\"value\":1594},\"distance\":{\"value\":25298}},{\"duration\":{\"value\":1385},\"distance\":{\"value\":22240}},{\"duration\":{\"value\":981},\"distance\":{\"value\":13268}}]},{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]},{\"elements\":[{\"duration\":{\"value\":956},\"distance\":{\"value\":16748}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":20718}},{\"duration\":{\"value\":1389},\"distance\":{\"value\":25208}},{\"duration\":{\"value\":1649},\"distance\":{\"value\":26520}},{\"duration\":{\"value\":1893},\"distance\":{\"value\":30988}},{\"duration\":{\"value\":1596},\"distance\":{\"value\":25942}},{\"duration\":{\"value\":2100},\"distance\":{\"value\":35194}},{\"duration\":{\"value\":1851},\"distance\":{\"value\":31655}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":26152}},{\"duration\":{\"value\":1193},\"distance\":{\"value\":17181}}]}]}"
	bytes, err := encode(testString)
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
	bytes, err := encode(testString)
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

	bytes, err := encode(string(marshal))
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
	//bytes, err = encode(string(marshal))
	//if err != nil {
	//	return
	//}
	//if err := os.WriteFile("./test/case3/test_02.bin", bytes, 0666); err != nil {
	//	log.Fatal(err)
	//}
}

func case4() {
	testString := "{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":223},\"distance\":{\"value\":2544}},{\"duration\":{\"value\":646},\"distance\":{\"value\":6473}},{\"duration\":{\"value\":718},\"distance\":{\"value\":8322}},{\"duration\":{\"value\":1116},\"distance\":{\"value\":14948}},{\"duration\":{\"value\":1325},\"distance\":{\"value\":20294}},{\"duration\":{\"value\":1218},\"distance\":{\"value\":16527}},{\"duration\":{\"value\":1722},\"distance\":{\"value\":25780}},{\"duration\":{\"value\":1594},\"distance\":{\"value\":25298}},{\"duration\":{\"value\":1385},\"distance\":{\"value\":22240}},{\"duration\":{\"value\":981},\"distance\":{\"value\":13268}}]},{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]},{\"elements\":[{\"duration\":{\"value\":956},\"distance\":{\"value\":16748}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":20718}},{\"duration\":{\"value\":1389},\"distance\":{\"value\":25208}},{\"duration\":{\"value\":1649},\"distance\":{\"value\":26520}},{\"duration\":{\"value\":1893},\"distance\":{\"value\":30988}},{\"duration\":{\"value\":1596},\"distance\":{\"value\":25942}},{\"duration\":{\"value\":2100},\"distance\":{\"value\":35194}},{\"duration\":{\"value\":1851},\"distance\":{\"value\":31655}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":26152}},{\"duration\":{\"value\":1193},\"distance\":{\"value\":17181}}]}]}"
	bytes, err := encode(testString)
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

func case5() {
	testString := "{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":-1},\"distance\":{\"value\":-1}},{\"duration\":{\"value\":646},\"distance\":{\"value\":6473}},{\"duration\":{\"value\":718},\"distance\":{\"value\":8322}},{\"duration\":{\"value\":-1},\"distance\":{\"value\":-1}},{\"duration\":{\"value\":1325},\"distance\":{\"value\":20294}},{\"duration\":{\"value\":1218},\"distance\":{\"value\":16527}},{\"duration\":{\"value\":1722},\"distance\":{\"value\":25780}},{\"duration\":{\"value\":1594},\"distance\":{\"value\":25298}},{\"duration\":{\"value\":1385},\"distance\":{\"value\":22240}},{\"duration\":{\"value\":981},\"distance\":{\"value\":13268}}]},{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]},{\"elements\":[{\"duration\":{\"value\":956},\"distance\":{\"value\":16748}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":20718}},{\"duration\":{\"value\":1389},\"distance\":{\"value\":25208}},{\"duration\":{\"value\":1649},\"distance\":{\"value\":26520}},{\"duration\":{\"value\":1893},\"distance\":{\"value\":30988}},{\"duration\":{\"value\":1596},\"distance\":{\"value\":25942}},{\"duration\":{\"value\":2100},\"distance\":{\"value\":35194}},{\"duration\":{\"value\":1851},\"distance\":{\"value\":31655}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":26152}},{\"duration\":{\"value\":1193},\"distance\":{\"value\":17181}}]}]}"
	bytes, err := encode(testString)
	if err != nil {
		return
	}

	if err := os.WriteFile("./test/case5/1", bytes, 0666); err != nil {
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
	if err := os.WriteFile("./test/case5/"+META_FILE_NAME, marshal, 0666); err != nil {
		log.Fatal(err)
	}
}

func case6() {
	testString := "{\"status\":\"Ok\",\"rows\":[{\"elements\":[{\"duration\":{\"value\":-1},\"distance\":{\"value\":-1}},{\"duration\":{\"value\":646},\"distance\":{\"value\":6473}},{\"duration\":{\"value\":718},\"distance\":{\"value\":8322}},{\"duration\":{\"value\":-1},\"distance\":{\"value\":-1}},{\"duration\":{\"value\":1325},\"distance\":{\"value\":20294}},{\"duration\":{\"value\":1218},\"distance\":{\"value\":16527}},{\"duration\":{\"value\":1722},\"distance\":{\"value\":25780}},{\"duration\":{\"value\":1594},\"distance\":{\"value\":25298}},{\"duration\":{\"value\":1385},\"distance\":{\"value\":22240}},{\"duration\":{\"value\":981},\"distance\":{\"value\":13268}}]},{\"elements\":[{\"duration\":{\"value\":699},\"distance\":{\"value\":6981}},{\"duration\":{\"value\":436},\"distance\":{\"value\":2684}},{\"duration\":{\"value\":510},\"distance\":{\"value\":4523}},{\"duration\":{\"value\":996},\"distance\":{\"value\":9511}},{\"duration\":{\"value\":1243},\"distance\":{\"value\":17614}},{\"duration\":{\"value\":1245},\"distance\":{\"value\":16061}},{\"duration\":{\"value\":1749},\"distance\":{\"value\":25314}},{\"duration\":{\"value\":1621},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1716},\"distance\":{\"value\":20792}},{\"duration\":{\"value\":1153},\"distance\":{\"value\":13495}}]},{\"elements\":[{\"duration\":{\"value\":907},\"distance\":{\"value\":13318}},{\"duration\":{\"value\":758},\"distance\":{\"value\":7696}},{\"duration\":{\"value\":404},\"distance\":{\"value\":4519}},{\"duration\":{\"value\":470},\"distance\":{\"value\":5339}},{\"duration\":{\"value\":715},\"distance\":{\"value\":9807}},{\"duration\":{\"value\":736},\"distance\":{\"value\":7328}},{\"duration\":{\"value\":1317},\"distance\":{\"value\":19558}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":19077}},{\"duration\":{\"value\":1334},\"distance\":{\"value\":20623}},{\"duration\":{\"value\":1066},\"distance\":{\"value\":14690}}]},{\"elements\":[{\"duration\":{\"value\":1068},\"distance\":{\"value\":15949}},{\"duration\":{\"value\":968},\"distance\":{\"value\":11479}},{\"duration\":{\"value\":620},\"distance\":{\"value\":7545}},{\"duration\":{\"value\":348},\"distance\":{\"value\":2266}},{\"duration\":{\"value\":611},\"distance\":{\"value\":6702}},{\"duration\":{\"value\":637},\"distance\":{\"value\":4845}},{\"duration\":{\"value\":1383},\"distance\":{\"value\":17463}},{\"duration\":{\"value\":1476},\"distance\":{\"value\":19001}},{\"duration\":{\"value\":1622},\"distance\":{\"value\":20547}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":19980}}]},{\"elements\":[{\"duration\":{\"value\":1424},\"distance\":{\"value\":21412}},{\"duration\":{\"value\":1324},\"distance\":{\"value\":16942}},{\"duration\":{\"value\":1000},\"distance\":{\"value\":13103}},{\"duration\":{\"value\":775},\"distance\":{\"value\":9043}},{\"duration\":{\"value\":430},\"distance\":{\"value\":2922}},{\"duration\":{\"value\":829},\"distance\":{\"value\":9954}},{\"duration\":{\"value\":1238},\"distance\":{\"value\":12573}},{\"duration\":{\"value\":1527},\"distance\":{\"value\":23412}},{\"duration\":{\"value\":1673},\"distance\":{\"value\":24957}},{\"duration\":{\"value\":1651},\"distance\":{\"value\":26635}}]},{\"elements\":[{\"duration\":{\"value\":1340},\"distance\":{\"value\":26137}},{\"duration\":{\"value\":1405},\"distance\":{\"value\":20324}},{\"duration\":{\"value\":1079},\"distance\":{\"value\":18142}},{\"duration\":{\"value\":1189},\"distance\":{\"value\":20474}},{\"duration\":{\"value\":1101},\"distance\":{\"value\":19749}},{\"duration\":{\"value\":785},\"distance\":{\"value\":11597}},{\"duration\":{\"value\":941},\"distance\":{\"value\":10730}},{\"duration\":{\"value\":450},\"distance\":{\"value\":4974}},{\"duration\":{\"value\":632},\"distance\":{\"value\":6831}},{\"duration\":{\"value\":717},\"distance\":{\"value\":10176}}]},{\"elements\":[{\"duration\":{\"value\":1355},\"distance\":{\"value\":22861}},{\"duration\":{\"value\":1729},\"distance\":{\"value\":20342}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":24832}},{\"duration\":{\"value\":1707},\"distance\":{\"value\":27164}},{\"duration\":{\"value\":1619},\"distance\":{\"value\":26439}},{\"duration\":{\"value\":1303},\"distance\":{\"value\":18287}},{\"duration\":{\"value\":1470},\"distance\":{\"value\":18719}},{\"duration\":{\"value\":597},\"distance\":{\"value\":6736}},{\"duration\":{\"value\":201},\"distance\":{\"value\":2177}},{\"duration\":{\"value\":718},\"distance\":{\"value\":7812}}]},{\"elements\":[{\"duration\":{\"value\":714},\"distance\":{\"value\":11634}},{\"duration\":{\"value\":896},\"distance\":{\"value\":10838}},{\"duration\":{\"value\":952},\"distance\":{\"value\":14242}},{\"duration\":{\"value\":1208},\"distance\":{\"value\":16847}},{\"duration\":{\"value\":1453},\"distance\":{\"value\":21316}},{\"duration\":{\"value\":1155},\"distance\":{\"value\":16269}},{\"duration\":{\"value\":1659},\"distance\":{\"value\":25521}},{\"duration\":{\"value\":1379},\"distance\":{\"value\":18279}},{\"duration\":{\"value\":1118},\"distance\":{\"value\":12665}},{\"duration\":{\"value\":592},\"distance\":{\"value\":5827}}]},{\"elements\":[{\"duration\":{\"value\":1133},\"distance\":{\"value\":17947}},{\"duration\":{\"value\":1543},\"distance\":{\"value\":23408}},{\"duration\":{\"value\":1571},\"distance\":{\"value\":26606}},{\"duration\":{\"value\":1827},\"distance\":{\"value\":29211}},{\"duration\":{\"value\":1897},\"distance\":{\"value\":34447}},{\"duration\":{\"value\":1581},\"distance\":{\"value\":26296}},{\"duration\":{\"value\":1748},\"distance\":{\"value\":26727}},{\"duration\":{\"value\":1390},\"distance\":{\"value\":20949}},{\"duration\":{\"value\":1135},\"distance\":{\"value\":15445}},{\"duration\":{\"value\":845},\"distance\":{\"value\":8820}}]},{\"elements\":[{\"duration\":{\"value\":956},\"distance\":{\"value\":16748}},{\"duration\":{\"value\":1364},\"distance\":{\"value\":20718}},{\"duration\":{\"value\":1389},\"distance\":{\"value\":25208}},{\"duration\":{\"value\":1649},\"distance\":{\"value\":26520}},{\"duration\":{\"value\":1893},\"distance\":{\"value\":30988}},{\"duration\":{\"value\":1596},\"distance\":{\"value\":25942}},{\"duration\":{\"value\":2100},\"distance\":{\"value\":35194}},{\"duration\":{\"value\":1851},\"distance\":{\"value\":31655}},{\"duration\":{\"value\":1597},\"distance\":{\"value\":26152}},{\"duration\":{\"value\":1193},\"distance\":{\"value\":17181}}]}]}"
	bytes, err := encode(testString)
	if err != nil {
		return
	}

	if err := os.WriteFile("./test/case6/1", bytes, 0666); err != nil {
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
	task.version = Int32Binary

	marshal, err := json.Marshal(task)
	if err != nil {
		return
	}
	if err := os.WriteFile("./test/case6/"+META_FILE_NAME, marshal, 0666); err != nil {
		log.Fatal(err)
	}
}
