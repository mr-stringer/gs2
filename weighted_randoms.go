package main

import (
	"math/rand"

	"github.com/jmcvetta/randutil"
)

//GetWeightedSals returns a slice of weighted salutations
func GetWeightedSals() []randutil.Choice {
	sals := []randutil.Choice{
		{Weight: 10, Item: "Mr"},
		{Weight: 5, Item: "Miss"},
		{Weight: 2, Item: "Mrs"},
		{Weight: 2, Item: "Ms"},
		{Weight: 1, Item: "Dr"},
	}
	return sals
}

//GetWeightedFirstIntial returns a slice of weighted First Initials - the weighting is pretty arbitrary :/
func GetWeightedFirstIntial() []randutil.Choice {
	finit := []randutil.Choice{
		{Weight: 10, Item: "A"},
		{Weight: 10, Item: "B"},
		{Weight: 7, Item: "C"},
		{Weight: 7, Item: "D"},
		{Weight: 3, Item: "E"},
		{Weight: 5, Item: "F"},
		{Weight: 3, Item: "G"},
		{Weight: 2, Item: "H"},
		{Weight: 2, Item: "I"},
		{Weight: 8, Item: "J"},
		{Weight: 5, Item: "K"},
		{Weight: 4, Item: "L"},
		{Weight: 6, Item: "M"},
		{Weight: 7, Item: "N"},
		{Weight: 3, Item: "O"},
		{Weight: 8, Item: "P"},
		{Weight: 1, Item: "Q"},
		{Weight: 7, Item: "R"},
		{Weight: 8, Item: "S"},
		{Weight: 7, Item: "T"},
		{Weight: 2, Item: "U"},
		{Weight: 1, Item: "V"},
		{Weight: 4, Item: "W"},
		{Weight: 1, Item: "X"},
		{Weight: 1, Item: "Y"},
		{Weight: 1, Item: "Z"},
	}
	return finit
}

