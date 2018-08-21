package internal

import (
	"fmt"
	"image"
	"image/color"
)

// Material Design colors
var (
	// Red50          = color.RGBA{0xff, 0xeb, 0xee, 0xff} // rgb(255, 235, 238)
	// Red100         = color.RGBA{0xff, 0xcd, 0xd2, 0xff} // rgb(255, 205, 210)
	// Red200         = color.RGBA{0xef, 0x9a, 0x9a, 0xff} // rgb(239, 154, 154)
	// Red300         = color.RGBA{0xe5, 0x73, 0x73, 0xff} // rgb(229, 115, 115)
	Red400 = color.RGBA{0xef, 0x53, 0x50, 0xff} // rgb(239, 83, 80)
	Red500 = color.RGBA{0xf4, 0x43, 0x36, 0xff} // rgb(244, 67, 54)
	Red600 = color.RGBA{0xe5, 0x39, 0x35, 0xff} // rgb(229, 57, 53)
	// Red700 = color.RGBA{0xd3, 0x2f, 0x2f, 0xff} // rgb(211, 47, 47)
	// Red800         = color.RGBA{0xc6, 0x28, 0x28, 0xff} // rgb(198, 40, 40)
	// Red900         = color.RGBA{0xb7, 0x1c, 0x1c, 0xff} // rgb(183, 28, 28)
	// RedA100        = color.RGBA{0xff, 0x8a, 0x80, 0xff} // rgb(255, 138, 128)
	// RedA200        = color.RGBA{0xff, 0x52, 0x52, 0xff} // rgb(255, 82, 82)
	// RedA400        = color.RGBA{0xff, 0x17, 0x44, 0xff} // rgb(255, 23, 68)
	// RedA700        = color.RGBA{0xd5, 0x00, 0x00, 0xff} // rgb(213, 0, 0)
	// Pink50         = color.RGBA{0xfc, 0xe4, 0xec, 0xff} // rgb(252, 228, 236)
	// Pink100        = color.RGBA{0xf8, 0xbb, 0xd0, 0xff} // rgb(248, 187, 208)
	// Pink200        = color.RGBA{0xf4, 0x8f, 0xb1, 0xff} // rgb(244, 143, 177)
	// Pink300        = color.RGBA{0xf0, 0x62, 0x92, 0xff} // rgb(240, 98, 146)
	Pink400 = color.RGBA{0xec, 0x40, 0x7a, 0xff} // rgb(236, 64, 122)
	Pink500 = color.RGBA{0xe9, 0x1e, 0x63, 0xff} // rgb(233, 30, 99)
	Pink600 = color.RGBA{0xd8, 0x1b, 0x60, 0xff} // rgb(216, 27, 96)
	// Pink700 = color.RGBA{0xc2, 0x18, 0x5b, 0xff} // rgb(194, 24, 91)
	// Pink800        = color.RGBA{0xad, 0x14, 0x57, 0xff} // rgb(173, 20, 87)
	// Pink900        = color.RGBA{0x88, 0x0e, 0x4f, 0xff} // rgb(136, 14, 79)
	// PinkA100       = color.RGBA{0xff, 0x80, 0xab, 0xff} // rgb(255, 128, 171)
	// PinkA200       = color.RGBA{0xff, 0x40, 0x81, 0xff} // rgb(255, 64, 129)
	// PinkA400       = color.RGBA{0xf5, 0x00, 0x57, 0xff} // rgb(245, 0, 87)
	// PinkA700       = color.RGBA{0xc5, 0x11, 0x62, 0xff} // rgb(197, 17, 98)
	// Purple50       = color.RGBA{0xf3, 0xe5, 0xf5, 0xff} // rgb(243, 229, 245)
	// Purple100      = color.RGBA{0xe1, 0xbe, 0xe7, 0xff} // rgb(225, 190, 231)
	// Purple200      = color.RGBA{0xce, 0x93, 0xd8, 0xff} // rgb(206, 147, 216)
	// Purple300      = color.RGBA{0xba, 0x68, 0xc8, 0xff} // rgb(186, 104, 200)
	Purple400 = color.RGBA{0xab, 0x47, 0xbc, 0xff} // rgb(171, 71, 188)
	Purple500 = color.RGBA{0x9c, 0x27, 0xb0, 0xff} // rgb(156, 39, 176)
	Purple600 = color.RGBA{0x8e, 0x24, 0xaa, 0xff} // rgb(142, 36, 170)
	// Purple700      = color.RGBA{0x7b, 0x1f, 0xa2, 0xff} // rgb(123, 31, 162)
	// Purple800      = color.RGBA{0x6a, 0x1b, 0x9a, 0xff} // rgb(106, 27, 154)
	// Purple900      = color.RGBA{0x4a, 0x14, 0x8c, 0xff} // rgb(74, 20, 140)
	// PurpleA100     = color.RGBA{0xea, 0x80, 0xfc, 0xff} // rgb(234, 128, 252)
	// PurpleA200     = color.RGBA{0xe0, 0x40, 0xfb, 0xff} // rgb(224, 64, 251)
	// PurpleA400     = color.RGBA{0xd5, 0x00, 0xf9, 0xff} // rgb(213, 0, 249)
	// PurpleA700     = color.RGBA{0xaa, 0x00, 0xff, 0xff} // rgb(170, 0, 255)
	// DeepPurple50   = color.RGBA{0xed, 0xe7, 0xf6, 0xff} // rgb(237, 231, 246)
	// DeepPurple100  = color.RGBA{0xd1, 0xc4, 0xe9, 0xff} // rgb(209, 196, 233)
	// DeepPurple200  = color.RGBA{0xb3, 0x9d, 0xdb, 0xff} // rgb(179, 157, 219)
	// DeepPurple300  = color.RGBA{0x95, 0x75, 0xcd, 0xff} // rgb(149, 117, 205)
	DeepPurple400 = color.RGBA{0x7e, 0x57, 0xc2, 0xff} // rgb(126, 87, 194)
	DeepPurple500 = color.RGBA{0x67, 0x3a, 0xb7, 0xff} // rgb(103, 58, 183)
	DeepPurple600 = color.RGBA{0x5e, 0x35, 0xb1, 0xff} // rgb(94, 53, 177)
	// DeepPurple700  = color.RGBA{0x51, 0x2d, 0xa8, 0xff} // rgb(81, 45, 168)
	// DeepPurple800  = color.RGBA{0x45, 0x27, 0xa0, 0xff} // rgb(69, 39, 160)
	// DeepPurple900  = color.RGBA{0x31, 0x1b, 0x92, 0xff} // rgb(49, 27, 146)
	// DeepPurpleA100 = color.RGBA{0xb3, 0x88, 0xff, 0xff} // rgb(179, 136, 255)
	// DeepPurpleA200 = color.RGBA{0x7c, 0x4d, 0xff, 0xff} // rgb(124, 77, 255)
	// DeepPurpleA400 = color.RGBA{0x65, 0x1f, 0xff, 0xff} // rgb(101, 31, 255)
	// DeepPurpleA700 = color.RGBA{0x62, 0x00, 0xea, 0xff} // rgb(98, 0, 234)
	// Indigo50       = color.RGBA{0xe8, 0xea, 0xf6, 0xff} // rgb(232, 234, 246)
	// Indigo100      = color.RGBA{0xc5, 0xca, 0xe9, 0xff} // rgb(197, 202, 233)
	// Indigo200      = color.RGBA{0x9f, 0xa8, 0xda, 0xff} // rgb(159, 168, 218)
	// Indigo300      = color.RGBA{0x79, 0x86, 0xcb, 0xff} // rgb(121, 134, 203)
	Indigo400 = color.RGBA{0x5c, 0x6b, 0xc0, 0xff} // rgb(92, 107, 192)
	Indigo500 = color.RGBA{0x3f, 0x51, 0xb5, 0xff} // rgb(63, 81, 181)
	Indigo600 = color.RGBA{0x39, 0x49, 0xab, 0xff} // rgb(57, 73, 171)
	// Indigo700      = color.RGBA{0x30, 0x3f, 0x9f, 0xff} // rgb(48, 63, 159)
	// Indigo800      = color.RGBA{0x28, 0x35, 0x93, 0xff} // rgb(40, 53, 147)
	// Indigo900      = color.RGBA{0x1a, 0x23, 0x7e, 0xff} // rgb(26, 35, 126)
	// IndigoA100     = color.RGBA{0x8c, 0x9e, 0xff, 0xff} // rgb(140, 158, 255)
	// IndigoA200     = color.RGBA{0x53, 0x6d, 0xfe, 0xff} // rgb(83, 109, 254)
	// IndigoA400     = color.RGBA{0x3d, 0x5a, 0xfe, 0xff} // rgb(61, 90, 254)
	// IndigoA700     = color.RGBA{0x30, 0x4f, 0xfe, 0xff} // rgb(48, 79, 254)
	// Blue50         = color.RGBA{0xe3, 0xf2, 0xfd, 0xff} // rgb(227, 242, 253)
	// Blue100        = color.RGBA{0xbb, 0xde, 0xfb, 0xff} // rgb(187, 222, 251)
	// Blue200        = color.RGBA{0x90, 0xca, 0xf9, 0xff} // rgb(144, 202, 249)
	// Blue300        = color.RGBA{0x64, 0xb5, 0xf6, 0xff} // rgb(100, 181, 246)
	Blue400 = color.RGBA{0x42, 0xa5, 0xf5, 0xff} // rgb(66, 165, 245)
	Blue500 = color.RGBA{0x21, 0x96, 0xf3, 0xff} // rgb(33, 150, 243)
	Blue600 = color.RGBA{0x1e, 0x88, 0xe5, 0xff} // rgb(30, 136, 229)
	// Blue700        = color.RGBA{0x19, 0x76, 0xd2, 0xff} // rgb(25, 118, 210)
	// Blue800        = color.RGBA{0x15, 0x65, 0xc0, 0xff} // rgb(21, 101, 192)
	// Blue900        = color.RGBA{0x0d, 0x47, 0xa1, 0xff} // rgb(13, 71, 161)
	// BlueA100       = color.RGBA{0x82, 0xb1, 0xff, 0xff} // rgb(130, 177, 255)
	// BlueA200       = color.RGBA{0x44, 0x8a, 0xff, 0xff} // rgb(68, 138, 255)
	// BlueA400       = color.RGBA{0x29, 0x79, 0xff, 0xff} // rgb(41, 121, 255)
	// BlueA700       = color.RGBA{0x29, 0x62, 0xff, 0xff} // rgb(41, 98, 255)
	// LightBlue50    = color.RGBA{0xe1, 0xf5, 0xfe, 0xff} // rgb(225, 245, 254)
	// LightBlue100   = color.RGBA{0xb3, 0xe5, 0xfc, 0xff} // rgb(179, 229, 252)
	// LightBlue200   = color.RGBA{0x81, 0xd4, 0xfa, 0xff} // rgb(129, 212, 250)
	// LightBlue300   = color.RGBA{0x4f, 0xc3, 0xf7, 0xff} // rgb(79, 195, 247)
	LightBlue400 = color.RGBA{0x29, 0xb6, 0xf6, 0xff} // rgb(41, 182, 246)
	LightBlue500 = color.RGBA{0x03, 0xa9, 0xf4, 0xff} // rgb(3, 169, 244)
	LightBlue600 = color.RGBA{0x03, 0x9b, 0xe5, 0xff} // rgb(3, 155, 229)
	// LightBlue700   = color.RGBA{0x02, 0x88, 0xd1, 0xff} // rgb(2, 136, 209)
	// LightBlue800   = color.RGBA{0x02, 0x77, 0xbd, 0xff} // rgb(2, 119, 189)
	// LightBlue900   = color.RGBA{0x01, 0x57, 0x9b, 0xff} // rgb(1, 87, 155)
	// LightBlueA100  = color.RGBA{0x80, 0xd8, 0xff, 0xff} // rgb(128, 216, 255)
	// LightBlueA200  = color.RGBA{0x40, 0xc4, 0xff, 0xff} // rgb(64, 196, 255)
	// LightBlueA400  = color.RGBA{0x00, 0xb0, 0xff, 0xff} // rgb(0, 176, 255)
	// LightBlueA700  = color.RGBA{0x00, 0x91, 0xea, 0xff} // rgb(0, 145, 234)
	// Cyan50         = color.RGBA{0xe0, 0xf7, 0xfa, 0xff} // rgb(224, 247, 250)
	// Cyan100        = color.RGBA{0xb2, 0xeb, 0xf2, 0xff} // rgb(178, 235, 242)
	// Cyan200        = color.RGBA{0x80, 0xde, 0xea, 0xff} // rgb(128, 222, 234)
	// Cyan300        = color.RGBA{0x4d, 0xd0, 0xe1, 0xff} // rgb(77, 208, 225)
	Cyan400 = color.RGBA{0x26, 0xc6, 0xda, 0xff} // rgb(38, 198, 218)
	Cyan500 = color.RGBA{0x00, 0xbc, 0xd4, 0xff} // rgb(0, 188, 212)
	Cyan600 = color.RGBA{0x00, 0xac, 0xc1, 0xff} // rgb(0, 172, 193)
	// Cyan700        = color.RGBA{0x00, 0x97, 0xa7, 0xff} // rgb(0, 151, 167)
	// Cyan800        = color.RGBA{0x00, 0x83, 0x8f, 0xff} // rgb(0, 131, 143)
	// Cyan900        = color.RGBA{0x00, 0x60, 0x64, 0xff} // rgb(0, 96, 100)
	// CyanA100       = color.RGBA{0x84, 0xff, 0xff, 0xff} // rgb(132, 255, 255)
	// CyanA200       = color.RGBA{0x18, 0xff, 0xff, 0xff} // rgb(24, 255, 255)
	// CyanA400       = color.RGBA{0x00, 0xe5, 0xff, 0xff} // rgb(0, 229, 255)
	// CyanA700       = color.RGBA{0x00, 0xb8, 0xd4, 0xff} // rgb(0, 184, 212)
	// Teal50         = color.RGBA{0xe0, 0xf2, 0xf1, 0xff} // rgb(224, 242, 241)
	// Teal100        = color.RGBA{0xb2, 0xdf, 0xdb, 0xff} // rgb(178, 223, 219)
	// Teal200        = color.RGBA{0x80, 0xcb, 0xc4, 0xff} // rgb(128, 203, 196)
	// Teal300        = color.RGBA{0x4d, 0xb6, 0xac, 0xff} // rgb(77, 182, 172)
	Teal400 = color.RGBA{0x26, 0xa6, 0x9a, 0xff} // rgb(38, 166, 154)
	Teal500 = color.RGBA{0x00, 0x96, 0x88, 0xff} // rgb(0, 150, 136)
	Teal600 = color.RGBA{0x00, 0x89, 0x7b, 0xff} // rgb(0, 137, 123)
	// Teal700        = color.RGBA{0x00, 0x79, 0x6b, 0xff} // rgb(0, 121, 107)
	// Teal800        = color.RGBA{0x00, 0x69, 0x5c, 0xff} // rgb(0, 105, 92)
	// Teal900        = color.RGBA{0x00, 0x4d, 0x40, 0xff} // rgb(0, 77, 64)
	// TealA100       = color.RGBA{0xa7, 0xff, 0xeb, 0xff} // rgb(167, 255, 235)
	// TealA200       = color.RGBA{0x64, 0xff, 0xda, 0xff} // rgb(100, 255, 218)
	// TealA400       = color.RGBA{0x1d, 0xe9, 0xb6, 0xff} // rgb(29, 233, 182)
	// TealA700       = color.RGBA{0x00, 0xbf, 0xa5, 0xff} // rgb(0, 191, 165)
	// Green50        = color.RGBA{0xe8, 0xf5, 0xe9, 0xff} // rgb(232, 245, 233)
	// Green100       = color.RGBA{0xc8, 0xe6, 0xc9, 0xff} // rgb(200, 230, 201)
	// Green200       = color.RGBA{0xa5, 0xd6, 0xa7, 0xff} // rgb(165, 214, 167)
	// Green300       = color.RGBA{0x81, 0xc7, 0x84, 0xff} // rgb(129, 199, 132)
	Green400 = color.RGBA{0x66, 0xbb, 0x6a, 0xff} // rgb(102, 187, 106)
	Green500 = color.RGBA{0x4c, 0xaf, 0x50, 0xff} // rgb(76, 175, 80)
	Green600 = color.RGBA{0x43, 0xa0, 0x47, 0xff} // rgb(67, 160, 71)
	// Green700       = color.RGBA{0x38, 0x8e, 0x3c, 0xff} // rgb(56, 142, 60)
	// Green800       = color.RGBA{0x2e, 0x7d, 0x32, 0xff} // rgb(46, 125, 50)
	// Green900       = color.RGBA{0x1b, 0x5e, 0x20, 0xff} // rgb(27, 94, 32)
	// GreenA100      = color.RGBA{0xb9, 0xf6, 0xca, 0xff} // rgb(185, 246, 202)
	// GreenA200      = color.RGBA{0x69, 0xf0, 0xae, 0xff} // rgb(105, 240, 174)
	// GreenA400      = color.RGBA{0x00, 0xe6, 0x76, 0xff} // rgb(0, 230, 118)
	// GreenA700      = color.RGBA{0x00, 0xc8, 0x53, 0xff} // rgb(0, 200, 83)
	// LightGreen50   = color.RGBA{0xf1, 0xf8, 0xe9, 0xff} // rgb(241, 248, 233)
	// LightGreen100  = color.RGBA{0xdc, 0xed, 0xc8, 0xff} // rgb(220, 237, 200)
	// LightGreen200  = color.RGBA{0xc5, 0xe1, 0xa5, 0xff} // rgb(197, 225, 165)
	// LightGreen300  = color.RGBA{0xae, 0xd5, 0x81, 0xff} // rgb(174, 213, 129)
	LightGreen400 = color.RGBA{0x9c, 0xcc, 0x65, 0xff} // rgb(156, 204, 101)
	LightGreen500 = color.RGBA{0x8b, 0xc3, 0x4a, 0xff} // rgb(139, 195, 74)
	LightGreen600 = color.RGBA{0x7c, 0xb3, 0x42, 0xff} // rgb(124, 179, 66)
	// LightGreen700  = color.RGBA{0x68, 0x9f, 0x38, 0xff} // rgb(104, 159, 56)
	// LightGreen800  = color.RGBA{0x55, 0x8b, 0x2f, 0xff} // rgb(85, 139, 47)
	// LightGreen900  = color.RGBA{0x33, 0x69, 0x1e, 0xff} // rgb(51, 105, 30)
	// LightGreenA100 = color.RGBA{0xcc, 0xff, 0x90, 0xff} // rgb(204, 255, 144)
	// LightGreenA200 = color.RGBA{0xb2, 0xff, 0x59, 0xff} // rgb(178, 255, 89)
	// LightGreenA400 = color.RGBA{0x76, 0xff, 0x03, 0xff} // rgb(118, 255, 3)
	// LightGreenA700 = color.RGBA{0x64, 0xdd, 0x17, 0xff} // rgb(100, 221, 23)
	// Lime50         = color.RGBA{0xf9, 0xfb, 0xe7, 0xff} // rgb(249, 251, 231)
	// Lime100        = color.RGBA{0xf0, 0xf4, 0xc3, 0xff} // rgb(240, 244, 195)
	// Lime200        = color.RGBA{0xe6, 0xee, 0x9c, 0xff} // rgb(230, 238, 156)
	// Lime300        = color.RGBA{0xdc, 0xe7, 0x75, 0xff} // rgb(220, 231, 117)
	Lime400 = color.RGBA{0xd4, 0xe1, 0x57, 0xff} // rgb(212, 225, 87)
	Lime500 = color.RGBA{0xcd, 0xdc, 0x39, 0xff} // rgb(205, 220, 57)
	Lime600 = color.RGBA{0xc0, 0xca, 0x33, 0xff} // rgb(192, 202, 51)
	// Lime700        = color.RGBA{0xaf, 0xb4, 0x2b, 0xff} // rgb(175, 180, 43)
	// Lime800        = color.RGBA{0x9e, 0x9d, 0x24, 0xff} // rgb(158, 157, 36)
	// Lime900        = color.RGBA{0x82, 0x77, 0x17, 0xff} // rgb(130, 119, 23)
	// LimeA100       = color.RGBA{0xf4, 0xff, 0x81, 0xff} // rgb(244, 255, 129)
	// LimeA200       = color.RGBA{0xee, 0xff, 0x41, 0xff} // rgb(238, 255, 65)
	// LimeA400       = color.RGBA{0xc6, 0xff, 0x00, 0xff} // rgb(198, 255, 0)
	// LimeA700       = color.RGBA{0xae, 0xea, 0x00, 0xff} // rgb(174, 234, 0)
	// Yellow50       = color.RGBA{0xff, 0xfd, 0xe7, 0xff} // rgb(255, 253, 231)
	// Yellow100      = color.RGBA{0xff, 0xf9, 0xc4, 0xff} // rgb(255, 249, 196)
	// Yellow200      = color.RGBA{0xff, 0xf5, 0x9d, 0xff} // rgb(255, 245, 157)
	// Yellow300      = color.RGBA{0xff, 0xf1, 0x76, 0xff} // rgb(255, 241, 118)
	// Yellow400      = color.RGBA{0xff, 0xee, 0x58, 0xff} // rgb(255, 238, 88)
	Yellow500 = color.RGBA{0xff, 0xeb, 0x3b, 0xff} // rgb(255, 235, 59)
	Yellow600 = color.RGBA{0xfd, 0xd8, 0x35, 0xff} // rgb(253, 216, 53)
	Yellow700 = color.RGBA{0xfb, 0xc0, 0x2d, 0xff} // rgb(251, 192, 45)
	// Yellow800      = color.RGBA{0xf9, 0xa8, 0x25, 0xff} // rgb(249, 168, 37)
	// Yellow900      = color.RGBA{0xf5, 0x7f, 0x17, 0xff} // rgb(245, 127, 23)
	// YellowA100     = color.RGBA{0xff, 0xff, 0x8d, 0xff} // rgb(255, 255, 141)
	// YellowA200     = color.RGBA{0xff, 0xff, 0x00, 0xff} // rgb(255, 255, 0)
	// YellowA400     = color.RGBA{0xff, 0xea, 0x00, 0xff} // rgb(255, 234, 0)
	// YellowA700     = color.RGBA{0xff, 0xd6, 0x00, 0xff} // rgb(255, 214, 0)
	// Amber50        = color.RGBA{0xff, 0xf8, 0xe1, 0xff} // rgb(255, 248, 225)
	// Amber100       = color.RGBA{0xff, 0xec, 0xb3, 0xff} // rgb(255, 236, 179)
	// Amber200       = color.RGBA{0xff, 0xe0, 0x82, 0xff} // rgb(255, 224, 130)
	// Amber300       = color.RGBA{0xff, 0xd5, 0x4f, 0xff} // rgb(255, 213, 79)
	Amber400 = color.RGBA{0xff, 0xca, 0x28, 0xff} // rgb(255, 202, 40)
	Amber500 = color.RGBA{0xff, 0xc1, 0x07, 0xff} // rgb(255, 193, 7)
	Amber600 = color.RGBA{0xff, 0xb3, 0x00, 0xff} // rgb(255, 179, 0)
	// Amber700       = color.RGBA{0xff, 0xa0, 0x00, 0xff} // rgb(255, 160, 0)
	// Amber800       = color.RGBA{0xff, 0x8f, 0x00, 0xff} // rgb(255, 143, 0)
	// Amber900       = color.RGBA{0xff, 0x6f, 0x00, 0xff} // rgb(255, 111, 0)
	// AmberA100      = color.RGBA{0xff, 0xe5, 0x7f, 0xff} // rgb(255, 229, 127)
	// AmberA200      = color.RGBA{0xff, 0xd7, 0x40, 0xff} // rgb(255, 215, 64)
	// AmberA400      = color.RGBA{0xff, 0xc4, 0x00, 0xff} // rgb(255, 196, 0)
	// AmberA700      = color.RGBA{0xff, 0xab, 0x00, 0xff} // rgb(255, 171, 0)
	// Orange50       = color.RGBA{0xff, 0xf3, 0xe0, 0xff} // rgb(255, 243, 224)
	// Orange100      = color.RGBA{0xff, 0xe0, 0xb2, 0xff} // rgb(255, 224, 178)
	// Orange200      = color.RGBA{0xff, 0xcc, 0x80, 0xff} // rgb(255, 204, 128)
	// Orange300      = color.RGBA{0xff, 0xb7, 0x4d, 0xff} // rgb(255, 183, 77)
	Orange400 = color.RGBA{0xff, 0xa7, 0x26, 0xff} // rgb(255, 167, 38)
	Orange500 = color.RGBA{0xff, 0x98, 0x00, 0xff} // rgb(255, 152, 0)
	Orange600 = color.RGBA{0xfb, 0x8c, 0x00, 0xff} // rgb(251, 140, 0)
	// Orange700      = color.RGBA{0xf5, 0x7c, 0x00, 0xff} // rgb(245, 124, 0)
	// Orange800      = color.RGBA{0xef, 0x6c, 0x00, 0xff} // rgb(239, 108, 0)
	// Orange900      = color.RGBA{0xe6, 0x51, 0x00, 0xff} // rgb(230, 81, 0)
	// OrangeA100     = color.RGBA{0xff, 0xd1, 0x80, 0xff} // rgb(255, 209, 128)
	// OrangeA200     = color.RGBA{0xff, 0xab, 0x40, 0xff} // rgb(255, 171, 64)
	// OrangeA400     = color.RGBA{0xff, 0x91, 0x00, 0xff} // rgb(255, 145, 0)
	// OrangeA700     = color.RGBA{0xff, 0x6d, 0x00, 0xff} // rgb(255, 109, 0)
	// DeepOrange50   = color.RGBA{0xfb, 0xe9, 0xe7, 0xff} // rgb(251, 233, 231)
	// DeepOrange100  = color.RGBA{0xff, 0xcc, 0xbc, 0xff} // rgb(255, 204, 188)
	// DeepOrange200  = color.RGBA{0xff, 0xab, 0x91, 0xff} // rgb(255, 171, 145)
	// DeepOrange300  = color.RGBA{0xff, 0x8a, 0x65, 0xff} // rgb(255, 138, 101)
	DeepOrange400 = color.RGBA{0xff, 0x70, 0x43, 0xff} // rgb(255, 112, 67)
	DeepOrange500 = color.RGBA{0xff, 0x57, 0x22, 0xff} // rgb(255, 87, 34)
	DeepOrange600 = color.RGBA{0xf4, 0x51, 0x1e, 0xff} // rgb(244, 81, 30)
	// DeepOrange700  = color.RGBA{0xe6, 0x4a, 0x19, 0xff} // rgb(230, 74, 25)
	// DeepOrange800  = color.RGBA{0xd8, 0x43, 0x15, 0xff} // rgb(216, 67, 21)
	// DeepOrange900  = color.RGBA{0xbf, 0x36, 0x0c, 0xff} // rgb(191, 54, 12)
	// DeepOrangeA100 = color.RGBA{0xff, 0x9e, 0x80, 0xff} // rgb(255, 158, 128)
	// DeepOrangeA200 = color.RGBA{0xff, 0x6e, 0x40, 0xff} // rgb(255, 110, 64)
	// DeepOrangeA400 = color.RGBA{0xff, 0x3d, 0x00, 0xff} // rgb(255, 61, 0)
	// DeepOrangeA700 = color.RGBA{0xdd, 0x2c, 0x00, 0xff} // rgb(221, 44, 0)
	// Brown50        = color.RGBA{0xef, 0xeb, 0xe9, 0xff} // rgb(239, 235, 233)
	// Brown100       = color.RGBA{0xd7, 0xcc, 0xc8, 0xff} // rgb(215, 204, 200)
	// Brown200       = color.RGBA{0xbc, 0xaa, 0xa4, 0xff} // rgb(188, 170, 164)
	// Brown300       = color.RGBA{0xa1, 0x88, 0x7f, 0xff} // rgb(161, 136, 127)
	Brown400 = color.RGBA{0x8d, 0x6e, 0x63, 0xff} // rgb(141, 110, 99)
	Brown500 = color.RGBA{0x79, 0x55, 0x48, 0xff} // rgb(121, 85, 72)
	Brown600 = color.RGBA{0x6d, 0x4c, 0x41, 0xff} // rgb(109, 76, 65)
	// Brown700       = color.RGBA{0x5d, 0x40, 0x37, 0xff} // rgb(93, 64, 55)
	// Brown800       = color.RGBA{0x4e, 0x34, 0x2e, 0xff} // rgb(78, 52, 46)
	// Brown900       = color.RGBA{0x3e, 0x27, 0x23, 0xff} // rgb(62, 39, 35)
	// Grey50         = color.RGBA{0xfa, 0xfa, 0xfa, 0xff} // rgb(250, 250, 250)
	// Grey100        = color.RGBA{0xf5, 0xf5, 0xf5, 0xff} // rgb(245, 245, 245)
	// Grey200        = color.RGBA{0xee, 0xee, 0xee, 0xff} // rgb(238, 238, 238)
	// Grey300        = color.RGBA{0xe0, 0xe0, 0xe0, 0xff} // rgb(224, 224, 224)
	Grey400 = color.RGBA{0xbd, 0xbd, 0xbd, 0xff} // rgb(189, 189, 189)
	Grey500 = color.RGBA{0x9e, 0x9e, 0x9e, 0xff} // rgb(158, 158, 158)
	Grey600 = color.RGBA{0x75, 0x75, 0x75, 0xff} // rgb(117, 117, 117)
	// Grey700        = color.RGBA{0x61, 0x61, 0x61, 0xff} // rgb(97, 97, 97)
	// Grey800        = color.RGBA{0x42, 0x42, 0x42, 0xff} // rgb(66, 66, 66)
	// Grey900        = color.RGBA{0x21, 0x21, 0x21, 0xff} // rgb(33, 33, 33)
	// BlueGrey50     = color.RGBA{0xec, 0xef, 0xf1, 0xff} // rgb(236, 239, 241)
	// BlueGrey100    = color.RGBA{0xcf, 0xd8, 0xdc, 0xff} // rgb(207, 216, 220)
	// BlueGrey200    = color.RGBA{0xb0, 0xbe, 0xc5, 0xff} // rgb(176, 190, 197)
	// BlueGrey300    = color.RGBA{0x90, 0xa4, 0xae, 0xff} // rgb(144, 164, 174)
	BlueGrey400 = color.RGBA{0x78, 0x90, 0x9c, 0xff} // rgb(120, 144, 156)
	BlueGrey500 = color.RGBA{0x60, 0x7d, 0x8b, 0xff} // rgb(96, 125, 139)
	BlueGrey600 = color.RGBA{0x54, 0x6e, 0x7a, 0xff} // rgb(84, 110, 122)
	// BlueGrey700    = color.RGBA{0x45, 0x5a, 0x64, 0xff} // rgb(69, 90, 100)
	// BlueGrey800    = color.RGBA{0x37, 0x47, 0x4f, 0xff} // rgb(55, 71, 79)
	// BlueGrey900    = color.RGBA{0x26, 0x32, 0x38, 0xff} // rgb(38, 50, 56)
	// Black          = color.RGBA{0x00, 0x00, 0x00, 0xff} // rgb(0, 0, 0)
	// White          = color.RGBA{0xff, 0xff, 0xff, 0xff} // rgb(255, 255, 255)
)

