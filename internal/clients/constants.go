package clients

import "time"

// Constants for the default client
const (
	DefaultTimeout        = 5 * time.Second
	DefaultCacheTimeout   = 5 * time.Minute
	DefaultClientProtocol = "HTTP"
)

// Constants for the PubChem client
const (
	BasePubChemURL        = "https://pubchem.ncbi.nlm.nih.gov/rest/pug"
	PubChemTimeout        = 10 * time.Second
	PubChemCacheTimeout   = 5 * time.Minute
	PubChemClientProtocol = "HTTP"
)
