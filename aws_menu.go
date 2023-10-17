package aws_menu

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/nchillal/aws_profiles"
)

func PrintAwsProfileMenu() (string, error) {
	// Get list of profiles configured
	profiles, err := aws_profiles.ListAWSProfiles()
	if err != nil {
		return "", err
	}
	profileSearcher := func(input string, index int) bool {
		profile := profiles[index]
		name := strings.Replace(strings.ToLower(profile), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	// Create a Select template with custom formatting
	profileTemplate := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "🔥 {{ . | cyan }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "\U0001F336 {{ . | red | cyan }}",
	}

	// Prompt profiles
	promptProfile := promptui.Select{
		Label:        "Select AWS Profile",
		Items:        profiles,
		Size:         len(profiles),
		HideSelected: true,
		Templates:    profileTemplate,
		Searcher:     profileSearcher,
	}

	_, awsProfile, err := promptProfile.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return awsProfile, nil
}

func PrintAwsRegionMenu(awsProfile string) (string, error) {
	regions, err := aws_profiles.ListAWSRegions(awsProfile)
	if err != nil {
		fmt.Println(err)
	}
	regionSearcher := func(input string, index int) bool {
		region := regions[index]
		name := strings.Replace(strings.ToLower(region), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	// Create a Select template with custom formatting
	regionTemplate := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "🔥 {{ . | cyan }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "\U0001F336 {{ . | red | cyan }}",
	}

	// Prompt regions
	promptRegion := promptui.Select{
		Label:        "Select AWS Regions",
		Items:        regions,
		Size:         len(regions),
		HideSelected: true,
		Templates:    regionTemplate,
		Searcher:     regionSearcher,
	}

	_, awsRegion, err := promptRegion.Run()

	if err != nil {
		return "", nil
	}

	return awsRegion, nil
}