//GetWeightedSurnames returns a slice of weighted surnames.  Surnames are best on the most common UK surnames
//Stringer is added at the bottom as a nod to the original author
func GetWeightedSurnames() []randutil.Choice {

	sns := []randutil.Choice{
		{Weight: 10, Item: "Smith"},
		{Weight: 10, Item: "Jones"},
		{Weight: 10, Item: "Williams"},
		{Weight: 10, Item: "Brown"},
		{Weight: 10, Item: "Taylor"},
		{Weight: 10, Item: "Davies"},
		{Weight: 10, Item: "Wilson"},
		{Weight: 10, Item: "Evans"},
		{Weight: 10, Item: "Thomas"},
		{Weight: 10, Item: "Johnson"},
		{Weight: 10, Item: "Roberts"},
		{Weight: 9, Item: "Walker"},
		{Weight: 9, Item: "Wright"},
		{Weight: 9, Item: "Thompson"},
		{Weight: 9, Item: "Robinson"},
		{Weight: 9, Item: "White"},
		{Weight: 9, Item: "Hughes"},
		{Weight: 9, Item: "Edwards"},
		{Weight: 9, Item: "Hall"},
		{Weight: 9, Item: "Green"},
		{Weight: 9, Item: "Martin"},
		{Weight: 9, Item: "Wood"},
		{Weight: 8, Item: "Lewis"},
		{Weight: 8, Item: "Harris"},
		{Weight: 8, Item: "Clarke"},
		{Weight: 8, Item: "Jackson"},
		{Weight: 8, Item: "Clark"},
		{Weight: 8, Item: "Turner"},
		{Weight: 8, Item: "Scott"},
		{Weight: 8, Item: "Hill"},
		{Weight: 8, Item: "Moore"},
		{Weight: 8, Item: "Cooper"},
		{Weight: 8, Item: "Ward"},
		{Weight: 7, Item: "Morris"},
		{Weight: 7, Item: "King"},
		{Weight: 7, Item: "Watson"},
		{Weight: 7, Item: "Harrison"},
		{Weight: 7, Item: "Morgan"},
		{Weight: 7, Item: "Baker"},
		{Weight: 7, Item: "Young"},
		{Weight: 7, Item: "Patel"},
		{Weight: 7, Item: "Allen"},
		{Weight: 7, Item: "Anderson"},
		{Weight: 7, Item: "Mitchell"},
		{Weight: 6, Item: "Phillips"},
		{Weight: 6, Item: "James"},
		{Weight: 6, Item: "Campbell"},
		{Weight: 6, Item: "Bell"},
		{Weight: 6, Item: "Lee"},
		{Weight: 6, Item: "Kelly"},
		{Weight: 6, Item: "Parker"},
		{Weight: 6, Item: "Davis"},
		{Weight: 6, Item: "Bennett"},
		{Weight: 6, Item: "Miller"},
		{Weight: 6, Item: "Price"},
		{Weight: 5, Item: "Shaw"},
		{Weight: 5, Item: "Cook"},
		{Weight: 5, Item: "Simpson"},
		{Weight: 5, Item: "Griffiths"},
		{Weight: 5, Item: "Richardson"},
		{Weight: 5, Item: "Stewart"},
		{Weight: 5, Item: "Marshall"},
		{Weight: 5, Item: "Collins"},
		{Weight: 5, Item: "Carter"},
		{Weight: 5, Item: "Bailey"},
		{Weight: 5, Item: "Murphy"},
		{Weight: 4, Item: "Gray"},
		{Weight: 4, Item: "Murray"},
		{Weight: 4, Item: "Cox"},
		{Weight: 4, Item: "Adams"},
		{Weight: 4, Item: "Richards"},
		{Weight: 4, Item: "Graham"},
		{Weight: 4, Item: "Ellis"},
		{Weight: 4, Item: "Wilkinson"},
		{Weight: 4, Item: "Foster"},
		{Weight: 4, Item: "Robertson"},
		{Weight: 4, Item: "Chapman"},
		{Weight: 3, Item: "Russell"},
		{Weight: 3, Item: "Mason"},
		{Weight: 3, Item: "Webb"},
		{Weight: 3, Item: "Powell"},
		{Weight: 3, Item: "Rogers"},
		{Weight: 3, Item: "Gibson"},
		{Weight: 3, Item: "Hunt"},
		{Weight: 3, Item: "Holmes"},
		{Weight: 3, Item: "Mills"},
		{Weight: 3, Item: "Owen"},
		{Weight: 3, Item: "Palmer"},
		{Weight: 2, Item: "Matthews"},
		{Weight: 2, Item: "Reid"},
		{Weight: 2, Item: "Thomson"},
		{Weight: 2, Item: "Fisher"},
		{Weight: 2, Item: "Barnes"},
		{Weight: 2, Item: "Knight"},
		{Weight: 2, Item: "Lloyd"},
		{Weight: 2, Item: "Harvey"},
		{Weight: 2, Item: "Jenkins"},
		{Weight: 2, Item: "Barker"},
		{Weight: 2, Item: "Butler"},
		{Weight: 1, Item: "Stringer"},
	}
	return sns
}

