package utils

import (
    "unicode"
)

func IsValidTopicName(name string) bool {
    for _, char := range name {
        if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
            return false
        }
    }
    return true
}