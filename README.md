# Orc Library

Orc library generates Orc names, greetings, and weapons.

## Usage

``` go
package main

import (
	"fmt"

	"github.com/alimac/orc"
)

func main() {
	fmt.Printf("%s greets you with \"%s!\", holding a %s \n",
		orc.GenerateName(), orc.GenerateGreeting(), orc.GenerateWeapon())
}
```

## Example output

```
Guldug greets you with "Time to die!", holding a DismalTrowel
```