//GetWeightedStreetNames returns a slice of Street Names.  Some of the street names of common in the UK however, some are made up.
func GetWeightedStreetNames() []randutil.Choice {
	sn := []randutil.Choice{
		{Weight: 10, Item: "Station Road"},
		{Weight: 10, Item: "Main Street"},
		{Weight: 10, Item: "Park Road"},
		{Weight: 10, Item: "High Street"},
		{Weight: 10, Item: "Church Road"},
		{Weight: 10, Item: "Church Street"},
		{Weight: 9, Item: "London Road"},
		{Weight: 9, Item: "Victoria Road"},
		{Weight: 9, Item: "Green Lane"},
		{Weight: 9, Item: "Manor Road"},
		{Weight: 9, Item: "Church Lane"},
		{Weight: 9, Item: "Park Avenue"},
		{Weight: 8, Item: "The Avenue"},
		{Weight: 8, Item: "The Crescent"},
		{Weight: 8, Item: "Queens Road"},
		{Weight: 8, Item: "New Road"},
		{Weight: 8, Item: "Grange Road"},
		{Weight: 8, Item: "Kings Road"},
		{Weight: 7, Item: "Kingsway"},
		{Weight: 7, Item: "Windsor Road"},
		{Weight: 7, Item: "Highfield Road"},
		{Weight: 7, Item: "Mill Lane"},
		{Weight: 7, Item: "Alexander Road"},
		{Weight: 7, Item: "York Road"},
		{Weight: 6, Item: "St. John’s Road"},
		{Weight: 6, Item: "Main Road"},
		{Weight: 6, Item: "Broadway"},
		{Weight: 6, Item: "King Street"},
		{Weight: 6, Item: "The Green"},
		{Weight: 6, Item: "Springfield Road"},
		{Weight: 5, Item: "George Street"},
		{Weight: 5, Item: "Park Lane"},
		{Weight: 5, Item: "Victoria Street"},
		{Weight: 5, Item: "Albert Road"},
		{Weight: 5, Item: "Queensway"},
		{Weight: 5, Item: "New Street"},
		{Weight: 4, Item: "Queen Street"},
		{Weight: 4, Item: "West Street"},
		{Weight: 4, Item: "North Street"},
		{Weight: 4, Item: "Manchester Road"},
		{Weight: 4, Item: "The Grove"},
		{Weight: 4, Item: "Richmond Road"},
		{Weight: 3, Item: "Grove Road"},
		{Weight: 3, Item: "South Street"},
		{Weight: 2, Item: "School Lane"},
		{Weight: 2, Item: "The Drive"},
		{Weight: 1, Item: "North Road"},
		{Weight: 1, Item: "Stanley Road"},
		{Weight: 1, Item: "Chester Road"},
		{Weight: 1, Item: "Mill Road"},
		{Weight: 1, Item: "Bepoke Close"},
		{Weight: 1, Item: "Chicken Way"},
		{Weight: 1, Item: "Fish End"},
		{Weight: 1, Item: "Electric Avenue"},
		{Weight: 1, Item: "Big Street"},
	}
	return sn
}

