package entry

// FromWordPairCsv erzeugt ein neues Entry-Objekt aus einem String,
// der ein Wortpaar enthält. Das deutsche und das englische Wort sind
// durch ein Komma getrennt.
// Gibt es keines oder mehrere Kommas im String, wird ein leerer Eintrag zurückgegeben.
func FromWordPairCsv(pair string) Entry {

	string_de := ""
	string_en := ""

	cmflg := false

	for _, speci := range pair {

		//error catchen
		if string(speci) == " " {

			return Entry{}

		}

		if string(speci) == "," {

			// error catchen

			if cmflg {

				return Entry{}

			}

			cmflg = true
			continue

		}

		if cmflg {

			string_en += string(speci)

		} else {

			string_de += string(speci)

		}

	}

	var entry Entry

	entry.de = string_de
	entry.en = string_en

	return entry
}
