package main

type LevelImage struct {
	ID   string
	X, Y int
}

var LevelImages = []LevelImage{
	{"small tree", 7245, 545},
	{"small tree", 7800, -789},
	{"small tree", 9032, -794},
	{"small tree", 5665, 282},
	{"small tree", 5149, 305},
	{"small tree", 4646, 307},
	{"small tree", 3803, 28},
	{"small tree", 1997, 428},
	{"small tree", 3033, 422},
	{"small tree", 2771, 425},
	{"small tree", 1420, 424},
	{"small tree", 640, 427},
	{"small tree", 991, 426},
	{"small tree", 317, 426},
	{"big tree", 8369, -905},
	{"big tree", 7479, 428},
	{"big tree", 6101, 139},
	{"big tree", 5054, 189},
	{"big tree", 4251, 184},
	{"big tree", 3421, -91},
	{"big tree", 2604, 312},
	{"big tree", 1692, 320},
	{"big tree", 1078, 315},
	{"big tree", 412, 315},
	{"huge tree", -6, 89},
	{"huge tree", 1482, 89},
	{"huge tree", 2155, 87},
	{"huge tree", 4402, -36},
	{"huge tree", 7773, -1131},
	{"huge tree", 8749, -1131},
	{"ground center 2", 338, 535},
	{"ground center 1", 399, 535},
	{"ground center 3", 278, 535},
	{"ground center 1", 216, 535},
	{"ground center 2", 155, 535},
	{"ground center 1", 460, 535},
	{"ground center 2", 521, 535},
	{"ground center 1", 582, 534},
	{"ground center 2", 643, 534},
	{"ground center 3", 705, 533},
	{"ground center 2", 765, 533},
	{"ground center 1", 826, 532},
	{"ground center 2", 887, 532},
	{"ground center 1", 948, 531},
	{"ground center 1", 1009, 531},
	{"ground center 2", 1070, 532},
	{"ground center 3", 1132, 532},
	{"ground center 2", 1192, 533},
	{"ground center 2", 1253, 533},
	{"ground center 1", 1314, 532},
	{"ground center 2", 1375, 532},
	{"ground center 2", 1860, 532},
	{"ground center 1", 1677, 532},
	{"ground center 2", 1799, 532},
	{"ground center 1", 3006, 248},
	{"huge tree", -474, 535},
	{"grass left", -355, -75},
	{"grass right", -355, 9},
	{"grass center 1", -358, 82},
	{"grass center 2", -356, 154},
	{"grass center 3", -357, 229},
	{"small tree", -531, 381},
	{"big tree", -330, 290},
	{"ground left", -392, -276},
	{"ground center 1", -243, -290},
	{"ground center 2", -233, -173},
	{"ground center 3", -249, -400},
	{"ground right", -107, -288},
	{"grass center 3", 1682, 529},
	{"grass center 3", 940, 529},
	{"grass center 3", 2690, 122},
	{"grass center 3", 410, 529},
	{"grass center 3", 304, 529},
	{"grass center 3", 251, 529},
	{"grass center 3", 198, 529},
	{"grass center 2", -346, 164},
	{"grass center 2", -346, 164},
	{"grass center 2", -346, 164},
	{"grass center 2", 1152, 529},
	{"grass center 2", 2680, 53},
	{"grass center 2", 675, 529},
	{"grass center 2", 357, 529},
	{"grass center 1", 463, 529},
	{"grass center 3", 516, 529},
	{"grass center 3", 569, 529},
	{"grass center 3", 622, 529},
	{"grass center 3", 728, 529},
	{"grass center 3", 781, 529},
	{"grass center 1", 887, 529},
	{"grass center 2", 834, 529},
	{"ground right", 3078, 535},
	{"ground center 1", 1436, 532},
	{"ground center 2", 1921, 532},
	{"ground center 2", 3018, 534},
	{"ground center 1", 2470, 534},
	{"ground center 1", 2531, 534},
	{"ground center 1", 2592, 534},
	{"ground center 1", 2714, 534},
	{"ground center 1", 2896, 534},
	{"ground center 2", 1982, 532},
	{"ground center 2", 2165, 533},
	{"ground center 2", 2287, 534},
	{"ground center 2", 2409, 534},
	{"ground center 2", 2775, 534},
	{"ground center 2", 2835, 534},
	{"ground center 3", 1498, 532},
	{"ground center 3", 1556, 532},
	{"ground center 3", 1739, 532},
	{"ground center 3", 2105, 532},
	{"ground center 3", 2658, 364},
	{"ground center 3", 2654, 534},
	{"ground center 3", 2958, 534},
	{"ground center 1", 1616, 532},
	{"ground center 1", 2043, 532},
	{"ground center 1", 2226, 533},
	{"ground center 1", 2348, 534},
	{"grass center 1", -348, 92},
	{"grass center 1", -348, 92},
	{"grass center 1", 2688, -14},
	{"grass right", -345, 19},
	{"grass right", -345, 19},
	{"grass right", 3086, 527},
	{"grass right", 2674, -89},
	{"grass left", -345, -65},
	{"grass left", -345, -65},
	{"grass left", -345, -65},
	{"grass left", 2671, -156},
	{"ground right", 2443, -76},
	{"ground left", 3065, -29},
	{"ground left", 2576, 364},
	{"ground right", 2714, 365},
	{"ground center 1", 2622, 364},
	{"grass left", 2565, 360},
	{"grass center 1", 2617, 360},
	{"grass right", 2720, 360},
	{"grass center 3", 2668, 360},
	{"ground right", 3045, 250},
	{"ground left", 2911, 249},
	{"ground center 2", 2945, 249},
	{"grass left", 2901, 246},
	{"grass right", 3059, 246},
	{"grass center 3", 2951, 246},
	{"grass center 2", 3004, 246},
	{"grass center 3", 1311, 529},
	{"grass center 3", 1205, 529},
	{"grass center 3", 1099, 529},
	{"grass center 3", 1046, 529},
	{"grass center 3", 993, 529},
	{"grass center 1", 1629, 529},
	{"grass center 1", 1576, 529},
	{"grass center 1", 1258, 529},
	{"grass center 3", 1523, 529},
	{"grass center 3", 1470, 529},
	{"grass center 3", 1417, 529},
	{"grass center 3", 1364, 529},
	{"grass center 3", -347, 239},
	{"grass center 3", -347, 239},
	{"grass center 3", -347, 239},
	{"grass center 3", 3035, 527},
	{"grass center 3", 1947, 529},
	{"grass center 3", 1894, 529},
	{"grass center 3", 1841, 529},
	{"grass center 3", 1735, 529},
	{"grass center 2", 2636, 528},
	{"grass center 2", 1788, 529},
	{"grass center 1", 2689, 528},
	{"grass center 1", 2000, 529},
	{"grass center 3", 2159, 529},
	{"grass center 3", 2106, 529},
	{"grass center 3", 2053, 529},
	{"grass center 3", 2938, 527},
	{"grass center 3", 2740, 528},
	{"grass center 3", 2583, 529},
	{"grass center 3", 2477, 529},
	{"grass center 3", 2265, 529},
	{"grass center 3", 2212, 529},
	{"grass center 3", 2318, 529},
	{"grass center 3", 2371, 529},
	{"grass center 1", 2424, 529},
	{"grass center 3", 2530, 529},
	{"grass center 3", 2791, 528},
	{"grass center 1", 2842, 528},
	{"grass center 3", 2890, 527},
	{"grass center 3", 2986, 527},
	{"ground left", 3389, 132},
	{"ground left", 4226, 409},
	{"ground left", 5531, 386},
	{"ground left", 6025, 357},
	{"ground left", 6575, 317},
	{"ground left", 7156, 647},
	{"ground left", 7329, 290},
	{"ground left", 7829, 439},
	{"ground left", 7654, 49},
	{"ground left", 7264, -160},
	{"ground left", 7709, -354},
	{"ground left", 7762, -688},
	{"ground left", 7225, -536},
}