//GetWeightedTownNames returns a slice of fictional Town names which were created by https://www.name-generator.org.uk/town/
//The given weighting are arbitrary
func GetWeightedTownNames() []randutil.Choice {
	tn := []randutil.Choice{
		{Weight: 1, Item: "Fern"},
		{Weight: 1, Item: "Neyof With Docknage"},
		{Weight: 1, Item: "South Shotwich Upon Holmehull"},
		{Weight: 1, Item: "Theler"},
		{Weight: 1, Item: "Hingstoke"},
		{Weight: 1, Item: "Lytlei"},
		{Weight: 1, Item: "East Leyil"},
		{Weight: 1, Item: "Royal Crahe"},
		{Weight: 1, Item: "Stonehouserell"},
		{Weight: 1, Item: "North Dulayl"},
		{Weight: 1, Item: "East Grasan"},
		{Weight: 1, Item: "Leighblaiseton"},
		{Weight: 1, Item: "Sprofarnchester"},
		{Weight: 1, Item: "Kempphammedport"},
		{Weight: 1, Item: "Wicklymenkridge"},
		{Weight: 3, Item: "Bar"},
		{Weight: 1, Item: "Royal Mexrnaholt"},
		{Weight: 1, Item: "Fhamhartfe"},
		{Weight: 1, Item: "Wedveford"},
		{Weight: 1, Item: "Pshamarcaster"},
		{Weight: 1, Item: "St Swadeking"},
		{Weight: 1, Item: "Dentherni"},
		{Weight: 1, Item: "North Naillut"},
		{Weight: 1, Item: "West Tishya Upon Sutraunds"},
		{Weight: 1, Item: "West Stot"},
		{Weight: 1, Item: "North Terlon"},
		{Weight: 1, Item: "Roegridge Castle"},
		{Weight: 1, Item: "East Smerewoulds"},
		{Weight: 1, Item: "Royal Thirskhexham"},
		{Weight: 1, Item: "Keynessallport"},
		{Weight: 1, Item: "Safnchel"},
		{Weight: 1, Item: "Gravelischester"},
		{Weight: 1, Item: "Landto"},
		{Weight: 1, Item: "Bridheathbriggtels"},
		{Weight: 1, Item: "West Weisodsal"},
		{Weight: 1, Item: "St Ninchsteig"},
		{Weight: 1, Item: "Shieldsngarulford"},
		{Weight: 1, Item: "Royal Totwain With Vilso"},
		{Weight: 1, Item: "Royal Borneci Upon Tido"},
		{Weight: 1, Item: "Stanloshead"},
		{Weight: 1, Item: "Bridqueenbeach"},
		{Weight: 2, Item: "Neldale"},
		{Weight: 1, Item: "Swanssitbrouburgh"},
		{Weight: 1, Item: "St Linam"},
		{Weight: 1, Item: "Chibrookhedge"},
		{Weight: 1, Item: "South Yatemere"},
		{Weight: 1, Item: "West Thaxtedness"},
		{Weight: 1, Item: "Hawesfast"},
		{Weight: 1, Item: "St Patlade With Bajwis"},
		{Weight: 1, Item: "Peylang Wells"},
		{Weight: 1, Item: "Nellfrawood"},
		{Weight: 1, Item: "St Haypen"},
		{Weight: 1, Item: "New Surbecclesstur"},
		{Weight: 1, Item: "Nashmadale"},
		{Weight: 1, Item: "South Belstreet"},
		{Weight: 1, Item: "Sopoashamhampton"},
		{Weight: 1, Item: "Charl"},
		{Weight: 1, Item: "Selgre Upon Highhen"},
		{Weight: 1, Item: "Shingrstone"},
		{Weight: 1, Item: "South Whita"},
		{Weight: 1, Item: "Twhistfaington"},
		{Weight: 1, Item: "Godpreechworth"},
		{Weight: 1, Item: "North Ring"},
		{Weight: 1, Item: "Manfrith Upon Guisspils"},
		{Weight: 1, Item: "Dissroyal"},
		{Weight: 1, Item: "West Perthdown"},
		{Weight: 1, Item: "New Lonstgravesen"},
		{Weight: 3, Item: "Coldta"},
		{Weight: 1, Item: "South Basbrack"},
		{Weight: 1, Item: "New Tilbotrdon"},
		{Weight: 1, Item: "West Nchintyne"},
		{Weight: 1, Item: "Bridchettchipp"},
		{Weight: 1, Item: "Bridworthkesmin"},
		{Weight: 1, Item: "West Wotwim"},
		{Weight: 2, Item: "St Padwar"},
		{Weight: 1, Item: "Mallgis"},
		{Weight: 1, Item: "Stonebogcot"},
		{Weight: 1, Item: "St Bergains"},
		{Weight: 1, Item: "Ntagebrix Under Veywalburgh"},
		{Weight: 1, Item: "East Sportdearne"},
		{Weight: 1, Item: "New Winfrod"},
		{Weight: 1, Item: "Attdeal"},
		{Weight: 2, Item: "South Jor"},
		{Weight: 1, Item: "Azstalbray"},
		{Weight: 1, Item: "Barnesorbridge"},
		{Weight: 1, Item: "Royal Chormunds-In-Vyhad"},
		{Weight: 1, Item: "Endbingville"},
		{Weight: 1, Item: "New Leeksysett"},
		{Weight: 1, Item: "St Harties With Didwards"},
		{Weight: 1, Item: "Bridryero"},
		{Weight: 1, Item: "New Swor"},
		{Weight: 1, Item: "Stocksdalm"},
		{Weight: 1, Item: "North Stroodmir"},
		{Weight: 1, Item: "East Wenrybir"},
		{Weight: 1, Item: "Os Heath"},
		{Weight: 1, Item: "New Rnardmarsh"},
		{Weight: 1, Item: "Tenhors-By-The-Sea"},
		{Weight: 1, Item: "New Wich"},
		{Weight: 1, Item: "Wyenorthfield"},
		{Weight: 1, Item: "New Rksworthstainesde"},
	}
	return tn
}

//GetWeightedDiscount returns a slice of customer discounts
func GetWeightedDiscount() []randutil.Choice {
	disco := []randutil.Choice{
		{Weight: 50, Item: 0}, /*Not sure if this will cause a divide by 0 error later on!*/
		{Weight: 30, Item: 5},
		{Weight: 5, Item: 10},
		{Weight: 1, Item: 15},
	}
	return disco
}

//RandInt returns a random integer.  The argument min sets the minumum value of the returned int whilst the max sets the maximum value.
func RandInt(min int, max int) int {
	return rand.Intn(max-min) + min
}
