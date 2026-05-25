package clients

type Substances struct {
	PCSubstances []struct {
		Sid struct {
			ID      int `json:"id"`
			Version int `json:"version"`
		} `json:"sid"`
		Source struct {
			Db struct {
				Name     string `json:"name"`
				SourceID struct {
					Str string `json:"str"`
				} `json:"source_id"`
			} `json:"db"`
		} `json:"source"`
		Synonyms []string `json:"synonyms"`
		Comment  []string `json:"comment,omitempty"`
		Xref     []struct {
			Dburl string `json:"dburl,omitempty"`
			Rn    string `json:"rn,omitempty"`
			Sburl string `json:"sburl,omitempty"`
			Regid string `json:"regid,omitempty"`
		} `json:"xref"`
		Compound []struct {
			ID struct {
				Type int `json:"type"`
			} `json:"id,omitempty"`
			Atoms struct {
				Aid     []int `json:"aid"`
				Element []int `json:"element"`
			} `json:"atoms,omitempty"`
			Bonds struct {
				Aid1  []int `json:"aid1"`
				Aid2  []int `json:"aid2"`
				Order []int `json:"order"`
			} `json:"bonds,omitempty"`
			Coords []struct {
				Type       []int `json:"type"`
				Aid        []int `json:"aid"`
				Conformers []struct {
					X []float64 `json:"x"`
					Y []float64 `json:"y"`
				} `json:"conformers"`
			} `json:"coords,omitempty"`
			Charge int `json:"charge,omitempty"`
			ID0    struct {
				Type int `json:"type"`
				ID   struct {
					Cid int `json:"cid"`
				} `json:"id"`
			} `json:"id,omitempty"`
		} `json:"compound,omitempty"`
	} `json:"PC_Substances"`
}

type Compounds struct {
	PCCompounds []struct {
		ID struct {
			ID struct {
				Cid int `json:"cid"`
			} `json:"id"`
		} `json:"id"`
		Atoms struct {
			Aid     []int `json:"aid"`
			Element []int `json:"element"`
		} `json:"atoms"`
		Bonds struct {
			Aid1  []int `json:"aid1"`
			Aid2  []int `json:"aid2"`
			Order []int `json:"order"`
		} `json:"bonds"`
		Coords []struct {
			Type       []int `json:"type"`
			Aid        []int `json:"aid"`
			Conformers []struct {
				X []int `json:"x"`
				Y []int `json:"y"`
			} `json:"conformers"`
		} `json:"coords"`
		Charge int `json:"charge"`
		Props  []struct {
			Urn struct {
				Label          string `json:"label"`
				Name           string `json:"name"`
				Datatype       int    `json:"datatype"`
				Parameters     string `json:"parameters,omitempty"`
				Implementation string `json:"implementation,omitempty"`
				Version        string `json:"version,omitempty"`
				Software       string `json:"software,omitempty"`
				Source         string `json:"source,omitempty"`
				Release        string `json:"release"`
			} `json:"urn"`
			Value struct {
				Ival int `json:"ival"`
			} `json:"value"`
		} `json:"props"`
		Count struct {
			HeavyAtom       int `json:"heavy_atom"`
			AtomChiral      int `json:"atom_chiral"`
			AtomChiralDef   int `json:"atom_chiral_def"`
			AtomChiralUndef int `json:"atom_chiral_undef"`
			BondChiral      int `json:"bond_chiral"`
			BondChiralDef   int `json:"bond_chiral_def"`
			BondChiralUndef int `json:"bond_chiral_undef"`
			IsotopeAtom     int `json:"isotope_atom"`
			CovalentUnit    int `json:"covalent_unit"`
			Tautomers       int `json:"tautomers"`
		} `json:"count"`
	} `json:"PC_Compounds"`
}
