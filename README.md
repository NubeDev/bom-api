# bom-api
Australian Weather Data (using bom.gov.au)

Example

```golang
search, geo, err := client.SearchByTown("Helensburgh", "NSW")
if err != nil {
  fmt.Print(err)
}
```
