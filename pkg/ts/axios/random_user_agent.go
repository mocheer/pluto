package axios

import (
	"fmt"
	"math/rand"
	"strings"
)

var uaGens = []func() string{
	genChromeUA,
	genEdgeUA,
}

var uaGensMobile = []func() string{
	genMobilePixel7UA,
	genMobilePixel6UA,
	genMobilePixel5UA,
	genMobilePixel4UA,
	genMobileNexus10UA,
}

func RandomUserAgent() string {
	return uaGens[rand.Intn(len(uaGens))]()
}

// RandomMobileUserAgent 每次请求生成一个随机的移动端浏览器用户代理
func RandomMobileUserAgent() string {
	return uaGensMobile[rand.Intn(len(uaGensMobile))]()

}

var chromeVersions = []string{
	// 2023 年 5 月 1 日之后发布的版本。
	// 参考：https://chromereleases.googleblog.com/2023/05/stable-channel-update-for-desktop.html
	"113.0.5672.63",
	"113.0.5672.64",

	// 参考：https://chromereleases.googleblog.com/2023/05/stable-channel-update-for-desktop_8.html
	"113.0.5672.92",
	"113.0.5672.93",
}

var edgeVersions = []string{
	// 注意：仅列出 2022 年 6 月 1 日之后发布的版本。
	// 数据来源：https://learn.microsoft.com/en-us/deployedge/microsoft-edge-release-schedule

	// 2023 年
	"109.0.0.0,109.0.1518.49",
	"110.0.0.0,110.0.1587.41",
	"111.0.0.0,111.0.1661.41",
	"112.0.0.0,112.0.1722.34",
	"113.0.0.0,113.0.1774.3",
}

var pixel7AndroidVersions = []string{
	// 数据来源：
	// - https://developer.android.com/about/versions
	// - https://source.android.com/docs/setup/about/build-numbers#source-code-tags-and-builds
	"13",
}

var pixel6AndroidVersions = []string{
	// 数据来源：
	// - https://developer.android.com/about/versions
	// - https://source.android.com/docs/setup/about/build-numbers#source-code-tags-and-builds
	"12",
	"13",
}

var pixel5AndroidVersions = []string{
	// 数据来源：
	// - https://developer.android.com/about/versions
	// - https://source.android.com/docs/setup/about/build-numbers#source-code-tags-and-builds
	"11",
	"12",
	"13",
}

var pixel4AndroidVersions = []string{
	// 数据来源：
	// - https://developer.android.com/about/versions
	// - https://source.android.com/docs/setup/about/build-numbers#source-code-tags-and-builds
	"10",
	"11",
	"12",
	"13",
}

var nexus10AndroidVersions = []string{
	// 数据来源：
	// - https://developer.android.com/about/versions
	// - https://source.android.com/docs/setup/about/build-numbers#source-code-tags-and-builds
	"4.4.2",
	"4.4.4",
	"5.0",
	"5.0.1",
	"5.0.2",
	"5.1",
	"5.1.1",
}

var nexus10Builds = []string{
	// 数据来源：https://source.android.com/docs/setup/about/build-numbers#source-code-tags-and-builds

	"LMY49M", // android-5.1.1_r38（Lollipop 棒棒糖）
	"LMY49J", // android-5.1.1_r37（Lollipop 棒棒糖）
	"LMY49I", // android-5.1.1_r36（Lollipop 棒棒糖）
	"LMY49H", // android-5.1.1_r35（Lollipop 棒棒糖）
	"LMY49G", // android-5.1.1_r34（Lollipop 棒棒糖）
	"LMY49F", // android-5.1.1_r33（Lollipop 棒棒糖）
	"LMY48Z", // android-5.1.1_r30（Lollipop 棒棒糖）
	"LMY48X", // android-5.1.1_r25（Lollipop 棒棒糖）
	"LMY48T", // android-5.1.1_r19（Lollipop 棒棒糖）
	"LMY48M", // android-5.1.1_r14（Lollipop 棒棒糖）
	"LMY48I", // android-5.1.1_r9（Lollipop 棒棒糖）
	"LMY47V", // android-5.1.1_r1（Lollipop 棒棒糖）
	"LMY47D", // android-5.1.0_r1（Lollipop 棒棒糖）
	"LRX22G", // android-5.0.2_r1（Lollipop 棒棒糖）
	"LRX22C", // android-5.0.1_r1（Lollipop 棒棒糖）
	"LRX21P", // android-5.0.0_r4.0.1（Lollipop 棒棒糖）
	"KTU84P", // android-4.4.4_r1（KitKat 奇巧）
	"KTU84L", // android-4.4.3_r1（KitKat 奇巧）
	"KOT49H", // android-4.4.2_r1（KitKat 奇巧）
	"KOT49E", // android-4.4.1_r1（KitKat 奇巧）
	"KRT16S", // android-4.4_r1.2（KitKat 奇巧）
	"JWR66Y", // android-4.3_r1.1（Jelly Bean 果冻豆）
	"JWR66V", // android-4.3_r1（Jelly Bean 果冻豆）
	"JWR66N", // android-4.3_r0.9.1（Jelly Bean 果冻豆）
	"JDQ39 ", // android-4.2.2_r1（Jelly Bean 果冻豆）
	"JOP40F", // android-4.2.1_r1.1（Jelly Bean 果冻豆）
	"JOP40D", // android-4.2.1_r1（Jelly Bean 果冻豆）
	"JOP40C", // android-4.2_r1（Jelly Bean 果冻豆）
}

