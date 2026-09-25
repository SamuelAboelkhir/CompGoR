# CompGoR
- The name "CompGoR" comes from the fact that this is going to be a `[comp]rehensive scientific toolkit` built with Go and R (so far)
	- Namig is not my strong suit. The name is subject to change

## Motivation
- This project is going to be a reflection of my education, of everything I know
- It will combine
	- Biology
	- Chemistry
	- Data sciences
	- Bioinformatics
	- Medical sciences
	- Full stack development
	- Low level concepts
	- High level concepts
	- Polyglot architecture
- And yet, I have no idea what I'm even building, I just know that I want to build it
- I expect this project to live with me for a quite a while
## Quick Start
- To install this tool run `go install github.com/SamuelAboelkhir/CompGoR`

## Usage
- This tool is usable both as a normal CLI or a REPL
- To use it as a REPL just type `CompGoR` with no additional arguments then type `help`
``` shell
❯ CompGoR
CompGoR > help
Welcome to the CompGoR CLI!
Available Commands:
-------------------
showTable: Renders a periodic table using the system's default image renderer.

fetchJSONData: Fetches generic data of JSON format.

buildPubChemQuery: Takes a domain, namespace, identifier, an optional operation,
	and an output type, and queries the PubChem API for a matching
	chemical substance or compound
	Example: domain: compound,
		 namespace: name,
		 identifier: hydrogen,
		 operation: <optional, ENTER to skip>,
		 format: JSON

help: Prints this help menu

CompGoR >
```
- Or as a CLI
``` shell
❯ CompGoR help
Welcome to the CompGoR CLI!
Available Commands:
-------------------
buildPubChemQuery: Takes a domain, namespace, identifier, an optional operation,
	and an output type, and queries the PubChem API for a matching
	chemical substance or compound
	Example: domain: compound,
		 namespace: name,
		 identifier: hydrogen,
		 operation: <optional, ENTER to skip>,
		 format: JSON

help: Prints this help menu

showTable: Renders a periodic table using the system's default image renderer.

fetchJSONData: Fetches generic data of JSON format.
```
- CLI usage was more of an after thought, so the intended usage is as a REPL

## The stack
- As of 03-05-2026
	- Go 1.26.2
	- R 4.5.3
	- RStudio
```R
> version
               _                           
platform       x86_64-pc-linux-gnu         
arch           x86_64                      
os             linux-gnu                   
system         x86_64, linux-gnu           
status                                     
major          4                           
minor          5.3                         
year           2026                        
month          03                          
day            11                          
svn rev        89597                       
language       R                           
version.string R version 4.5.3 (2026-03-11)
nickname       Reassured Reassurer
```

## Current execution
- The project currently is a repl CLI tool that utilizes Go as an HTTP client querying PubChem's REST API to fetch chemical data
- The core function so far is a query builder that takes user input and assembles it into a valid PubChem URL where the data is fetched from
- The app also features 
	- A caching system that make use of Go's concurrency
	- A generic API client (I call it the postman) that can fetch any generic data
	- Wrapper functions around the generic API client that properly type the data
	- A command interface that any new command must satisfy
	- A commands struct that stores a map of all the available commands

## Future ideas
- Add a postgres database that optionally stores data based on user input
- A dockerized R server in a small private network with the main Go system for that handles data analytics
- Various scientific tools
- A clients interface that allows for extending the app with clients for connecting with different data sources
- A generic query builder, and dedicated query builders for very specific data sources (like PubChem)
- Proper database connection handling (retry logic, disconnection logic)
- Request rate limiting, as to not accidentally crash any API

## Contributing
- If you care to contribute then feel free to clone the repo an go wild
- This software is fairly opinionated, and personal, but anyone with a science background and enough technical skill to offer something is welcome, especially since the tool is still in its infancy
- If you want to offer technical or scientific feedback, you're also more than welcome
