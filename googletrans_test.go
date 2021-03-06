package googletransx

import (
	"errors"
	"testing"

	"github.com/stretchr/objx"
)

var smallTestSlice = []string{"Closet Cabinetry", "Entrance Foyer", "Pantry", "Split Bedroom", "Walk-In Closet(s)", "Family Room", "Florida Room", "Media Room", "4371 Casper Ct, Hollywood FL 33021"}
var largeTestSlice = []string{"Attached", "Three Or More Stories", "West", "Central", "Electric", "THIS HOME IS ONE OF KIND ITS A MUST SEE: BEAUTIFUL 3 STORIES SINGLE FAMILIY, PREMIUM SIZE CORNER LOT: 5946 SQFT. FULLY FURNISHED by ARTEFACTO, WOLF APPLIANCES, POOL, ROOFTOP DESIGNED FOR FAMILY ENTERTAINMENT, SUMMER KITCHEN AND JACUZZI INCLUDED.", "Dishwasher", "Disposal", "Dryer", "Electric Water Heater", "Ice Maker", "Microwave", "Other Equipment/Appliances", "Water Purifier", "Refrigerator", "Self Cleaning Oven", "Wall Oven", "Washer", "Central Air", "Concrete Block Construction", "Concrete Block With Brick", "Built-In Grill", "Lighting", "Open Balcony", "Ceramic Floor", "Wood", "Less Than 1/4 Acre Lot", "Utility Room/Laundry", "First Floor Entry", "Bar", "Closet Cabinetry", "Cooking Island", "Volume Ceilings", "Walk-In Closet(s)", "10595 NW 68th Ter, Doral FL 33178", "Covered", "Driveway", "Residential", "Single Family Residence", "Awning(s)", "Open Porch", "Patio", "Above Ground", "Concrete", "GRAND FLORIDIAN ESTATES,The Mansions", "Attached", "North", "Central", "Cash only. Spacious 1 story , 2 bedrooms 1 bath . Great location for access to major road.\\r\\nPerfect investment opportunity in the heart of Miami. TENANT OCCUPIED. PLEASE DO NOT DISTURB TENANTS.", "Microwave", "Electric Range", "Refrigerator", "Central Air", "CBS Construction", "Fruit Trees", "Tile", "Less Than 1/4 Acre Lot", "First Floor Entry", "30 NE 116th St, Miami FL 33161", "Driveway", "Residential", "Single Family Residence", "Flat Tile", "LA PALOMA", "Detached", "One Story", "East", "Central", "Electric", "Huge Lot 64,904 SqFt & Main & Guest House offers Total of 7 Bedrooms & 4 Bathrooms! Front Home: 4 Bed, 2 Bath, Modern Kitchen with Granite Countertops & New Appliances, Impact Windows & New AC. Living Room, Dining Room, Family Room, Media Room, Pool with lots of Entertaining space, Terrace with Grill, New 2018 Metal Roof & Deck! Guest House/ 2nd Home: 3 Bed, 2 Bath, 2 Modern Kitchens & 2 Living/ Dining Room. Fruit Trees: Avocado, Coconut, Mango & More. Sliding Gate entrance into the Property. NO HOA!!!\\r\\n(Virtual Tour Links) Main House: Huge Lot 64,904 SqFt & Main & Guest House offers Total of 7 Bedrooms & 4 Bathrooms!\\r\\nGuest House Part 2: https://my.matterport.com/show/?m=ddN79T9EZn6&mls=1   \\r\\nMay the Force be with you!", "Dryer", "Electric Water Heater", "Ice Maker", "Microwave", "Electric Range", "Refrigerator", "Self Cleaning Oven", "Washer", "Ceiling Fan(s)", "Central Air", "Electric", "Concrete Block Construction", "CBS Construction", "Built-In Grill", "Fruit Trees", "Lighting", "Tile", "Vinyl", "1 To Less Than 2 Acre Lot", "Utility Room/Laundry", "First Floor Entry", "Built-in Features", "Closet Cabinetry", "Family Room", "Florida Room", "Media Room", "29900 SW 170th Ave, Homestead FL 33030", "Guest House", "Circular Driveway", "Driveway", "Residential", "Single Family Residence", "Deck", "Patio", "In Ground", "Metal Roof", "Shingle", "Redland", "Hurricane Shutters", "Main Entrance", "Pathway to the main house", "Top view of the pool", "Main house front view", "Main house side view", "Pool view from the main house", "Entrance to the Main house", "Main house outside area", "Main house main entrance", "Main House Living room", "Main House Living room", "Main House Living Room 2", "Main House Dining table", "Main House Kitchen Area", "Main House Kitchen area", "Main House Kitchen Area", "Main House Laundry area", "Main House Bathroom 1", "Main House Master Bedroom ", "Main House Master Bedroom", "Master Bedroom Bathroom", "Master Bedroom Bathroom", "Main House Bedroom 1", "Main House Bedroom 1", "Main House Bedroom 2", "Main House Office room", "Guest House (Section 1) Living area", "Guest House (Section 1) Kitchen area", "Guest House (Section 1) Kitchen area", "Guest House (Section 1) Bedroom", "Guest House (Section 1) Bathroom", "Guest Room (Section 2) Living room", "Guest House (Section 2) Living area", "Guest Roon (Section 2) Kitchen area", "Guest Room (Section 2) Kitchen", "Guest Room (Section 2) Bedroom", "Guest Room (Section 2) Bathroom", "Guest Room (Section 2) Garage", "Top View", "Top View", "Top View", "Guest House", "Attached", "One Story", "East", "Central", "Welcome to this delightful, updated home in Spring Garden: a unique, historic tree-filled neighborhood. This 2 Bed / 2 Bath residence features a straightforward floor plan, open kitchen plus an office room. The front & back yards offer plenty of space to play, relax or add a pool. Upgrades include new impact windows/doors, refinished hardwood floors, new tile floor in master, new heavy gauge copper plumbing, renovated kitchen and the home was recently painted inside and out. Sited in a tranquil, canopied area of Miami, Spring Garden residents enjoy peaceful living in the dynamic Miami River District. The community is characterized by its broad streets and drives, large lot sizes, shade trees and two neighborhood parks. You're conveniently located close to Downtown, Miami Beach and Wynwood.", "Dryer", "Refrigerator", "Washer", "Central Air", "CBS Construction", "Vinyl", "Wood", "Less Than 1/4 Acre Lot", "No Additional Rooms", "920 NW 10th Ave, Miami FL 33136", "Driveway", "Residential", "Single Family Residence", "Deck", "COUNTRY CLUB ADDN", "Detached", "Two Story", "East", "3/2 Home in central area of Miami Close to Downtown. This house needs TLC Offers welcome.", "Other Equipment/Appliances", "Less Than 1/4 Acre Lot", "First Floor Entry", "1410 NW 25th Ave, Miami FL 33125", "Driveway", "Residential", "Single Family Residence", "Shingle", "MUSA ISLE", "Detached", "One Story", "East", "Central", "Electric", "Up to $10, 000 towards closing cost with preferred lender & title! NO HOA fees or CDD Tax. Our development of modern (Semi-Custom) single family homes in the Redlands area. Take advantage of pre-construction pricing now! Home features include: - BELLA MARIE ESTATES STARTS IN THE 469k-24 X 24 TILES APPLIANCES o REFRIGERATOR o OVEN RANGE o MICROWAVE - ROMAN-TUB  - ALARM - 2 CAR GARAGE - IMPACT WINDOWS - OVER 15, 000 S/F LOTS - NO CDD - NO ASSOCIATION - TILE ROOF - PAVED DRIVEWAY - SOFT CLOSING WOODEN CABINETS - TANKLESS WATER HEATERS-IMPACT WINDOWS AVAILABLE.", "Dishwasher", "Disposal", "Electric Range", "Refrigerator", "Central Air", "Electric", "Concrete Block Construction", "CBS Construction", "None", "Tile", "1/4 To Less Than 1/2 Acre Lot", "Laundry Tub", "Utility Room/Laundry", "First Floor Entry", "Closet Cabinetry", "Cooking Island", "Family Room", "29149 SW 165 TER, Homestead FL 33030", "Driveway", "Residential", "Single Family Residence", "Barrel Roof", "Bella Marie Estates", "Detached", "One Story", "Mediterranean", "South", "Central", "Electric", "ONE STORY CORNER HOME WITH A GREAT FLOOR PLAN IN THE FALLS AREA. THE FOYER LEADS TO THE FORMAL LIVING & DINING ROOMS W/HIGH CEILINGS & FRENCH TERRACOTA FLOORS THROUGHOUT. THE SPACIOUS EAT-IN KITCHEN, WITH GRANITE COUNTERS AND WOOD CABINETS, OVERLOOKS THE POOL AND PATIO AND IT\\u2019S NEXT TO THE LARGE FAMILY ROOM. A BIG GUEST BEDROOM AND BATHROOM ARE ADJACENT TO THE FAMILY ROOM. THE BEDROOM WING HAS 3 BEDROOMS, 2 BATHROOMS PLUS THE SPACIOUS MAIN SUITE WITH A LARGE BATHRM, BIG WALK-IN CLOSET & PRIVATE LANAI. SPACIOUS COVERED TERRACE, LARGE POOL, CABANA BATH & ALL-FENCED GREAT BACKYARD FOR PRIVACY.  THIS PROPERTY ALSO OFFERS A PORTE-COCHERE, 2-CAR GARAGE, 2 DRIVEWAYS, IT\\u2019S GATED ALL AROUND. AS PER SELLER\\u2019S REQUEST, PROSPECTIVE BUYERS MUST PROVIDE PRE-APPROVAL LETTER FROM BANK FOR SHOWINGS, PLEASE.", "Dishwasher", "Dryer", "Electric Water Heater", "Microwave",
	"Electric Range", "Refrigerator", "Wall Oven", "Washer", "Ceiling Fan(s)", "Central Air", "Electric", "CBS Construction", "Ceramic Floor", "1/4 To Less Than 1/2 Acre Lot", "Corner Lot", "West Of US 1", "First Floor Entry", "Cooking Island", "Entrance Foyer", "Pantry", "Vaulted Ceiling(s)", "Walk-In Closet(s)", "Family Room", "10381 SW 141st St, Miami FL 33176", "Attached Carport", "Circular Driveway", "Paver Block", "Residential", "Single Family Residence", "Open Porch", "Patio", "In Ground", "Curved/S-Tile Roof", "LUSIA R SUB", "Blinds", "Drapes", "Attached", "Two Story", "West", "Electric", "An exclusive 5BD/4.5BA cul-de-sac home with pristine landscaping in a 24 hr. guard gated community. The home has a ground floor master bedroom with an oversized walk in closet. There are 3 BD + loft area upstairs. A separate guest house with a full bath room. A pool with a large patio great for entertaining and BBQ's. The community has a tennis court, childrens' playground, private marina and is located directly across the street from the sandy white beaches of Sunny Isles. New ACs installed, completely repaved patio.", "Dishwasher", "Disposal", "Dryer", "Electric Water Heater", "Water Heater Leased", "Water Softener Owned", "Ceiling Fan(s)", "Central Air", "Zoned", "CBS Construction", "Open Balcony", "Marble", "Wood", "Less Than 1/4 Acre Lot", "Utility Room/Laundry", "First Floor Entry", "Built-in Features", "Closet Cabinetry", "Cooking Island", "Walk-In Closet(s)", "Family Room", "Garage Converted", "19425 39th Ave, Sunny Isles Beach FL 33160", "Guest House", "Driveway", "Residential", "Single Family Residence", "Patio", "In Ground", "Barrel Roof", "GOLDEN GATE ESTATES & MAR", "Detached", "One Story", "North", "Central", "Electric", "BEST VALUE IN THE AREA. 4 BED/2 BATH HOME. KITCHEN & BATHROOMS UPDATED ABOUT 5 YEARS AGO. LARGE BACK YARD. CURRENTLY LEASED FOR $1,600 MONTH BUT CAN BE DELIVERED WITH VACANT POSSESSION. LOCATED IN MIAMI-DADE OPPORTUNITY ZONE. DO NOT DISTURB TENANTS. SEE SHOWINGTIME FOR SPECIAL SHOWING AVAILABILITY.", "Dishwasher", "Electric Range", "Refrigerator", "Washer", "Ceiling Fan(s)", "Central Air", "Electric", "Lighting", "Ceramic Floor", "Less Than 1/4 Acre Lot", "Utility Room/Laundry", "First Floor Entry", "1418 NW 103rd St, Miami FL 33147", "Driveway", "No Trucks/Trailers", "Residential", "Single Family Residence", "Shingle", "MIAMI PARK SEC 2", "Attached", "One Story", "North", "Central", "Electric", "BEAUTIFUL  4 BEDROOM 4 BATH HOME IN GUARD GATED EXCLUSIVE COMMUNITY OAK FOREST. THIS SPACIOUS HOME IS OVER 3500 SQ FT OF LIVING SPACE SITTING ON A 14,000 SQ FT LOT,  A HUGE POOL AND ENORMOUS BACK YARD IS A PERFRCT ENVIRONMENT FOR FAMILY FUN.  ALL BEDROOMS HAVE BEEN UPDATED ALONG WITH MODERN BATHROOMS. THE MASTER BATH IS TRULY EXCEPTIONAL AND COMPLETELY MODERN. THE LIVING ROOM IS A PERFECT PLACE FOR ENTERTAINING ALONG WITH A BIG KITCHEN WITH NEW TOP OF THE LINE APPLIANCES. A HOME OFFICE SITS OUTSIDE THE LIVING ROOM AND THE CONVERTED GARAGE IS A PERFECT MOVIE ROOM OR MAIDS QUARTERS. BEAUTIFUL LUSH LANDSCAPING SURROUNDS THIS INCREDIBLE HOME.", "Dishwasher", "Dryer", "Microwave", "Refrigerator", "Ceiling Fan(s)", "Central Air", "Electric", "Concrete Block Construction", "Lighting", "Marble", "Tile", "1/4 To Less Than 1/2 Acre Lot", "Utility Room/Laundry", "First Floor Entry", "Built-in Features", "Den/Library/Office", "Family Room", "Garage Converted", "Maid/In-Law Quarters", "19940 NE 21st Ct, Miami FL 33179", "Circular Driveway", "Residential", "Single Family Residence", "In Ground", "Barrel Roof", "HIGHLAND OAKS", "Attached", "Two Story", "West", "Electric", "Beautifu, cozy Art Deco home. Live steps from Lincoln Rd and fine restaurants from this beautifully furnished and newly renovated home. Residence has spacious master bedroom upstairs with oversized walk in closets. Master Bath has enclosed shower and freestanding bath. On the ground level you have 1 bedroom with spacious living room , den area with 2 more beds for guests or kids, and island kitchen with breakfast area. Enclosed garage and private gust cottage complete this lovely home.Art Deco Charm in this brand new 2-story home with Guest House and room for a pool", "Trash Compactor", "Dishwasher", "Dryer", "Electric Water Heater", "Ice Maker", "Microwave", "Other Equipment/Appliances", "Refrigerator", "Self Cleaning Oven", "Washer", "Ceiling Fan(s)", "Central Air", "Ceramic Floor", "Less Than 1/4 Acre Lot", "First Floor Entry", "Cooking Island", "Walk-In Closet(s)", "Den/Library/Office", "Family Room", "Storage", "1729 Jefferson Ave, Miami Beach FL 33139", "Guest House", "Covered", "Residential", "Single Family Residence", "Concrete", "GOLF COURSE SUB AMD PL", "Detached", "One Story", "South", "Central", "GREAT DEAL in this highly sought after East Hollywood area. This property is priced to sell fast. Easy to show.", "Dishwasher", "Electric Water Heater", "Microwave", "Electric Range", "Refrigerator", "Central Air", "CBS Construction", "Terrazzo", "Tile", "Less Than 1/4 Acre Lot", "Washer/Dryer Hook-Up", "Pantry", "Split Bedroom", "Florida Room", "807 Tyler St, Hollywood FL 33019", "On Street", "Residential", "Single Family Residence", "Flat Tile", "HOLLYWOOD LAKES SEC", "Detached", "One Story", "Southwest", "Central", "Electric", "HOME WITH POOL IN THE REDLANDS. Over 1 full acre. Extremely private location, fenced and gated forest. Featuring 3 bedrooms 2.5 baths + additional room can be use as an office. Over 2700 Sq Ft of living area, entertain in grand style in this beautiful home. Open airy layout with vaulted ceilings surrounded by premium finishes and hardwood flooring throughout to include modern fixtures, shutters, solar panels, new AC, cameras, security alarm, fireplace, heated pool, screened in patio, laundry room, utility room, bedrooms boast a sleek masterful design with generously sized walk-in closets. Large master bedroom suite offers a relaxing separate sitting area for enjoyment, which also leads to the outdoor terrace and pool area. No Association. No CDD. Bring your boat, RV, no restrictions.", "Dishwasher", "Disposal", "Ice Maker", "Gas Range", "Refrigerator", "Self Cleaning Oven", "Solar Hot Water", "Water Softener Owned", "Ceiling Fan(s)", "Central Air", "Electric", "Frame With Stucco", "Stone", "Lighting", "Marble", "Wood", "1 To Less Than 2 Acre Lot", "Laundry Tub", "Utility Room/Laundry", "Washer/Dryer Hook-Up", "Entrance Foyer", "Volume Ceilings", "Walk-In Closet(s)", "Den/Library/Office", "Family Room", "Great Room", "Loft", "Storage", "17080 SW 284th St, Homestead FL 33030", "Shed(s)", "Circular Driveway", "Rv/Boat Parking", "Unpaved", "Residential", "Single Family Residence", "Open Porch", "Patio", "Screened Porch", "In Ground", "REDLANDS", "Partial Accordian Shutters", "Partial Panel Shutters/Awnings", "Hurricane Shutters", "Skylight", "Blinds", "Drapes", "Sliding", "Attached", "One Story", "East", "Central", "Electric", "BREATHTAKINGLY GORGEOUS!!  BEAUTIFULLY REMODELED RANCH STYLE HOME IN EMERALD HILLS WITH LUSH LANDSCAPE OVERLOOKING A GRAND LAKE AND GOLF COURSE.  SOME OF THE FEATURES THIS BEAUTIFUL HOME BOASTS ARE...  GRAND FLOOR PLAN, IMPACT DOORS AND WINDOWS, LUXURY SLATE TILE FLOORS, UPGRADED ITALIAN KITCHEN  CABINETS, HUGE BEDROOMS WITH WALK IN CLOSETS.  SPECTACULAR HOME!!!", "Dishwasher", "Disposal", "Dryer", "Electric Water Heater", "Microwave", "Electric Range", "Refrigerator", "Self Cleaning Oven", "Washer", "Ceiling Fan(s)", "Central Air", "Electric", "Concrete Block Construction", "CBS Construction", "Barbeque", "Built-In Grill", "Lighting", "Slate", "Less Than 1/4 Acre Lot",
	"Utility Room/Laundry", "First Floor Entry", "Closet Cabinetry", "Entrance Foyer", "Pantry", "Split Bedroom", "Walk-In Closet(s)", "Family Room", "Florida Room", "Media Room", "4371 Casper Ct, Hollywood FL 33021", "Circular Driveway", "Driveway", "Residential", "Single Family Residence", "Open Porch", "Patio", "Glass Enclosed", "In Ground", "Heated", "Shingle", "HOLLYWOOD HILLS NORTH SEC", "Garden Apartment", "Central", "Be the first to live in the newly renovated 1/1 in the heart of South Beach with parking just blocks from the beach and shopping, Ocean Drive, Flamingo Park, Lincoln Road, nightlife and more! Perfect for first-time home buyers, investors, or as a  pied-\\u00e0-terre. See one of the lowest condo dues in Miami Beach at $325 a month. Pets and rentals allowed twice a year. The condo association has reserved, 10-year inspection completed, and the building is in the process of approvals for an exterior update. Get in now and be the first to live in this gorgeous unit \\u2014 Stainless steel refrigerator on order. 1 step from Ocean Drive & Lincoln Road,", "Microwave", "Electric Range", "Refrigerator", "Ceiling Fan(s)", "Central Air", "Concrete Block Construction", "Brick", "Ceramic Floor", "First Floor Entry", "Cooking Island", "Pantry", "719 Euclid Ave # 6, Miami Beach FL 33139", "1 Space", "Residential", "Stock Cooperative", "719 EUCLID APTS CONDOMINI", "Detached", "One Story", "North", "Central", "Unapproved short sale \\r\\nThis home needs some updating and is awaiting bank approval for price", "Electric Range", "Central Air", "CBS Construction", "Ceramic Floor", "Less Than 1/4 Acre Lot", "First Floor Entry", "4920 NW 192nd St, Miami Gardens FL 33055", "Circular Driveway", "Residential", "Single Family Residence", "Shingle", "MIAMI GARDENS MANOR SEC 4", "Attached", "Two Story", "North", "Central", "Welcome to Boca's Hidden Gem. \\r\\nLocation! Location! Location! This home won't last long!\\r\\nA+ Rated Schools (Elementary, Middle and High school) This Home Features 3 large Bedroom, Private fenced pool, private fenced front patio 2 car garage 2nd floor balcony. Renovated Bathroom, Brand New A/C, Brand New Travertine Patio, Brand New Pool renovation includes(2,000 sq feet Premium Country Classic Travertine, Diamond Brite Finish, Pool Jets, Pumps Cleaning Robot and remote lights. More Pictures coming Soon. And Much More.", "Dishwasher", "Disposal", "Dryer", "Electric Water Heater", "Microwave", "Electric Range", "Refrigerator", "Washer", "Central Air", "Concrete Block Construction", "Open Balcony", "Outdoor Shower", "Ceramic Floor", "Marble", "Less Than 1/4 Acre Lot", "First Floor Entry", "Closet Cabinetry", "Vaulted Ceiling(s)", "Den/Library/Office", "7638 W Sierra Ter W, Boca Raton FL 33433", "Driveway", "Residential", "Single Family Residence", "Open Porch", "Patio", "In Ground", "Shingle", "SIERRA DEL MAR", "Detached", "Two Story", "Northwest", "If you ever dreamed of having a massive, private property with your own helicopter pad.\\r\\nHere is your chance, Unique ranch style home 5 Bedroom / 5 Bathroom with the only privet helicopter pad in the area plus landing strip. Rare opportunity!  An amazing chance for privet pilots or Helicopter owners.", "Dishwasher", "Ceiling Fan(s)", "Central Air", "Concrete Block Construction", "Barbeque", "Built-In Grill", "Ceramic Floor", "3/4 To Less Than 1 Acre Lot", "First Floor Entry", "Garage Converted", "16880 SW 59th Ct, Southwest Ranches FL 33331", "Circular Driveway", "Driveway", "Residential", "Single Family Residence", "Deck", "Curved/S-Tile Roof", "CHAMBERS LAND CO SUB", "Detached", "One Story", "South", "Central", "Electric", "Beautiful & spacious 4 bedroom, 3 bath home located in the desirable resort style community of North Gate at Keys Gate. This lovely home features 2,073 sf of living area, open kitchen with granite counters, new GE appliances (\"Graphite\"), tile floors in main living areas, wood laminate floors in bedrooms, NEW BARREL TILE ROOF 2018 with warranty transferable to new owner, inside laundry and screened in covered patio overlooking the lake. One of the bathrooms is handicapped accessible. This gorgeous home has been meticulously maintained by its original owner and will not last. The community clubhouse features pool, hot tub, tennis courts, meeting room, library, billiard room, gym and much more! Call today!", "Dishwasher", "Disposal", "Dryer", "Electric Water Heater", "Microwave", "Electric Range", "Refrigerator", "Washer", "Central Air", "Electric", "Concrete Block Construction", "CBS Construction", "Carpet", "Tile", "Wood", "Less Than 1/4 Acre Lot", "Utility Room/Laundry", "First Floor Entry", "Split Bedroom", "Vaulted Ceiling(s)", "Walk-In Closet(s)", "Family Room", "2350 SE 5th Ct, Homestead FL 33033", "Driveway", "Residential", "Single Family Residence", "Patio", "Screened Porch", "Barrel Roof", "KEYS-GATE NO 1", "Complete Accordian Shutters", "Hurricane Shutters", "Single Hung Metal", "Detached", "One Story", "West", "Electric", "Fairly updated with a 2 Year old Roof and 3 year old 5 tons A/C Unit. Main House consist of a 2/1 with living area and Garage that is now being used as a laundry room/ storage,  rear is a 1/1 with living area and back yard Centrally located in a quite area, minutes from SB, Down Town Miami. Recently upgraded, A/C, kitchen bathrooms, tile floors. This one is a turn key property. Currently generating $2500 per moth. NO FHA, NOR CONVENTIONAL.", "Dryer", "Electric Water Heater", "Central Air", "CBS Construction", "Ceramic Floor", "Less Than 1/4 Acre Lot", "In Garage", "First Floor Entry", "Garage Converted", "Studio Apartment", "8925 NW 12th Ave, Miami FL 33150", "Covered", "Residential", "Single Family Residence", "Patio", "Shingle", "CRAVERO LAKE SHORE ESTATE"}