var osStrings = []string{
	// macOS - High Sierra
	"Macintosh; Intel Mac OS X 10_13",
	"Macintosh; Intel Mac OS X 10_13_1",
	"Macintosh; Intel Mac OS X 10_13_2",
	"Macintosh; Intel Mac OS X 10_13_3",
	"Macintosh; Intel Mac OS X 10_13_4",
	"Macintosh; Intel Mac OS X 10_13_5",
	"Macintosh; Intel Mac OS X 10_13_6",

	// macOS - Mojave
	"Macintosh; Intel Mac OS X 10_14",
	"Macintosh; Intel Mac OS X 10_14_1",
	"Macintosh; Intel Mac OS X 10_14_2",
	"Macintosh; Intel Mac OS X 10_14_3",
	"Macintosh; Intel Mac OS X 10_14_4",
	"Macintosh; Intel Mac OS X 10_14_5",
	"Macintosh; Intel Mac OS X 10_14_6",

	// macOS - Catalina
	"Macintosh; Intel Mac OS X 10_15",
	"Macintosh; Intel Mac OS X 10_15_1",
	"Macintosh; Intel Mac OS X 10_15_2",
	"Macintosh; Intel Mac OS X 10_15_3",
	"Macintosh; Intel Mac OS X 10_15_4",
	"Macintosh; Intel Mac OS X 10_15_5",
	"Macintosh; Intel Mac OS X 10_15_6",
	"Macintosh; Intel Mac OS X 10_15_7",

	// macOS - Big Sur
	"Macintosh; Intel Mac OS X 11_0",
	"Macintosh; Intel Mac OS X 11_0_1",
	"Macintosh; Intel Mac OS X 11_1",
	"Macintosh; Intel Mac OS X 11_2",
	"Macintosh; Intel Mac OS X 11_2_1",
	"Macintosh; Intel Mac OS X 11_2_2",
	"Macintosh; Intel Mac OS X 11_2_3",
	"Macintosh; Intel Mac OS X 11_3",
	"Macintosh; Intel Mac OS X 11_3_1",
	"Macintosh; Intel Mac OS X 11_4",
	"Macintosh; Intel Mac OS X 11_5",
	"Macintosh; Intel Mac OS X 11_5_1",
	"Macintosh; Intel Mac OS X 11_5_2",
	"Macintosh; Intel Mac OS X 11_6",
	"Macintosh; Intel Mac OS X 11_6_1",
	"Macintosh; Intel Mac OS X 11_6_2",
	"Macintosh; Intel Mac OS X 11_6_3",
	"Macintosh; Intel Mac OS X 11_6_4",
	"Macintosh; Intel Mac OS X 11_6_5",
	"Macintosh; Intel Mac OS X 11_6_6",
	"Macintosh; Intel Mac OS X 11_6_7",
	"Macintosh; Intel Mac OS X 11_6_8",
	"Macintosh; Intel Mac OS X 11_7",
	"Macintosh; Intel Mac OS X 11_7_1",
	"Macintosh; Intel Mac OS X 11_7_2",
	"Macintosh; Intel Mac OS X 11_7_3",
	"Macintosh; Intel Mac OS X 11_7_4",
	"Macintosh; Intel Mac OS X 11_7_5",
	"Macintosh; Intel Mac OS X 11_7_6",

	// macOS - Monterey
	"Macintosh; Intel Mac OS X 12_0",
	"Macintosh; Intel Mac OS X 12_0_1",
	"Macintosh; Intel Mac OS X 12_1",
	"Macintosh; Intel Mac OS X 12_2",
	"Macintosh; Intel Mac OS X 12_2_1",
	"Macintosh; Intel Mac OS X 12_3",
	"Macintosh; Intel Mac OS X 12_3_1",
	"Macintosh; Intel Mac OS X 12_4",
	"Macintosh; Intel Mac OS X 12_5",
	"Macintosh; Intel Mac OS X 12_5_1",
	"Macintosh; Intel Mac OS X 12_6",
	"Macintosh; Intel Mac OS X 12_6_1",
	"Macintosh; Intel Mac OS X 12_6_2",
	"Macintosh; Intel Mac OS X 12_6_3",
	"Macintosh; Intel Mac OS X 12_6_4",
	"Macintosh; Intel Mac OS X 12_6_5",

	// macOS - Ventura
	"Macintosh; Intel Mac OS X 13_0",
	"Macintosh; Intel Mac OS X 13_0_1",
	"Macintosh; Intel Mac OS X 13_1",
	"Macintosh; Intel Mac OS X 13_2",
	"Macintosh; Intel Mac OS X 13_2_1",
	"Macintosh; Intel Mac OS X 13_3",
	"Macintosh; Intel Mac OS X 13_3_1",

	// Windows
	"Windows NT 10.0; Win64; x64",
	"Windows NT 5.1",
	"Windows NT 6.1; WOW64",
	"Windows NT 6.1; Win64; x64",

	// Linux
	"X11; Linux x86_64",
}

