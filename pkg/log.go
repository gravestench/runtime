package pkg

import (
	"fmt"
	"hash/fnv"
	"image/color"
	"math"
	"math/rand"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// newLogger creates a factory function that generates a
// zerolog.Logger instance with a prefix and randomized colors based on the
// provided name. The luminosity delta between the colors is greater than 0.5.
func newLogger(name string) *zerolog.Logger {
	c1 := getRandomColor(name)
	c2 := getContrastingColor(c1, name)

	foregroundEscape := getRGBEscapeSequence(c2, false)
	backgroundEscape := getRGBEscapeSequence(c1, true)
	emoji := getRandomFoodEmoji(name)
	resetEscape := "\x1b[0m"

	format := fmt.Sprintf(
		"[%s%s%s %s %s]",
		foregroundEscape,
		backgroundEscape,
		emoji,
		name,
		resetEscape,
	)

	writer := zerolog.ConsoleWriter{
		Out: os.Stdout,
		FormatMessage: func(input any) string {
			return fmt.Sprintf("%s: %s", format, input)
		},
	}

	logger := log.Output(writer).With().Logger()

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return &logger
}

// getRGBEscapeSequence returns the ANSI escape sequence for the given color.Color.
func getRGBEscapeSequence(c color.Color, foreground bool) string {
	r, g, b, _ := c.RGBA()

	if foreground {
		return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r>>8, g>>8, b>>8)
	}

	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r>>8, g>>8, b>>8)
}

// getContrastingColor returns a contrasting color based on the input baseColor
// and the name string.
func getContrastingColor(baseColor color.Color, name string) color.Color {
	baseRGB := convertToRGBA(baseColor)
	baseLuminosity := getLuminosity(baseRGB)

	// Use a consistent seed based on the hash of the name string
	seed := hashString(name)
	rng := rand.New(rand.NewSource(int64(seed)))

	// Generate random RGB values until a contrasting color is found
	for {
		// Generate random RGB values
		r := rng.Intn(256)
		g := rng.Intn(256)
		b := rng.Intn(256)

		contrastColor := color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		contrastRGB := convertToRGBA(contrastColor)
		contrastLuminosity := getLuminosity(contrastRGB)

		// Calculate the luminosity delta
		luminosityDelta := math.Abs(contrastLuminosity - baseLuminosity)

		if luminosityDelta > 0.5 {
			return contrastColor
		}
	}
}