func TestDo(t *testing.T) {
	params := TranslateParams{
		Src:  "auto",
		Dest: "zh-CN",
		Text: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. ",
	}
	transData, err := defaultTranslator.do(params)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", transData)
}

func TestTranslate(t *testing.T) {
	params := TranslateParams{
		Src:  "auto",
		Dest: "zh-CN",
		Text: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. ",
	}
	translated, err := defaultTranslator.Translate(params)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", translated)
}

func TestDetect(t *testing.T) {
	text := "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. "
	detected, err := defaultTranslator.Detect(text)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", detected)
}

func TestExtractTextsFromHTML(t *testing.T) {
	htmlsource := "<div>Some<span>test</span></div>"
	texts := extractTextsFromHTML(htmlsource)
	if len(texts) != 2 {
		t.Error(errors.New("Incorrect count of extracted texts"))
	}
}

func TestTranslateHTML(t *testing.T) {
	input := "<div>Example<span>example</span></div>"
	result, err := Translate(TranslateParams{
		Text:     input,
		Src:      "auto",
		Dest:     "ru",
		MimeType: "text/html",
	})
	if err != nil {
		t.Error(err)
	}
	if result.Text != "<div>пример<span>пример</span></div>" {
		t.Error(errors.New("HTML translation is incorrect"))
		t.Error(result.Text)
	}
}

