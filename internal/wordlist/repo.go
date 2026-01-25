package wordlist

import (
	"fmt"
	"strings"
)

var Words = []string{
	"the", "quick", "brown", "fox", "jumps", "over", "lazy", "dog",
	"hello", "world", "programming", "keyboard", "computer", "screen",
	"typing", "practice", "speed", "accuracy", "fingers", "hands",
	"apple", "banana", "orange", "grape", "mango", "cherry", "lemon",
	"house", "garden", "window", "door", "floor", "ceiling", "wall",
	"river", "mountain", "ocean", "forest", "desert", "valley", "hill",
	"running", "walking", "jumping", "swimming", "flying", "climbing",
	"happy", "sad", "angry", "excited", "calm", "nervous", "brave",
	"table", "chair", "lamp", "book", "pencil", "paper", "notebook",
	"morning", "evening", "night", "afternoon", "sunrise", "sunset",
	"coffee", "tea", "water", "juice", "milk", "bread", "butter",
	"music", "dance", "song", "rhythm", "melody", "harmony", "beat",
	"cloud", "rain", "snow", "wind", "storm", "thunder", "lightning",
	"friend", "family", "parent", "child", "brother", "sister", "cousin",
	"school", "teacher", "student", "lesson", "homework", "exam", "grade",
	"travel", "journey", "adventure", "explore", "discover", "wander",
	"create", "design", "build", "develop", "improve", "enhance", "grow",
	"simple", "complex", "easy", "difficult", "basic", "advanced", "expert",
	"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "theta",
	"pizza", "pasta", "salad", "soup", "sandwich", "burger", "fries",
}

func DebugWords() {
	text := strings.Join(Words, " ")
	for i, char := range text {
		fmt.Printf("%d%c ", i, char)
	}
}
