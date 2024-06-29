package pensioninsurancenumber

type AreaCode uint8

// https://www.gesetze-im-internet.de/vkvv/anlage.html
// 2023-01-21
const (
	// 0
	// 1
	AreaMecklenburgWesternPomerania = 2
	AreaThuringia                   = 3
	AreaBrandenburg                 = 4
	// 5
	// 6
	// 7
	AreaSaxonyAnhalt                = 8
	AreaSaxony                      = 9
	AreaHanover                     = 10
	AreaWestphalia                  = 11
	AreaHesse                       = 12
	AreaRhineProvince               = 13
	AreaUpperBavaria                = 14
	AreaLowerBavariaUpperPalatinate = 15
	AreaRhinelandPalatinate         = 16
	AreaSaarland                    = 17
	AreaUpperAndMiddleFranconia     = 18
	AreaHamburg                     = 19
	AreaLowerFranconia              = 20
	AreaSwabia                      = 21
	// 22
	AreaWurttemberg       = 23
	AreaBaden             = 24
	AreaBerlin            = 25
	AreaSchleswigHolstein = 26
	// 27
	AreaOldenburgBremen = 28
	AreaBraunschweig    = 29
	// 30
	// 31
	// 32
	// 33
	// 34
	// 35
	// 36
	// 37
	AreaKnappschaftBahnSee_Bahn          = 38
	AreaKnappschaftBahnSee_See           = 39
	AreaZulagenstelleFuerAltersvermoegen = 40

	// 41
	AreaBundMecklenburgWesternPomerania = 2 + 40
	AreaBundThuringia                   = 3 + 40
	AreaBundBrandenburg                 = 4 + 40
	// 45
	// 46
	// 47
	AreaBundSaxonyAnhalt                = 8 + 40
	AreaBundSaxony                      = 9 + 40
	AreaBundHanover                     = 10 + 40
	AreaBundWestphalia                  = 11 + 40
	AreaBundHesse                       = 12 + 40
	AreaBundRhineProvince               = 13 + 40
	AreaBundUpperBavaria                = 14 + 40
	AreaBundLowerBavariaUpperPalatinate = 15 + 40
	AreaBundRhinelandPalatinate         = 16 + 40
	AreaBundSaarland                    = 17 + 40
	AreaBundUpperAndMiddleFranconia     = 18 + 40
	AreaBundHamburg                     = 19 + 40
	AreaBundLowerFranconia              = 20 + 40
	AreaBundSwabia                      = 21 + 40
	// 62
	AreaBundWurttemberg       = 23 + 40
	AreaBundBaden             = 24 + 40
	AreaBundBerlin            = 25 + 40
	AreaBundSchleswigHolstein = 26 + 40
	// 67
	AreaBundOldenburgBremen = 28 + 40
	AreaBundBraunschweig    = 29 + 40
	// 70
	// 71
	// 72
	// 73
	// 74
	// 75
	// 76
	// 77
	AreaBundKnappschaftBahnSee_Bahn = 38 + 40
	AreaBundKnappschaftBahnSee_See  = 39 + 40

	AreaKnappschaftBahnSeeBerBreHamLoSaxSchleHolWestphalia = 80 // ik the name is whack, anything else would've been too long tho
	AreaKnappschaftBahnSeeHesseRhineProvinz                = 81
	AreaKnappschaftBahnSeeBaWuBavariaRhinePalatSaar        = 82
	// 83
	// 84
	// 85
	// 86
	// 87
	// 88
	AreaKnappschaftBahnSeeBrandMeckWesPomSaxAnhSaxThu = 89
)

