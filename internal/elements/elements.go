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

func SkillBlocks(skills []string) string {
	result := `
	  <div class="max-w-3xl mx-auto px-4 py-4">
	    <h2 class="font-semibold">Skills:</h2>
	    <div class="flex flex-wrap gap-2">
		`
	skillWrap := func(s string) string {
		start := `
	      	<p class="block text-center bg-gray-200 dark:bg-gray-700 text-gray-900 dark:text-gray-100 text-lg py-2 px-4 rounded-lg">
		`
		end := "</p>"

		return start + s + end

	}
	for _, skill := range skills {
		result += skillWrap(skill)
	}

	return result + "</div>"
}