// 生成 Chrome 浏览器用户代理（桌面端）
//
// 示例："Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36"
func genChromeUA() string {
	version := chromeVersions[rand.Intn(len(chromeVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version)
}

// 生成 Microsoft Edge 浏览器用户代理（桌面端）
//
// 示例："User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36 Edg/90.0.818.39"
func genEdgeUA() string {
	version := edgeVersions[rand.Intn(len(edgeVersions))]
	chromeVersion := strings.Split(version, ",")[0]
	edgeVersion := strings.Split(version, ",")[1]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36 Edg/%s", os, chromeVersion, edgeVersion)
}

// 生成 Pixel 7 浏览器用户代理（移动端）
//
// 示例：Mozilla/5.0 (Linux; Android 13; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36
func genMobilePixel7UA() string {
	android := pixel7AndroidVersions[rand.Intn(len(pixel7AndroidVersions))]
	chrome := chromeVersions[rand.Intn(len(chromeVersions))]
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", android, chrome)
}

// 生成 Pixel 6 浏览器用户代理（移动端）
//
// 示例："Mozilla/5.0 (Linux; Android 13; Pixel 6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36"
func genMobilePixel6UA() string {
	android := pixel6AndroidVersions[rand.Intn(len(pixel6AndroidVersions))]
	chrome := chromeVersions[rand.Intn(len(chromeVersions))]
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; Pixel 6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", android, chrome)
}

// 生成 Pixel 5 浏览器用户代理（移动端）
//
// 示例："Mozilla/5.0 (Linux; Android 13; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36"
func genMobilePixel5UA() string {
	android := pixel5AndroidVersions[rand.Intn(len(pixel5AndroidVersions))]
	chrome := chromeVersions[rand.Intn(len(chromeVersions))]
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", android, chrome)
}

// 生成 Pixel 4 浏览器用户代理（移动端）
//
// 示例："Mozilla/5.0 (Linux; Android 13; Pixel 4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36"
func genMobilePixel4UA() string {
	android := pixel4AndroidVersions[rand.Intn(len(pixel4AndroidVersions))]
	chrome := chromeVersions[rand.Intn(len(chromeVersions))]
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; Pixel 4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", android, chrome)
}

// 生成 Nexus 10 浏览器用户代理（移动端）
//
// 示例："Mozilla/5.0 (Linux; Android 5.1.1; Nexus 10 Build/LMY48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.91 Safari/537.36"
func genMobileNexus10UA() string {
	build := nexus10Builds[rand.Intn(len(nexus10Builds))]
	android := nexus10AndroidVersions[rand.Intn(len(nexus10AndroidVersions))]
	chrome := chromeVersions[rand.Intn(len(chromeVersions))]
	return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; Nexus 10 Build/%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", android, build, chrome)
}