func (c AreaCode) String() string {
	// https://de.wikipedia.org/wiki/Versicherungsnummer#Aufbau_der_Ziffern_von_der_Bereichsnummer_bis_zur_Seriennummer
	// 2023-01-21

	switch c {
	case AreaMecklenburgWesternPomerania:
		return "Deutsche Rentenversicherung Nord (Mecklenburg-Vorpommern)"
	case AreaThuringia:
		return "Deutsche Rentenversicherung Mitteldeutschland (Thüringen)"
	case AreaBrandenburg:
		return "Deutsche Rentenversicherung Berlin-Brandenburg (Brandenburg)"
	case AreaSaxonyAnhalt:
		return "Deutsche Rentenversicherung Mitteldeutschland (Sachsen-Anhalt)"
	case AreaSaxony:
		return "Deutsche Rentenversicherung Mitteldeutschland (Sachsen)"
	case AreaHanover:
		return "Deutsche Rentenversicherung Braunschweig-Hannover (Hannover)"
	case AreaWestphalia:
		return "Deutsche Rentenversicherung Westfalen"
	case AreaHesse:
		return "Deutsche Rentenversicherung AreaHesse"
	case AreaRhineProvince:
		return "Deutsche Rentenversicherung Rheinland (Rheinprovinz)"
	case AreaUpperBavaria:
		return "Deutsche Rentenversicherung Bayern-Süd (Oberbayern)"
	case AreaLowerBavariaUpperPalatinate:
		return "Deutsche Rentenversicherung Bayern-Süd (Niederbayern-Oberpfalz)"
	case AreaRhinelandPalatinate:
		return "Deutsche Rentenversicherung Rheinland-Pfalz"
	case AreaSaarland:
		return "Deutsche Rentenversicherung Saarland"
	case AreaUpperAndMiddleFranconia:
		return "Deutsche Rentenversicherung Nordbayern (Ober- und Mittelfranken)"
	case AreaHamburg:
		return "Deutsche Rentenversicherung Nord (Hamburg)"
	case AreaLowerFranconia:
		return "Deutsche Rentenversicherung Nordbayern (Unterfranken)"
	case AreaSwabia:
		return "Deutsche Rentenversicherung Schwaben"
	case AreaWurttemberg:
		return "Deutsche Rentenversicherung Baden-Württemberg (Württemberg)"
	case AreaBaden:
		return "Deutsche Rentenversicherung Baden-Württemberg (Baden)"
	case AreaBerlin:
		return "Deutsche Rentenversicherung Berlin-Brandenburg (Berlin)"
	case AreaSchleswigHolstein:
		return "Deutsche Rentenversicherung Nord (Schleswig-Holstein)"
	case AreaOldenburgBremen:
		return "Deutsche Rentenversicherung Oldenburg-Bremen"
	case AreaBraunschweig:
		return "Deutsche Rentenversicherung Braunschweig-Hannover (Braunschweig)"
	case AreaKnappschaftBahnSee_Bahn:
		return "Deutsche Rentenversicherung Knappschaft-Bahn-See (Wirtschaftsbereich Bahn)"
	case AreaKnappschaftBahnSee_See:
		return "Deutsche Rentenversicherung Knappschaft-Bahn-See (Wirtschaftsbereich Seefahrt)"
	case AreaZulagenstelleFuerAltersvermoegen:
		return "Zentrale Zulagenstelle für Altersvermögen"
	case AreaBundMecklenburgWesternPomerania, AreaBundThuringia,
		AreaBundBrandenburg, AreaBundSaxonyAnhalt,
		AreaBundSaxony, AreaBundHanover,
		AreaBundWestphalia, AreaBundHesse, AreaBundRhineProvince,
		AreaBundUpperBavaria, AreaBundLowerBavariaUpperPalatinate,
		AreaBundRhinelandPalatinate, AreaBundSaarland, AreaBundUpperAndMiddleFranconia,
		AreaBundHamburg, AreaBundLowerFranconia, AreaBundSwabia,
		AreaBundWurttemberg, AreaBundBaden,
		AreaBundBerlin, AreaBundSchleswigHolstein,
		AreaBundOldenburgBremen, AreaBundBraunschweig,
		AreaBundKnappschaftBahnSee_Bahn, AreaBundKnappschaftBahnSee_See:
		return "Deutsche Rentenversicherung Bund"
	case AreaKnappschaftBahnSeeBerBreHamLoSaxSchleHolWestphalia:
		return "Deutsche Rentenversicherung Knappschaft-Bahn-See (Berlin, Bremen, Hamburg, Niedersachsen, Westfalen und Schleswig-Holstein)"
	case AreaKnappschaftBahnSeeHesseRhineProvinz:
		return "Deutsche Rentenversicherung Knappschaft-Bahn-See (AreaHesse und Rheinprovinz)"
	case AreaKnappschaftBahnSeeBaWuBavariaRhinePalatSaar:
		return "Deutsche Rentenversicherung Knappschaft-Bahn-See (Baden-Württemberg, Bayern, Rheinland-Pfalz und Saarland)"
	case AreaKnappschaftBahnSeeBrandMeckWesPomSaxAnhSaxThu:
		return "Deutsche Rentenversicherung Knappschaft-Bahn-See (Brandenburg, Mecklenburg-Vorpommern, Sachsen-Anhalt, Sachsen und Thüringen)"
	default:
		return "invalid"
	}
}

func (c AreaCode) IsValid() bool {
	switch c {
	case AreaMecklenburgWesternPomerania:
	case AreaThuringia:
	case AreaBrandenburg:
	case AreaSaxonyAnhalt:
	case AreaSaxony:
	case AreaHanover:
	case AreaWestphalia:
	case AreaHesse:
	case AreaRhineProvince:
	case AreaUpperBavaria:
	case AreaLowerBavariaUpperPalatinate:
	case AreaRhinelandPalatinate:
	case AreaSaarland:
	case AreaUpperAndMiddleFranconia:
	case AreaHamburg:
	case AreaLowerFranconia:
	case AreaSwabia:
	case AreaWurttemberg:
	case AreaBaden:
	case AreaBerlin:
	case AreaSchleswigHolstein:
	case AreaOldenburgBremen:
	case AreaBraunschweig:
	case AreaKnappschaftBahnSee_Bahn:
	case AreaKnappschaftBahnSee_See:
	case AreaZulagenstelleFuerAltersvermoegen:
	case AreaBundMecklenburgWesternPomerania:
	case AreaBundThuringia:
	case AreaBundBrandenburg:
	case AreaBundSaxonyAnhalt:
	case AreaBundSaxony:
	case AreaBundHanover:
	case AreaBundWestphalia:
	case AreaBundHesse:
	case AreaBundRhineProvince:
	case AreaBundUpperBavaria:
	case AreaBundLowerBavariaUpperPalatinate:
	case AreaBundRhinelandPalatinate:
	case AreaBundSaarland:
	case AreaBundUpperAndMiddleFranconia:
	case AreaBundHamburg:
	case AreaBundLowerFranconia:
	case AreaBundSwabia:
	case AreaBundWurttemberg:
	case AreaBundBaden:
	case AreaBundBerlin:
	case AreaBundSchleswigHolstein:
	case AreaBundOldenburgBremen:
	case AreaBundBraunschweig:
	case AreaBundKnappschaftBahnSee_Bahn:
	case AreaBundKnappschaftBahnSee_See:
	case AreaKnappschaftBahnSeeBerBreHamLoSaxSchleHolWestphalia:
	case AreaKnappschaftBahnSeeHesseRhineProvinz:
	case AreaKnappschaftBahnSeeBaWuBavariaRhinePalatSaar:
	case AreaKnappschaftBahnSeeBrandMeckWesPomSaxAnhSaxThu:
	default:
		return false
	}

	return true
}
