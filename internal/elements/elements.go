package elements

func ConsoleMessages(lines []string) string {
	result := `
		<div
		  id="console"
		  class="block bg-black border-2 border-gray-500 rounded-sm max-w-xl h-48"
		>
		`
	messageConsoleLine := func(msg string) string {
		start := `<div class="flex">
			  <p class="font-bold mx-2">></p>
			  <p class="text-green-500">`
		end := `</p></div>`

		return start + msg + end

	}
	for _, line := range lines {
		result += messageConsoleLine(line)
	}

	return result + "</div>"
}
