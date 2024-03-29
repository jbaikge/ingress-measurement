package main

import (
	"testing"
	"time"
)

var tests = []struct {
	Operation   string
	Measurement int
	Start       time.Time
	Original    string
	Padded      string
	OTP         string
	Encrypted   string
	MD5         string
}{
	// {
	// 	Operation:   "Minotaur",
	// 	Measurement: 1,
	// 	Start:       time.Date(2013, 8, 1, 2, 0, 0, 0, time.Local),
	// 	Original:    "TWO O CLOCK FOUR MINUTES AND THIRTY EIGHT SECONDS",
	// 	Padded:      "TWOXOXCLOCKXFOURXMINUTESXANDXTHIRTYXEIGHTXSECONDS",
	// 	OTP:         "ZPCDKCAPANHLJTXFBNZEJOHZELDOJPOLPVGLXNMLPBKPNPBQJ",
	// 	Encrypted:   "SLQAYZCAOPRIOHRWYZHRDHLRBLQRGIVTGOEIBVSSIYCTPDOTB",
	// 	MD5:         "e31492102271d3d2dfb46e12bb3643f0",
	// },
	// {
	// 	Operation:   "Minotaur",
	// 	Measurement: 2,
	// 	Start:       time.Date(2013, 8, 1, 3, 0, 0, 0, time.Local),
	// 	Original:    " THREE O  CLOCK ONE MINUTE AND FIFTY FIVE SECONDS",
	// 	Padded:      "XTHREEXOXXCLOCKXONEXMINUTEXANDXFIFTYXFIVEXSECONDS",
	// 	OTP:         "VBSPJANJEQMGHQDYQNYDJVULFTCGQTSDKGVXRSDFCHABJMZXV",
	// 	Encrypted:   "SUZGNEKXBNORVSNVEACAVDHFYXZGDWPISLOVOXLAGESFLAMAN",
	// 	MD5:         "6b7ca401083b34628db954c1e8b1c5de",
	// },
	// {
	// 	Operation:   "Minotaur",
	// 	Measurement: 3,
	// 	Start:       time.Date(2013, 8, 1, 4, 0, 0, 0, time.Local),
	// 	Original:    " FOUR O  CLOCK  THREE  MINUTES AND TWENTY SECONDS",
	// 	Padded:      "XFOURXOXXCLOCKXXTHREEXXMINUTESXANDXTWENTYXSECONDS",
	// 	OTP:         "OISLXYTPMVBXENULQEXVIKUFGAHODYOZWRTEZQJXTXRWQCBHZ",
	// 	Encrypted:   "LNGFOVHMJXMLGXRIJLOZMHRRONBHHQLZJUQXVUWQRUJASQOKR",
	// 	MD5:         "f0f8fd5d99d227caaa3d02bad43179cd",
	// },
	// {
	// 	Operation:   "Minotaur",
	// 	Measurement: 4,
	// 	Start:       time.Date(2013, 8, 1, 5, 0, 0, 0, time.Local),
	// 	Original:    "FIVE O CLOCK TWO MINUTES AND TWENTY SEVEN SECONDS",
	// 	Padded:      "FIVEXOXCLOCKXTWOXMINUTESXANDXTWENTYXSEVENXSECONDS",
	// 	OTP:         "KZUQKFADYLAGSRZMFKTQYGPDDJXIHDKGIWATFBLGZZEFWPWWN",
	// 	Encrypted:   "PHPUHTXFJZCQPKVACWBDSZTVAJKLEWGKVPYQXFGKMWWJYDJZF",
	// 	MD5:         "029d2072bd788bac6ebfcf431c590001",
	// },
	// {
	// 	Operation:   "Cassandra - Sydney",
	// 	Measurement: 1,
	// 	Start:       time.Date(2013, 8, 1, 1, 1, 34, 0, time.Local),
	// 	Original:    "ONE  MINUTE   THIRTY FOUR  SECONDS   PAST  ONE  O  CLOCK   ",
	// 	Padded:      "ONEXXMINUTEXXXTHIRTYXFOURXXSECONDSXXXPASTXXONEXXOXXCLOCKXXX",
	// 	OTP:         "LWIAXSXZNPDJAKBKSQAVRXDHFDDKVAMECERVZDSOHCEPTLEMNLFPRHNMDQG",
	// 	Encrypted:   "ZJMXUEFMHIHGXHURAHTTOCRBWAACZCARFWOSWSSGAZBDGPBJBICRCVPWAND",
	// 	MD5:         "64389AE19DEBAB3B7E5D668C751FD545",
	// },
	// {
	// 	Operation:   "Cassandra - Sydney",
	// 	Measurement: 2,
	// 	Start:       time.Date(2013, 8, 1, 2, 0, 0, 0, time.Local),
	// 	Original:    "  TWO  O  CLOCK AND FOUR  MINUTES  AND FIFTY FIVE SECONDS  ",
	// 	Padded:      "XXTWOXXOXXCLOCKXANDXFOURXXMINUTESXXANDXFIFTYXFIVEXSECONDSXX",
	// 	OTP:         "YMBGTLBMLYGNKAFAYTQMJNFMFGBEYFQURUGAHVCDUIXTTOFZSNPVBGQEWWE",
	// 	Encrypted:   "VJUCHIYAIVIYYCPXYGTJOBZDCDNMLZJYJRDAUYZICNQRQTNUWKHZDUDHOTB",
	// 	MD5:         "544444F7085EC9C2BD6152C25E808E2D",
	// },
	// {
	// 	Operation:   "Cassandra - Sydney",
	// 	Measurement: 3,
	// 	Start:       time.Date(2013, 8, 1, 3, 0, 0, 0, time.Local),
	// 	Original:    "   THREE  MINUTES   FIFTY TWO  SECONDS PAST THREE O CLOCK  ",
	// 	Padded:      "XXXTHREEXXMINUTESXXXFIFTYXTWOXXSECONDSXPASTXTHREEXOXCLOCKXX",
	// 	OTP:         "XQMQHNDRAKICVDGBAQTIIKJPYLMSPIJPPMAEQNQBLOSVXXBVSLOYONKQIJB",
	// 	Encrypted:   "UNJJOEHVXHUKIXZFSNQFNSOIWIFODFGHTOORTFNQLGLSQESZWICVQYYSSGY",
	// 	MD5:         "15BF98576D8B1D2D4A441643C4550323",
	// },
	// {
	// 	Operation:   "Cassandra - Milan",
	// 	Measurement: 1,
	// 	Start:       time.Date(2013, 8, 1, 7, 0, 0, 0, time.Local),
	// 	Original:    "  FIFTY  FIVE SECONDS   AND   THREE  MINUTES  AFTER   SEVEN",
	// 	Padded:      "XXFIFTYXXFIVEXSECONDSXXXANDXXXTHREEXXMINUTESXXAFTERXXXSEVEN",
	// 	OTP:         "TRTRBRYFRYVXMARIWQMSAQBZAZPFQFFHDIZDXTZDUDHLFSFUANESJGZQJIK",
	// 	Encrypted:   "QOYZGKWCODDSQXJMYEZVSNYWAMSCNCYOUMDAUFHQOWLDCPFZTRVPGDRUEMX",
	// 	MD5:         "0E83F2DFE0B755CCF2CA0413F48732C1",
	// },
	// {
	// 	Operation:   "Cassandra - Milan",
	// 	Measurement: 2,
	// 	Start:       time.Date(2013, 8, 1, 8, 0, 0, 0, time.Local),
	// 	Original:    "  EIGHT O CLOCK   TWO  MINUTES AND   FIFTY   FOUR   SECONDS",
	// 	Padded:      "XXEIGHTXOXCLOCKXXXTWOXXMINUTESXANDXXXFIFTYXXXFOURXXXSECONDS",
	// 	OTP:         "PCLYGZFEHWUMGWYUFXYVUZBQVNWLLOYRFEQNQAZDKEKUHURHNIHTNINSQZD",
	// 	Encrypted:   "MZPGMGYBVTWXUYIRCURRIWYCDAQEPGVRSHNKNFHIDCHREZFBEFEQFMPGDCV",
	// 	MD5:         "3BA87D382EE137FC9D6737437F57B550",
	// },
	// {
	// 	Operation:   "Cassandra - Milan",
	// 	Measurement: 3,
	// 	Start:       time.Date(2013, 8, 1, 9, 0, 0, 0, time.Local),
	// 	Original:    "MEASUREMENT THREE  IS  AT  NINE  O THREE AND THIRTY SECONDS",
	// 	Padded:      "MEASUREMENTXTHREEXXISXXATXXNINEXXOXTHREEXANDXTHIRTYXSECONDS",
	// 	OTP:         "STBIPKMHJWTPGRIJQUELXPUUDVBEMMIDOLNZSNPHJOOKZPANTFBKOGEKAAV",
	// 	Encrypted:   "EXBAJBQTNJMMZYZNURBTPMRUWSYRUZMALZKSZETLGOBNWIHVKYZHGKGYNDN",
	// 	MD5:         "13084003CF687FA0FEE313161767BF07",
	// },
	// {
	// 	Operation:   "Cassandra - Tokyo",
	// 	Measurement: 1,
	// 	Start:       time.Date(2013, 8, 1, 6, 0, 0, 0, time.Local),
	// 	Padded:      "XXSIXXXOXXCLOCKXANDXXFOURXMINUTESXXANDXXTHIRTEENXXXSECONDSX",
	// 	Encrypted:   "LSEFLMCXCDNLEBGOLTAAGUDKJJIEQYLREGGHVHSXDNZAQGJGVSZAWZPVUBX",
	// 	MD5:         "430A199D8124796826030643ED15E8F3",
	// },
	// {
	// 	Operation:   "Cassandra - Tokyo",
	// 	Measurement: 2,
	// 	Start:       time.Date(2013, 8, 1, 7, 0, 0, 0, time.Local),
	// 	Padded:      "XXXXXTHIRTEENXXSECONDSXXXPASTXXXXSEVENXXOXXCLOCKXXXSHARPXXX",
	// 	Encrypted:   "DPZDDKNEEJUPKNILMNSWYZWGSAREHNQSDAGSIKLPLAUSXYVTHIMVKNCLUOG",
	// 	MD5:         "DAB80D468DB085319D4D23DF1669CCFF",
	// },
	{
		Operation:   "Cassandra - Tokyo",
		Measurement: 3,
		Start:       time.Date(2013, 8, 1, 8, 0, 0, 0, time.Local),
		Padded:      "XXXTWOXMINUTESXXXXANDXSEVENXXXXSECONDSXXXPASTXXEIGHTXXXPMXX",
		Encrypted:   "QOQUAEMANBBXGRTRGNYMAUJOFZFJPTVTYSZEUHIMAMGMCQLAYMEXQELFGJZ",
		MD5:         "DB55C2E1C350396EC9D59DA94532EF63",
	},
}

func TestPackageFind(t *testing.T) {
	for _, test := range tests {
		t.Logf("Testing Operation: %s [%d]", test.Operation, test.Measurement)
		for _, f := range Formats {
			p, err := NewPackage(
				f.Format,
				test.Measurement,
				test.Start,
				test.Encrypted,
				test.MD5,
			)
			if err != nil {
				t.Fatal(err)
			}
			if p.Find() {
				t.Logf("Found OTP: %s", p.OTP)
				break
			}
		}
	}
}
