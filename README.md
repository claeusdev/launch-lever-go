## Manage your feature toggles easily.

### Usage
```go
func main() {
	flagJson := `[
		{
			"Name": "DarkMode",
			"Description": "Enables the dark mode theme across the application",
			"Status": "Enabled"
		},
		{
			"Name": "FeatureX",
			"Description": "Activates Feature X in the application",
			"Status": "Disabled"
		}
	]`
	getFlagByName := LaunchLever(flagJson)
	var darkMode Toggle
	if toggle, found := getFlagByName("DarkMode"); found {
		darkMode = toggle
	}
	fmt.Println("Hello, 世界", darkMode.Status)
}

```