// convertToRGBA converts a color.Color to color.RGBA.
func convertToRGBA(c color.Color) color.RGBA {
	if rgba, ok := c.(color.RGBA); ok {
		return rgba
	}
	r, g, b, a := c.RGBA()
	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

// getLuminosity calculates the luminosity value of an RGB color.
func getLuminosity(c color.RGBA) float64 {
	r, g, b := float64(c.R)/255.0, float64(c.G)/255.0, float64(c.B)/255.0
	return 0.2126*r + 0.7152*g + 0.0722*b
}

// getRandomColor generates a random color.Color based on the input string.
func getRandomColor(input string) color.Color {
	// Initialize the hash function
	hash := fnv.New32a()
	hash.Write([]byte(input))
	hashValue := hash.Sum32()

	// Use the hash value as the seed for random number generation
	rng := rand.New(rand.NewSource(int64(hashValue)))

	// Generate random RGB values
	r := rng.Intn(256)
	g := rng.Intn(256)
	b := rng.Intn(256)

	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

// hashString generates a hash value for the given input string.
func hashString(s string) uint32 {
	hash := uint32(5381)
	for i := 0; i < len(s); i++ {
		hash = (hash << 5) + hash + uint32(s[i])
	}
	return hash
}

// getRandomFoodEmoji selects a random food emoji based on the input name string.
func getRandomFoodEmoji(name string) string {
	// Convert the name string to lowercase and remove spaces
	name = strings.ToLower(strings.ReplaceAll(name, " ", ""))

	// Calculate the hash value of the name string
	hash := fnv.New32a()
	hash.Write([]byte(name))
	hashValue := hash.Sum32()

	// Use the hash value as the seed for random number generation
	rng := rand.New(rand.NewSource(int64(hashValue)))

	// List of food emojis
	foodEmojis := []string{
		"🍇", "🍈", "🍉", "🍊", "🍋", "🍌", "🍍", "🥭", "🍎", "🍏",
		"🍐", "🍑", "🍒", "🍓", "🫐", "🥝", "🍅", "🫒", "🥥", "🥑",
		"🍆", "🥔", "🥕", "🌽", "🌶️", "🫑", "🥒", "🥬", "🥦", "🧄",
		"🧅", "🍄", "🥜", "🌰", "🍞", "🥐", "🥖", "🫓", "🥨", "🥯",
		"🥞", "🧇", "🧀", "🍖", "🍗", "🥩", "🍔", "🍟", "🍕", "🌭",
		"🥪", "🌮", "🌯", "🫔", "🥙", "🧆", "🥚", "🍳", "🥘", "🍲",
		"🫕", "🥣", "🥗", "🍿", "🧈", "🧂", "🥫", "🍱", "🍘", "🍙",
		"🍚", "🍛", "🍜", "🍝", "🍠", "🍢", "🍣", "🍤", "🍥", "🥮",
		"🍡", "🥟", "🥠", "🥡", "🦀", "🦞", "🦐", "🦑", "🦪", "🍦",
		"🍧", "🍨", "🍩", "🍪", "🎂", "🍰", "🧁", "🥧", "🍫", "🍬",
		"🍭", "🍮", "🍯", "🍼", "🥛", "☕", "🍵", "🍶", "🍾", "🍷",
		"🍸", "🍹", "🍺", "🍻", "🥂", "🥃", "🥤", "🧃", "🧉", "🧊",
		"🍴", "🍽️", "🥣", "🥡", "🥢", "🥄", "🫕", "🍲", "🥘", "🍝",
		"🥗", "🍛", "🍜", "🥪", "🥟", "🌭", "🍔", "🥠", "🍕", "🍟",
		"🍚", "🌮", "🌯", "🍣", "🍱", "🍢", "🍱", "🍣", "🥠", "🍘",
		"🍥", "🍙", "🍚", "🍛", "🍜", "🍝", "🍠", "🍡", "🍢", "🍣",
		"🍤", "🍥", "🍡", "🍿", "🥫", "🍚", "🍛", "🍜", "🍝", "🍞",
		"🍟", "🍠", "🍡", "🍢", "🍣", "🍤", "🍥", "🥮", "🍦", "🍧",
		"🍨", "🍩", "🍪", "🎂", "🍰", "🧁", "🥧", "🍫", "🍬", "🍭",
		"🍮", "🍯", "🍼", "🥛", "☕", "🍵", "🍶", "🍾", "🍷", "🍸",
		"🍹", "🍺", "🍻", "🥂", "🥃", "🥤", "🧃", "🧉", "🧊", "🥢",
		"🥄", "🥂", "🍾", "🥃", "🥤", "🧃", "🧉", "🧊", "🍴", "🍽️",
		"🥣", "🥡", "🥢", "🥄", "🍷", "🍸", "🍹", "🍺", "🍻", "🥂",
		"🥃", "🥤", "🧃", "🧉", "🧊", "🍺", "🍻", "🥂", "🥃", "🥤",
		"🧃", "🧉", "🧊", "🍺", "🍻", "🥂", "🥃", "🥤", "🧃", "🧉",
		"🧊", "🍺", "🍻", "🥂", "🥃", "🥤", "🧃", "🧉", "🧊", "🍺",
		"🍻", "🥂", "🥃", "🥤", "🧃", "🧉", "🧊", "🍺", "🍻", "🥂",
		"🥃", "🥤", "🧃", "🧉", "🧊",
	}

	// Select a random food emoji
	index := rng.Intn(len(foodEmojis))
	return foodEmojis[index]
}
