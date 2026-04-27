// Package testset defines the set of common pathfinding test cases
package testset

import (
	"github.com/TheBizzle/PathFindingCore-Golang/internal/interpreter"
)

type PathingMapString = interpreter.PathingMapString

type PathingMapTest struct {
	Dist   *float64
	MapStr PathingMapString
}

var Tests = []PathingMapTest{
	testMap1, testMap2, testMap3, testMap4, testMap5, testMap6, testMap7, testMap8, testMap9,
	testMap10, testMap11, testMap12, testMap13, testMap14, testMap15, testMap16, testMap17, testMap18,
	testMap19, testMap20, testMap21, testMap22, testMap23, testMap24, testMap25, testMap26, testMap27,
	testMap28, testMap29, testMap30, testMap31, testMap32, testMap33, testMap34, testMap35, testMap36,
	testMap37, testMap38, testMap39,
}

func pms(delim string, contents string) PathingMapString {
	return PathingMapString{Contents: contents[1:], Delim: delim + "\n"}
}

var testMap1 = PathingMapTest{new(14.0), pms("akjshdkjashldjaksdhljakds", "*             G")}

var testMap2 = PathingMapTest{new(2.0), pms("asdf", `
 *asdf
G asdf`)}

var testMap3 = PathingMapTest{nil, pms("|", `
 %  *|
OG% %|
%%   |`)}

var testMap4 = PathingMapTest{new(6.0), pms("|", `
 %  *|
OG% %|
%    |`)}

var testMap5 = PathingMapTest{new(39.0), pms("|", `
               |
           *   |
               |
               |
%%%%%%%%%%     |
        GD     |
D DDDDDDDD     |
  D D    D     |
 DD      D     |
    D DDDD     |
DDDDD    D     |
    DDDD D     |
               |
               |
               `)}

var testMap6 = PathingMapTest{new(61.0), pms("|", `
               |
           *   |
         O%%%%%|
               |
%%%%%%%%%%%%%% |
        GD     |
D DDDDDDDD %%%%|
  D D    D     |
 DD      D%%%% |
    D DDDD     |
DDDDD    D     |
    DDDD D     |
       % %     |
       % %     |
               `)}

var testMap7 = PathingMapTest{nil, pms("|", "*DG")}

var testMap8 = PathingMapTest{new(14.0), pms("|", "G             *")}

var testMap9 = PathingMapTest{new(14.0), pms("|", `
*|
 |
 |
 |
 |
 |
 |
 |
 |
 |
 |
 |
 |
 |
G`)}

var testMap10 = PathingMapTest{new(14.0), pms("|", `
G|
 |
 |
 |
 |
 |
 |
 |
 |
 |
 |
 |
 |
 |
*`)}

var testMap11 = PathingMapTest{new(7.0), pms("|", "       *      G")}

var testMap12 = PathingMapTest{new(8.0), pms("|", `
 |
 |
 |
 |
 |
 |
*|
 |
 |
 |
 |
 |
 |
 |
G`)}

var testMap13 = PathingMapTest{new(14.0), pms("|", `
*             G|
               |
               |
               |
               `)}

var testMap14 = PathingMapTest{new(14.0), pms("|", `
G             *|
               |
               |
               |
               `)}

var testMap15 = PathingMapTest{new(14.0), pms("|", `
               |
               |
               |
               |
*             G`)}

var testMap16 = PathingMapTest{new(14.0), pms("|", `
               |
               |
               |
               |
G             *`)}

var testMap17 = PathingMapTest{new(14.0), pms("|", `
               |
               |
*             G|
               |
               `)}

var testMap18 = PathingMapTest{new(4.0), pms("|", `
*              |
               |
               |
               |
G              `)}

var testMap19 = PathingMapTest{new(4.0), pms("|", `
G              |
               |
               |
               |
*              `)}

var testMap20 = PathingMapTest{new(4.0), pms("|", `
              *|
               |
               |
               |
              G`)}

