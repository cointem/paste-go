package sql

import (
	"fmt"
	"regexp"
	"strings"

	"paste-go/pkg/parser"
	"paste-go/pkg/schema"
)

type SQLParser struct{}

func NewSQLParser() parser.Parser {
	return &SQLParser{}
}

func (p *SQLParser) Name() string {
	return "sql"
}

func (p *SQLParser) CanParse(content string) bool {
	content = strings.TrimSpace(strings.ToUpper(content))
	return strings.Contains(content, "CREATE TABLE")
}

func (p *SQLParser) Parse(content string) (*schema.Struct, error) {
	// 1. Get Table Name
	tableNameParams := regexp.MustCompile(`(?i)CREATE\s+TABLE\s+["` + "`" + `]?(\w+)["` + "`" + `]?`)
	matches := tableNameParams.FindStringSubmatch(content)
	if len(matches) < 2 {
		return nil, fmt.Errorf("could not find table name")
	}
	tableName := matches[1]

	result := &schema.Struct{
		Name:   toPascalCase(tableName),
		Fields: []schema.Field{},
	}

	// 2. Extract Columns
	startIdx := strings.Index(content, "(")
	endIdx := strings.LastIndex(content, ")")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("invalid sql syntax")
	}

	columnSection := content[startIdx+1 : endIdx]
	lines := strings.Split(columnSection, ",")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		upperLine := strings.ToUpper(line)
		if line == "" || strings.HasPrefix(upperLine, "PRIMARY") || strings.HasPrefix(upperLine, "KEY") || strings.HasPrefix(upperLine, "CONSTRAINT") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		colName := strings.Trim(parts[0], "`\"")
		sqlType := strings.ToUpper(parts[1])

		field := schema.Field{
			Name:         toPascalCase(colName),
			OriginalName: colName,
			Kind:         schema.KindString, // Default
		}

		if strings.Contains(sqlType, "INT") {
			field.Kind = schema.KindInt
		} else if strings.Contains(sqlType, "BOOL") {
			field.Kind = schema.KindBool
		} else if strings.Contains(sqlType, "FLOAT") || strings.Contains(sqlType, "DOUBLE") || strings.Contains(sqlType, "DECIMAL") {
			field.Kind = schema.KindFloat
		} else if strings.Contains(sqlType, "TIME") || strings.Contains(sqlType, "DATE") {
			field.Kind = schema.KindTime
		}

		result.Fields = append(result.Fields, field)
	}

	return result, nil
}

func toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
		}
	}
	return strings.Join(parts, "")
}
