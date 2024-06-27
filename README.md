## Manage your feature toggles easily.

### Usage
```go
func main() {
	
	isEnabled := LaunchLever(flagJson)
	dark_mode := isEnabled("DarkMode")
	fmt.Println("Is dark mode enabled?", enabled)
}

```