// Map of Color Names to color
var Map = make(map[string][3]color.RGBA)

type point image.Point

type line struct {
	p1 point
	p2 point
}

func init() {
	Map["red"] = [3]color.RGBA{Red400, Red500, Red600}
	Map["pink"] = [3]color.RGBA{Pink400, Pink500, Pink600}
	Map["purple"] = [3]color.RGBA{Purple400, Purple500, Purple600}
	Map["deeppurple"] = [3]color.RGBA{DeepPurple400, DeepPurple500, DeepPurple600}
	Map["indigo"] = [3]color.RGBA{Indigo400, Indigo500, Indigo600}
	Map["blue"] = [3]color.RGBA{Blue400, Blue500, Blue600}
	Map["lightblue"] = [3]color.RGBA{LightBlue400, LightBlue500, LightBlue600}
	Map["cyan"] = [3]color.RGBA{Cyan400, Cyan500, Cyan600}
	Map["teal"] = [3]color.RGBA{Teal400, Teal500, Teal600}
	Map["green"] = [3]color.RGBA{Green400, Green500, Green600}
	Map["lime"] = [3]color.RGBA{Lime400, Lime500, Lime600}
	Map["yellow"] = [3]color.RGBA{Yellow500, Yellow600, Yellow700}
	Map["amber"] = [3]color.RGBA{Amber400, Amber500, Amber600}
	Map["orange"] = [3]color.RGBA{Orange400, Orange500, Orange600}
	Map["deeporange"] = [3]color.RGBA{DeepOrange400, DeepOrange500, DeepOrange600}
	Map["brown"] = [3]color.RGBA{Brown400, Brown500, Brown600}
	Map["grey"] = [3]color.RGBA{Grey400, Grey500, Grey600}
	Map["bluegrey"] = [3]color.RGBA{BlueGrey400, BlueGrey500, BlueGrey600}
}

