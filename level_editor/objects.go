package main

type LevelObject struct {
	X, Y, W, H int
	Solid      bool
}

var LevelObjects = []LevelObject{
	{175, -608, 29, 1192, true},
	{204, 537, 2933, 47, true},
	{2915, 254, 190, 38, false},
	{3396, 136, 659, 41, false},
	{2581, 368, 194, 38, false},
	{4231, 413, 1118, 47, true},
	{4374, 268, 216, 159, true},
	{5537, 391, 275, 47, true},
	{6029, 362, 277, 47, true},
	{6581, 322, 382, 47, true},
	{7162, 653, 662, 47, true},
	{7334, 296, 352, 47, false},
	{7661, 54, 249, 45, false},
	{7269, -155, 251, 48, false},
	{7836, 443, 267, 50, false},
	{7714, -350, 254, 45, false},
	{7231, -531, 228, 43, false},
	{7767, -682, 1585, 44, true},
}
