## Manage your feature toggles easily.

### Usage
```go
func main() {
	
	isEnabled := LaunchLever(./path/to/json)
	dark_mode := isEnabled("DarkMode")
	fmt.Println("Is dark mode enabled?", enabled)
}

```