// testBulkTranslate is a generic method for testing bulk translates
func testBulkTranslate(t *testing.T, input []string) {
	params := []TranslateParams{}
	for _, input := range input {
		params = append(params, TranslateParams{
			Src:  "auto",
			Dest: "zh-CN",
			Text: input,
		})
	}
	results, err := defaultTranslator.BulkTranslate(params)
	if err != nil {
		t.Error(err)
	}
	if len(results) != len(input) {
		t.Error(errors.New("Incorrect count of translated results"))
	}
}

// TestBulkTranslateSmall is a wrapper around generic testBulkTranslate,
// provides small test slice as input
func TestBulkTranslateSmall(t *testing.T) {
	testBulkTranslate(t, smallTestSlice)
}

// TestBulkTranslateLarge is a wrapper around generic testBulkTranslate,
// provides large test slice as input
func TestBulkTranslateLarge(t *testing.T) {
	testBulkTranslate(t, largeTestSlice)
}

func TestTranslateInterface(t *testing.T) {
	input := map[string]interface{}{
		"A": map[string]interface{}{
			"B": "I'm a test",
			"D": []string{"Example", "Example"},
		},
		"C": "Example",
	}
	params := TranslateParams{
		Src:  "en",
		Dest: "ru",
	}
	fields := []TranslateField{
		{
			Src:    "A.B",
			Dest:   "A.B_ru",
			Params: params,
		},
		{
			Src:    "A.D",
			Dest:   "A.D_ru",
			Params: params,
		},
	}
	output, err := TranslateInterface(input, fields)
	if err != nil {
		t.Fatal(err)
	}
	// Validate
	m := objx.New(output.(map[string]interface{}))
	if m.Get("A.B_ru").Str() != "Я тест" {
		t.Errorf("A.B_ru is incorrect. Got %s instead", m.Get("A.B_ru").Str())
	}
	if m.Get("A.D_ru[0]").Str() != "пример" {
		t.Errorf("A.D_ru[0] is incorrect. Got %s instead", m.Get("A.D_ru[0]").Str())
	}
	if m.Get("A.D_ru[1]").Str() != "пример" {
		t.Errorf("A.D_ru[1] is incorrect. Got %s instead", m.Get("A.D_ru[1]").Str())
	}
}
