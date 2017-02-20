package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	rnd         = rand.NewSource(time.Now().UnixNano())
	letterRunes = []rune(letterBytes + "éèä®ŷÏÄÃÒÉ⁄™‹›")
	days        = []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}
	months      = []time.Month{time.January, time.February, time.March, time.April, time.May, time.June, time.July, time.August, time.September, time.October, time.November, time.December}
	// year: 2006, month: 01, day: 02
	datePatterns = []string{"20060102", "2006/01/02", "2006-01-02", "2006.01.02", "20060201", "2006/02/01", "2006-02-01", "2006.02.01"}
	// hour: 15, minute: 04, second: 05
	timePatterns = []string{"15:04:05"}
)

// RandStringBytesMaskImprSrc ASCII
func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A rnd.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, rnd.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rnd.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// RandStringRunes UTF-8
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandNumerical(n int) string {
	return strconv.Itoa(rand.Intn(n) * rand.Intn(10))
}

func RandDate() func(int) string {
	pattern := randomFrom(datePatterns)
	return func(int) string {
		return time.Date(1000+rand.Intn(1017), months[rand.Intn(len(months))], int(days[rand.Intn(len(days))]), 0, 0, 0, 0, time.UTC).Format(pattern)
	}
}

func RandTime() func(int) string {
	pattern := randomFrom(timePatterns)
	return func(int) string {
		return time.Date(2017, time.January, 15, rand.Intn(24), rand.Intn(60), rand.Intn(60), 0, time.UTC).Format(pattern)
	}
}

func RandTimestamp() func(int) string {
	date := RandDate()
	time := RandTime()
	return func(n int) string {
		return fmt.Sprintf("%s %s", date(n), time(n))
	}
}

func randomFrom(sl []string) string {
	return sl[rand.Intn(len(sl))]
}