// getColor get a RGBA color
func getColor(colorName *string) ([3]color.RGBA, error) {
	colors, ok := Map[*colorName]
	if !ok {
		return [3]color.RGBA{}, fmt.Errorf("Color %s could not be found", *colorName)
	}
	return colors, nil
}

func SetImageColor(pic *image.RGBA, bgColor *string) error {

	width := pic.Bounds().Max.X
	height := pic.Bounds().Max.Y

	leftLine := line{point{width / 4, 0}, point{width / 20, height}}
	rightLine := line{point{width, height - 2*height/3}, point{width - width/3, height}}

	colors, err := getColor(bgColor)
	if err != nil {
		return err
	}

	leftColor := colors[0]
	middleColor := colors[1]
	rightColor := colors[2]

	// Update the colour of each pixel
	for x := pic.Bounds().Min.X; x < pic.Bounds().Max.X; x++ {
		for y := pic.Bounds().Min.Y; y < pic.Bounds().Max.Y; y++ {

			p := point{x, y}
			if isLeftOfLine(leftLine, p) {
				pic.Set(x, y, leftColor)
			} else if isRightOfLine(rightLine, p) {
				pic.Set(x, y, rightColor)
			} else {
				pic.Set(x, y, middleColor)
			}
		}
	}
	return nil
}

// isRightOfLine checks if a point is to the right of the line
func isRightOfLine(l line, p point) bool {
	return !isLeftOfLine(l, p)
}

// isLeftOfLine checks if a point is to the left of the line
func isLeftOfLine(l line, p point) bool {
	d := (p.X-l.p1.X)*(l.p2.Y-l.p1.Y) - (p.Y-l.p1.Y)*(l.p2.X-l.p1.X)
	return d < 0
}