var testMap21 = PathingMapTest{new(4.0), pms("|", `
              G|
               |
               |
               |
              *`)}

var testMap22 = PathingMapTest{new(4.0), pms("|", `
       *       |
               |
               |
               |
       G       `)}

var testMap23 = PathingMapTest{new(4.0), pms("|", `
       G       |
               |
               |
               |
       *       `)}

var testMap24 = PathingMapTest{new(18.0), pms("|", `
              G|
               |
               |
               |
*              `)}

var testMap25 = PathingMapTest{new(18.0), pms("|", `
G              |
               |
               |
               |
              *`)}

var testMap26 = PathingMapTest{new(9.0), pms("|", `
G              |
               |
       *       |
               |
               `)}

var testMap27 = PathingMapTest{new(20.0), pms("|", `
GD DD   D      |
   DD  D  D D  |
 D      D      |
    D  D     D |
 D  D      D  *`)}

var testMap28 = PathingMapTest{new(4.0), pms("|", `
              G|
             D |
             D |
             D |
             D*`)}

var testMap29 = PathingMapTest{new(32.0), pms("|", `
G              |
               |
               |
DDDDDDDDDDDDDD |
*              `)}

var testMap30 = PathingMapTest{new(15.0), pms("|", `
      D        |
      D        |
      D*D      |
      DDD      |
G              `)}

var testMap31 = PathingMapTest{new(13.0), pms("|", `
               |
      D D      |
      D*D      |
      DDD      |
G              `)}

var testMap32 = PathingMapTest{new(13.0), pms("|", `
        D      |
      D D      |
      D*D      |
      DDD      |
G              `)}

var testMap33 = PathingMapTest{new(9.0), pms("|", `
      D        |
      D        |
      D*D      |
      D D      |
G              `)}

var testMap34 = PathingMapTest{new(9.0), pms("|", `
      D        |
      D        |
       *D      |
      DDD      |
G              `)}

var testMap35 = PathingMapTest{nil, pms("|", `
               |
      DDD      |
      D*D      |
      DDD      |
G              `)}

var testMap36 = PathingMapTest{nil, pms("|", `
                                              |
                                              |
                                              |
                 DDDDDDDDDDDDD                |
                 D    D  D   D                |
                 D  D        D                |
                 D         D D                |
                 D D         D                |
                 D    *      D                |
                 D          DD                |
                 D   D       D                |
                 DD  D   D   D                |
                 DDDDDDDDDDDDD                |
                                              |
       G                                      |
                                              |
                                              `)}

var testMap37 = PathingMapTest{nil, pms("|", `
               |
      DDD      |
      DGD      |
      DDD      |
*              `)}

var testMap38 = PathingMapTest{nil, pms("|", `
                                              |
                                              |
                                              |
                 DDDDDDDDDDDDD                |
                 D    D  D   D                |
                 D  D        D                |
                 D         D D                |
                 D D         D                |
                 D    G      D                |
                 D          DD                |
                 D   D       D                |
                 DD  D   D   D                |
                 DDDDDDDDDDDDD                |
                                              |
       *                                      |
                                              |
                                              `)}

var testMap39 = PathingMapTest{nil, pms("|", `
                                              |
 *                                            |
                DDDDDDDDDDDDDDDDDDDDDDDDDDDDDD|
                DDDDDDDDDDDDDDDDDDDDDDDDDDDDDD|
                DD    D  D                    |
                DD  D    f   DDDDDDDDDDDDDDDD |
                DD       f D DD             D |
                DD D      fffDD            D  |
                DD    G      DD             D |
                DD   D      DDD DDDDDDDDDDDDD |
                DD   D   D   DD DD            |
 DDDDDDDDDDDDDDDDDDDDDDDDDDDDDD DDDDDDDDDDDDD |
 DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD D D   D   D |
                          D   D   D D D D D D |
DDDDDDDDDDDDDDDDDDDDDDDDD   D     D   D   D   |
DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD|
                        DDDDDDDDDDDDDDDDDDDDDD`)}
