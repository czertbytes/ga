package ga

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

var (
	ErrParserContentWrongFormat = errors.New("parser: Content has wrong format!")
)

type Parser struct {
	value string
}

func NewParser(value string) *Parser {
	return &Parser{
		value: value,
	}
}

// Returns Timestamp and UserId from Apache log payload
// Log format is following:
// 127.0.0.1 - hostname [<timestamp>] "GET /<userId>/foo/bar HTTP/1.1" 200 2326
func (self *Parser) ParseTimestampUserId() (int64, int, error) {
	parts := strings.Split(self.value, " ")
	if len(parts) != 9 {
		log.Printf("parser: Expected 9 parts but got %d", len(parts))
		return 0, 0, ErrParserContentWrongFormat
	}

	timestamp, err := self.parseTimestamp(parts[3])
	if err != nil {
		return 0, 0, err
	}

	userId, err := self.parseUserId(parts[5])
	if err != nil {
		return 0, 0, err
	}

	return timestamp, userId, nil
}

func (self *Parser) parseTimestamp(value string) (int64, error) {
	timestampPart := strings.Trim(value, "[]")
	timestamp, err := strconv.ParseInt(timestampPart, 10, 64)
	if err != nil {
		log.Printf("parser: Parsing '%s' as timestamp failed! Error: %s", timestampPart, err.Error())
		return 0, ErrParserContentWrongFormat
	}

	if timestamp < 0 {
		log.Printf("parser: Parsing '%s' as timestamp failed! Error: can't be negative", timestampPart)
		return 0, ErrParserContentWrongFormat
	}

	return timestamp, err
}

func (self *Parser) parseUserId(value string) (int, error) {
	userIdParts := strings.Split(value, "/")
	if len(userIdParts) < 1 {
		log.Printf("parser: Parsing '%s' as userId failed! Error: wrong path", value)
		return 0, ErrParserContentWrongFormat
	}

	userId, err := strconv.ParseInt(userIdParts[1], 10, 64)
	if err != nil {
		log.Printf("parser: Parsing '%s' as userId failed! Error: %s", userIdParts[1], err.Error())
		return 0, ErrParserContentWrongFormat
	}

	if userId < 0 {
		log.Printf("parser: Parsing '%s' as userId failed! Error: can't be negative", userIdParts[1])
		return 0, ErrParserContentWrongFormat
	}

	return int(userId), nil
}
